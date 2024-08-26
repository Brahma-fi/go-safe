package encoders

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type TightPack struct {
	arguments [][]byte
}

func (pack *TightPack) AddAddress(address common.Address) error {
	addr := common.NewMixedcaseAddress(address)
	pack.arguments = append(pack.arguments, addr.Address().Bytes())
	return nil
}

func (pack *TightPack) AddBytes(data []byte) error {
	pack.arguments = append(pack.arguments, data)
	return nil
}

func (pack *TightPack) AddUint256(data *big.Int) error {
	sigUint256Type, _ := abi.NewType("uint256", "", nil)
	inputArgs := abi.Arguments{
		{
			Type: sigUint256Type,
		},
	}
	packedUints, err := inputArgs.Pack(data)
	if err != nil {
		return err
	}
	pack.arguments = append(pack.arguments, packedUints)
	return nil
}

func (pack *TightPack) AddUint8(data uint8) error {
	pack.arguments = append(pack.arguments, []byte{data})
	return nil
}

func (pack *TightPack) AddUint32(data uint32) error {
	encodedUint32 := make([]byte, 4)
	binary.BigEndian.PutUint32(encodedUint32, data)
	pack.arguments = append(pack.arguments, encodedUint32)
	return nil
}

func (pack *TightPack) Pack() ([]byte, error) {
	var encoded []byte
	for _, v := range pack.arguments {
		encoded = append(encoded, v...)
	}
	return encoded, nil
}

func NewPackBuilder() *TightPack {
	return &TightPack{arguments: [][]byte{}}
}
