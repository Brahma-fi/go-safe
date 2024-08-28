package encoders

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
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

func Test(t *testing.T) {
	builder := NewPackBuilder()
	val := new(big.Int).SetInt64(0)
	data, _ := hexutil.Decode("0xa9059cbb000000000000000000000000c20e3a951e03edd9e17c027324a07d2841cdf983000000000000000000000000000000000000000000000000016345785d8a0000")
	assert.NoError(t, builder.AddAddress(common.HexToAddress("0x02ABBDbAaa7b1BB64B5c878f7ac17f8DDa169532")))
	assert.NoError(t, builder.AddUint256(val))
	assert.NoError(t, builder.AddBytes(data))
	assert.NoError(t, builder.AddUint8(0))
	assert.NoError(t, builder.AddAddress(common.HexToAddress("0x4fadb18339be16d57f27a8a65d78ffecf29037d0")))
	assert.NoError(t, builder.AddUint256(val))
	assert.NoError(t, builder.AddBytes(common.HexToHash("0x954c8850b002325bd02c8cffb07b6226ff77b3bd16f865827a63ca510ff1559a").Bytes()))
	assert.NoError(t, builder.AddUint32(1694514609))
	p, _ := builder.Pack()
	assert.Equal(
		t,
		"0x02abbdbaaa7b1bb64b5c878f7ac17f8dda1695320000000000000000000000000000000000000000000000000000000000000000a9059cbb000000000000000000000000c20e3a951e03edd9e17c027324a07d2841cdf983000000000000000000000000000000000000000000000000016345785d8a0000004fadb18339be16d57f27a8a65d78ffecf29037d00000000000000000000000000000000000000000000000000000000000000000954c8850b002325bd02c8cffb07b6226ff77b3bd16f865827a63ca510ff1559a65003db1",
		hexutil.Encode(p), "packed value mismatch",
	)
}
