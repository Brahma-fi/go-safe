package wallet

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/rs/zerolog/log"
)

func TestGetPublicKeyAndCheckWalletAddress(t *testing.T) {
	var hashedText = accounts.TextHash([]byte("msg1"))
	var sig, _ = getSignature("0x835529e4c546e3cd906052cbe2fda65a012d104d0472fd2f2a0dc75f22a5e2737fc70c806344d48948b69339854041ab469d8825ee91491948dcf58204d613ff01")

	hashedText = append(hashedText, []byte("A")...)
	var _, err = getPublicKeyAndCheckWalletAddress(hashedText, sig, "any")
	if err == nil {
		log.Error().Err(err).Msg("failed to get public key")
		t.Error("TestGetPublicKeyAndCheckWalletAddress() error == nil and it should fail")
	}
}

func TestCheckSignature(t *testing.T) {
	var hashedText = accounts.TextHash([]byte("msg1"))
	var sig, _ = getSignature("0x835529e4c546e3cd906052cbe2fda65a012d104d0472fd2f2a0dc75f22a5e2737fc70c806344d48948b69339854041ab469d8825ee91491948dcf58204d613ff01")
	var sigPublicKey, _ = getPublicKeyAndCheckWalletAddress(hashedText, sig, "0x16A798262b8f77AF783a3681BB497d4CD73EBb0A")

	hashedText = append(hashedText, []byte("A")...)
	var err = checkSignature(sigPublicKey, hashedText, sig)
	if err == nil {
		log.Error().Err(err).Msg("failed to get public key")
		t.Error("TestcheckSignature() error == nil and it should fail")
	}
}
