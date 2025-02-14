package types

import (
	"context"
	"math/big"

	"github.com/Brahma-fi/go-safe/contracts/multicall"
	"github.com/ethereum/go-ethereum/common"
)

type AddressRegistry interface {
	GetAddressByChainID(chainID int64, key string,) (common.Address, error)
}

type MultiCaller interface {
	Aggregate3(
		ctx context.Context,
		calls []multicall.Multicall3Call3,
		block *big.Int,
		chainID int64,
	) ([]multicall.Multicall3Result, error)
}
