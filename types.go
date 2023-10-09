package safe

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	Operation() uint8
}

type InternalTxn struct {
	Operation uint8          `json:"operation"`
	To        common.Address `json:"to"`
	Value     *big.Int       `json:"value"`
	Data      hexutil.Bytes  `json:"data"`
}

type SafeMultiSigInput struct {
	InternalTxn    `mapsstructure:",squash"`
	Operation      uint8          `json:"operation"`
	SafeTxGas      *big.Int       `json:"safeTxGas"`
	BaseGas        *big.Int       `json:"baseGas"`
	GasPrice       *big.Int       `json:"gasPrice"`
	GasToken       common.Address `json:"gasToken"`
	RefundReceiver common.Address `json:"refundReceiver"`
	Signatures     hexutil.Bytes  `json:"signatures"`
}

type SafeMultiSigEvent struct {
	InternalTxn
	SafeMultiSigInput
	AdditionalInfo hexutil.Bytes `json:"additionalInfo"`
}
