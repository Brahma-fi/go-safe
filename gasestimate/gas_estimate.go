package gasestimate

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/Brahma-fi/go-safe/pkg/logger"
	"github.com/Brahma-fi/go-safe/types"
	"github.com/Brahma-fi/go-safe/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/go-resty/resty/v2"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	SafeBufferLimit     = 0.1
	SimulateTxnAccessor = "SimulateTxnAccessor"
)

var (
	ErrInvalidChainID = errors.New("invalid chainID")
)

// common resty client
// safe on concurrent use
var commonClient = resty.New()

type gasEstimatorFactory interface {
	RetryableGasEstimator(chainID int64) (ethereum.GasEstimator, error)
}

type Clients interface {
	ethereum.BlockNumberReader
	ethereum.ChainIDReader
	ethereum.GasEstimator
	ethereum.ChainStateReader
	ethereum.GasPricer
	ethereum.GasPricer1559
	ethereum.TransactionReader
	bind.ContractCaller
	bind.DeployBackend
}

type estimateGasResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

type ethTransaction struct {
	From string `json:"from,omitempty"`
	To   string `json:"to"`
	Data string `json:"data"`
}

type ethTransactionWithGas struct {
	ethTransaction
	GasPrice string `json:"gasPrice,omitempty"`
	GasLimit string `json:"gas,omitempty"`
}

type estimateGasRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

type Estimation struct {
	gasEstimatorFactory gasEstimatorFactory
	addressRegistry     types.AddressRegistry

	safeAbi     *abi.ABI
	accessorAbi *abi.ABI
	clientURLs  map[int64]string // map[chainID]ethRpcURL
	logger      logger.Logger
}

func NewGasEstimation(
	clientFactory gasEstimatorFactory,
	addRegistry types.AddressRegistry,
	safeAbi *abi.ABI,
	accessorAbi *abi.ABI,
	clientURLs map[int64]string,
	logger logger.Logger,
) *Estimation {
	return &Estimation{
		gasEstimatorFactory: clientFactory,
		addressRegistry:     addRegistry,
		safeAbi:             safeAbi,
		accessorAbi:         accessorAbi,
		clientURLs:          clientURLs,
		logger:              logger,
	}
}

// EstimateSafeGasv1_3_0 this estimates the max gas limit that should be given in form of safeTxGas and is compatible
// with safe version 1.3.0
// see https://github.com/safe-global/safe-core-sdk/blob/7959821ab08c96cf3babb9ed906c01d644ac49f4/packages/protocol-kit/src/utils/transactions/gas.ts#L17
//
//nolint:lll
func (g *Estimation) EstimateSafeGasv1_3_0(ctx context.Context, safeTxn *types.SafeTx) (uint64, error) {
	chainID := (*big.Int)(safeTxn.ChainId).Int64()

	safeAddress := safeTxn.Safe.Address()
	encoded, err := g.safeAbi.Pack(
		"requiredTxGas",
		safeTxn.To.Address(), (*big.Int)(&safeTxn.Value), ([]byte)(*safeTxn.Data),
		safeTxn.Operation,
	)
	if err != nil {
		return 0, err
	}
	data, err := g.rawEstimateGasCall(
		ethTransaction{
			From: safeAddress.Hex(),
			To:   safeAddress.Hex(),
			Data: hexutil.Encode(encoded),
		},
		chainID,
	)
	if err != nil {
		return 0, err
	}

	// the required length should be min 64 (uint256)
	if len(data) < 64 {
		if g.logger != nil {
			g.logger.Warn(
				"invalid response from eth_gasEstimate",
				logger.Str("safeAddress", safeAddress.Hex()),
				logger.Str("data", hexutil.Encode(encoded)),
			)
		}
		return g.estimateGasViaEthClient(ctx, safeAddress, safeAddress, safeTxn.Value, safeTxn.Data, chainID)
	}

	// extract last 64 bytes and convert to bigInt which is the used gas
	// see https://github.com/safe-global/safe-core-sdk/blob/7959821ab08c96cf3babb9ed906c01d644ac49f4/packages/protocol-kit/src/utils/transactions/gas.ts#L42
	substr := data[len(data)-64:]
	hexInt, err := hexutil.Decode("0x" + substr)
	if err != nil {
		if g.logger != nil {
			g.logger.Warn(
				"failed to convert resp to big.Int", logger.Err(err),
				logger.Str("safeAddress", safeAddress.Hex()),
				logger.Str("data", hexutil.Encode(encoded)),
			)
		}
		return g.estimateGasViaEthClient(ctx, safeAddress, safeAddress, safeTxn.Value, safeTxn.Data, chainID)
	}

	// refer https://github.com/safe-global/safe-core-sdk/blob/7959821ab08c96cf3babb9ed906c01d644ac49f4/packages/protocol-kit/src/utils/transactions/gas.ts#L42
	txGasEstimation := new(big.Int).SetBytes(hexInt).Uint64() + 10000
	dataGasCost := estimateDataGasCost(encoded)

	// refer https://github.com/safe-global/safe-core-sdk/blob/7959821ab08c96cf3babb9ed906c01d644ac49f4/packages/protocol-kit/src/utils/transactions/gas.ts#L48
	additionalGas := uint64(10000)
	for i := 0; i < 10; i++ {
		response, callRPCErr := g.rawEthCall(
			ethTransactionWithGas{
				ethTransaction: ethTransaction{
					From: safeAddress.Hex(),
					To:   safeAddress.Hex(),
					Data: hexutil.Encode(encoded),
				},
				GasLimit: hexutil.EncodeUint64(txGasEstimation + dataGasCost + additionalGas),
				GasPrice: "0x0",
			},
			chainID,
		)
		if callRPCErr != nil {
			return 0, callRPCErr
		}
		if response != "" {
			break
		}
		txGasEstimation += additionalGas
		additionalGas *= 2
	}
	// add 25K buffer for moderated account transactions
	return txGasEstimation + dataGasCost + additionalGas + 25000, nil
}

// EstimateSafeGasv1_4_0 this estimates the max gas limit that should be given in form of safeTxGas and is compatible
// with safe version 1.4.0
func (g *Estimation) EstimateSafeGasv1_4_0(_ context.Context, safeTxn *types.SafeTx) (uint64, error) {
	chainID := (*big.Int)(safeTxn.ChainId).Int64()

	simAddress, err := g.addressRegistry.GetAddressByChainID(chainID, SimulateTxnAccessor)
	if err != nil {
		return 0, err
	}
	safeAddress := safeTxn.Safe.Address()
	//nolint:lll
	// see https://github.com/safe-global/safe-contracts/blob/7a77545f288361893313af23194988731ee95261/test/accessors/SimulateTxAccessor.spec.ts#L70
	encoded, err := g.accessorAbi.Pack(
		"simulate",
		safeTxn.To.Address(), (*big.Int)(&safeTxn.Value), ([]byte)(*safeTxn.Data),
		safeTxn.Operation,
	)
	if err != nil {
		return 0, err
	}

	simMixedCaseAddress := common.NewMixedcaseAddress(simAddress)
	//nolint:lll
	// see https://github.com/safe-global/safe-contracts/blob/7a77545f288361893313af23194988731ee95261/contracts/common/StorageAccessible.sol#L40
	simulateAndRevert, err := g.safeAbi.Pack(
		"simulateAndRevert",
		simMixedCaseAddress.Address(),
		encoded,
	)
	if err != nil {
		return 0, err
	}
	data, err := g.rawEthCall(
		ethTransactionWithGas{
			ethTransaction: ethTransaction{
				To:   safeAddress.Hex(),
				Data: hexutil.Encode(simulateAndRevert),
			},
		},
		chainID,
	)
	if err != nil {
		return 0, err
	}
	// remove extra text in the revert data
	data = strings.ReplaceAll(data, "Reverted ", "")
	// decode the hex as bytes
	decoded, err := hexutil.Decode(data)
	if err != nil {
		return 0, err
	}
	// the return data is bool(success), bytes(response) from the simulateAndRevert return
	// the response is the return of simulate from `SimulateTxAccessor` which is
	// abi.encode(uint256(estimate),bool(success),bytes(returnData))
	// hence we safely read the estimate as from 64 to 96
	// as [success](32),[response.length](32),[estimate](32),[success](32),[returnData](variable)
	gasUsed, err := utils.Slice(decoded, 64, 96)
	if err != nil {
		return 0, err
	}
	gasLimit := new(big.Int).SetBytes(gasUsed).Uint64()
	// add a `SafeBufferLimit`x buffer to the actual computed value
	return gasLimit + uint64(SafeBufferLimit*float64(gasLimit)), nil
}

// EstimateSafeGas this estimates the max gas limit that should be given in form of safeTxGas
// it selects the appropriate function according to the safe version
func (g *Estimation) EstimateSafeGas(ctx context.Context, safeTxn *types.SafeTx) (uint64, error) {
	return g.EstimateSafeGasv1_4_0(ctx, safeTxn)
}

func estimateDataGasCost(data []byte) (cost uint64) {
	for _, v := range data {
		if v == 0 {
			cost += 4
		} else {
			cost += 16
		}
	}

	return cost
}

// this client calls the rpc using the resty client
// this is done because the client.ethClient and rpc.Client throw errors using error.message
// this does not return the data given in Error object which is used to recover the gasUsed from the safe
// see https://github.com/safe-global/safe-contracts/blob/186a21a74b327f17fc41217a927dea7064f74604/contracts/GnosisSafe.sol#L315
//
//nolint:lll
func (g *Estimation) rawEstimateGasCall(transaction ethTransaction, chainID int64) (string, error) {
	rpcURL, err := g.getURL(chainID)
	if err != nil {
		return "", err
	}

	resp := &estimateGasResponse{}
	_, err = commonClient.SetBaseURL(rpcURL).R().SetBody(
		&estimateGasRequest{
			Jsonrpc: "2.0",
			Method:  "eth_estimateGas",
			Params:  []interface{}{transaction, "latest"},
			ID:      1,
		},
	).SetResult(resp).Post("/")
	if err != nil || resp == nil {
		return "", err
	}
	return resp.Error.Data, nil
}

// similar to the above but this does eth_call with extra gas and gasPrice
// see https://github.com/safe-global/safe-contracts/blob/186a21a74b327f17fc41217a927dea7064f74604/contracts/GnosisSafe.sol#L315
//
//nolint:lll
func (g *Estimation) rawEthCall(transaction ethTransactionWithGas, chainID int64) (string, error) {
	rpcURL, err := g.getURL(chainID)
	if err != nil {
		return "", err
	}

	resp := &estimateGasResponse{}
	_, err = commonClient.SetBaseURL(rpcURL).R().SetBody(
		&estimateGasRequest{
			Jsonrpc: "2.0",
			Method:  "eth_call",
			Params:  []interface{}{transaction, "latest"},
			ID:      1,
		},
	).SetResult(resp).Post("/")
	if err != nil || resp == nil {
		return "", err
	}
	return resp.Error.Data, nil
}

// estimateGasViaEthClient is the fallback if the required requiredTxGas fails to be called
// this just computes the gas using ethClient.client
func (g *Estimation) estimateGasViaEthClient(
	ctx context.Context,
	from common.Address,
	to common.Address,
	value math.Decimal256,
	data *hexutil.Bytes,
	chainID int64,
) (uint64, error) {
	if data == nil {
		return 0, errors.New("invalid data")
	}

	ethClient, err := g.gasEstimatorFactory.RetryableGasEstimator(chainID)
	if err != nil {
		return 0, err
	}

	return ethClient.EstimateGas(
		ctx, ethereum.CallMsg{
			From:  from,
			To:    &to,
			Value: (*big.Int)(&value),
			Data:  *data,
		},
	)
}

func (g *Estimation) getURL(chainID int64) (string, error) {
	rpcURL, ok := g.clientURLs[chainID]
	if !ok {
		return "", fmt.Errorf("%w, %d", ErrInvalidChainID, chainID)
	}

	return rpcURL, nil
}
