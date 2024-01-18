package metadata

import (
	"context"
	"math/big"

	"github.com/Brahma-fi/console-transaction-builder/addresses"
	"github.com/Brahma-fi/console-transaction-builder/contracts/multicall3"
)

type addressRegistry interface {
	AddressProvider(chainID int64) (addresses.AddressProvider, error)
}

type multiCaller interface {
	Aggregate3(
		ctx context.Context,
		calls []multicall3.Multicall3Call3,
		block *big.Int,
		chainID int64,
	) ([]multicall3.Multicall3Result, error)
}
