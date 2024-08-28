// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletregistry

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

// WalletregistryMetaData contains all meta data concerning the Walletregistry contract.
var WalletregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddressProvider\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IsSubAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subAccount\",\"type\":\"address\"}],\"name\":\"RegisterSubAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"RegisterWallet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addressProvider\",\"outputs\":[{\"internalType\":\"contractAddressProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressProviderTarget\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executorRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"getSubAccountsForWallet\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWallet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"policyRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_subAccount\",\"type\":\"address\"}],\"name\":\"registerSubAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"subAccount\",\"type\":\"address\"}],\"name\":\"subAccountToWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"walletToSubAccountList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"subAccountList\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WalletregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletregistryMetaData.ABI instead.
var WalletregistryABI = WalletregistryMetaData.ABI

// Walletregistry is an auto generated Go binding around an Ethereum contract.
type Walletregistry struct {
	WalletregistryCaller     // Read-only binding to the contract
	WalletregistryTransactor // Write-only binding to the contract
	WalletregistryFilterer   // Log filterer for contract events
}

// WalletregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletregistrySession struct {
	Contract     *Walletregistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletregistryCallerSession struct {
	Contract *WalletregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WalletregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletregistryTransactorSession struct {
	Contract     *WalletregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WalletregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletregistryRaw struct {
	Contract *Walletregistry // Generic contract binding to access the raw methods on
}

// WalletregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletregistryCallerRaw struct {
	Contract *WalletregistryCaller // Generic read-only contract binding to access the raw methods on
}

// WalletregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletregistryTransactorRaw struct {
	Contract *WalletregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletregistry creates a new instance of Walletregistry, bound to a specific deployed contract.
func NewWalletregistry(address common.Address, backend bind.ContractBackend) (*Walletregistry, error) {
	contract, err := bindWalletregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Walletregistry{WalletregistryCaller: WalletregistryCaller{contract: contract}, WalletregistryTransactor: WalletregistryTransactor{contract: contract}, WalletregistryFilterer: WalletregistryFilterer{contract: contract}}, nil
}

// NewWalletregistryCaller creates a new read-only instance of Walletregistry, bound to a specific deployed contract.
func NewWalletregistryCaller(address common.Address, caller bind.ContractCaller) (*WalletregistryCaller, error) {
	contract, err := bindWalletregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletregistryCaller{contract: contract}, nil
}

// NewWalletregistryTransactor creates a new write-only instance of Walletregistry, bound to a specific deployed contract.
func NewWalletregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletregistryTransactor, error) {
	contract, err := bindWalletregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletregistryTransactor{contract: contract}, nil
}

// NewWalletregistryFilterer creates a new log filterer instance of Walletregistry, bound to a specific deployed contract.
func NewWalletregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletregistryFilterer, error) {
	contract, err := bindWalletregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletregistryFilterer{contract: contract}, nil
}

// bindWalletregistry binds a generic wrapper to an already deployed contract.
func bindWalletregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Walletregistry *WalletregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Walletregistry.Contract.WalletregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Walletregistry *WalletregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Walletregistry.Contract.WalletregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Walletregistry *WalletregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Walletregistry.Contract.WalletregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Walletregistry *WalletregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Walletregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Walletregistry *WalletregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Walletregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Walletregistry *WalletregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Walletregistry.Contract.contract.Transact(opts, method, params...)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Walletregistry *WalletregistryCaller) AddressProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "addressProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Walletregistry *WalletregistrySession) AddressProvider() (common.Address, error) {
	return _Walletregistry.Contract.AddressProvider(&_Walletregistry.CallOpts)
}

// AddressProvider is a free data retrieval call binding the contract method 0x2954018c.
//
// Solidity: function addressProvider() view returns(address)
func (_Walletregistry *WalletregistryCallerSession) AddressProvider() (common.Address, error) {
	return _Walletregistry.Contract.AddressProvider(&_Walletregistry.CallOpts)
}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_Walletregistry *WalletregistryCaller) AddressProviderTarget(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "addressProviderTarget")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_Walletregistry *WalletregistrySession) AddressProviderTarget() (common.Address, error) {
	return _Walletregistry.Contract.AddressProviderTarget(&_Walletregistry.CallOpts)
}

// AddressProviderTarget is a free data retrieval call binding the contract method 0x21b1e480.
//
// Solidity: function addressProviderTarget() view returns(address)
func (_Walletregistry *WalletregistryCallerSession) AddressProviderTarget() (common.Address, error) {
	return _Walletregistry.Contract.AddressProviderTarget(&_Walletregistry.CallOpts)
}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_Walletregistry *WalletregistryCaller) ExecutorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "executorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_Walletregistry *WalletregistrySession) ExecutorRegistry() (common.Address, error) {
	return _Walletregistry.Contract.ExecutorRegistry(&_Walletregistry.CallOpts)
}

// ExecutorRegistry is a free data retrieval call binding the contract method 0xb1cebbe0.
//
// Solidity: function executorRegistry() view returns(address)
func (_Walletregistry *WalletregistryCallerSession) ExecutorRegistry() (common.Address, error) {
	return _Walletregistry.Contract.ExecutorRegistry(&_Walletregistry.CallOpts)
}

// GetSubAccountsForWallet is a free data retrieval call binding the contract method 0x281f60c1.
//
// Solidity: function getSubAccountsForWallet(address _wallet) view returns(address[])
func (_Walletregistry *WalletregistryCaller) GetSubAccountsForWallet(opts *bind.CallOpts, _wallet common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "getSubAccountsForWallet", _wallet)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSubAccountsForWallet is a free data retrieval call binding the contract method 0x281f60c1.
//
// Solidity: function getSubAccountsForWallet(address _wallet) view returns(address[])
func (_Walletregistry *WalletregistrySession) GetSubAccountsForWallet(_wallet common.Address) ([]common.Address, error) {
	return _Walletregistry.Contract.GetSubAccountsForWallet(&_Walletregistry.CallOpts, _wallet)
}

// GetSubAccountsForWallet is a free data retrieval call binding the contract method 0x281f60c1.
//
// Solidity: function getSubAccountsForWallet(address _wallet) view returns(address[])
func (_Walletregistry *WalletregistryCallerSession) GetSubAccountsForWallet(_wallet common.Address) ([]common.Address, error) {
	return _Walletregistry.Contract.GetSubAccountsForWallet(&_Walletregistry.CallOpts, _wallet)
}

// IsWallet is a free data retrieval call binding the contract method 0xce5570ec.
//
// Solidity: function isWallet(address ) view returns(bool)
func (_Walletregistry *WalletregistryCaller) IsWallet(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "isWallet", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWallet is a free data retrieval call binding the contract method 0xce5570ec.
//
// Solidity: function isWallet(address ) view returns(bool)
func (_Walletregistry *WalletregistrySession) IsWallet(arg0 common.Address) (bool, error) {
	return _Walletregistry.Contract.IsWallet(&_Walletregistry.CallOpts, arg0)
}

// IsWallet is a free data retrieval call binding the contract method 0xce5570ec.
//
// Solidity: function isWallet(address ) view returns(bool)
func (_Walletregistry *WalletregistryCallerSession) IsWallet(arg0 common.Address) (bool, error) {
	return _Walletregistry.Contract.IsWallet(&_Walletregistry.CallOpts, arg0)
}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_Walletregistry *WalletregistryCaller) PolicyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "policyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_Walletregistry *WalletregistrySession) PolicyRegistry() (common.Address, error) {
	return _Walletregistry.Contract.PolicyRegistry(&_Walletregistry.CallOpts)
}

// PolicyRegistry is a free data retrieval call binding the contract method 0x1c4dd7d0.
//
// Solidity: function policyRegistry() view returns(address)
func (_Walletregistry *WalletregistryCallerSession) PolicyRegistry() (common.Address, error) {
	return _Walletregistry.Contract.PolicyRegistry(&_Walletregistry.CallOpts)
}

// SubAccountToWallet is a free data retrieval call binding the contract method 0x0db142f2.
//
// Solidity: function subAccountToWallet(address subAccount) view returns(address wallet)
func (_Walletregistry *WalletregistryCaller) SubAccountToWallet(opts *bind.CallOpts, subAccount common.Address) (common.Address, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "subAccountToWallet", subAccount)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubAccountToWallet is a free data retrieval call binding the contract method 0x0db142f2.
//
// Solidity: function subAccountToWallet(address subAccount) view returns(address wallet)
func (_Walletregistry *WalletregistrySession) SubAccountToWallet(subAccount common.Address) (common.Address, error) {
	return _Walletregistry.Contract.SubAccountToWallet(&_Walletregistry.CallOpts, subAccount)
}

// SubAccountToWallet is a free data retrieval call binding the contract method 0x0db142f2.
//
// Solidity: function subAccountToWallet(address subAccount) view returns(address wallet)
func (_Walletregistry *WalletregistryCallerSession) SubAccountToWallet(subAccount common.Address) (common.Address, error) {
	return _Walletregistry.Contract.SubAccountToWallet(&_Walletregistry.CallOpts, subAccount)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Walletregistry *WalletregistryCaller) WalletRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "walletRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Walletregistry *WalletregistrySession) WalletRegistry() (common.Address, error) {
	return _Walletregistry.Contract.WalletRegistry(&_Walletregistry.CallOpts)
}

// WalletRegistry is a free data retrieval call binding the contract method 0xab7aa6ad.
//
// Solidity: function walletRegistry() view returns(address)
func (_Walletregistry *WalletregistryCallerSession) WalletRegistry() (common.Address, error) {
	return _Walletregistry.Contract.WalletRegistry(&_Walletregistry.CallOpts)
}

// WalletToSubAccountList is a free data retrieval call binding the contract method 0x588d65fb.
//
// Solidity: function walletToSubAccountList(address wallet, uint256 ) view returns(address subAccountList)
func (_Walletregistry *WalletregistryCaller) WalletToSubAccountList(opts *bind.CallOpts, wallet common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Walletregistry.contract.Call(opts, &out, "walletToSubAccountList", wallet, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletToSubAccountList is a free data retrieval call binding the contract method 0x588d65fb.
//
// Solidity: function walletToSubAccountList(address wallet, uint256 ) view returns(address subAccountList)
func (_Walletregistry *WalletregistrySession) WalletToSubAccountList(wallet common.Address, arg1 *big.Int) (common.Address, error) {
	return _Walletregistry.Contract.WalletToSubAccountList(&_Walletregistry.CallOpts, wallet, arg1)
}

// WalletToSubAccountList is a free data retrieval call binding the contract method 0x588d65fb.
//
// Solidity: function walletToSubAccountList(address wallet, uint256 ) view returns(address subAccountList)
func (_Walletregistry *WalletregistryCallerSession) WalletToSubAccountList(wallet common.Address, arg1 *big.Int) (common.Address, error) {
	return _Walletregistry.Contract.WalletToSubAccountList(&_Walletregistry.CallOpts, wallet, arg1)
}

// RegisterSubAccount is a paid mutator transaction binding the contract method 0x13784dec.
//
// Solidity: function registerSubAccount(address _wallet, address _subAccount) returns()
func (_Walletregistry *WalletregistryTransactor) RegisterSubAccount(opts *bind.TransactOpts, _wallet common.Address, _subAccount common.Address) (*types.Transaction, error) {
	return _Walletregistry.contract.Transact(opts, "registerSubAccount", _wallet, _subAccount)
}

// RegisterSubAccount is a paid mutator transaction binding the contract method 0x13784dec.
//
// Solidity: function registerSubAccount(address _wallet, address _subAccount) returns()
func (_Walletregistry *WalletregistrySession) RegisterSubAccount(_wallet common.Address, _subAccount common.Address) (*types.Transaction, error) {
	return _Walletregistry.Contract.RegisterSubAccount(&_Walletregistry.TransactOpts, _wallet, _subAccount)
}

// RegisterSubAccount is a paid mutator transaction binding the contract method 0x13784dec.
//
// Solidity: function registerSubAccount(address _wallet, address _subAccount) returns()
func (_Walletregistry *WalletregistryTransactorSession) RegisterSubAccount(_wallet common.Address, _subAccount common.Address) (*types.Transaction, error) {
	return _Walletregistry.Contract.RegisterSubAccount(&_Walletregistry.TransactOpts, _wallet, _subAccount)
}

// RegisterWallet is a paid mutator transaction binding the contract method 0x33ff495a.
//
// Solidity: function registerWallet() returns()
func (_Walletregistry *WalletregistryTransactor) RegisterWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Walletregistry.contract.Transact(opts, "registerWallet")
}

// RegisterWallet is a paid mutator transaction binding the contract method 0x33ff495a.
//
// Solidity: function registerWallet() returns()
func (_Walletregistry *WalletregistrySession) RegisterWallet() (*types.Transaction, error) {
	return _Walletregistry.Contract.RegisterWallet(&_Walletregistry.TransactOpts)
}

// RegisterWallet is a paid mutator transaction binding the contract method 0x33ff495a.
//
// Solidity: function registerWallet() returns()
func (_Walletregistry *WalletregistryTransactorSession) RegisterWallet() (*types.Transaction, error) {
	return _Walletregistry.Contract.RegisterWallet(&_Walletregistry.TransactOpts)
}

// WalletregistryRegisterSubAccountIterator is returned from FilterRegisterSubAccount and is used to iterate over the raw logs and unpacked data for RegisterSubAccount events raised by the Walletregistry contract.
type WalletregistryRegisterSubAccountIterator struct {
	Event *WalletregistryRegisterSubAccount // Event containing the contract specifics and raw log

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
func (it *WalletregistryRegisterSubAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletregistryRegisterSubAccount)
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
		it.Event = new(WalletregistryRegisterSubAccount)
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
func (it *WalletregistryRegisterSubAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletregistryRegisterSubAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletregistryRegisterSubAccount represents a RegisterSubAccount event raised by the Walletregistry contract.
type WalletregistryRegisterSubAccount struct {
	Wallet     common.Address
	SubAccount common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegisterSubAccount is a free log retrieval operation binding the contract event 0xacb9c63c1050d69261995b89893bf1e5e1fe8f53647db9747212d022f59bea84.
//
// Solidity: event RegisterSubAccount(address indexed wallet, address indexed subAccount)
func (_Walletregistry *WalletregistryFilterer) FilterRegisterSubAccount(opts *bind.FilterOpts, wallet []common.Address, subAccount []common.Address) (*WalletregistryRegisterSubAccountIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var subAccountRule []interface{}
	for _, subAccountItem := range subAccount {
		subAccountRule = append(subAccountRule, subAccountItem)
	}

	logs, sub, err := _Walletregistry.contract.FilterLogs(opts, "RegisterSubAccount", walletRule, subAccountRule)
	if err != nil {
		return nil, err
	}
	return &WalletregistryRegisterSubAccountIterator{contract: _Walletregistry.contract, event: "RegisterSubAccount", logs: logs, sub: sub}, nil
}

// WatchRegisterSubAccount is a free log subscription operation binding the contract event 0xacb9c63c1050d69261995b89893bf1e5e1fe8f53647db9747212d022f59bea84.
//
// Solidity: event RegisterSubAccount(address indexed wallet, address indexed subAccount)
func (_Walletregistry *WalletregistryFilterer) WatchRegisterSubAccount(opts *bind.WatchOpts, sink chan<- *WalletregistryRegisterSubAccount, wallet []common.Address, subAccount []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var subAccountRule []interface{}
	for _, subAccountItem := range subAccount {
		subAccountRule = append(subAccountRule, subAccountItem)
	}

	logs, sub, err := _Walletregistry.contract.WatchLogs(opts, "RegisterSubAccount", walletRule, subAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletregistryRegisterSubAccount)
				if err := _Walletregistry.contract.UnpackLog(event, "RegisterSubAccount", log); err != nil {
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

// ParseRegisterSubAccount is a log parse operation binding the contract event 0xacb9c63c1050d69261995b89893bf1e5e1fe8f53647db9747212d022f59bea84.
//
// Solidity: event RegisterSubAccount(address indexed wallet, address indexed subAccount)
func (_Walletregistry *WalletregistryFilterer) ParseRegisterSubAccount(log types.Log) (*WalletregistryRegisterSubAccount, error) {
	event := new(WalletregistryRegisterSubAccount)
	if err := _Walletregistry.contract.UnpackLog(event, "RegisterSubAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletregistryRegisterWalletIterator is returned from FilterRegisterWallet and is used to iterate over the raw logs and unpacked data for RegisterWallet events raised by the Walletregistry contract.
type WalletregistryRegisterWalletIterator struct {
	Event *WalletregistryRegisterWallet // Event containing the contract specifics and raw log

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
func (it *WalletregistryRegisterWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletregistryRegisterWallet)
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
		it.Event = new(WalletregistryRegisterWallet)
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
func (it *WalletregistryRegisterWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletregistryRegisterWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletregistryRegisterWallet represents a RegisterWallet event raised by the Walletregistry contract.
type WalletregistryRegisterWallet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegisterWallet is a free log retrieval operation binding the contract event 0x3de0c4f89978c8c9ac219df0e8813f0b4a3df2fad979b3e3607cbe6c0ce6c764.
//
// Solidity: event RegisterWallet(address indexed wallet)
func (_Walletregistry *WalletregistryFilterer) FilterRegisterWallet(opts *bind.FilterOpts, wallet []common.Address) (*WalletregistryRegisterWalletIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Walletregistry.contract.FilterLogs(opts, "RegisterWallet", walletRule)
	if err != nil {
		return nil, err
	}
	return &WalletregistryRegisterWalletIterator{contract: _Walletregistry.contract, event: "RegisterWallet", logs: logs, sub: sub}, nil
}

// WatchRegisterWallet is a free log subscription operation binding the contract event 0x3de0c4f89978c8c9ac219df0e8813f0b4a3df2fad979b3e3607cbe6c0ce6c764.
//
// Solidity: event RegisterWallet(address indexed wallet)
func (_Walletregistry *WalletregistryFilterer) WatchRegisterWallet(opts *bind.WatchOpts, sink chan<- *WalletregistryRegisterWallet, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Walletregistry.contract.WatchLogs(opts, "RegisterWallet", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletregistryRegisterWallet)
				if err := _Walletregistry.contract.UnpackLog(event, "RegisterWallet", log); err != nil {
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

// ParseRegisterWallet is a log parse operation binding the contract event 0x3de0c4f89978c8c9ac219df0e8813f0b4a3df2fad979b3e3607cbe6c0ce6c764.
//
// Solidity: event RegisterWallet(address indexed wallet)
func (_Walletregistry *WalletregistryFilterer) ParseRegisterWallet(log types.Log) (*WalletregistryRegisterWallet, error) {
	event := new(WalletregistryRegisterWallet)
	if err := _Walletregistry.contract.UnpackLog(event, "RegisterWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
