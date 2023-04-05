package safe

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SafeMultiSendRequest struct {
	Transactions []Transaction  `json:"transactions"`
	ChainId      int64          `json:"chain_id"`
	From         common.Address `json:"from"`
}

type Transaction interface {
	From() common.Address
	To() common.Address
	CallData() string
	Value() *big.Int
}
