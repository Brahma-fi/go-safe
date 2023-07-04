package safe

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMultiSendData(t *testing.T) {
	subData, _ := hexutil.Decode(
		"0x8D80FF0A000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000CE0079D664098E9A97C40695522B9568A67EC752A26600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024BE300447000000000000000000000000000000000000000000000000006A94D74F430000000405D9D1443DFB051D5E8E231E41C911DC8393A40000000000000000000000000000000000000000000000000018838370F340000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	)
	_, err := ParseMultiSendData(subData)
	if err != nil {
		t.Error(err)
	}
}

func FuzzParseMultiSendData(f *testing.F) {
	subData, _ := hexutil.Decode(
		"0x8D80FF0A000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000CE0079D664098E9A97C40695522B9568A67EC752A26600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024BE300447000000000000000000000000000000000000000000000000006A94D74F430000000405D9D1443DFB051D5E8E231E41C911DC8393A40000000000000000000000000000000000000000000000000018838370F340000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	)
	f.Add(subData)
	f.Fuzz(
		func(t *testing.T, data []byte) {
			out, err := ParseMultiSendData(data)
			if err != nil {
				t.Logf("%q %v", out, err)
			}
		},
	)
}

func TestParseExecTransactionWithGelatoRelay(t *testing.T) {
	gelatoRealyCallData, _ := hexutil.Decode("0x4522589f0000000000000000000000000000000000000000000000000000000000000020f49e86987339087684f483739f2b28c974cdd2f262bf169b0429154c3a924f0a000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000004e0000000000000000000000000000000000000000000000000000000000000056000000000000000000000000075ba5af8effdcfca32e1e288806d54277d1fde99000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000649c584100000000000000000000000000000000000000000000000000000000649c5a9900000000000000000000000000000000000000000000000000000000000003a427be3c2400000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000ae75b29ade678372d77a8b41225654138a7e6ff1000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001f49e86987339087684f483739f2b28c974cdd2f262bf169b0429154c3a924f0a0000000000000000000000000000000000000000000000000000000000000005000000000000000000000000ac98de41fae7c7394e97fbe6fe48fbb8107b63c8000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000002246a761202000000000000000000000000ac98de41fae7c7394e97fbe6fe48fbb8107b63c8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a8fe000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ae75b29ade678372d77a8b41225654138a7e6ff10000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000827dd81428c9f574157c65b73b3253edf0dfc03571f7e07fbb006b9ae69e9b8d8d63253ed267cbed5bdb15f1a6b8a37aee1f902de83b93958bd764f19ecf8783c71c35bec4c0542070c5629e4225531ff9ee32e55684fa4175846b46fc2deaa2528458fabd4bb4d648464b7485282b2814b361a2579c03bc94870c2a6bddb7526bd01c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000041f3a27759d7cfa1a521110e7b5a91a7faaccb9bf37dc29f90bb88c754b66f63c64ef83dfa01892d04ee3f3417c2441c6a0034e2a6e5fc08ed5db31ff937abb8dd1c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000411f2d347795997bf64b60dd2843ff579072dda6a0c30e0e1f65ddd314023b57b21a6633467ffd622d27d8bacb6cee55c6cf553d53e2550c7add3b88a2ac1776e71b00000000000000000000000000000000000000000000000000000000000000")
	execCallData, err := ExtractExecTransactionCallData(gelatoRealyCallData)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(
		t,
		"0x000000000000000000000000ac98de41fae7c7394e97fbe6fe48fbb8107b63c8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a8fe000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ae75b29ade678372d77a8b41225654138a7e6ff10000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000827dd81428c9f574157c65b73b3253edf0dfc03571f7e07fbb006b9ae69e9b8d8d63253ed267cbed5bdb15f1a6b8a37aee1f902de83b93958bd764f19ecf8783c71c35bec4c0542070c5629e4225531ff9ee32e55684fa4175846b46fc2deaa2528458fabd4bb4d648464b7485282b2814b361a2579c03bc94870c2a6bddb7526bd01c000000000000000000000000000000000000000000000000000000000000",
		hexutil.Encode(execCallData), "failed to extract correct calldata",
	)
}
func TestParseExecTransaction(t *testing.T) {
	callData, _ := hexutil.Decode("0x6a761202000000000000000000000000e9f47d5ee5b73326e1eb9361630105e8ca386874000000000000000000000000000000000000000000000000006a94d74f43000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c605000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ae75b29ade678372d77a8b41225654138a7e6ff100000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004174f62e97665e6a43e595cae51f880e712b2480a16bdd2e3087782361311807481784e65339c9f30a07a6b119ae395c75a8e32ae2ab93c68dd29521333352a0813600000000000000000000000000000000000000000000000000000000000000")
	execCallData, err := ExtractExecTransactionCallData(callData)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(
		t,
		"0x000000000000000000000000e9f47d5ee5b73326e1eb9361630105e8ca386874000000000000000000000000000000000000000000000000006a94d74f43000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c605000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ae75b29ade678372d77a8b41225654138a7e6ff100000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004174f62e97665e6a43e595cae51f880e712b2480a16bdd2e3087782361311807481784e65339c9f30a07a6b119ae395c75a8e32ae2ab93c68dd29521333352a0813600000000000000000000000000000000000000000000000000000000000000",
		hexutil.Encode(execCallData), "failed to extract correct calldata",
	)
}
