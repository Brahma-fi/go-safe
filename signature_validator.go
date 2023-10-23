package safe

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/subtle"
	"errors"

	"github.com/Brahma-fi/console-transaction-builder/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/goccy/go-json"
)

var (
	ErrInvalidWalletAddress       = errors.New("invalid wallet address")      // bad request error code 400
	ErrInvalidSignature           = errors.New("invalid signature")           // bad request error code 400
	ErrPubKeyRecovery             = errors.New("public key recovery failed")  // bad request error code 400
	ErrVerifySigner               = errors.New("failed to verify signer")     // bad request error code 400
	ErrFailedToGetRelayAuthSigner = errors.New("relay auth signer not found") // bad request error code 400
	ErrNilValue                   = errors.New("value is nil")                // bad request error code 400
	ErrFailedToVerifySignature    = errors.New("failed to verify signature")  // bad request error code 400

)

type keyManager interface {
	SignHash(message common.Hash) ([]byte, error)
	GetPublicKey() *ecdsa.PublicKey
}

func VerifyTypedDataSignature(eoa, sigHex string, request apitypes.TypedData) error {
	if !common.IsHexAddress(eoa) {
		return ErrInvalidWalletAddress
	}

	var sig, err = getSignature(sigHex)
	if err != nil {
		return err
	}

	sigHash, err := GetTypedDataHash(request)
	if err != nil {
		return err
	}

	sigPublicKey, err := getPublicKeyAndCheckWalletAddress(sigHash.Bytes(), sig, eoa)
	if err != nil {
		return err
	}

	return checkSignature(sigPublicKey, sigHash.Bytes(), sig)
}

func VerifyPersonalSignSignature(eoa, sigHex, message string) error {
	if !common.IsHexAddress(eoa) {
		return ErrInvalidWalletAddress
	}

	var sig, err = getSignature(sigHex)
	if err != nil {
		return err
	}
	sigHash := GetMessageHash(message)

	sigPublicKey, err := getPublicKeyAndCheckWalletAddress(sigHash.Bytes(), sig, eoa)
	if err != nil {
		return err
	}

	return checkSignature(sigPublicKey, sigHash.Bytes(), sig)
}

func checkSignature(sigPublicKey, hashedText, sig []byte) error {
	var signatureNoRecoverID = sig[:len(sig)-1] // remove recovery id
	if !crypto.VerifySignature(sigPublicKey, hashedText, signatureNoRecoverID) {
		return ErrInvalidSignature
	}

	return nil
}

func getPublicKeyAndCheckWalletAddress(hashedText, sig []byte, eoa string) ([]byte, error) {
	var sigPublicKey, err = crypto.Ecrecover(hashedText, sig)
	if err != nil {
		return nil, ErrPubKeyRecovery
	}

	pubKey, err := crypto.UnmarshalPubkey(sigPublicKey)
	if err != nil {
		return nil, err
	}

	var recoveredAddr = crypto.PubkeyToAddress(*pubKey)

	if subtle.ConstantTimeCompare(common.HexToAddress(eoa).Bytes(), recoveredAddr.Bytes()) == 0 {
		return nil, ErrVerifySigner
	}

	return sigPublicKey, nil
}

func getSignature(sigHex string) ([]byte, error) {
	var sig, err = hexutil.Decode(sigHex)
	if err != nil {
		return nil, ErrInvalidSignature
	}

	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	return sig, nil
}

func SignSafeTxnData(signer keyManager, tx *types.SafeTx) ([]byte, error) {
	if signer == nil {
		return nil, ErrFailedToGetRelayAuthSigner
	}

	if tx == nil {
		return nil, ErrNilValue
	}

	data, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}

	return signer.SignHash(crypto.Keccak256Hash(data))
}

func VerifySignedSafeTxnData(signer *ecdsa.PublicKey, sig string, tx *types.SafeTx) error {
	if tx == nil || sig == "" {
		return ErrNilValue
	}

	sigBytes, err := hexutil.Decode(sig)
	if err != nil {
		return err
	}

	data, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(crypto.Keccak256Hash(data).Bytes(), sigBytes)
	if err != nil {
		return err
	}

	if !bytes.Equal(crypto.FromECDSAPub(sigPublicKeyECDSA), crypto.FromECDSAPub(signer)) {
		return ErrFailedToVerifySignature
	}

	return nil
}
