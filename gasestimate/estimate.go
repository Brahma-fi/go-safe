package gasestimate

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

const (
	DefaultMaxGasLimit   = uint64(6400000)
	MaxBinarySearchDepth = 10000
)

var (
	ErrEstimateTransactionLimitFailed = errors.New("failed to estimate safe transaction gas")
)

type EstimateSafeTransactionGasRequest struct {
	From        common.Address
	To          common.Address
	CallData    []byte
	Value       *big.Int
	MaxGasLimit uint64
}

type EstimateSafeTransactionClient interface {
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}

// EstimateSafeTransactionGasLimit estimates the gas limit for a safe transaction.
// It uses a binary search algorithm to find the optimal gas limit for the transaction.
// And finds the optimal gas at which the transaction succeeds by checking the return of the eth_call
// if encoded callData is execTransactionFromModule or execTransaction and call is successful safe returns (true) boolean as return
// https://github.com/safe-global/safe-contracts/blob/1cfa95710057e33832600e6b9ad5ececca8f7839/contracts/SafeL2.sol#L12-L13
func EstimateSafeTransactionGasLimit(
	ctx context.Context, client EstimateSafeTransactionClient,
	request EstimateSafeTransactionGasRequest,
) (
	uint64, error,
) {
	if request.MaxGasLimit <= 0 {
		request.MaxGasLimit = DefaultMaxGasLimit
	}

	var hi = request.MaxGasLimit
	ethereumMsg := ethereum.CallMsg{
		From:  request.From,
		To:    &request.To,
		Value: request.Value,
		Data:  request.CallData,
	}
	// estimateGas gives us the baseGas at which the transaction can be executed
	lo, err := client.EstimateGas(ctx, ethereumMsg)
	if err != nil {
		return 0, err
	}

	// Create a helper to check if a gas allowance results in a successful transaction
	executable := func(gas uint64) (bool, error) {
		ethereumMsg.Gas = gas
		// call the safe contract at latest block
		result, err := client.CallContract(ctx, ethereumMsg, nil)
		if err != nil {
			return true, err
		}
		outBool := new(big.Int).SetBytes(result)
		return outBool.Int64() != 1, nil
	}

	// check if the execution results in true from the client estimateGas
	baseFailed, err := executable(lo)
	if err != nil {
		return 0, err
	}

	if !baseFailed {
		return ethereumMsg.Gas, nil
	}

	// base limit does not guarantee a successful execution of transaction
	// perform a binary search to find gasLimit at which
	// safe exec outputs true
	for lo+1 < hi {
		mid := (hi + lo) / 2
		// if the difference between the new calculated mid and high
		// is less than the configured depth, then the hi is used as the final value
		// this prevents excessive rpc call and gives estimated gas in
		// the range of  (actual limit - MAX_BINARY_SEARCH_DEPTH),(actual limit + MAX_BINARY_SEARCH_DEPTH)
		if hi-mid < MaxBinarySearchDepth {
			ethereumMsg.Gas = hi
			break
		}

		failed, err := executable(mid)
		// not a valid transaction, as og implementation state "don't struggle anymore"
		if err != nil {
			return 0, err
		}

		if failed {
			lo = mid
		} else {
			hi = mid
		}
	}

	if ethereumMsg.Gas == request.MaxGasLimit {
		return 0, ErrEstimateTransactionLimitFailed
	}

	return ethereumMsg.Gas, nil
}
