package safe

import (
	"github.com/ethereum/go-ethereum/common"
)

type SafeMultiSendRequest struct {
	Transactions []*Transaction `json:"transactions"`
	ChainId      int64          `json:"chain_id"`
	From         common.Address `json:"from"`
}

type Transaction struct {
	Calldata string         `json:"calldata"`
	To       common.Address `json:"to"`
	Value    string         `json:"value"`
}
