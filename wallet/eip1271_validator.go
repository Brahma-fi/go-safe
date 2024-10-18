package wallet

import (
	"context"
	"fmt"

	eip1271validator "github.com/Brahma-fi/go-safe/contracts/validators"
	"github.com/Brahma-fi/go-safe/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

const (
	//MagicValueHex bytes4(keccak256("isValidSignature(bytes32,bytes)")
	MagicValueHex = "0x1626ba7e"
)

func isValidSignature(
	ctx context.Context,
	caller bind.ContractCaller,
	signer common.Address,
	hash common.Hash,
	signature []byte,
) error {
	validator, err := eip1271validator.NewEip1271validatorCaller(signer, caller)
	if err != nil {
		return err
	}

	magicValue, err := validator.IsValidSignature(&bind.CallOpts{Context: ctx}, hash, signature)
	switch {
	case err != nil:
		return err
	case hexutil.Encode(magicValue[:]) != MagicValueHex:
		return fmt.Errorf("failed to verify magicValue want=%s found=%s", MagicValueHex, hexutil.Encode(magicValue[:]))
	}

	return nil
}

func VerifyTypedDataSignatureWithEIP1271(
	ctx context.Context,
	caller bind.ContractCaller,
	signer common.Address,
	request apitypes.TypedData,
	signature []byte,
) error {
	hash, err := utils.GetTypedDataHash(request)
	if err != nil {
		return err
	}

	return isValidSignature(ctx, caller, signer, hash, signature)
}

func VerifyPersonalSignatureWithEIP1271(
	ctx context.Context,
	caller bind.ContractCaller,
	signer common.Address,
	message string,
	signature []byte,
) error {
	return isValidSignature(ctx, caller, signer, utils.GetMessageHash(message), signature)
}

func GetSafeMessageDigest(
	message common.Hash,
	chainID int64,
	safeAddress common.Address,
) (*common.Hash, error) {
	safeTransactionTypedData := &apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"SafeMessage": []apitypes.Type{
				{Name: "message", Type: "bytes"},
			},
		},
		PrimaryType: "SafeMessage",
		Domain: apitypes.TypedDataDomain{
			ChainId:           math.NewHexOrDecimal256(chainID),
			VerifyingContract: safeAddress.Hex(),
		},
		Message: map[string]interface{}{
			"message": message.Hex(),
		},
	}

	domainSeparator, err := safeTransactionTypedData.HashStruct("EIP712Domain", safeTransactionTypedData.Domain.Map())
	if err != nil {
		return nil, fmt.Errorf("eip712domain hash struct: %w", err)
	}

	typedDataHash, err := safeTransactionTypedData.HashStruct(
		safeTransactionTypedData.PrimaryType,
		safeTransactionTypedData.Message,
	)
	if err != nil {
		return nil, fmt.Errorf("primary type hash struct: %w", err)
	}

	sigHash := crypto.Keccak256Hash([]byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash))))
	return &sigHash, nil
}
