package decoders

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"go-safe/types"
	"go-safe/utils"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	ADDRESS_SIZE = 20
	UINT256_SIZE = 32
	UINT8_SIZE   = 1
)

func ParseExecTransactionCallData(calldata []byte) (execCallData []byte, err error) {
	// encode to hex string
	encodedCallData := hexutil.Encode(calldata)
	// selector index searches for execTransaction selector inside the calldata
	selectorIndex := strings.Index(encodedCallData, "6a761202")
	if selectorIndex == -1 {
		return nil, errors.New("no execSignature selector found")
	}
	// if the selectorIndex is less than the length of the encoded uint256
	// then the callData is encoded directly
	if selectorIndex < UINT256_SIZE {
		// return the actual call data without the selector
		return hexutil.Decode(fmt.Sprintf("0x%s", encodedCallData[selectorIndex+8:]))
	}
	var encodedCallDataLength []byte
	encodedCallDataLength, err = hexutil.Decode(
		fmt.Sprintf(
			"0x%s", encodedCallData[selectorIndex-32:selectorIndex],
		),
	)
	if err != nil {
		return nil, err
	}
	callDataLength := new(big.Int).SetBytes(encodedCallDataLength)
	if callDataLength.Uint64() == 0 {
		return nil, errors.New("parsed callDataLength is zero")
	}
	// length is encoded in hex, hence double the length to slice from the hex-encoded string
	encodedCallDataEndIndex := selectorIndex + int(callDataLength.Int64()*2)
	if len(encodedCallData) < encodedCallDataEndIndex {
		return nil, errors.New("parsed callDataLength is invalid")
	}

	// get the actual call data without the selector
	return hexutil.Decode(
		fmt.Sprintf(
			"0x%s", encodedCallData[selectorIndex+8:encodedCallDataEndIndex],
		),
	)
}

func ParseMultiSendData(data []byte) ([]types.InternalTxn, error) {
	if functionSig, err := utils.Slice(data, 0, 4); err != nil {
		return nil, err
	} else {
		if common.Bytes2Hex(functionSig) != "8d80ff0a" {
			return nil, nil
		}
	}
	var encodedLen *big.Int
	if length, err := utils.Slice(data, 36, 36+UINT256_SIZE); err != nil {
		return nil, err
	} else {
		encodedLen = new(big.Int).SetBytes(length)
	}
	// no encoded data found so return with nothing
	if encodedLen.Int64() == 0 {
		return nil, nil
	}
	var multiSendPacked []byte
	var internalTransactions []types.InternalTxn
	var err error
	if multiSendPacked, _, err = utils.ReadNAndShift(data, 36+UINT256_SIZE, int(encodedLen.Int64())); err != nil {
		return nil, err
	}
	baseOffset := 0
	for {
		if baseOffset >= len(multiSendPacked) {
			break
		}
		var internalTxn types.InternalTxn
		internalTxn, baseOffset, err = parseInternalTransaction(multiSendPacked, baseOffset)
		if err != nil {
			return nil, err
		}
		internalTransactions = append(internalTransactions, internalTxn)
	}
	return internalTransactions, err
}

func parseInternalTransaction(multiSendPacked []byte, baseOffset int) (types.InternalTxn, int, error) {
	var operation []byte
	var err error
	internalTxn := types.InternalTxn{}
	operation, baseOffset, err = utils.ReadNAndShift(multiSendPacked, baseOffset, UINT8_SIZE)
	if err != nil {
		return internalTxn, 0, err
	}
	internalTxn.Operation = operation[0]
	var toAddress []byte
	toAddress, baseOffset, err = utils.ReadNAndShift(multiSendPacked, baseOffset, ADDRESS_SIZE)
	if err != nil {
		return internalTxn, 0, err
	}
	internalTxn.To = common.BytesToAddress(toAddress)
	var value []byte
	value, baseOffset, err = utils.ReadNAndShift(multiSendPacked, baseOffset, UINT256_SIZE)
	if err != nil {
		return internalTxn, 0, err
	}
	internalTxn.Value = new(big.Int).SetBytes(value)
	var datalen []byte
	datalen, baseOffset, err = utils.ReadNAndShift(multiSendPacked, baseOffset, UINT256_SIZE)
	if err != nil {
		return internalTxn, 0, err
	}
	dataLen := new(big.Int).SetBytes(datalen).Int64()
	if dataLen == 0 {
		return internalTxn, baseOffset, nil
	}
	internalTxn.Data, baseOffset, err = utils.ReadNAndShift(multiSendPacked, baseOffset, int(dataLen))
	if err != nil {
		return internalTxn, 0, err
	}
	return internalTxn, baseOffset, nil
}
