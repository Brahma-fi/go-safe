package metadata

import (
	"context"
	"encoding/hex"
	"errors"
	"math/big"
	"regexp"
	"strings"

	"github.com/lastdotnet/go-safe/contracts/multicall"
	"github.com/lastdotnet/go-safe/contracts/safe"
	"github.com/lastdotnet/go-safe/contracts/walletregistry"
	"github.com/lastdotnet/go-safe/types"
	
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
)

const (
	WalletRegistryAddress = "walletRegistryAddress"
)

var (
	ErrMalformedResponse = errors.New("malformed response received")
)

type SafeMetadataService struct {
	addressRegistry types.AddressRegistry
	multiCaller     types.MultiCaller

	safeAbi           *abi.ABI
	walletRegistryAbi *abi.ABI
	re                *regexp.Regexp
}

func NewSafeMetadataService(
	addrRegistry types.AddressRegistry,
	mc types.MultiCaller,
) (*SafeMetadataService, error) {
	walletRegistryAbi, err := abi.JSON(strings.NewReader(walletregistry.WalletregistryMetaData.ABI))
	if err != nil {
		return nil, err
	}

	safeAbi, err := abi.JSON(strings.NewReader(safe.SafeMetaData.ABI))
	if err != nil {
		return nil, err
	}

	return &SafeMetadataService{
		addressRegistry:   addrRegistry,
		multiCaller:       mc,
		safeAbi:           &safeAbi,
		walletRegistryAbi: &walletRegistryAbi,
		re:                regexp.MustCompile("[^\u0020-\u007E]|\\t|\\n|\u0000"),
	}, nil
}

func (s *SafeMetadataService) prepareSafeMulticall(
	safes []common.Address,
	walletRegistry common.Address,
) ([]multicall.Multicall3Call3, error) {
	var multiCalls []multicall.Multicall3Call3

	getOwnersCallData, err := s.safeAbi.Pack("getOwners")
	if err != nil {
		return nil, err
	}

	getThresholdCallData, err := s.safeAbi.Pack("getThreshold")
	if err != nil {
		return nil, err
	}

	getNonceCallData, err := s.safeAbi.Pack("nonce")
	if err != nil {
		return nil, err
	}

	getVersionCallData, err := s.safeAbi.Pack("VERSION")
	if err != nil {
		return nil, err
	}

	guardStorageSlot := new(big.Int).SetBytes(
		common.HexToHash("0x4a204f620c8c5ccdca3fd54d003badd85ba500436a431f0cbda4f558c93c34c8").Bytes(),
	)

	getGuardCallData, err := s.safeAbi.Pack("getStorageAt", guardStorageSlot, new(big.Int).SetInt64(1))
	if err != nil {
		return nil, err
	}

	for _, safe := range safes {
		isWalletRegisteredCallData, packErr := s.walletRegistryAbi.Pack("isWallet", safe)
		if packErr != nil {
			return nil, packErr
		}

		moderatedAccountOwner, packErr := s.walletRegistryAbi.Pack(
			"subAccountToWallet", safe,
		)
		if packErr != nil {
			return nil, packErr
		}

		multiCalls = append(
			multiCalls,
			multicall.Multicall3Call3{
				Target:       safe,
				AllowFailure: true,
				CallData:     getOwnersCallData,
			},
			multicall.Multicall3Call3{
				Target:       safe,
				AllowFailure: true,
				CallData:     getThresholdCallData,
			},
			multicall.Multicall3Call3{
				Target:       safe,
				AllowFailure: true,
				CallData:     getNonceCallData,
			},
			multicall.Multicall3Call3{
				Target:       walletRegistry,
				AllowFailure: true,
				CallData:     isWalletRegisteredCallData,
			},
			multicall.Multicall3Call3{
				Target:       walletRegistry,
				AllowFailure: true,
				CallData:     moderatedAccountOwner,
			},
			multicall.Multicall3Call3{
				Target:       safe,
				AllowFailure: true,
				CallData:     getGuardCallData,
			},
			multicall.Multicall3Call3{
				Target:       safe,
				AllowFailure: true,
				CallData:     getVersionCallData,
			},
		)
	}

	return multiCalls, nil
}

func (s *SafeMetadataService) processSafeMulticallResponse(
	multiCallResponse []multicall.Multicall3Result,
	safes []common.Address,
) ([]types.Metadata, error) {
	var response []types.Metadata
	multicallIterator := 0
	multiSendCallLength := 7

	for _, address := range safes {
		if !multiCallResponse[multicallIterator].Success ||
			!multiCallResponse[multicallIterator+1].Success ||
			!multiCallResponse[multicallIterator+2].Success {
			log.Error().Any("response", multiCallResponse).Msg("invalid multicall response ")
			return nil, ErrMalformedResponse
		}

		ownersEncoded := multiCallResponse[multicallIterator].ReturnData
		unpack, err := s.safeAbi.Unpack("getOwners", ownersEncoded)
		if err != nil {
			return nil, err
		}

		if len(unpack) == 0 {
			log.Error().Int("length", len(unpack)).Msg("invalid unpacked response length")
			return nil, ErrMalformedResponse
		}

		owners, ok := unpack[0].([]common.Address)
		if !ok {
			log.Error().Int("length", len(unpack)).Msg("failed to unmarshal")
			return nil, ErrMalformedResponse
		}

		isConsole := false
		if isWallet := new(big.Int).SetBytes(multiCallResponse[multicallIterator+3].ReturnData).Uint64(); isWallet == 1 {
			isConsole = true
		}

		moderatedAccountOwner := common.BytesToAddress(multiCallResponse[multicallIterator+4].ReturnData)

		guardAddress := common.HexToAddress(common.BytesToHash(multiCallResponse[multicallIterator+5].ReturnData).Hex())
		versionBytes, _ := hex.DecodeString(common.Bytes2Hex(multiCallResponse[multicallIterator+6].ReturnData))
		version := s.re.ReplaceAllLiteralString(string(versionBytes), "")

		response = append(
			response, types.Metadata{
				Version:               strings.Trim(version, " "),
				SafeAddress:           address,
				Owners:                owners,
				Threshold:             new(big.Int).SetBytes(multiCallResponse[multicallIterator+1].ReturnData).Uint64(),
				Nonce:                 new(big.Int).SetBytes(multiCallResponse[multicallIterator+2].ReturnData).Uint64(),
				IsConsole:             isConsole,
				ModeratedAccountOwner: moderatedAccountOwner,
				GuardAddress:          guardAddress,
			},
		)
		multicallIterator += multiSendCallLength
	}
	return response, nil
}

func (s *SafeMetadataService) GetSafeMetadata(
	ctx context.Context,
	safes []common.Address,
	chainID int64,
) ([]types.Metadata, error) {
	walletRegistry, err := s.addressRegistry.GetAddressByChainID(chainID, WalletRegistryAddress)
	if err != nil {
		return nil, err
	}

	multiCalls, err := s.prepareSafeMulticall(safes, walletRegistry)
	if err != nil {
		return nil, err
	}

	multiCallResponse, err := s.multiCaller.Aggregate3(ctx, multiCalls, nil, chainID)
	if err != nil {
		return nil, err
	}

	if len(multiCallResponse) != len(safes)*7 {
		return nil, ErrMalformedResponse
	}

	return s.processSafeMulticallResponse(multiCallResponse, safes)
}
