package safe

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SafeMultiSendRequest struct {
	Transactions []*SendTxnRequest `json:"transactions"`
	ChainId      int64             `json:"chain_id"`
	From         common.Address    `json:"from"`
}
type SafeTxnRequest struct {
	Operation      uint8
	SafeAddress    common.Address
	To             common.Address
	Value          *big.Int
	Data           []byte
	Nonce          *big.Int
	ChainId        *big.Int
	GasToken       common.Address
	RefundReceiver common.Address
}

type SendTxnRequest struct {
	ChainId  int64          `json:"chainId"`
	Calldata string         `json:"calldata"`
	To       common.Address `json:"to"`
	Value    string         `json:"value"`
}
