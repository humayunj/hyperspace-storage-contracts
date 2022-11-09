// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StorageNodeContract

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

// StorageNodeContractMetaData contains all meta data concerning the StorageNodeContract contract.
var StorageNodeContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_TLSCert\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"_HOST\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fileMerkleRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiryTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"segmentIndex\",\"type\":\"uint32\"}],\"name\":\"EvProveStorage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fileMerkleRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EvValidationExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fileMerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"segmentIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EvValidationSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"HOST\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TLSCert\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRootHash\",\"type\":\"bytes32\"}],\"name\":\"computeKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumStorageNode.CallerType\",\"name\":\"callerType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"fileSize\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"timerStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timerEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"proveTimeoutLength\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"concludeTimeoutLength\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"segmentsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"concludeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRootHash\",\"type\":\"bytes32\"}],\"name\":\"finishTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockedCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mappingLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mappingsList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"processValidation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"fileRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"segmentIndex\",\"type\":\"uint32\"}],\"name\":\"validateStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"validationExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StorageNodeContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageNodeContractMetaData.ABI instead.
var StorageNodeContractABI = StorageNodeContractMetaData.ABI

// StorageNodeContract is an auto generated Go binding around an Ethereum contract.
type StorageNodeContract struct {
	StorageNodeContractCaller     // Read-only binding to the contract
	StorageNodeContractTransactor // Write-only binding to the contract
	StorageNodeContractFilterer   // Log filterer for contract events
}

// StorageNodeContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageNodeContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageNodeContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageNodeContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageNodeContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageNodeContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageNodeContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageNodeContractSession struct {
	Contract     *StorageNodeContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StorageNodeContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageNodeContractCallerSession struct {
	Contract *StorageNodeContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// StorageNodeContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageNodeContractTransactorSession struct {
	Contract     *StorageNodeContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// StorageNodeContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageNodeContractRaw struct {
	Contract *StorageNodeContract // Generic contract binding to access the raw methods on
}

// StorageNodeContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageNodeContractCallerRaw struct {
	Contract *StorageNodeContractCaller // Generic read-only contract binding to access the raw methods on
}

// StorageNodeContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageNodeContractTransactorRaw struct {
	Contract *StorageNodeContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageNodeContract creates a new instance of StorageNodeContract, bound to a specific deployed contract.
func NewStorageNodeContract(address common.Address, backend bind.ContractBackend) (*StorageNodeContract, error) {
	contract, err := bindStorageNodeContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageNodeContract{StorageNodeContractCaller: StorageNodeContractCaller{contract: contract}, StorageNodeContractTransactor: StorageNodeContractTransactor{contract: contract}, StorageNodeContractFilterer: StorageNodeContractFilterer{contract: contract}}, nil
}

// NewStorageNodeContractCaller creates a new read-only instance of StorageNodeContract, bound to a specific deployed contract.
func NewStorageNodeContractCaller(address common.Address, caller bind.ContractCaller) (*StorageNodeContractCaller, error) {
	contract, err := bindStorageNodeContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageNodeContractCaller{contract: contract}, nil
}

// NewStorageNodeContractTransactor creates a new write-only instance of StorageNodeContract, bound to a specific deployed contract.
func NewStorageNodeContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageNodeContractTransactor, error) {
	contract, err := bindStorageNodeContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageNodeContractTransactor{contract: contract}, nil
}

// NewStorageNodeContractFilterer creates a new log filterer instance of StorageNodeContract, bound to a specific deployed contract.
func NewStorageNodeContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageNodeContractFilterer, error) {
	contract, err := bindStorageNodeContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageNodeContractFilterer{contract: contract}, nil
}

// bindStorageNodeContract binds a generic wrapper to an already deployed contract.
func bindStorageNodeContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageNodeContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageNodeContract *StorageNodeContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageNodeContract.Contract.StorageNodeContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageNodeContract *StorageNodeContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.StorageNodeContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageNodeContract *StorageNodeContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.StorageNodeContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageNodeContract *StorageNodeContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageNodeContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageNodeContract *StorageNodeContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageNodeContract *StorageNodeContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.contract.Transact(opts, method, params...)
}

// HOST is a free data retrieval call binding the contract method 0x49f289dc.
//
// Solidity: function HOST() view returns(string)
func (_StorageNodeContract *StorageNodeContractCaller) HOST(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StorageNodeContract.contract.Call(opts, &out, "HOST")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// HOST is a free data retrieval call binding the contract method 0x49f289dc.
//
// Solidity: function HOST() view returns(string)
func (_StorageNodeContract *StorageNodeContractSession) HOST() (string, error) {
	return _StorageNodeContract.Contract.HOST(&_StorageNodeContract.CallOpts)
}

// HOST is a free data retrieval call binding the contract method 0x49f289dc.
//
// Solidity: function HOST() view returns(string)
func (_StorageNodeContract *StorageNodeContractCallerSession) HOST() (string, error) {
	return _StorageNodeContract.Contract.HOST(&_StorageNodeContract.CallOpts)
}

// TLSCert is a free data retrieval call binding the contract method 0xb3752fa3.
//
// Solidity: function TLSCert() view returns(bytes)
func (_StorageNodeContract *StorageNodeContractCaller) TLSCert(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StorageNodeContract.contract.Call(opts, &out, "TLSCert")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// TLSCert is a free data retrieval call binding the contract method 0xb3752fa3.
//
// Solidity: function TLSCert() view returns(bytes)
func (_StorageNodeContract *StorageNodeContractSession) TLSCert() ([]byte, error) {
	return _StorageNodeContract.Contract.TLSCert(&_StorageNodeContract.CallOpts)
}

// TLSCert is a free data retrieval call binding the contract method 0xb3752fa3.
//
// Solidity: function TLSCert() view returns(bytes)
func (_StorageNodeContract *StorageNodeContractCallerSession) TLSCert() ([]byte, error) {
	return _StorageNodeContract.Contract.TLSCert(&_StorageNodeContract.CallOpts)
}

// ComputeKey is a free data retrieval call binding the contract method 0x0568b4d7.
//
// Solidity: function computeKey(address userAddress, bytes32 merkleRootHash) pure returns(bytes32)
func (_StorageNodeContract *StorageNodeContractCaller) ComputeKey(opts *bind.CallOpts, userAddress common.Address, merkleRootHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StorageNodeContract.contract.Call(opts, &out, "computeKey", userAddress, merkleRootHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeKey is a free data retrieval call binding the contract method 0x0568b4d7.
//
// Solidity: function computeKey(address userAddress, bytes32 merkleRootHash) pure returns(bytes32)
func (_StorageNodeContract *StorageNodeContractSession) ComputeKey(userAddress common.Address, merkleRootHash [32]byte) ([32]byte, error) {
	return _StorageNodeContract.Contract.ComputeKey(&_StorageNodeContract.CallOpts, userAddress, merkleRootHash)
}

// ComputeKey is a free data retrieval call binding the contract method 0x0568b4d7.
//
// Solidity: function computeKey(address userAddress, bytes32 merkleRootHash) pure returns(bytes32)
func (_StorageNodeContract *StorageNodeContractCallerSession) ComputeKey(userAddress common.Address, merkleRootHash [32]byte) ([32]byte, error) {
	return _StorageNodeContract.Contract.ComputeKey(&_StorageNodeContract.CallOpts, userAddress, merkleRootHash)
}

// LockedCollateral is a free data retrieval call binding the contract method 0xb952cc4a.
//
// Solidity: function lockedCollateral() view returns(uint256)
func (_StorageNodeContract *StorageNodeContractCaller) LockedCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageNodeContract.contract.Call(opts, &out, "lockedCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedCollateral is a free data retrieval call binding the contract method 0xb952cc4a.
//
// Solidity: function lockedCollateral() view returns(uint256)
func (_StorageNodeContract *StorageNodeContractSession) LockedCollateral() (*big.Int, error) {
	return _StorageNodeContract.Contract.LockedCollateral(&_StorageNodeContract.CallOpts)
}

// LockedCollateral is a free data retrieval call binding the contract method 0xb952cc4a.
//
// Solidity: function lockedCollateral() view returns(uint256)
func (_StorageNodeContract *StorageNodeContractCallerSession) LockedCollateral() (*big.Int, error) {
	return _StorageNodeContract.Contract.LockedCollateral(&_StorageNodeContract.CallOpts)
}

// MappingLength is a free data retrieval call binding the contract method 0x116766a6.
//
// Solidity: function mappingLength() view returns(uint256)
func (_StorageNodeContract *StorageNodeContractCaller) MappingLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageNodeContract.contract.Call(opts, &out, "mappingLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappingLength is a free data retrieval call binding the contract method 0x116766a6.
//
// Solidity: function mappingLength() view returns(uint256)
func (_StorageNodeContract *StorageNodeContractSession) MappingLength() (*big.Int, error) {
	return _StorageNodeContract.Contract.MappingLength(&_StorageNodeContract.CallOpts)
}

// MappingLength is a free data retrieval call binding the contract method 0x116766a6.
//
// Solidity: function mappingLength() view returns(uint256)
func (_StorageNodeContract *StorageNodeContractCallerSession) MappingLength() (*big.Int, error) {
	return _StorageNodeContract.Contract.MappingLength(&_StorageNodeContract.CallOpts)
}

// MappingsList is a free data retrieval call binding the contract method 0xca88afbc.
//
// Solidity: function mappingsList(uint256 ) view returns(bytes32)
func (_StorageNodeContract *StorageNodeContractCaller) MappingsList(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StorageNodeContract.contract.Call(opts, &out, "mappingsList", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MappingsList is a free data retrieval call binding the contract method 0xca88afbc.
//
// Solidity: function mappingsList(uint256 ) view returns(bytes32)
func (_StorageNodeContract *StorageNodeContractSession) MappingsList(arg0 *big.Int) ([32]byte, error) {
	return _StorageNodeContract.Contract.MappingsList(&_StorageNodeContract.CallOpts, arg0)
}

// MappingsList is a free data retrieval call binding the contract method 0xca88afbc.
//
// Solidity: function mappingsList(uint256 ) view returns(bytes32)
func (_StorageNodeContract *StorageNodeContractCallerSession) MappingsList(arg0 *big.Int) ([32]byte, error) {
	return _StorageNodeContract.Contract.MappingsList(&_StorageNodeContract.CallOpts, arg0)
}

// Verify is a free data retrieval call binding the contract method 0x21fb335c.
//
// Solidity: function verify(bytes32[] proof, bytes32 root, bytes32 leaf, uint256 index) pure returns(bool)
func (_StorageNodeContract *StorageNodeContractCaller) Verify(opts *bind.CallOpts, proof [][32]byte, root [32]byte, leaf [32]byte, index *big.Int) (bool, error) {
	var out []interface{}
	err := _StorageNodeContract.contract.Call(opts, &out, "verify", proof, root, leaf, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x21fb335c.
//
// Solidity: function verify(bytes32[] proof, bytes32 root, bytes32 leaf, uint256 index) pure returns(bool)
func (_StorageNodeContract *StorageNodeContractSession) Verify(proof [][32]byte, root [32]byte, leaf [32]byte, index *big.Int) (bool, error) {
	return _StorageNodeContract.Contract.Verify(&_StorageNodeContract.CallOpts, proof, root, leaf, index)
}

// Verify is a free data retrieval call binding the contract method 0x21fb335c.
//
// Solidity: function verify(bytes32[] proof, bytes32 root, bytes32 leaf, uint256 index) pure returns(bool)
func (_StorageNodeContract *StorageNodeContractCallerSession) Verify(proof [][32]byte, root [32]byte, leaf [32]byte, index *big.Int) (bool, error) {
	return _StorageNodeContract.Contract.Verify(&_StorageNodeContract.CallOpts, proof, root, leaf, index)
}

// ConcludeTransaction is a paid mutator transaction binding the contract method 0x02d12776.
//
// Solidity: function concludeTransaction(uint8 callerType, address userAddress, bytes32 merkleRootHash, uint32 fileSize, uint256 timerStart, uint256 timerEnd, uint64 proveTimeoutLength, uint64 concludeTimeoutLength, uint32 segmentsCount, uint256 bidAmount) payable returns()
func (_StorageNodeContract *StorageNodeContractTransactor) ConcludeTransaction(opts *bind.TransactOpts, callerType uint8, userAddress common.Address, merkleRootHash [32]byte, fileSize uint32, timerStart *big.Int, timerEnd *big.Int, proveTimeoutLength uint64, concludeTimeoutLength uint64, segmentsCount uint32, bidAmount *big.Int) (*types.Transaction, error) {
	return _StorageNodeContract.contract.Transact(opts, "concludeTransaction", callerType, userAddress, merkleRootHash, fileSize, timerStart, timerEnd, proveTimeoutLength, concludeTimeoutLength, segmentsCount, bidAmount)
}

// ConcludeTransaction is a paid mutator transaction binding the contract method 0x02d12776.
//
// Solidity: function concludeTransaction(uint8 callerType, address userAddress, bytes32 merkleRootHash, uint32 fileSize, uint256 timerStart, uint256 timerEnd, uint64 proveTimeoutLength, uint64 concludeTimeoutLength, uint32 segmentsCount, uint256 bidAmount) payable returns()
func (_StorageNodeContract *StorageNodeContractSession) ConcludeTransaction(callerType uint8, userAddress common.Address, merkleRootHash [32]byte, fileSize uint32, timerStart *big.Int, timerEnd *big.Int, proveTimeoutLength uint64, concludeTimeoutLength uint64, segmentsCount uint32, bidAmount *big.Int) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.ConcludeTransaction(&_StorageNodeContract.TransactOpts, callerType, userAddress, merkleRootHash, fileSize, timerStart, timerEnd, proveTimeoutLength, concludeTimeoutLength, segmentsCount, bidAmount)
}

// ConcludeTransaction is a paid mutator transaction binding the contract method 0x02d12776.
//
// Solidity: function concludeTransaction(uint8 callerType, address userAddress, bytes32 merkleRootHash, uint32 fileSize, uint256 timerStart, uint256 timerEnd, uint64 proveTimeoutLength, uint64 concludeTimeoutLength, uint32 segmentsCount, uint256 bidAmount) payable returns()
func (_StorageNodeContract *StorageNodeContractTransactorSession) ConcludeTransaction(callerType uint8, userAddress common.Address, merkleRootHash [32]byte, fileSize uint32, timerStart *big.Int, timerEnd *big.Int, proveTimeoutLength uint64, concludeTimeoutLength uint64, segmentsCount uint32, bidAmount *big.Int) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.ConcludeTransaction(&_StorageNodeContract.TransactOpts, callerType, userAddress, merkleRootHash, fileSize, timerStart, timerEnd, proveTimeoutLength, concludeTimeoutLength, segmentsCount, bidAmount)
}

// FinishTransaction is a paid mutator transaction binding the contract method 0x438eb46e.
//
// Solidity: function finishTransaction(address userAddress, bytes32 merkleRootHash) returns()
func (_StorageNodeContract *StorageNodeContractTransactor) FinishTransaction(opts *bind.TransactOpts, userAddress common.Address, merkleRootHash [32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.contract.Transact(opts, "finishTransaction", userAddress, merkleRootHash)
}

// FinishTransaction is a paid mutator transaction binding the contract method 0x438eb46e.
//
// Solidity: function finishTransaction(address userAddress, bytes32 merkleRootHash) returns()
func (_StorageNodeContract *StorageNodeContractSession) FinishTransaction(userAddress common.Address, merkleRootHash [32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.FinishTransaction(&_StorageNodeContract.TransactOpts, userAddress, merkleRootHash)
}

// FinishTransaction is a paid mutator transaction binding the contract method 0x438eb46e.
//
// Solidity: function finishTransaction(address userAddress, bytes32 merkleRootHash) returns()
func (_StorageNodeContract *StorageNodeContractTransactorSession) FinishTransaction(userAddress common.Address, merkleRootHash [32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.FinishTransaction(&_StorageNodeContract.TransactOpts, userAddress, merkleRootHash)
}

// ProcessValidation is a paid mutator transaction binding the contract method 0xe172c5cb.
//
// Solidity: function processValidation(address userAddress, bytes32 rootHash, bytes data, bytes32[] proof) returns(bool)
func (_StorageNodeContract *StorageNodeContractTransactor) ProcessValidation(opts *bind.TransactOpts, userAddress common.Address, rootHash [32]byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.contract.Transact(opts, "processValidation", userAddress, rootHash, data, proof)
}

// ProcessValidation is a paid mutator transaction binding the contract method 0xe172c5cb.
//
// Solidity: function processValidation(address userAddress, bytes32 rootHash, bytes data, bytes32[] proof) returns(bool)
func (_StorageNodeContract *StorageNodeContractSession) ProcessValidation(userAddress common.Address, rootHash [32]byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.ProcessValidation(&_StorageNodeContract.TransactOpts, userAddress, rootHash, data, proof)
}

// ProcessValidation is a paid mutator transaction binding the contract method 0xe172c5cb.
//
// Solidity: function processValidation(address userAddress, bytes32 rootHash, bytes data, bytes32[] proof) returns(bool)
func (_StorageNodeContract *StorageNodeContractTransactorSession) ProcessValidation(userAddress common.Address, rootHash [32]byte, data []byte, proof [][32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.ProcessValidation(&_StorageNodeContract.TransactOpts, userAddress, rootHash, data, proof)
}

// ValidateStorage is a paid mutator transaction binding the contract method 0x5d206fda.
//
// Solidity: function validateStorage(address userAddress, bytes32 fileRootHash, uint32 segmentIndex) returns()
func (_StorageNodeContract *StorageNodeContractTransactor) ValidateStorage(opts *bind.TransactOpts, userAddress common.Address, fileRootHash [32]byte, segmentIndex uint32) (*types.Transaction, error) {
	return _StorageNodeContract.contract.Transact(opts, "validateStorage", userAddress, fileRootHash, segmentIndex)
}

// ValidateStorage is a paid mutator transaction binding the contract method 0x5d206fda.
//
// Solidity: function validateStorage(address userAddress, bytes32 fileRootHash, uint32 segmentIndex) returns()
func (_StorageNodeContract *StorageNodeContractSession) ValidateStorage(userAddress common.Address, fileRootHash [32]byte, segmentIndex uint32) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.ValidateStorage(&_StorageNodeContract.TransactOpts, userAddress, fileRootHash, segmentIndex)
}

// ValidateStorage is a paid mutator transaction binding the contract method 0x5d206fda.
//
// Solidity: function validateStorage(address userAddress, bytes32 fileRootHash, uint32 segmentIndex) returns()
func (_StorageNodeContract *StorageNodeContractTransactorSession) ValidateStorage(userAddress common.Address, fileRootHash [32]byte, segmentIndex uint32) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.ValidateStorage(&_StorageNodeContract.TransactOpts, userAddress, fileRootHash, segmentIndex)
}

// ValidationExpired is a paid mutator transaction binding the contract method 0x22f44022.
//
// Solidity: function validationExpired(address userAddress, bytes32 rootHash) returns()
func (_StorageNodeContract *StorageNodeContractTransactor) ValidationExpired(opts *bind.TransactOpts, userAddress common.Address, rootHash [32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.contract.Transact(opts, "validationExpired", userAddress, rootHash)
}

// ValidationExpired is a paid mutator transaction binding the contract method 0x22f44022.
//
// Solidity: function validationExpired(address userAddress, bytes32 rootHash) returns()
func (_StorageNodeContract *StorageNodeContractSession) ValidationExpired(userAddress common.Address, rootHash [32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.ValidationExpired(&_StorageNodeContract.TransactOpts, userAddress, rootHash)
}

// ValidationExpired is a paid mutator transaction binding the contract method 0x22f44022.
//
// Solidity: function validationExpired(address userAddress, bytes32 rootHash) returns()
func (_StorageNodeContract *StorageNodeContractTransactorSession) ValidationExpired(userAddress common.Address, rootHash [32]byte) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.ValidationExpired(&_StorageNodeContract.TransactOpts, userAddress, rootHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address target) returns()
func (_StorageNodeContract *StorageNodeContractTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _StorageNodeContract.contract.Transact(opts, "withdraw", amount, target)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address target) returns()
func (_StorageNodeContract *StorageNodeContractSession) Withdraw(amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.Withdraw(&_StorageNodeContract.TransactOpts, amount, target)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 amount, address target) returns()
func (_StorageNodeContract *StorageNodeContractTransactorSession) Withdraw(amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _StorageNodeContract.Contract.Withdraw(&_StorageNodeContract.TransactOpts, amount, target)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StorageNodeContract *StorageNodeContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageNodeContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StorageNodeContract *StorageNodeContractSession) Receive() (*types.Transaction, error) {
	return _StorageNodeContract.Contract.Receive(&_StorageNodeContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StorageNodeContract *StorageNodeContractTransactorSession) Receive() (*types.Transaction, error) {
	return _StorageNodeContract.Contract.Receive(&_StorageNodeContract.TransactOpts)
}

// StorageNodeContractEvProveStorageIterator is returned from FilterEvProveStorage and is used to iterate over the raw logs and unpacked data for EvProveStorage events raised by the StorageNodeContract contract.
type StorageNodeContractEvProveStorageIterator struct {
	Event *StorageNodeContractEvProveStorage // Event containing the contract specifics and raw log

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
func (it *StorageNodeContractEvProveStorageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageNodeContractEvProveStorage)
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
		it.Event = new(StorageNodeContractEvProveStorage)
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
func (it *StorageNodeContractEvProveStorageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageNodeContractEvProveStorageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageNodeContractEvProveStorage represents a EvProveStorage event raised by the StorageNodeContract contract.
type StorageNodeContractEvProveStorage struct {
	UserAddress        common.Address
	FileMerkleRootHash [32]byte
	Timestamp          *big.Int
	ExpiryTimestamp    *big.Int
	SegmentIndex       uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEvProveStorage is a free log retrieval operation binding the contract event 0x083b6024bc1dcbf2784878665777050c94a4eaceac2f9cb0688b8a93d25f49b3.
//
// Solidity: event EvProveStorage(address userAddress, bytes32 fileMerkleRootHash, uint256 timestamp, uint256 expiryTimestamp, uint32 segmentIndex)
func (_StorageNodeContract *StorageNodeContractFilterer) FilterEvProveStorage(opts *bind.FilterOpts) (*StorageNodeContractEvProveStorageIterator, error) {

	logs, sub, err := _StorageNodeContract.contract.FilterLogs(opts, "EvProveStorage")
	if err != nil {
		return nil, err
	}
	return &StorageNodeContractEvProveStorageIterator{contract: _StorageNodeContract.contract, event: "EvProveStorage", logs: logs, sub: sub}, nil
}

// WatchEvProveStorage is a free log subscription operation binding the contract event 0x083b6024bc1dcbf2784878665777050c94a4eaceac2f9cb0688b8a93d25f49b3.
//
// Solidity: event EvProveStorage(address userAddress, bytes32 fileMerkleRootHash, uint256 timestamp, uint256 expiryTimestamp, uint32 segmentIndex)
func (_StorageNodeContract *StorageNodeContractFilterer) WatchEvProveStorage(opts *bind.WatchOpts, sink chan<- *StorageNodeContractEvProveStorage) (event.Subscription, error) {

	logs, sub, err := _StorageNodeContract.contract.WatchLogs(opts, "EvProveStorage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageNodeContractEvProveStorage)
				if err := _StorageNodeContract.contract.UnpackLog(event, "EvProveStorage", log); err != nil {
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

// ParseEvProveStorage is a log parse operation binding the contract event 0x083b6024bc1dcbf2784878665777050c94a4eaceac2f9cb0688b8a93d25f49b3.
//
// Solidity: event EvProveStorage(address userAddress, bytes32 fileMerkleRootHash, uint256 timestamp, uint256 expiryTimestamp, uint32 segmentIndex)
func (_StorageNodeContract *StorageNodeContractFilterer) ParseEvProveStorage(log types.Log) (*StorageNodeContractEvProveStorage, error) {
	event := new(StorageNodeContractEvProveStorage)
	if err := _StorageNodeContract.contract.UnpackLog(event, "EvProveStorage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageNodeContractEvValidationExpiredIterator is returned from FilterEvValidationExpired and is used to iterate over the raw logs and unpacked data for EvValidationExpired events raised by the StorageNodeContract contract.
type StorageNodeContractEvValidationExpiredIterator struct {
	Event *StorageNodeContractEvValidationExpired // Event containing the contract specifics and raw log

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
func (it *StorageNodeContractEvValidationExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageNodeContractEvValidationExpired)
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
		it.Event = new(StorageNodeContractEvValidationExpired)
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
func (it *StorageNodeContractEvValidationExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageNodeContractEvValidationExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageNodeContractEvValidationExpired represents a EvValidationExpired event raised by the StorageNodeContract contract.
type StorageNodeContractEvValidationExpired struct {
	UserAddress        common.Address
	FileMerkleRootHash [32]byte
	Timestamp          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterEvValidationExpired is a free log retrieval operation binding the contract event 0x22cd95a98f968866886bf2e9f6c8af31f17c78f37500bf0b0ca6725de29c7abb.
//
// Solidity: event EvValidationExpired(address userAddress, bytes32 fileMerkleRootHash, uint256 timestamp)
func (_StorageNodeContract *StorageNodeContractFilterer) FilterEvValidationExpired(opts *bind.FilterOpts) (*StorageNodeContractEvValidationExpiredIterator, error) {

	logs, sub, err := _StorageNodeContract.contract.FilterLogs(opts, "EvValidationExpired")
	if err != nil {
		return nil, err
	}
	return &StorageNodeContractEvValidationExpiredIterator{contract: _StorageNodeContract.contract, event: "EvValidationExpired", logs: logs, sub: sub}, nil
}

// WatchEvValidationExpired is a free log subscription operation binding the contract event 0x22cd95a98f968866886bf2e9f6c8af31f17c78f37500bf0b0ca6725de29c7abb.
//
// Solidity: event EvValidationExpired(address userAddress, bytes32 fileMerkleRootHash, uint256 timestamp)
func (_StorageNodeContract *StorageNodeContractFilterer) WatchEvValidationExpired(opts *bind.WatchOpts, sink chan<- *StorageNodeContractEvValidationExpired) (event.Subscription, error) {

	logs, sub, err := _StorageNodeContract.contract.WatchLogs(opts, "EvValidationExpired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageNodeContractEvValidationExpired)
				if err := _StorageNodeContract.contract.UnpackLog(event, "EvValidationExpired", log); err != nil {
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

// ParseEvValidationExpired is a log parse operation binding the contract event 0x22cd95a98f968866886bf2e9f6c8af31f17c78f37500bf0b0ca6725de29c7abb.
//
// Solidity: event EvValidationExpired(address userAddress, bytes32 fileMerkleRootHash, uint256 timestamp)
func (_StorageNodeContract *StorageNodeContractFilterer) ParseEvValidationExpired(log types.Log) (*StorageNodeContractEvValidationExpired, error) {
	event := new(StorageNodeContractEvValidationExpired)
	if err := _StorageNodeContract.contract.UnpackLog(event, "EvValidationExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageNodeContractEvValidationSubmittedIterator is returned from FilterEvValidationSubmitted and is used to iterate over the raw logs and unpacked data for EvValidationSubmitted events raised by the StorageNodeContract contract.
type StorageNodeContractEvValidationSubmittedIterator struct {
	Event *StorageNodeContractEvValidationSubmitted // Event containing the contract specifics and raw log

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
func (it *StorageNodeContractEvValidationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageNodeContractEvValidationSubmitted)
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
		it.Event = new(StorageNodeContractEvValidationSubmitted)
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
func (it *StorageNodeContractEvValidationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageNodeContractEvValidationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageNodeContractEvValidationSubmitted represents a EvValidationSubmitted event raised by the StorageNodeContract contract.
type StorageNodeContractEvValidationSubmitted struct {
	UserAddress    common.Address
	FileMerkleRoot [32]byte
	SegmentIndex   uint32
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEvValidationSubmitted is a free log retrieval operation binding the contract event 0x125690aeed72600ce0497b2ecc092d1b1b925f244d54906a195066af49c8b0b7.
//
// Solidity: event EvValidationSubmitted(address userAddress, bytes32 fileMerkleRoot, uint32 segmentIndex, uint256 timestamp)
func (_StorageNodeContract *StorageNodeContractFilterer) FilterEvValidationSubmitted(opts *bind.FilterOpts) (*StorageNodeContractEvValidationSubmittedIterator, error) {

	logs, sub, err := _StorageNodeContract.contract.FilterLogs(opts, "EvValidationSubmitted")
	if err != nil {
		return nil, err
	}
	return &StorageNodeContractEvValidationSubmittedIterator{contract: _StorageNodeContract.contract, event: "EvValidationSubmitted", logs: logs, sub: sub}, nil
}

// WatchEvValidationSubmitted is a free log subscription operation binding the contract event 0x125690aeed72600ce0497b2ecc092d1b1b925f244d54906a195066af49c8b0b7.
//
// Solidity: event EvValidationSubmitted(address userAddress, bytes32 fileMerkleRoot, uint32 segmentIndex, uint256 timestamp)
func (_StorageNodeContract *StorageNodeContractFilterer) WatchEvValidationSubmitted(opts *bind.WatchOpts, sink chan<- *StorageNodeContractEvValidationSubmitted) (event.Subscription, error) {

	logs, sub, err := _StorageNodeContract.contract.WatchLogs(opts, "EvValidationSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageNodeContractEvValidationSubmitted)
				if err := _StorageNodeContract.contract.UnpackLog(event, "EvValidationSubmitted", log); err != nil {
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

// ParseEvValidationSubmitted is a log parse operation binding the contract event 0x125690aeed72600ce0497b2ecc092d1b1b925f244d54906a195066af49c8b0b7.
//
// Solidity: event EvValidationSubmitted(address userAddress, bytes32 fileMerkleRoot, uint32 segmentIndex, uint256 timestamp)
func (_StorageNodeContract *StorageNodeContractFilterer) ParseEvValidationSubmitted(log types.Log) (*StorageNodeContractEvValidationSubmitted, error) {
	event := new(StorageNodeContractEvValidationSubmitted)
	if err := _StorageNodeContract.contract.UnpackLog(event, "EvValidationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
