// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eip1271validator

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
	_ = abi.ConvertType
)

// Eip1271validatorMetaData contains all meta data concerning the Eip1271validator contract.
var Eip1271validatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Eip1271validatorABI is the input ABI used to generate the binding from.
// Deprecated: Use Eip1271validatorMetaData.ABI instead.
var Eip1271validatorABI = Eip1271validatorMetaData.ABI

// Eip1271validator is an auto generated Go binding around an Ethereum contract.
type Eip1271validator struct {
	Eip1271validatorCaller     // Read-only binding to the contract
	Eip1271validatorTransactor // Write-only binding to the contract
	Eip1271validatorFilterer   // Log filterer for contract events
}

// Eip1271validatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type Eip1271validatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eip1271validatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Eip1271validatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eip1271validatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Eip1271validatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Eip1271validatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Eip1271validatorSession struct {
	Contract     *Eip1271validator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Eip1271validatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Eip1271validatorCallerSession struct {
	Contract *Eip1271validatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Eip1271validatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Eip1271validatorTransactorSession struct {
	Contract     *Eip1271validatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Eip1271validatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type Eip1271validatorRaw struct {
	Contract *Eip1271validator // Generic contract binding to access the raw methods on
}

// Eip1271validatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Eip1271validatorCallerRaw struct {
	Contract *Eip1271validatorCaller // Generic read-only contract binding to access the raw methods on
}

// Eip1271validatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Eip1271validatorTransactorRaw struct {
	Contract *Eip1271validatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEip1271validator creates a new instance of Eip1271validator, bound to a specific deployed contract.
func NewEip1271validator(address common.Address, backend bind.ContractBackend) (*Eip1271validator, error) {
	contract, err := bindEip1271validator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eip1271validator{Eip1271validatorCaller: Eip1271validatorCaller{contract: contract}, Eip1271validatorTransactor: Eip1271validatorTransactor{contract: contract}, Eip1271validatorFilterer: Eip1271validatorFilterer{contract: contract}}, nil
}

// NewEip1271validatorCaller creates a new read-only instance of Eip1271validator, bound to a specific deployed contract.
func NewEip1271validatorCaller(address common.Address, caller bind.ContractCaller) (*Eip1271validatorCaller, error) {
	contract, err := bindEip1271validator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Eip1271validatorCaller{contract: contract}, nil
}

// NewEip1271validatorTransactor creates a new write-only instance of Eip1271validator, bound to a specific deployed contract.
func NewEip1271validatorTransactor(address common.Address, transactor bind.ContractTransactor) (*Eip1271validatorTransactor, error) {
	contract, err := bindEip1271validator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Eip1271validatorTransactor{contract: contract}, nil
}

// NewEip1271validatorFilterer creates a new log filterer instance of Eip1271validator, bound to a specific deployed contract.
func NewEip1271validatorFilterer(address common.Address, filterer bind.ContractFilterer) (*Eip1271validatorFilterer, error) {
	contract, err := bindEip1271validator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Eip1271validatorFilterer{contract: contract}, nil
}

// bindEip1271validator binds a generic wrapper to an already deployed contract.
func bindEip1271validator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Eip1271validatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eip1271validator *Eip1271validatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eip1271validator.Contract.Eip1271validatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eip1271validator *Eip1271validatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eip1271validator.Contract.Eip1271validatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eip1271validator *Eip1271validatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eip1271validator.Contract.Eip1271validatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eip1271validator *Eip1271validatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Eip1271validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eip1271validator *Eip1271validatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eip1271validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eip1271validator *Eip1271validatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eip1271validator.Contract.contract.Transact(opts, method, params...)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Eip1271validator *Eip1271validatorCaller) IsValidSignature(opts *bind.CallOpts, _hash [32]byte, _signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Eip1271validator.contract.Call(opts, &out, "isValidSignature", _hash, _signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Eip1271validator *Eip1271validatorSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Eip1271validator.Contract.IsValidSignature(&_Eip1271validator.CallOpts, _hash, _signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 _hash, bytes _signature) view returns(bytes4)
func (_Eip1271validator *Eip1271validatorCallerSession) IsValidSignature(_hash [32]byte, _signature []byte) ([4]byte, error) {
	return _Eip1271validator.Contract.IsValidSignature(&_Eip1271validator.CallOpts, _hash, _signature)
}
