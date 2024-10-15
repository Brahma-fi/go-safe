package wallet

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestVerifySignedSafeSignatureWithEIP1271(t *testing.T) {
	client, err := ethclient.Dial("https://base.llamarpc.com")
	assert.NoError(t, err, "failed to get client")
	message := "hello"
	safe := common.HexToAddress("0xd1f745f0d14918a2c1e31153f2492891e4526ea4")
	signature, err := hexutil.Decode("0x2f535137e7484489f293b25edc35e29bd5c59451dfdbd51b93d5060ae201869a054f380cf700b335e311392de3a674e5899b4c924e02b2d820dca196ecc200a601")
	assert.NoError(t, err, "failed to parse signature")
	if signature[crypto.RecoveryIDOffset] == 0 || signature[crypto.RecoveryIDOffset] == 1 {
		signature[crypto.RecoveryIDOffset] += 27 // Transform yellow paper V from 27/28 to 0/1

	}

	assert.NoError(t, VerifyPersonalSignatureWithEIP1271(
		context.Background(),
		client,
		safe,
		message,
		signature,
	), "failed to verify sig")
}
