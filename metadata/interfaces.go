package metadata

import (
	"context"
	"math/big"

	"github.com/Brahma-fi/go-safe/contracts/multicall"
	"github.com/ethereum/go-ethereum/common"
)

type addressProvider interface {
	GetAddress(key string) (common.Address, error)
}

type addressRegistry interface {
	AddressProvider(chainID int64) (addressProvider, error)
}

type multiCaller interface {
	Aggregate3(
		ctx context.Context,
		calls []multicall.Multicall3Call3,
		block *big.Int,
		chainID int64,
	) ([]multicall.Multicall3Result, error)
}
