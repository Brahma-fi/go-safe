package utils

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// GetTypedDataHash returns the hash of the fully encoded eip-712 %%value%% for %%types%% with %%domain%%.
// https://github.com/ethers-io/ethers.js/blob/c5cb7cd71d9a12b8feeec4fd956d0a416b0be32f/src.ts/hash/typed-data.ts#L491
func GetTypedDataHash(typedData apitypes.TypedData) (common.Hash, error) {
	// EIP-712 typed data marshalling
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return common.Hash{}, fmt.Errorf("eip712domain hash struct: %w", err)
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return common.Hash{}, fmt.Errorf("primary type hash struct: %w", err)
	}

	// add magic string prefix
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	return common.BytesToHash(crypto.Keccak256(rawData)), nil
}

// GetMessageHash gets EIP-191 personal-sign message digest to sign
// https://github.com/ethers-io/ethers.js/blob/c5cb7cd71d9a12b8feeec4fd956d0a416b0be32f/src.ts/hash/message.ts#L35-L36
func GetMessageHash(message string) common.Hash {
	return common.BytesToHash(
		crypto.Keccak256(
			[]byte(fmt.Sprintf(
				"\x19Ethereum Signed Message:\n%d%s", len(message), message,
			)),
		),
	)
}
