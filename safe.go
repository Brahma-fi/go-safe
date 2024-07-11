package safe

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/Brahma-fi/console-transaction-builder/contracts/safe"
	builder "github.com/Brahma-fi/console-transaction-builder/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"

	"github.com/Brahma-fi/go-safe/contracts/safe_vOneDotThreeBinding"
)

func GetSignedSafeTxn(safeTxn *core.GnosisSafeTx, signatures [][]byte) error {
	//create new packer for solidityPack call
	packer := NewPackBuilder()
	for _, signature := range signatures {
		//add all individual signatures
		err := packer.AddBytes(signature)
		if err != nil {
			return err
		}
	}
	//pack it !!!
	packedSignature, err := packer.Pack()
	if err != nil {
		return err
	}
	safeTxn.Signature = packedSignature
	return nil
}

func GetSafeSignatureDomainHash(safeTxn *core.GnosisSafeTx) ([]byte, error) {
	typedData := safeTxn.ToTypedData()
	domainHash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return nil, err
	}
	return domainHash, nil
}

func GetEncodedExecTransaction(safeTxn *core.GnosisSafeTx, abi *abi.ABI) ([]byte, error) {
	/*
		function execTransaction(
			address to,
			uint256 value,
			bytes calldata data,
			Enum.Operation operation,
			uint256 safeTxGas,
			uint256 baseGas,
			uint256 gasPrice,
			address gasToken,
			address payable refundReceiver,
			bytes memory signatures
		) public payable virtual returns (bool success)
	*/
	return abi.Pack(
		"execTransaction", safeTxn.To.Address(), (*big.Int)(&safeTxn.Value),
		([]byte)(*safeTxn.Data), safeTxn.Operation, &safeTxn.SafeTxGas, &safeTxn.BaseGas,
		(*big.Int)(&safeTxn.GasPrice), safeTxn.GasToken, safeTxn.RefundReceiver, ([]byte)(safeTxn.Signature),
	)
}

// GetApprovedHashSafeTxn see https://github.com/safe-global/safe-contracts/blob/main/contracts/Safe.sol#L317
func GetApprovedHashSafeTxn(safeTxn *core.GnosisSafeTx, owner common.Address) error {
	packer := NewPackBuilder()
	// The signature format is a compact form of:
	//   {bytes32 r}{bytes32 s}{uint8 v}
	// Compact means, uint8 is not padded to 32 bytes.
	// r bytes32
	err := packer.AddBytes(AddressToBytes32(owner))
	if err != nil {
		return err
	}
	// s bytes32
	err = packer.AddBytes(make([]byte, 32))
	if err != nil {
		return err
	}
	// v uint8
	err = packer.AddUint8(1)
	if err != nil {
		return err
	}
	//pack it !!!
	packedSignature, err := packer.Pack()
	if err != nil {
		return err
	}
	safeTxn.Signature = packedSignature
	return nil
}

func PackTransactions(request *SafeMultiSendRequest) ([]byte, *big.Int, error) {
	packer := NewPackBuilder()
	totalValue := new(big.Int).SetInt64(0)
	for _, txn := range request.Transactions {
		callData := common.Hex2Bytes(txn.CallData())
		rawTransaction, err := PackTxn(txn.Operation(), txn.To(), txn.Value(), callData)
		if err != nil {
			return nil, nil, err
		}
		err = packer.AddBytes(rawTransaction)
		if err != nil {
			return nil, nil, err
		}
		totalValue.And(totalValue, txn.Value())
	}
	packed, err := packer.Pack()
	if err != nil {
		return nil, nil, err
	}
	return packed, totalValue, nil
}

func PackTxn(operation uint8, toAddress common.Address, value *big.Int, callData []byte) ([]byte, error) {
	packer := NewPackBuilder()
	err := packer.AddUint8(operation)
	if err != nil {
		return nil, err
	}
	err = packer.AddAddress(toAddress)
	if err != nil {
		return nil, err
	}
	err = packer.AddUint256(value)
	if err != nil {
		return nil, err
	}
	err = packer.AddUint256(new(big.Int).SetInt64(int64(len(callData))))
	if err != nil {
		return nil, err
	}
	err = packer.AddBytes(callData)
	if err != nil {
		return nil, err
	}
	return packer.Pack()
}

func GetEncodedMultiSendTransaction(callData []byte, abi *abi.ABI) ([]byte, error) {
	/*
		function multiSend(
			bytes calldata data,
		)
	*/
	return abi.Pack("multiSend", callData)
}

func AddressToBytes32(address common.Address) []byte {
	return append(make([]byte, 12), address.Bytes()...)
}

func GetModuleTransaction(callData []byte, to common.Address, value *big.Int, operation uint8) ([]byte, error) {
	/*
	   function execTransactionFromModule(
	       address to,
	       uint256 value,
	       bytes memory data,
	       Enum.Operation operation
	   )
	*/

	parsedAbi, err := abi.JSON(strings.NewReader(safe_vOneDotThreeBinding.SafeVOneDotThreeBindingMetaData.ABI))
	if err != nil {
		fmt.Println("failed to parse gearboxAbi", err)
		return nil, err
	}
	safeAbi := &parsedAbi
	return safeAbi.Pack("execTransactionFromModule", to, value, callData, operation)
}

func GetSafeNonce(ctx context.Context, client bind.ContractCaller, safeAddress common.Address) (*big.Int, error) {
	userSafe, err := safe.NewSafeCaller(safeAddress, client)
	if err != nil {
		return nil, err
	}
	return userSafe.Nonce(&bind.CallOpts{
		Context: ctx,
	})
}

func IsValidOwner(
	ctx context.Context,
	client bind.ContractCaller,
	safeAddress common.Address,
	owner common.Address,
) (bool, error) {
	userSafe, err := safe.NewSafeCaller(safeAddress, client)
	if err != nil {
		return false, err
	}
	return userSafe.IsOwner(&bind.CallOpts{Context: ctx}, owner)
}

func GetThreshold(
	ctx context.Context,
	client bind.ContractCaller,
	safeAddress common.Address,
) (*big.Int, error) {
	userSafe, err := safe.NewSafeCaller(safeAddress, client)
	if err != nil {
		return nil, err
	}
	return userSafe.GetThreshold(&bind.CallOpts{
		Context: ctx,
	})
}

func GetTransactionHash(
	ctx context.Context,
	client bind.ContractCaller,
	safeAddress common.Address, txn *builder.SafeTx,
) (
	common.Hash, error,
) {
	userSafe, err := safe.NewSafeCaller(safeAddress, client)
	if err != nil {
		return common.HexToHash(""), err
	}

	// GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int)
	txnHash, err := userSafe.GetTransactionHash(
		&bind.CallOpts{Context: ctx}, txn.To.Address(), (*big.Int)(&txn.Value), ([]byte)(*txn.Data), txn.Operation, &txn.SafeTxGas, &txn.BaseGas,
		(*big.Int)(&txn.GasPrice), txn.GasToken, txn.RefundReceiver, &txn.Nonce,
	)
	if err != nil {
		return common.HexToHash(""), err
	}
	return txnHash, nil
}
