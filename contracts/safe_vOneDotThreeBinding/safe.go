// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package safe_vOneDotThreeBinding

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SafeVOneDotThreeBindingMetaData contains all meta data concerning the SafeVOneDotThreeBinding contract.
var SafeVOneDotThreeBindingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"approvedHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ApproveHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"ChangedFallbackHandler\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"ChangedGuard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ChangedThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"DisabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"EnabledModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"ExecutionFromModuleSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"ExecutionSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"RemovedOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initializer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"}],\"name\":\"SafeSetup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"SignMsg\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"addOwnerWithThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToApprove\",\"type\":\"bytes32\"}],\"name\":\"approveHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"approvedHashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"changeThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"requiredSignatures\",\"type\":\"uint256\"}],\"name\":\"checkNSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"checkSignatures\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"disableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"enableModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"encodeTransactionData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"execTransactionFromModuleReturnData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"start\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getModulesPaginated\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"array\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"next\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isModuleEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"}],\"name\":\"requiredTxGas\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"}],\"name\":\"setFallbackHandler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"setGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"fallbackHandler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentReceiver\",\"type\":\"address\"}],\"name\":\"setup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"signedMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"swapOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SafeVOneDotThreeBindingABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeVOneDotThreeBindingMetaData.ABI instead.
var SafeVOneDotThreeBindingABI = SafeVOneDotThreeBindingMetaData.ABI

// SafeVOneDotThreeBinding is an auto generated Go binding around an Ethereum contract.
type SafeVOneDotThreeBinding struct {
	SafeVOneDotThreeBindingCaller     // Read-only binding to the contract
	SafeVOneDotThreeBindingTransactor // Write-only binding to the contract
	SafeVOneDotThreeBindingFilterer   // Log filterer for contract events
}

// SafeVOneDotThreeBindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeVOneDotThreeBindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeVOneDotThreeBindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeVOneDotThreeBindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeVOneDotThreeBindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeVOneDotThreeBindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeVOneDotThreeBindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeVOneDotThreeBindingSession struct {
	Contract     *SafeVOneDotThreeBinding // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SafeVOneDotThreeBindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeVOneDotThreeBindingCallerSession struct {
	Contract *SafeVOneDotThreeBindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// SafeVOneDotThreeBindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeVOneDotThreeBindingTransactorSession struct {
	Contract     *SafeVOneDotThreeBindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// SafeVOneDotThreeBindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeVOneDotThreeBindingRaw struct {
	Contract *SafeVOneDotThreeBinding // Generic contract binding to access the raw methods on
}

// SafeVOneDotThreeBindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeVOneDotThreeBindingCallerRaw struct {
	Contract *SafeVOneDotThreeBindingCaller // Generic read-only contract binding to access the raw methods on
}

// SafeVOneDotThreeBindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeVOneDotThreeBindingTransactorRaw struct {
	Contract *SafeVOneDotThreeBindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeVOneDotThreeBinding creates a new instance of SafeVOneDotThreeBinding, bound to a specific deployed contract.
func NewSafeVOneDotThreeBinding(address common.Address, backend bind.ContractBackend) (*SafeVOneDotThreeBinding, error) {
	contract, err := bindSafeVOneDotThreeBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBinding{SafeVOneDotThreeBindingCaller: SafeVOneDotThreeBindingCaller{contract: contract}, SafeVOneDotThreeBindingTransactor: SafeVOneDotThreeBindingTransactor{contract: contract}, SafeVOneDotThreeBindingFilterer: SafeVOneDotThreeBindingFilterer{contract: contract}}, nil
}

// NewSafeVOneDotThreeBindingCaller creates a new read-only instance of SafeVOneDotThreeBinding, bound to a specific deployed contract.
func NewSafeVOneDotThreeBindingCaller(address common.Address, caller bind.ContractCaller) (*SafeVOneDotThreeBindingCaller, error) {
	contract, err := bindSafeVOneDotThreeBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingCaller{contract: contract}, nil
}

// NewSafeVOneDotThreeBindingTransactor creates a new write-only instance of SafeVOneDotThreeBinding, bound to a specific deployed contract.
func NewSafeVOneDotThreeBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeVOneDotThreeBindingTransactor, error) {
	contract, err := bindSafeVOneDotThreeBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingTransactor{contract: contract}, nil
}

// NewSafeVOneDotThreeBindingFilterer creates a new log filterer instance of SafeVOneDotThreeBinding, bound to a specific deployed contract.
func NewSafeVOneDotThreeBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeVOneDotThreeBindingFilterer, error) {
	contract, err := bindSafeVOneDotThreeBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingFilterer{contract: contract}, nil
}

// bindSafeVOneDotThreeBinding binds a generic wrapper to an already deployed contract.
func bindSafeVOneDotThreeBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeVOneDotThreeBindingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeVOneDotThreeBinding.Contract.SafeVOneDotThreeBindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SafeVOneDotThreeBindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SafeVOneDotThreeBindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeVOneDotThreeBinding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) VERSION() (string, error) {
	return _SafeVOneDotThreeBinding.Contract.VERSION(&_SafeVOneDotThreeBinding.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) VERSION() (string, error) {
	return _SafeVOneDotThreeBinding.Contract.VERSION(&_SafeVOneDotThreeBinding.CallOpts)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) ApprovedHashes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "approvedHashes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.ApprovedHashes(&_SafeVOneDotThreeBinding.CallOpts, arg0, arg1)
}

// ApprovedHashes is a free data retrieval call binding the contract method 0x7d832974.
//
// Solidity: function approvedHashes(address , bytes32 ) view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) ApprovedHashes(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.ApprovedHashes(&_SafeVOneDotThreeBinding.CallOpts, arg0, arg1)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) CheckNSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "checkNSignatures", dataHash, data, signatures, requiredSignatures)

	if err != nil {
		return err
	}

	return err

}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _SafeVOneDotThreeBinding.Contract.CheckNSignatures(&_SafeVOneDotThreeBinding.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckNSignatures is a free data retrieval call binding the contract method 0x12fb68e0.
//
// Solidity: function checkNSignatures(bytes32 dataHash, bytes data, bytes signatures, uint256 requiredSignatures) view returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) CheckNSignatures(dataHash [32]byte, data []byte, signatures []byte, requiredSignatures *big.Int) error {
	return _SafeVOneDotThreeBinding.Contract.CheckNSignatures(&_SafeVOneDotThreeBinding.CallOpts, dataHash, data, signatures, requiredSignatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) CheckSignatures(opts *bind.CallOpts, dataHash [32]byte, data []byte, signatures []byte) error {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "checkSignatures", dataHash, data, signatures)

	if err != nil {
		return err
	}

	return err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _SafeVOneDotThreeBinding.Contract.CheckSignatures(&_SafeVOneDotThreeBinding.CallOpts, dataHash, data, signatures)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x934f3a11.
//
// Solidity: function checkSignatures(bytes32 dataHash, bytes data, bytes signatures) view returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) CheckSignatures(dataHash [32]byte, data []byte, signatures []byte) error {
	return _SafeVOneDotThreeBinding.Contract.CheckSignatures(&_SafeVOneDotThreeBinding.CallOpts, dataHash, data, signatures)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) DomainSeparator() ([32]byte, error) {
	return _SafeVOneDotThreeBinding.Contract.DomainSeparator(&_SafeVOneDotThreeBinding.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) DomainSeparator() ([32]byte, error) {
	return _SafeVOneDotThreeBinding.Contract.DomainSeparator(&_SafeVOneDotThreeBinding.CallOpts)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) EncodeTransactionData(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "encodeTransactionData", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _SafeVOneDotThreeBinding.Contract.EncodeTransactionData(&_SafeVOneDotThreeBinding.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// EncodeTransactionData is a free data retrieval call binding the contract method 0xe86637db.
//
// Solidity: function encodeTransactionData(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) EncodeTransactionData(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([]byte, error) {
	return _SafeVOneDotThreeBinding.Contract.EncodeTransactionData(&_SafeVOneDotThreeBinding.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) GetChainId() (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.GetChainId(&_SafeVOneDotThreeBinding.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) GetChainId() (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.GetChainId(&_SafeVOneDotThreeBinding.CallOpts)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) GetModulesPaginated(opts *bind.CallOpts, start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "getModulesPaginated", start, pageSize)

	outstruct := new(struct {
		Array []common.Address
		Next  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Array = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Next = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _SafeVOneDotThreeBinding.Contract.GetModulesPaginated(&_SafeVOneDotThreeBinding.CallOpts, start, pageSize)
}

// GetModulesPaginated is a free data retrieval call binding the contract method 0xcc2f8452.
//
// Solidity: function getModulesPaginated(address start, uint256 pageSize) view returns(address[] array, address next)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) GetModulesPaginated(start common.Address, pageSize *big.Int) (struct {
	Array []common.Address
	Next  common.Address
}, error) {
	return _SafeVOneDotThreeBinding.Contract.GetModulesPaginated(&_SafeVOneDotThreeBinding.CallOpts, start, pageSize)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) GetOwners() ([]common.Address, error) {
	return _SafeVOneDotThreeBinding.Contract.GetOwners(&_SafeVOneDotThreeBinding.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) GetOwners() ([]common.Address, error) {
	return _SafeVOneDotThreeBinding.Contract.GetOwners(&_SafeVOneDotThreeBinding.CallOpts)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _SafeVOneDotThreeBinding.Contract.GetStorageAt(&_SafeVOneDotThreeBinding.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _SafeVOneDotThreeBinding.Contract.GetStorageAt(&_SafeVOneDotThreeBinding.CallOpts, offset, length)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) GetThreshold() (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.GetThreshold(&_SafeVOneDotThreeBinding.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) GetThreshold() (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.GetThreshold(&_SafeVOneDotThreeBinding.CallOpts)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) GetTransactionHash(opts *bind.CallOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "getTransactionHash", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _SafeVOneDotThreeBinding.Contract.GetTransactionHash(&_SafeVOneDotThreeBinding.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// GetTransactionHash is a free data retrieval call binding the contract method 0xd8d11f78.
//
// Solidity: function getTransactionHash(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, uint256 _nonce) view returns(bytes32)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) GetTransactionHash(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, _nonce *big.Int) ([32]byte, error) {
	return _SafeVOneDotThreeBinding.Contract.GetTransactionHash(&_SafeVOneDotThreeBinding.CallOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, _nonce)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) IsModuleEnabled(opts *bind.CallOpts, module common.Address) (bool, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "isModuleEnabled", module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _SafeVOneDotThreeBinding.Contract.IsModuleEnabled(&_SafeVOneDotThreeBinding.CallOpts, module)
}

// IsModuleEnabled is a free data retrieval call binding the contract method 0x2d9ad53d.
//
// Solidity: function isModuleEnabled(address module) view returns(bool)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) IsModuleEnabled(module common.Address) (bool, error) {
	return _SafeVOneDotThreeBinding.Contract.IsModuleEnabled(&_SafeVOneDotThreeBinding.CallOpts, module)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) IsOwner(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "isOwner", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) IsOwner(owner common.Address) (bool, error) {
	return _SafeVOneDotThreeBinding.Contract.IsOwner(&_SafeVOneDotThreeBinding.CallOpts, owner)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address owner) view returns(bool)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) IsOwner(owner common.Address) (bool, error) {
	return _SafeVOneDotThreeBinding.Contract.IsOwner(&_SafeVOneDotThreeBinding.CallOpts, owner)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) Nonce() (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.Nonce(&_SafeVOneDotThreeBinding.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) Nonce() (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.Nonce(&_SafeVOneDotThreeBinding.CallOpts)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCaller) SignedMessages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SafeVOneDotThreeBinding.contract.Call(opts, &out, "signedMessages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.SignedMessages(&_SafeVOneDotThreeBinding.CallOpts, arg0)
}

// SignedMessages is a free data retrieval call binding the contract method 0x5ae6bd37.
//
// Solidity: function signedMessages(bytes32 ) view returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingCallerSession) SignedMessages(arg0 [32]byte) (*big.Int, error) {
	return _SafeVOneDotThreeBinding.Contract.SignedMessages(&_SafeVOneDotThreeBinding.CallOpts, arg0)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) AddOwnerWithThreshold(opts *bind.TransactOpts, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "addOwnerWithThreshold", owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.AddOwnerWithThreshold(&_SafeVOneDotThreeBinding.TransactOpts, owner, _threshold)
}

// AddOwnerWithThreshold is a paid mutator transaction binding the contract method 0x0d582f13.
//
// Solidity: function addOwnerWithThreshold(address owner, uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) AddOwnerWithThreshold(owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.AddOwnerWithThreshold(&_SafeVOneDotThreeBinding.TransactOpts, owner, _threshold)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) ApproveHash(opts *bind.TransactOpts, hashToApprove [32]byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "approveHash", hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ApproveHash(&_SafeVOneDotThreeBinding.TransactOpts, hashToApprove)
}

// ApproveHash is a paid mutator transaction binding the contract method 0xd4d9bdcd.
//
// Solidity: function approveHash(bytes32 hashToApprove) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) ApproveHash(hashToApprove [32]byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ApproveHash(&_SafeVOneDotThreeBinding.TransactOpts, hashToApprove)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) ChangeThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "changeThreshold", _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ChangeThreshold(&_SafeVOneDotThreeBinding.TransactOpts, _threshold)
}

// ChangeThreshold is a paid mutator transaction binding the contract method 0x694e80c3.
//
// Solidity: function changeThreshold(uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) ChangeThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ChangeThreshold(&_SafeVOneDotThreeBinding.TransactOpts, _threshold)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) DisableModule(opts *bind.TransactOpts, prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "disableModule", prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.DisableModule(&_SafeVOneDotThreeBinding.TransactOpts, prevModule, module)
}

// DisableModule is a paid mutator transaction binding the contract method 0xe009cfde.
//
// Solidity: function disableModule(address prevModule, address module) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) DisableModule(prevModule common.Address, module common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.DisableModule(&_SafeVOneDotThreeBinding.TransactOpts, prevModule, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) EnableModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "enableModule", module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.EnableModule(&_SafeVOneDotThreeBinding.TransactOpts, module)
}

// EnableModule is a paid mutator transaction binding the contract method 0x610b5925.
//
// Solidity: function enableModule(address module) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) EnableModule(module common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.EnableModule(&_SafeVOneDotThreeBinding.TransactOpts, module)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) ExecTransaction(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "execTransaction", to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ExecTransaction(&_SafeVOneDotThreeBinding.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x6a761202.
//
// Solidity: function execTransaction(address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) ExecTransaction(to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ExecTransaction(&_SafeVOneDotThreeBinding.TransactOpts, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) ExecTransactionFromModule(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "execTransactionFromModule", to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ExecTransactionFromModule(&_SafeVOneDotThreeBinding.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModule is a paid mutator transaction binding the contract method 0x468721a7.
//
// Solidity: function execTransactionFromModule(address to, uint256 value, bytes data, uint8 operation) returns(bool success)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) ExecTransactionFromModule(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ExecTransactionFromModule(&_SafeVOneDotThreeBinding.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) ExecTransactionFromModuleReturnData(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "execTransactionFromModuleReturnData", to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ExecTransactionFromModuleReturnData(&_SafeVOneDotThreeBinding.TransactOpts, to, value, data, operation)
}

// ExecTransactionFromModuleReturnData is a paid mutator transaction binding the contract method 0x5229073f.
//
// Solidity: function execTransactionFromModuleReturnData(address to, uint256 value, bytes data, uint8 operation) returns(bool success, bytes returnData)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) ExecTransactionFromModuleReturnData(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.ExecTransactionFromModuleReturnData(&_SafeVOneDotThreeBinding.TransactOpts, to, value, data, operation)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) RemoveOwner(opts *bind.TransactOpts, prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "removeOwner", prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.RemoveOwner(&_SafeVOneDotThreeBinding.TransactOpts, prevOwner, owner, _threshold)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0xf8dc5dd9.
//
// Solidity: function removeOwner(address prevOwner, address owner, uint256 _threshold) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) RemoveOwner(prevOwner common.Address, owner common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.RemoveOwner(&_SafeVOneDotThreeBinding.TransactOpts, prevOwner, owner, _threshold)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) RequiredTxGas(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "requiredTxGas", to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.RequiredTxGas(&_SafeVOneDotThreeBinding.TransactOpts, to, value, data, operation)
}

// RequiredTxGas is a paid mutator transaction binding the contract method 0xc4ca3a9c.
//
// Solidity: function requiredTxGas(address to, uint256 value, bytes data, uint8 operation) returns(uint256)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) RequiredTxGas(to common.Address, value *big.Int, data []byte, operation uint8) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.RequiredTxGas(&_SafeVOneDotThreeBinding.TransactOpts, to, value, data, operation)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) SetFallbackHandler(opts *bind.TransactOpts, handler common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "setFallbackHandler", handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SetFallbackHandler(&_SafeVOneDotThreeBinding.TransactOpts, handler)
}

// SetFallbackHandler is a paid mutator transaction binding the contract method 0xf08a0323.
//
// Solidity: function setFallbackHandler(address handler) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) SetFallbackHandler(handler common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SetFallbackHandler(&_SafeVOneDotThreeBinding.TransactOpts, handler)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) SetGuard(opts *bind.TransactOpts, guard common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "setGuard", guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SetGuard(&_SafeVOneDotThreeBinding.TransactOpts, guard)
}

// SetGuard is a paid mutator transaction binding the contract method 0xe19a9dd9.
//
// Solidity: function setGuard(address guard) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) SetGuard(guard common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SetGuard(&_SafeVOneDotThreeBinding.TransactOpts, guard)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) Setup(opts *bind.TransactOpts, _owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "setup", _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.Setup(&_SafeVOneDotThreeBinding.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// Setup is a paid mutator transaction binding the contract method 0xb63e800d.
//
// Solidity: function setup(address[] _owners, uint256 _threshold, address to, bytes data, address fallbackHandler, address paymentToken, uint256 payment, address paymentReceiver) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) Setup(_owners []common.Address, _threshold *big.Int, to common.Address, data []byte, fallbackHandler common.Address, paymentToken common.Address, payment *big.Int, paymentReceiver common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.Setup(&_SafeVOneDotThreeBinding.TransactOpts, _owners, _threshold, to, data, fallbackHandler, paymentToken, payment, paymentReceiver)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) SimulateAndRevert(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "simulateAndRevert", targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SimulateAndRevert(&_SafeVOneDotThreeBinding.TransactOpts, targetContract, calldataPayload)
}

// SimulateAndRevert is a paid mutator transaction binding the contract method 0xb4faba09.
//
// Solidity: function simulateAndRevert(address targetContract, bytes calldataPayload) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) SimulateAndRevert(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SimulateAndRevert(&_SafeVOneDotThreeBinding.TransactOpts, targetContract, calldataPayload)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) SwapOwner(opts *bind.TransactOpts, prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.Transact(opts, "swapOwner", prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SwapOwner(&_SafeVOneDotThreeBinding.TransactOpts, prevOwner, oldOwner, newOwner)
}

// SwapOwner is a paid mutator transaction binding the contract method 0xe318b52b.
//
// Solidity: function swapOwner(address prevOwner, address oldOwner, address newOwner) returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) SwapOwner(prevOwner common.Address, oldOwner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.SwapOwner(&_SafeVOneDotThreeBinding.TransactOpts, prevOwner, oldOwner, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.Fallback(&_SafeVOneDotThreeBinding.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.Fallback(&_SafeVOneDotThreeBinding.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingSession) Receive() (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.Receive(&_SafeVOneDotThreeBinding.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingTransactorSession) Receive() (*types.Transaction, error) {
	return _SafeVOneDotThreeBinding.Contract.Receive(&_SafeVOneDotThreeBinding.TransactOpts)
}

// SafeVOneDotThreeBindingAddedOwnerIterator is returned from FilterAddedOwner and is used to iterate over the raw logs and unpacked data for AddedOwner events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingAddedOwnerIterator struct {
	Event *SafeVOneDotThreeBindingAddedOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingAddedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingAddedOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingAddedOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingAddedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingAddedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingAddedOwner represents a AddedOwner event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingAddedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddedOwner is a free log retrieval operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterAddedOwner(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingAddedOwnerIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingAddedOwnerIterator{contract: _SafeVOneDotThreeBinding.contract, event: "AddedOwner", logs: logs, sub: sub}, nil
}

// WatchAddedOwner is a free log subscription operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchAddedOwner(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingAddedOwner) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "AddedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingAddedOwner)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "AddedOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddedOwner is a log parse operation binding the contract event 0x9465fa0c962cc76958e6373a993326400c1c94f8be2fe3a952adfa7f60b2ea26.
//
// Solidity: event AddedOwner(address owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseAddedOwner(log types.Log) (*SafeVOneDotThreeBindingAddedOwner, error) {
	event := new(SafeVOneDotThreeBindingAddedOwner)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "AddedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingApproveHashIterator is returned from FilterApproveHash and is used to iterate over the raw logs and unpacked data for ApproveHash events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingApproveHashIterator struct {
	Event *SafeVOneDotThreeBindingApproveHash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingApproveHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingApproveHash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingApproveHash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingApproveHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingApproveHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingApproveHash represents a ApproveHash event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingApproveHash struct {
	ApprovedHash [32]byte
	Owner        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterApproveHash is a free log retrieval operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterApproveHash(opts *bind.FilterOpts, approvedHash [][32]byte, owner []common.Address) (*SafeVOneDotThreeBindingApproveHashIterator, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingApproveHashIterator{contract: _SafeVOneDotThreeBinding.contract, event: "ApproveHash", logs: logs, sub: sub}, nil
}

// WatchApproveHash is a free log subscription operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchApproveHash(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingApproveHash, approvedHash [][32]byte, owner []common.Address) (event.Subscription, error) {

	var approvedHashRule []interface{}
	for _, approvedHashItem := range approvedHash {
		approvedHashRule = append(approvedHashRule, approvedHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "ApproveHash", approvedHashRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingApproveHash)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ApproveHash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproveHash is a log parse operation binding the contract event 0xf2a0eb156472d1440255b0d7c1e19cc07115d1051fe605b0dce69acfec884d9c.
//
// Solidity: event ApproveHash(bytes32 indexed approvedHash, address indexed owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseApproveHash(log types.Log) (*SafeVOneDotThreeBindingApproveHash, error) {
	event := new(SafeVOneDotThreeBindingApproveHash)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ApproveHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingChangedFallbackHandlerIterator is returned from FilterChangedFallbackHandler and is used to iterate over the raw logs and unpacked data for ChangedFallbackHandler events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingChangedFallbackHandlerIterator struct {
	Event *SafeVOneDotThreeBindingChangedFallbackHandler // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingChangedFallbackHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingChangedFallbackHandler)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingChangedFallbackHandler)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingChangedFallbackHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingChangedFallbackHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingChangedFallbackHandler represents a ChangedFallbackHandler event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingChangedFallbackHandler struct {
	Handler common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChangedFallbackHandler is a free log retrieval operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterChangedFallbackHandler(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingChangedFallbackHandlerIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingChangedFallbackHandlerIterator{contract: _SafeVOneDotThreeBinding.contract, event: "ChangedFallbackHandler", logs: logs, sub: sub}, nil
}

// WatchChangedFallbackHandler is a free log subscription operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchChangedFallbackHandler(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingChangedFallbackHandler) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "ChangedFallbackHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingChangedFallbackHandler)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedFallbackHandler is a log parse operation binding the contract event 0x5ac6c46c93c8d0e53714ba3b53db3e7c046da994313d7ed0d192028bc7c228b0.
//
// Solidity: event ChangedFallbackHandler(address handler)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseChangedFallbackHandler(log types.Log) (*SafeVOneDotThreeBindingChangedFallbackHandler, error) {
	event := new(SafeVOneDotThreeBindingChangedFallbackHandler)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ChangedFallbackHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingChangedGuardIterator is returned from FilterChangedGuard and is used to iterate over the raw logs and unpacked data for ChangedGuard events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingChangedGuardIterator struct {
	Event *SafeVOneDotThreeBindingChangedGuard // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingChangedGuardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingChangedGuard)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingChangedGuard)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingChangedGuardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingChangedGuardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingChangedGuard represents a ChangedGuard event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingChangedGuard struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChangedGuard is a free log retrieval operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterChangedGuard(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingChangedGuardIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingChangedGuardIterator{contract: _SafeVOneDotThreeBinding.contract, event: "ChangedGuard", logs: logs, sub: sub}, nil
}

// WatchChangedGuard is a free log subscription operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchChangedGuard(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingChangedGuard) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "ChangedGuard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingChangedGuard)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedGuard is a log parse operation binding the contract event 0x1151116914515bc0891ff9047a6cb32cf902546f83066499bcf8ba33d2353fa2.
//
// Solidity: event ChangedGuard(address guard)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseChangedGuard(log types.Log) (*SafeVOneDotThreeBindingChangedGuard, error) {
	event := new(SafeVOneDotThreeBindingChangedGuard)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ChangedGuard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingChangedThresholdIterator is returned from FilterChangedThreshold and is used to iterate over the raw logs and unpacked data for ChangedThreshold events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingChangedThresholdIterator struct {
	Event *SafeVOneDotThreeBindingChangedThreshold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingChangedThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingChangedThreshold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingChangedThreshold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingChangedThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingChangedThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingChangedThreshold represents a ChangedThreshold event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingChangedThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChangedThreshold is a free log retrieval operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterChangedThreshold(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingChangedThresholdIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingChangedThresholdIterator{contract: _SafeVOneDotThreeBinding.contract, event: "ChangedThreshold", logs: logs, sub: sub}, nil
}

// WatchChangedThreshold is a free log subscription operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchChangedThreshold(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingChangedThreshold) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "ChangedThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingChangedThreshold)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseChangedThreshold is a log parse operation binding the contract event 0x610f7ff2b304ae8903c3de74c60c6ab1f7d6226b3f52c5161905bb5ad4039c93.
//
// Solidity: event ChangedThreshold(uint256 threshold)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseChangedThreshold(log types.Log) (*SafeVOneDotThreeBindingChangedThreshold, error) {
	event := new(SafeVOneDotThreeBindingChangedThreshold)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ChangedThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingDisabledModuleIterator is returned from FilterDisabledModule and is used to iterate over the raw logs and unpacked data for DisabledModule events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingDisabledModuleIterator struct {
	Event *SafeVOneDotThreeBindingDisabledModule // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingDisabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingDisabledModule)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingDisabledModule)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingDisabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingDisabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingDisabledModule represents a DisabledModule event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingDisabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisabledModule is a free log retrieval operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterDisabledModule(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingDisabledModuleIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingDisabledModuleIterator{contract: _SafeVOneDotThreeBinding.contract, event: "DisabledModule", logs: logs, sub: sub}, nil
}

// WatchDisabledModule is a free log subscription operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchDisabledModule(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingDisabledModule) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "DisabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingDisabledModule)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "DisabledModule", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisabledModule is a log parse operation binding the contract event 0xaab4fa2b463f581b2b32cb3b7e3b704b9ce37cc209b5fb4d77e593ace4054276.
//
// Solidity: event DisabledModule(address module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseDisabledModule(log types.Log) (*SafeVOneDotThreeBindingDisabledModule, error) {
	event := new(SafeVOneDotThreeBindingDisabledModule)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "DisabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingEnabledModuleIterator is returned from FilterEnabledModule and is used to iterate over the raw logs and unpacked data for EnabledModule events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingEnabledModuleIterator struct {
	Event *SafeVOneDotThreeBindingEnabledModule // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingEnabledModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingEnabledModule)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingEnabledModule)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingEnabledModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingEnabledModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingEnabledModule represents a EnabledModule event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingEnabledModule struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnabledModule is a free log retrieval operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterEnabledModule(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingEnabledModuleIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingEnabledModuleIterator{contract: _SafeVOneDotThreeBinding.contract, event: "EnabledModule", logs: logs, sub: sub}, nil
}

// WatchEnabledModule is a free log subscription operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchEnabledModule(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingEnabledModule) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "EnabledModule")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingEnabledModule)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "EnabledModule", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEnabledModule is a log parse operation binding the contract event 0xecdf3a3effea5783a3c4c2140e677577666428d44ed9d474a0b3a4c9943f8440.
//
// Solidity: event EnabledModule(address module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseEnabledModule(log types.Log) (*SafeVOneDotThreeBindingEnabledModule, error) {
	event := new(SafeVOneDotThreeBindingEnabledModule)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "EnabledModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingExecutionFailureIterator struct {
	Event *SafeVOneDotThreeBindingExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingExecutionFailure represents a ExecutionFailure event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingExecutionFailure struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterExecutionFailure(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingExecutionFailureIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingExecutionFailureIterator{contract: _SafeVOneDotThreeBinding.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingExecutionFailure) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingExecutionFailure)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFailure is a log parse operation binding the contract event 0x23428b18acfb3ea64b08dc0c1d296ea9c09702c09083ca5272e64d115b687d23.
//
// Solidity: event ExecutionFailure(bytes32 txHash, uint256 payment)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseExecutionFailure(log types.Log) (*SafeVOneDotThreeBindingExecutionFailure, error) {
	event := new(SafeVOneDotThreeBindingExecutionFailure)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingExecutionFromModuleFailureIterator is returned from FilterExecutionFromModuleFailure and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleFailure events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingExecutionFromModuleFailureIterator struct {
	Event *SafeVOneDotThreeBindingExecutionFromModuleFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingExecutionFromModuleFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingExecutionFromModuleFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingExecutionFromModuleFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingExecutionFromModuleFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingExecutionFromModuleFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingExecutionFromModuleFailure represents a ExecutionFromModuleFailure event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingExecutionFromModuleFailure struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleFailure is a free log retrieval operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterExecutionFromModuleFailure(opts *bind.FilterOpts, module []common.Address) (*SafeVOneDotThreeBindingExecutionFromModuleFailureIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingExecutionFromModuleFailureIterator{contract: _SafeVOneDotThreeBinding.contract, event: "ExecutionFromModuleFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleFailure is a free log subscription operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchExecutionFromModuleFailure(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingExecutionFromModuleFailure, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "ExecutionFromModuleFailure", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingExecutionFromModuleFailure)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFromModuleFailure is a log parse operation binding the contract event 0xacd2c8702804128fdb0db2bb49f6d127dd0181c13fd45dbfe16de0930e2bd375.
//
// Solidity: event ExecutionFromModuleFailure(address indexed module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseExecutionFromModuleFailure(log types.Log) (*SafeVOneDotThreeBindingExecutionFromModuleFailure, error) {
	event := new(SafeVOneDotThreeBindingExecutionFromModuleFailure)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ExecutionFromModuleFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingExecutionFromModuleSuccessIterator is returned from FilterExecutionFromModuleSuccess and is used to iterate over the raw logs and unpacked data for ExecutionFromModuleSuccess events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingExecutionFromModuleSuccessIterator struct {
	Event *SafeVOneDotThreeBindingExecutionFromModuleSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingExecutionFromModuleSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingExecutionFromModuleSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingExecutionFromModuleSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingExecutionFromModuleSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingExecutionFromModuleSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingExecutionFromModuleSuccess represents a ExecutionFromModuleSuccess event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingExecutionFromModuleSuccess struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExecutionFromModuleSuccess is a free log retrieval operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterExecutionFromModuleSuccess(opts *bind.FilterOpts, module []common.Address) (*SafeVOneDotThreeBindingExecutionFromModuleSuccessIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingExecutionFromModuleSuccessIterator{contract: _SafeVOneDotThreeBinding.contract, event: "ExecutionFromModuleSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionFromModuleSuccess is a free log subscription operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchExecutionFromModuleSuccess(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingExecutionFromModuleSuccess, module []common.Address) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "ExecutionFromModuleSuccess", moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingExecutionFromModuleSuccess)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionFromModuleSuccess is a log parse operation binding the contract event 0x6895c13664aa4f67288b25d7a21d7aaa34916e355fb9b6fae0a139a9085becb8.
//
// Solidity: event ExecutionFromModuleSuccess(address indexed module)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseExecutionFromModuleSuccess(log types.Log) (*SafeVOneDotThreeBindingExecutionFromModuleSuccess, error) {
	event := new(SafeVOneDotThreeBindingExecutionFromModuleSuccess)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ExecutionFromModuleSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingExecutionSuccessIterator is returned from FilterExecutionSuccess and is used to iterate over the raw logs and unpacked data for ExecutionSuccess events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingExecutionSuccessIterator struct {
	Event *SafeVOneDotThreeBindingExecutionSuccess // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingExecutionSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingExecutionSuccess)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingExecutionSuccess)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingExecutionSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingExecutionSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingExecutionSuccess represents a ExecutionSuccess event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingExecutionSuccess struct {
	TxHash  [32]byte
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecutionSuccess is a free log retrieval operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterExecutionSuccess(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingExecutionSuccessIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingExecutionSuccessIterator{contract: _SafeVOneDotThreeBinding.contract, event: "ExecutionSuccess", logs: logs, sub: sub}, nil
}

// WatchExecutionSuccess is a free log subscription operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchExecutionSuccess(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingExecutionSuccess) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "ExecutionSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingExecutionSuccess)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionSuccess is a log parse operation binding the contract event 0x442e715f626346e8c54381002da614f62bee8d27386535b2521ec8540898556e.
//
// Solidity: event ExecutionSuccess(bytes32 txHash, uint256 payment)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseExecutionSuccess(log types.Log) (*SafeVOneDotThreeBindingExecutionSuccess, error) {
	event := new(SafeVOneDotThreeBindingExecutionSuccess)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "ExecutionSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingRemovedOwnerIterator is returned from FilterRemovedOwner and is used to iterate over the raw logs and unpacked data for RemovedOwner events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingRemovedOwnerIterator struct {
	Event *SafeVOneDotThreeBindingRemovedOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingRemovedOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingRemovedOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingRemovedOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingRemovedOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingRemovedOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingRemovedOwner represents a RemovedOwner event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingRemovedOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemovedOwner is a free log retrieval operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterRemovedOwner(opts *bind.FilterOpts) (*SafeVOneDotThreeBindingRemovedOwnerIterator, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingRemovedOwnerIterator{contract: _SafeVOneDotThreeBinding.contract, event: "RemovedOwner", logs: logs, sub: sub}, nil
}

// WatchRemovedOwner is a free log subscription operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchRemovedOwner(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingRemovedOwner) (event.Subscription, error) {

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "RemovedOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingRemovedOwner)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedOwner is a log parse operation binding the contract event 0xf8d49fc529812e9a7c5c50e69c20f0dccc0db8fa95c98bc58cc9a4f1c1299eaf.
//
// Solidity: event RemovedOwner(address owner)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseRemovedOwner(log types.Log) (*SafeVOneDotThreeBindingRemovedOwner, error) {
	event := new(SafeVOneDotThreeBindingRemovedOwner)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "RemovedOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingSafeReceivedIterator is returned from FilterSafeReceived and is used to iterate over the raw logs and unpacked data for SafeReceived events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingSafeReceivedIterator struct {
	Event *SafeVOneDotThreeBindingSafeReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingSafeReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingSafeReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingSafeReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingSafeReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingSafeReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingSafeReceived represents a SafeReceived event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingSafeReceived struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSafeReceived is a free log retrieval operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterSafeReceived(opts *bind.FilterOpts, sender []common.Address) (*SafeVOneDotThreeBindingSafeReceivedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingSafeReceivedIterator{contract: _SafeVOneDotThreeBinding.contract, event: "SafeReceived", logs: logs, sub: sub}, nil
}

// WatchSafeReceived is a free log subscription operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchSafeReceived(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingSafeReceived, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "SafeReceived", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingSafeReceived)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "SafeReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSafeReceived is a log parse operation binding the contract event 0x3d0ce9bfc3ed7d6862dbb28b2dea94561fe714a1b4d019aa8af39730d1ad7c3d.
//
// Solidity: event SafeReceived(address indexed sender, uint256 value)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseSafeReceived(log types.Log) (*SafeVOneDotThreeBindingSafeReceived, error) {
	event := new(SafeVOneDotThreeBindingSafeReceived)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "SafeReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingSafeSetupIterator is returned from FilterSafeSetup and is used to iterate over the raw logs and unpacked data for SafeSetup events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingSafeSetupIterator struct {
	Event *SafeVOneDotThreeBindingSafeSetup // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingSafeSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingSafeSetup)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingSafeSetup)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingSafeSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingSafeSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingSafeSetup represents a SafeSetup event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingSafeSetup struct {
	Initiator       common.Address
	Owners          []common.Address
	Threshold       *big.Int
	Initializer     common.Address
	FallbackHandler common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSafeSetup is a free log retrieval operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterSafeSetup(opts *bind.FilterOpts, initiator []common.Address) (*SafeVOneDotThreeBindingSafeSetupIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingSafeSetupIterator{contract: _SafeVOneDotThreeBinding.contract, event: "SafeSetup", logs: logs, sub: sub}, nil
}

// WatchSafeSetup is a free log subscription operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchSafeSetup(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingSafeSetup, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "SafeSetup", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingSafeSetup)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "SafeSetup", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSafeSetup is a log parse operation binding the contract event 0x141df868a6331af528e38c83b7aa03edc19be66e37ae67f9285bf4f8e3c6a1a8.
//
// Solidity: event SafeSetup(address indexed initiator, address[] owners, uint256 threshold, address initializer, address fallbackHandler)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseSafeSetup(log types.Log) (*SafeVOneDotThreeBindingSafeSetup, error) {
	event := new(SafeVOneDotThreeBindingSafeSetup)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "SafeSetup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeVOneDotThreeBindingSignMsgIterator is returned from FilterSignMsg and is used to iterate over the raw logs and unpacked data for SignMsg events raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingSignMsgIterator struct {
	Event *SafeVOneDotThreeBindingSignMsg // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SafeVOneDotThreeBindingSignMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SafeVOneDotThreeBindingSignMsg)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SafeVOneDotThreeBindingSignMsg)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SafeVOneDotThreeBindingSignMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SafeVOneDotThreeBindingSignMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SafeVOneDotThreeBindingSignMsg represents a SignMsg event raised by the SafeVOneDotThreeBinding contract.
type SafeVOneDotThreeBindingSignMsg struct {
	MsgHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignMsg is a free log retrieval operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) FilterSignMsg(opts *bind.FilterOpts, msgHash [][32]byte) (*SafeVOneDotThreeBindingSignMsgIterator, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.FilterLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return &SafeVOneDotThreeBindingSignMsgIterator{contract: _SafeVOneDotThreeBinding.contract, event: "SignMsg", logs: logs, sub: sub}, nil
}

// WatchSignMsg is a free log subscription operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) WatchSignMsg(opts *bind.WatchOpts, sink chan<- *SafeVOneDotThreeBindingSignMsg, msgHash [][32]byte) (event.Subscription, error) {

	var msgHashRule []interface{}
	for _, msgHashItem := range msgHash {
		msgHashRule = append(msgHashRule, msgHashItem)
	}

	logs, sub, err := _SafeVOneDotThreeBinding.contract.WatchLogs(opts, "SignMsg", msgHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SafeVOneDotThreeBindingSignMsg)
				if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "SignMsg", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignMsg is a log parse operation binding the contract event 0xe7f4675038f4f6034dfcbbb24c4dc08e4ebf10eb9d257d3d02c0f38d122ac6e4.
//
// Solidity: event SignMsg(bytes32 indexed msgHash)
func (_SafeVOneDotThreeBinding *SafeVOneDotThreeBindingFilterer) ParseSignMsg(log types.Log) (*SafeVOneDotThreeBindingSignMsg, error) {
	event := new(SafeVOneDotThreeBindingSignMsg)
	if err := _SafeVOneDotThreeBinding.contract.UnpackLog(event, "SignMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
