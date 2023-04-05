package safe

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestTightPack_Pack(t *testing.T) {
	builder := NewPackBuilder()
	err := builder.AddUint8(uint8(0))
	assert.NoError(t, err, "failed to call AddUint8")
	err = builder.AddAddress(common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"))
	assert.NoError(t, err, "failed to call AddAddress")
	err = builder.AddUint256(new(big.Int).SetInt64(0))
	assert.NoError(t, err, "failed to call AddUint256")
	randomCallData, _ := hexutil.Decode("0x095ea7b30000000000000000000000001111111254fb6c44bac0bed2854e76f90643097dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	err = builder.AddUint256(new(big.Int).SetInt64(int64(len(randomCallData))))
	assert.NoError(t, err, "failed to call AddUint256")
	err = builder.AddBytes(randomCallData)
	assert.NoError(t, err, "failed to call AddBytes")
	packed, err := builder.Pack()
	assert.NoError(t, err, "failed to call Pack")
	t.Log("packed", hexutil.Encode(packed))
	assert.Equal(
		t,
		"0x002791bca1f2de4661ed88a30c99a7a9449aa8417400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000044095ea7b30000000000000000000000001111111254fb6c44bac0bed2854e76f90643097dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
		hexutil.Encode(packed), "packed data does not match",
	)
}
