package gasestimate

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	SomeErr = errors.New("im underwater")
)

type mockEstimateBackend struct {
	successLimit uint64
	baseEstimate uint64
}

func (m *mockEstimateBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return m.baseEstimate, nil
}

func (m *mockEstimateBackend) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) (
	[]byte, error,
) {
	if msg.Gas >= m.successLimit {
		return new(big.Int).SetInt64(1).Bytes(), nil
	}
	return new(big.Int).SetInt64(0).Bytes(), nil
}

type failingBackend struct {
}

func (m *failingBackend) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 0, SomeErr
}

func (m *failingBackend) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) (
	[]byte, error,
) {
	return nil, SomeErr
}

func TestEstimateSafeTransactionGasLimit(t *testing.T) {
	ctx := context.Background()
	type testParams struct {
		backend       EstimateSafeTransactionClient
		err           error
		params        EstimateSafeTransactionGasRequest
		expectedLimit uint64
	}
	params := []testParams{
		{
			params: EstimateSafeTransactionGasRequest{
				From:        common.HexToAddress(""),
				To:          common.HexToAddress(""),
				CallData:    []byte{},
				Value:       new(big.Int).SetInt64(0),
				MaxGasLimit: 200000,
			},
			backend: &mockEstimateBackend{
				successLimit: 150000,
				baseEstimate: 120000,
			},
			expectedLimit: 150000,
			err:           nil,
		},
		{
			params: EstimateSafeTransactionGasRequest{
				From:        common.HexToAddress(""),
				To:          common.HexToAddress(""),
				CallData:    []byte{},
				Value:       new(big.Int).SetInt64(0),
				MaxGasLimit: 200000,
			},
			backend: &mockEstimateBackend{
				successLimit: 400000,
				baseEstimate: 120000,
			},
			expectedLimit: 0,
			err:           ErrEstimateTransactionLimitFailed,
		},
		{
			params: EstimateSafeTransactionGasRequest{
				From:        common.HexToAddress(""),
				To:          common.HexToAddress(""),
				CallData:    []byte{},
				Value:       new(big.Int).SetInt64(0),
				MaxGasLimit: 400000,
			},
			backend: &mockEstimateBackend{
				successLimit: 250000,
				baseEstimate: 10,
			},
			expectedLimit: 250003,
			err:           nil,
		},
		{
			params: EstimateSafeTransactionGasRequest{
				From:        common.HexToAddress(""),
				To:          common.HexToAddress(""),
				CallData:    []byte{},
				Value:       new(big.Int).SetInt64(0),
				MaxGasLimit: 400000,
			},
			backend:       &failingBackend{},
			expectedLimit: 0,
			err:           SomeErr,
		},
	}
	for _, p := range params {
		estimate, err := EstimateSafeTransactionGasLimit(ctx, p.backend, p.params)
		if p.err != nil {
			assert.Error(t, err, "missing error")
			assert.Equal(t, err, p.err, "expected error but received different error")
			continue
		}
		assert.NoError(t, err, "error not expected")
		assert.Equal(t, p.expectedLimit, estimate, "estimated values are different")
	}
}
