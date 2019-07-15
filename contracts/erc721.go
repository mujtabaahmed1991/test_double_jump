// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DJTERC721ImplementationABI is the input ABI used to generate the binding from.
const DJTERC721ImplementationABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getToken\",\"outputs\":[{\"name\":\"mintedBy\",\"type\":\"address\"},{\"name\":\"mintedAt\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// DJTERC721Implementation is an auto generated Go binding around an Ethereum contract.
type DJTERC721Implementation struct {
	DJTERC721ImplementationCaller     // Read-only binding to the contract
	DJTERC721ImplementationTransactor // Write-only binding to the contract
	DJTERC721ImplementationFilterer   // Log filterer for contract events
}

// DJTERC721ImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type DJTERC721ImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DJTERC721ImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DJTERC721ImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DJTERC721ImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DJTERC721ImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DJTERC721ImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DJTERC721ImplementationSession struct {
	Contract     *DJTERC721Implementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DJTERC721ImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DJTERC721ImplementationCallerSession struct {
	Contract *DJTERC721ImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// DJTERC721ImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DJTERC721ImplementationTransactorSession struct {
	Contract     *DJTERC721ImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// DJTERC721ImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type DJTERC721ImplementationRaw struct {
	Contract *DJTERC721Implementation // Generic contract binding to access the raw methods on
}

// DJTERC721ImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DJTERC721ImplementationCallerRaw struct {
	Contract *DJTERC721ImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// DJTERC721ImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DJTERC721ImplementationTransactorRaw struct {
	Contract *DJTERC721ImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDJTERC721Implementation creates a new instance of DJTERC721Implementation, bound to a specific deployed contract.
func NewDJTERC721Implementation(address common.Address, backend bind.ContractBackend) (*DJTERC721Implementation, error) {
	contract, err := bindDJTERC721Implementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DJTERC721Implementation{DJTERC721ImplementationCaller: DJTERC721ImplementationCaller{contract: contract}, DJTERC721ImplementationTransactor: DJTERC721ImplementationTransactor{contract: contract}, DJTERC721ImplementationFilterer: DJTERC721ImplementationFilterer{contract: contract}}, nil
}

// NewDJTERC721ImplementationCaller creates a new read-only instance of DJTERC721Implementation, bound to a specific deployed contract.
func NewDJTERC721ImplementationCaller(address common.Address, caller bind.ContractCaller) (*DJTERC721ImplementationCaller, error) {
	contract, err := bindDJTERC721Implementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DJTERC721ImplementationCaller{contract: contract}, nil
}

// NewDJTERC721ImplementationTransactor creates a new write-only instance of DJTERC721Implementation, bound to a specific deployed contract.
func NewDJTERC721ImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*DJTERC721ImplementationTransactor, error) {
	contract, err := bindDJTERC721Implementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DJTERC721ImplementationTransactor{contract: contract}, nil
}

// NewDJTERC721ImplementationFilterer creates a new log filterer instance of DJTERC721Implementation, bound to a specific deployed contract.
func NewDJTERC721ImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*DJTERC721ImplementationFilterer, error) {
	contract, err := bindDJTERC721Implementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DJTERC721ImplementationFilterer{contract: contract}, nil
}

// bindDJTERC721Implementation binds a generic wrapper to an already deployed contract.
func bindDJTERC721Implementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DJTERC721ImplementationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DJTERC721Implementation *DJTERC721ImplementationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DJTERC721Implementation.Contract.DJTERC721ImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DJTERC721Implementation *DJTERC721ImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.DJTERC721ImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DJTERC721Implementation *DJTERC721ImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.DJTERC721ImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DJTERC721Implementation *DJTERC721ImplementationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DJTERC721Implementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DJTERC721Implementation *DJTERC721ImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DJTERC721Implementation *DJTERC721ImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DJTERC721Implementation.Contract.BalanceOf(&_DJTERC721Implementation.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) constant returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _DJTERC721Implementation.Contract.BalanceOf(&_DJTERC721Implementation.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _DJTERC721Implementation.Contract.GetApproved(&_DJTERC721Implementation.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _DJTERC721Implementation.Contract.GetApproved(&_DJTERC721Implementation.CallOpts, _tokenId)
}

// GetToken is a free data retrieval call binding the contract method 0xe4b50cb8.
//
// Solidity: function getToken(uint256 _tokenId) constant returns(address mintedBy, uint64 mintedAt)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) GetToken(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	MintedBy common.Address
	MintedAt uint64
}, error) {
	ret := new(struct {
		MintedBy common.Address
		MintedAt uint64
	})
	out := ret
	err := _DJTERC721Implementation.contract.Call(opts, out, "getToken", _tokenId)
	return *ret, err
}

// GetToken is a free data retrieval call binding the contract method 0xe4b50cb8.
//
// Solidity: function getToken(uint256 _tokenId) constant returns(address mintedBy, uint64 mintedAt)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) GetToken(_tokenId *big.Int) (struct {
	MintedBy common.Address
	MintedAt uint64
}, error) {
	return _DJTERC721Implementation.Contract.GetToken(&_DJTERC721Implementation.CallOpts, _tokenId)
}

// GetToken is a free data retrieval call binding the contract method 0xe4b50cb8.
//
// Solidity: function getToken(uint256 _tokenId) constant returns(address mintedBy, uint64 mintedAt)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) GetToken(_tokenId *big.Int) (struct {
	MintedBy common.Address
	MintedAt uint64
}, error) {
	return _DJTERC721Implementation.Contract.GetToken(&_DJTERC721Implementation.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _DJTERC721Implementation.Contract.IsApprovedForAll(&_DJTERC721Implementation.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) constant returns(bool)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _DJTERC721Implementation.Contract.IsApprovedForAll(&_DJTERC721Implementation.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) Name() (string, error) {
	return _DJTERC721Implementation.Contract.Name(&_DJTERC721Implementation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) Name() (string, error) {
	return _DJTERC721Implementation.Contract.Name(&_DJTERC721Implementation.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address owner)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address owner)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _DJTERC721Implementation.Contract.OwnerOf(&_DJTERC721Implementation.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) constant returns(address owner)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _DJTERC721Implementation.Contract.OwnerOf(&_DJTERC721Implementation.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "supportsInterface", _interfaceID)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _DJTERC721Implementation.Contract.SupportsInterface(&_DJTERC721Implementation.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) constant returns(bool)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _DJTERC721Implementation.Contract.SupportsInterface(&_DJTERC721Implementation.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) Symbol() (string, error) {
	return _DJTERC721Implementation.Contract.Symbol(&_DJTERC721Implementation.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) Symbol() (string, error) {
	return _DJTERC721Implementation.Contract.Symbol(&_DJTERC721Implementation.CallOpts)
}

// TokenIndexToApproved is a free data retrieval call binding the contract method 0xa8bd9c32.
//
// Solidity: function tokenIndexToApproved(uint256 ) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) TokenIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "tokenIndexToApproved", arg0)
	return *ret0, err
}

// TokenIndexToApproved is a free data retrieval call binding the contract method 0xa8bd9c32.
//
// Solidity: function tokenIndexToApproved(uint256 ) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) TokenIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _DJTERC721Implementation.Contract.TokenIndexToApproved(&_DJTERC721Implementation.CallOpts, arg0)
}

// TokenIndexToApproved is a free data retrieval call binding the contract method 0xa8bd9c32.
//
// Solidity: function tokenIndexToApproved(uint256 ) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) TokenIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _DJTERC721Implementation.Contract.TokenIndexToApproved(&_DJTERC721Implementation.CallOpts, arg0)
}

// TokenIndexToOwner is a free data retrieval call binding the contract method 0x1d36e06c.
//
// Solidity: function tokenIndexToOwner(uint256 ) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) TokenIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "tokenIndexToOwner", arg0)
	return *ret0, err
}

// TokenIndexToOwner is a free data retrieval call binding the contract method 0x1d36e06c.
//
// Solidity: function tokenIndexToOwner(uint256 ) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) TokenIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _DJTERC721Implementation.Contract.TokenIndexToOwner(&_DJTERC721Implementation.CallOpts, arg0)
}

// TokenIndexToOwner is a free data retrieval call binding the contract method 0x1d36e06c.
//
// Solidity: function tokenIndexToOwner(uint256 ) constant returns(address)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) TokenIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _DJTERC721Implementation.Contract.TokenIndexToOwner(&_DJTERC721Implementation.CallOpts, arg0)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) constant returns(uint256[] result)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "tokensOfOwner", _owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) constant returns(uint256[] result)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _DJTERC721Implementation.Contract.TokensOfOwner(&_DJTERC721Implementation.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) constant returns(uint256[] result)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _DJTERC721Implementation.Contract.TokensOfOwner(&_DJTERC721Implementation.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DJTERC721Implementation.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) TotalSupply() (*big.Int, error) {
	return _DJTERC721Implementation.Contract.TotalSupply(&_DJTERC721Implementation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationCallerSession) TotalSupply() (*big.Int, error) {
	return _DJTERC721Implementation.Contract.TotalSupply(&_DJTERC721Implementation.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.Approve(&_DJTERC721Implementation.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.Approve(&_DJTERC721Implementation.TransactOpts, _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DJTERC721Implementation.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationSession) Mint() (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.Mint(&_DJTERC721Implementation.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_DJTERC721Implementation *DJTERC721ImplementationTransactorSession) Mint() (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.Mint(&_DJTERC721Implementation.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.SafeTransferFrom(&_DJTERC721Implementation.TransactOpts, _from, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.SafeTransferFrom(&_DJTERC721Implementation.TransactOpts, _from, _to, _tokenId)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _DJTERC721Implementation.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.SetApprovalForAll(&_DJTERC721Implementation.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.SetApprovalForAll(&_DJTERC721Implementation.TransactOpts, _operator, _approved)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.Transfer(&_DJTERC721Implementation.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.Transfer(&_DJTERC721Implementation.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.TransferFrom(&_DJTERC721Implementation.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_DJTERC721Implementation *DJTERC721ImplementationTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _DJTERC721Implementation.Contract.TransferFrom(&_DJTERC721Implementation.TransactOpts, _from, _to, _tokenId)
}

// DJTERC721ImplementationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DJTERC721Implementation contract.
type DJTERC721ImplementationApprovalIterator struct {
	Event *DJTERC721ImplementationApproval // Event containing the contract specifics and raw log

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
func (it *DJTERC721ImplementationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DJTERC721ImplementationApproval)
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
		it.Event = new(DJTERC721ImplementationApproval)
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
func (it *DJTERC721ImplementationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DJTERC721ImplementationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DJTERC721ImplementationApproval represents a Approval event raised by the DJTERC721Implementation contract.
type DJTERC721ImplementationApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*DJTERC721ImplementationApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _DJTERC721Implementation.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DJTERC721ImplementationApprovalIterator{contract: _DJTERC721Implementation.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DJTERC721ImplementationApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _DJTERC721Implementation.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DJTERC721ImplementationApproval)
				if err := _DJTERC721Implementation.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 indexed _tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) ParseApproval(log types.Log) (*DJTERC721ImplementationApproval, error) {
	event := new(DJTERC721ImplementationApproval)
	if err := _DJTERC721Implementation.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DJTERC721ImplementationApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DJTERC721Implementation contract.
type DJTERC721ImplementationApprovalForAllIterator struct {
	Event *DJTERC721ImplementationApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DJTERC721ImplementationApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DJTERC721ImplementationApprovalForAll)
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
		it.Event = new(DJTERC721ImplementationApprovalForAll)
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
func (it *DJTERC721ImplementationApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DJTERC721ImplementationApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DJTERC721ImplementationApprovalForAll represents a ApprovalForAll event raised by the DJTERC721Implementation contract.
type DJTERC721ImplementationApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*DJTERC721ImplementationApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _DJTERC721Implementation.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &DJTERC721ImplementationApprovalForAllIterator{contract: _DJTERC721Implementation.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DJTERC721ImplementationApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _DJTERC721Implementation.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DJTERC721ImplementationApprovalForAll)
				if err := _DJTERC721Implementation.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed _owner, address indexed _operator, bool _approved)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) ParseApprovalForAll(log types.Log) (*DJTERC721ImplementationApprovalForAll, error) {
	event := new(DJTERC721ImplementationApprovalForAll)
	if err := _DJTERC721Implementation.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DJTERC721ImplementationMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the DJTERC721Implementation contract.
type DJTERC721ImplementationMintIterator struct {
	Event *DJTERC721ImplementationMint // Event containing the contract specifics and raw log

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
func (it *DJTERC721ImplementationMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DJTERC721ImplementationMint)
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
		it.Event = new(DJTERC721ImplementationMint)
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
func (it *DJTERC721ImplementationMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DJTERC721ImplementationMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DJTERC721ImplementationMint represents a Mint event raised by the DJTERC721Implementation contract.
type DJTERC721ImplementationMint struct {
	Owner   common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address owner, uint256 tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) FilterMint(opts *bind.FilterOpts) (*DJTERC721ImplementationMintIterator, error) {

	logs, sub, err := _DJTERC721Implementation.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &DJTERC721ImplementationMintIterator{contract: _DJTERC721Implementation.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address owner, uint256 tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *DJTERC721ImplementationMint) (event.Subscription, error) {

	logs, sub, err := _DJTERC721Implementation.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DJTERC721ImplementationMint)
				if err := _DJTERC721Implementation.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address owner, uint256 tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) ParseMint(log types.Log) (*DJTERC721ImplementationMint, error) {
	event := new(DJTERC721ImplementationMint)
	if err := _DJTERC721Implementation.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DJTERC721ImplementationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DJTERC721Implementation contract.
type DJTERC721ImplementationTransferIterator struct {
	Event *DJTERC721ImplementationTransfer // Event containing the contract specifics and raw log

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
func (it *DJTERC721ImplementationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DJTERC721ImplementationTransfer)
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
		it.Event = new(DJTERC721ImplementationTransfer)
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
func (it *DJTERC721ImplementationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DJTERC721ImplementationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DJTERC721ImplementationTransfer represents a Transfer event raised by the DJTERC721Implementation contract.
type DJTERC721ImplementationTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*DJTERC721ImplementationTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _DJTERC721Implementation.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DJTERC721ImplementationTransferIterator{contract: _DJTERC721Implementation.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DJTERC721ImplementationTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _DJTERC721Implementation.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DJTERC721ImplementationTransfer)
				if err := _DJTERC721Implementation.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 indexed _tokenId)
func (_DJTERC721Implementation *DJTERC721ImplementationFilterer) ParseTransfer(log types.Log) (*DJTERC721ImplementationTransfer, error) {
	event := new(DJTERC721ImplementationTransfer)
	if err := _DJTERC721Implementation.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
