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
	Data           *hexutil.Bytes          `json:"data"`
	ChainId        *math.HexOrDecimal256   `json:"chainId,omitempty"`
	To             common.MixedcaseAddress `json:"to"`
	Safe           common.MixedcaseAddress `json:"safe"`
	SafeTxGas      big.Int                 `json:"safeTxGas"`
	BaseGas        big.Int                 `json:"baseGas"`
	GasPrice       math.Decimal256         `json:"gasPrice"`
	Nonce          big.Int                 `json:"nonce"`
	Value          math.Decimal256         `json:"value"`
	InputExpHash   common.Hash             `json:"safeTxHash,omitempty"`
	GasToken       common.Address          `json:"gasToken"`
	RefundReceiver common.Address          `json:"refundReceiver"`
	Operation      uint8                   `json:"operation"`
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
	Value     *big.Int       `json:"value"`
	Data      hexutil.Bytes  `json:"data"`
	To        common.Address `json:"to"`
	Operation uint8          `json:"operation"`
}

type SafeMultiSigInput struct {
	SafeTxGas      *big.Int      `json:"safeTxGas"`
	BaseGas        *big.Int      `json:"baseGas"`
	GasPrice       *big.Int      `json:"gasPrice"`
	Signatures     hexutil.Bytes `json:"signatures"`
	InternalTxn    `mapstructure:",squash"`
	GasToken       common.Address `json:"gasToken"`
	RefundReceiver common.Address `json:"refundReceiver"`
}

type SafeMultiSigEvent struct {
	AdditionalInfo hexutil.Bytes `json:"additionalInfo"`
	InternalTxn
	SafeMultiSigInput
}

type Metadata struct {
	Version               string           `json:"version"`
	Owners                []common.Address `json:"owners"`
	Threshold             uint64           `json:"threshold"`
	Nonce                 uint64           `json:"nonce"`
	SafeAddress           common.Address   `json:"safeAddress"`
	ModeratedAccountOwner common.Address   `json:"ModeratedAccountOwner"`
	GuardAddress          common.Address   `json:"-"`
	IsConsole             bool             `json:"isConsole"`
}
