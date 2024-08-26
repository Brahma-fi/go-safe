package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
)

// SafeTx is derived from the core.GnosisSafeTx type and does not include the output
// fields which makes easy for the marshalling and unmarshalling
type SafeTx struct {
	// These fields are used both on input and output
	Safe           common.MixedcaseAddress `json:"safe"`
	To             common.MixedcaseAddress `json:"to"`
	Value          math.Decimal256         `json:"value"`
	GasPrice       math.Decimal256         `json:"gasPrice"`
	Data           *hexutil.Bytes          `json:"data"`
	Operation      uint8                   `json:"operation"`
	GasToken       common.Address          `json:"gasToken"`
	RefundReceiver common.Address          `json:"refundReceiver"`
	BaseGas        big.Int                 `json:"baseGas"`
	SafeTxGas      big.Int                 `json:"safeTxGas"`
	Nonce          big.Int                 `json:"nonce"`
	InputExpHash   common.Hash             `json:"safeTxHash,omitempty"`
	ChainId        *math.HexOrDecimal256   `json:"chainId,omitempty"`
}

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
	InternalTxn    `mapstructure:",squash"`
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

type Metadata struct {
	Version               string           `json:"version"`
	SafeAddress           common.Address   `json:"safeAddress"`
	Owners                []common.Address `json:"owners"`
	Threshold             uint64           `json:"threshold"`
	Nonce                 uint64           `json:"nonce"`
	IsConsole             bool             `json:"isConsole"`
	ModeratedAccountOwner common.Address   `json:"ModeratedAccountOwner"`
	GuardAddress          common.Address   `json:"-"`
}
