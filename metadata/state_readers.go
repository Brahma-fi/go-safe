package metadata

import (
	"context"
	"math/big"

	"github.com/Brahma-fi/go-safe/contracts/safe"
	"github.com/Brahma-fi/go-safe/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

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
	safeAddress common.Address, txn *types.SafeTx,
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
