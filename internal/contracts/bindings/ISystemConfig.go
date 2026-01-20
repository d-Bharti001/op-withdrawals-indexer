// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// ISystemConfigAddresses is an auto generated low-level Go binding around an user-defined struct.
type ISystemConfigAddresses struct {
	L1CrossDomainMessenger       common.Address
	L1ERC721Bridge               common.Address
	L1StandardBridge             common.Address
	OptimismPortal               common.Address
	OptimismMintableERC20Factory common.Address
}

// ISystemConfigMetaData contains all meta data concerning the ISystemConfig contract.
var ISystemConfigMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumISystemConfig.UpdateType\",\"name\":\"updateType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ConfigUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feature\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"FeatureSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchInbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batcherHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeGameFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"l1CrossDomainMessenger\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1ERC721Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1StandardBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"optimismPortal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"optimismMintableERC20Factory\",\"type\":\"address\"}],\"internalType\":\"structISystemConfig.Addresses\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isFeatureEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1CrossDomainMessenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1ERC721Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1StandardBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2ChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimismMintableERC20Factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"optimismPortal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	ID:  "ISystemConfig",
}

// ISystemConfig is an auto generated Go binding around an Ethereum contract.
type ISystemConfig struct {
	abi abi.ABI
}

// NewISystemConfig creates a new instance of ISystemConfig.
func NewISystemConfig() *ISystemConfig {
	parsed, err := ISystemConfigMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ISystemConfig{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ISystemConfig) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xffa1ad74.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function VERSION() view returns(uint256)
func (iSystemConfig *ISystemConfig) PackVERSION() []byte {
	enc, err := iSystemConfig.abi.Pack("VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xffa1ad74.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function VERSION() view returns(uint256)
func (iSystemConfig *ISystemConfig) TryPackVERSION() ([]byte, error) {
	return iSystemConfig.abi.Pack("VERSION")
}

// UnpackVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (iSystemConfig *ISystemConfig) UnpackVERSION(data []byte) (*big.Int, error) {
	out, err := iSystemConfig.abi.Unpack("VERSION", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBatchInbox is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdac6e63a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function batchInbox() view returns(address addr_)
func (iSystemConfig *ISystemConfig) PackBatchInbox() []byte {
	enc, err := iSystemConfig.abi.Pack("batchInbox")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBatchInbox is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdac6e63a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function batchInbox() view returns(address addr_)
func (iSystemConfig *ISystemConfig) TryPackBatchInbox() ([]byte, error) {
	return iSystemConfig.abi.Pack("batchInbox")
}

// UnpackBatchInbox is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdac6e63a.
//
// Solidity: function batchInbox() view returns(address addr_)
func (iSystemConfig *ISystemConfig) UnpackBatchInbox(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("batchInbox", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackBatcherHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe81b2c6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function batcherHash() view returns(bytes32)
func (iSystemConfig *ISystemConfig) PackBatcherHash() []byte {
	enc, err := iSystemConfig.abi.Pack("batcherHash")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBatcherHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe81b2c6d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function batcherHash() view returns(bytes32)
func (iSystemConfig *ISystemConfig) TryPackBatcherHash() ([]byte, error) {
	return iSystemConfig.abi.Pack("batcherHash")
}

// UnpackBatcherHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (iSystemConfig *ISystemConfig) UnpackBatcherHash(data []byte) ([32]byte, error) {
	out, err := iSystemConfig.abi.Unpack("batcherHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackDisputeGameFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2b4e617.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disputeGameFactory() view returns(address addr_)
func (iSystemConfig *ISystemConfig) PackDisputeGameFactory() []byte {
	enc, err := iSystemConfig.abi.Pack("disputeGameFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisputeGameFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2b4e617.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disputeGameFactory() view returns(address addr_)
func (iSystemConfig *ISystemConfig) TryPackDisputeGameFactory() ([]byte, error) {
	return iSystemConfig.abi.Pack("disputeGameFactory")
}

// UnpackDisputeGameFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address addr_)
func (iSystemConfig *ISystemConfig) UnpackDisputeGameFactory(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("disputeGameFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa39fac12.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddresses() view returns((address,address,address,address,address))
func (iSystemConfig *ISystemConfig) PackGetAddresses() []byte {
	enc, err := iSystemConfig.abi.Pack("getAddresses")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa39fac12.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAddresses() view returns((address,address,address,address,address))
func (iSystemConfig *ISystemConfig) TryPackGetAddresses() ([]byte, error) {
	return iSystemConfig.abi.Pack("getAddresses")
}

// UnpackGetAddresses is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa39fac12.
//
// Solidity: function getAddresses() view returns((address,address,address,address,address))
func (iSystemConfig *ISystemConfig) UnpackGetAddresses(data []byte) (ISystemConfigAddresses, error) {
	out, err := iSystemConfig.abi.Unpack("getAddresses", data)
	if err != nil {
		return *new(ISystemConfigAddresses), err
	}
	out0 := *abi.ConvertType(out[0], new(ISystemConfigAddresses)).(*ISystemConfigAddresses)
	return out0, nil
}

// PackGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x452a9320.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function guardian() view returns(address)
func (iSystemConfig *ISystemConfig) PackGuardian() []byte {
	enc, err := iSystemConfig.abi.Pack("guardian")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x452a9320.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function guardian() view returns(address)
func (iSystemConfig *ISystemConfig) TryPackGuardian() ([]byte, error) {
	return iSystemConfig.abi.Pack("guardian")
}

// UnpackGuardian is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (iSystemConfig *ISystemConfig) UnpackGuardian(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("guardian", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsFeatureEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47af267b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isFeatureEnabled(bytes32 ) view returns(bool)
func (iSystemConfig *ISystemConfig) PackIsFeatureEnabled(arg0 [32]byte) []byte {
	enc, err := iSystemConfig.abi.Pack("isFeatureEnabled", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsFeatureEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x47af267b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isFeatureEnabled(bytes32 ) view returns(bool)
func (iSystemConfig *ISystemConfig) TryPackIsFeatureEnabled(arg0 [32]byte) ([]byte, error) {
	return iSystemConfig.abi.Pack("isFeatureEnabled", arg0)
}

// UnpackIsFeatureEnabled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x47af267b.
//
// Solidity: function isFeatureEnabled(bytes32 ) view returns(bool)
func (iSystemConfig *ISystemConfig) UnpackIsFeatureEnabled(data []byte) (bool, error) {
	out, err := iSystemConfig.abi.Unpack("isFeatureEnabled", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackL1CrossDomainMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7119869.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function l1CrossDomainMessenger() view returns(address addr_)
func (iSystemConfig *ISystemConfig) PackL1CrossDomainMessenger() []byte {
	enc, err := iSystemConfig.abi.Pack("l1CrossDomainMessenger")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackL1CrossDomainMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7119869.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function l1CrossDomainMessenger() view returns(address addr_)
func (iSystemConfig *ISystemConfig) TryPackL1CrossDomainMessenger() ([]byte, error) {
	return iSystemConfig.abi.Pack("l1CrossDomainMessenger")
}

// UnpackL1CrossDomainMessenger is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address addr_)
func (iSystemConfig *ISystemConfig) UnpackL1CrossDomainMessenger(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("l1CrossDomainMessenger", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackL1ERC721Bridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4e8ddfa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function l1ERC721Bridge() view returns(address addr_)
func (iSystemConfig *ISystemConfig) PackL1ERC721Bridge() []byte {
	enc, err := iSystemConfig.abi.Pack("l1ERC721Bridge")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackL1ERC721Bridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4e8ddfa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function l1ERC721Bridge() view returns(address addr_)
func (iSystemConfig *ISystemConfig) TryPackL1ERC721Bridge() ([]byte, error) {
	return iSystemConfig.abi.Pack("l1ERC721Bridge")
}

// UnpackL1ERC721Bridge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address addr_)
func (iSystemConfig *ISystemConfig) UnpackL1ERC721Bridge(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("l1ERC721Bridge", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackL1StandardBridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x078f29cf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function l1StandardBridge() view returns(address addr_)
func (iSystemConfig *ISystemConfig) PackL1StandardBridge() []byte {
	enc, err := iSystemConfig.abi.Pack("l1StandardBridge")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackL1StandardBridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x078f29cf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function l1StandardBridge() view returns(address addr_)
func (iSystemConfig *ISystemConfig) TryPackL1StandardBridge() ([]byte, error) {
	return iSystemConfig.abi.Pack("l1StandardBridge")
}

// UnpackL1StandardBridge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address addr_)
func (iSystemConfig *ISystemConfig) UnpackL1StandardBridge(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("l1StandardBridge", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackL2ChainId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd6ae3cd5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function l2ChainId() view returns(uint256)
func (iSystemConfig *ISystemConfig) PackL2ChainId() []byte {
	enc, err := iSystemConfig.abi.Pack("l2ChainId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackL2ChainId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd6ae3cd5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function l2ChainId() view returns(uint256)
func (iSystemConfig *ISystemConfig) TryPackL2ChainId() ([]byte, error) {
	return iSystemConfig.abi.Pack("l2ChainId")
}

// UnpackL2ChainId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd6ae3cd5.
//
// Solidity: function l2ChainId() view returns(uint256)
func (iSystemConfig *ISystemConfig) UnpackL2ChainId(data []byte) (*big.Int, error) {
	out, err := iSystemConfig.abi.Unpack("l2ChainId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOptimismMintableERC20Factory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b7d7f0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function optimismMintableERC20Factory() view returns(address addr_)
func (iSystemConfig *ISystemConfig) PackOptimismMintableERC20Factory() []byte {
	enc, err := iSystemConfig.abi.Pack("optimismMintableERC20Factory")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOptimismMintableERC20Factory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b7d7f0a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function optimismMintableERC20Factory() view returns(address addr_)
func (iSystemConfig *ISystemConfig) TryPackOptimismMintableERC20Factory() ([]byte, error) {
	return iSystemConfig.abi.Pack("optimismMintableERC20Factory")
}

// UnpackOptimismMintableERC20Factory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address addr_)
func (iSystemConfig *ISystemConfig) UnpackOptimismMintableERC20Factory(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("optimismMintableERC20Factory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOptimismPortal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a49cb03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function optimismPortal() view returns(address addr_)
func (iSystemConfig *ISystemConfig) PackOptimismPortal() []byte {
	enc, err := iSystemConfig.abi.Pack("optimismPortal")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOptimismPortal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a49cb03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function optimismPortal() view returns(address addr_)
func (iSystemConfig *ISystemConfig) TryPackOptimismPortal() ([]byte, error) {
	return iSystemConfig.abi.Pack("optimismPortal")
}

// UnpackOptimismPortal is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0a49cb03.
//
// Solidity: function optimismPortal() view returns(address addr_)
func (iSystemConfig *ISystemConfig) UnpackOptimismPortal(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("optimismPortal", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOverhead is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c18c162.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function overhead() view returns(uint256)
func (iSystemConfig *ISystemConfig) PackOverhead() []byte {
	enc, err := iSystemConfig.abi.Pack("overhead")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOverhead is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c18c162.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function overhead() view returns(uint256)
func (iSystemConfig *ISystemConfig) TryPackOverhead() ([]byte, error) {
	return iSystemConfig.abi.Pack("overhead")
}

// UnpackOverhead is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (iSystemConfig *ISystemConfig) UnpackOverhead(data []byte) (*big.Int, error) {
	out, err := iSystemConfig.abi.Unpack("overhead", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function owner() view returns(address)
func (iSystemConfig *ISystemConfig) PackOwner() []byte {
	enc, err := iSystemConfig.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function owner() view returns(address)
func (iSystemConfig *ISystemConfig) TryPackOwner() ([]byte, error) {
	return iSystemConfig.abi.Pack("owner")
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (iSystemConfig *ISystemConfig) UnpackOwner(data []byte) (common.Address, error) {
	out, err := iSystemConfig.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function paused() view returns(bool)
func (iSystemConfig *ISystemConfig) PackPaused() []byte {
	enc, err := iSystemConfig.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paused() view returns(bool)
func (iSystemConfig *ISystemConfig) TryPackPaused() ([]byte, error) {
	return iSystemConfig.abi.Pack("paused")
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (iSystemConfig *ISystemConfig) UnpackPaused(data []byte) (bool, error) {
	out, err := iSystemConfig.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackStartBlock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48cd4cb1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function startBlock() view returns(uint256 startBlock_)
func (iSystemConfig *ISystemConfig) PackStartBlock() []byte {
	enc, err := iSystemConfig.abi.Pack("startBlock")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStartBlock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x48cd4cb1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function startBlock() view returns(uint256 startBlock_)
func (iSystemConfig *ISystemConfig) TryPackStartBlock() ([]byte, error) {
	return iSystemConfig.abi.Pack("startBlock")
}

// UnpackStartBlock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256 startBlock_)
func (iSystemConfig *ISystemConfig) UnpackStartBlock(data []byte) (*big.Int, error) {
	out, err := iSystemConfig.abi.Unpack("startBlock", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() pure returns(string)
func (iSystemConfig *ISystemConfig) PackVersion() []byte {
	enc, err := iSystemConfig.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function version() pure returns(string)
func (iSystemConfig *ISystemConfig) TryPackVersion() ([]byte, error) {
	return iSystemConfig.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (iSystemConfig *ISystemConfig) UnpackVersion(data []byte) (string, error) {
	out, err := iSystemConfig.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// ISystemConfigConfigUpdate represents a ConfigUpdate event raised by the ISystemConfig contract.
type ISystemConfigConfigUpdate struct {
	Version    *big.Int
	UpdateType uint8
	Data       []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const ISystemConfigConfigUpdateEventName = "ConfigUpdate"

// ContractEventName returns the user-defined event name.
func (ISystemConfigConfigUpdate) ContractEventName() string {
	return ISystemConfigConfigUpdateEventName
}

// UnpackConfigUpdateEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ConfigUpdate(uint256 indexed version, uint8 indexed updateType, bytes data)
func (iSystemConfig *ISystemConfig) UnpackConfigUpdateEvent(log *types.Log) (*ISystemConfigConfigUpdate, error) {
	event := "ConfigUpdate"
	if len(log.Topics) == 0 || log.Topics[0] != iSystemConfig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ISystemConfigConfigUpdate)
	if len(log.Data) > 0 {
		if err := iSystemConfig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iSystemConfig.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// ISystemConfigFeatureSet represents a FeatureSet event raised by the ISystemConfig contract.
type ISystemConfigFeatureSet struct {
	Feature [32]byte
	Enabled bool
	Raw     *types.Log // Blockchain specific contextual infos
}

const ISystemConfigFeatureSetEventName = "FeatureSet"

// ContractEventName returns the user-defined event name.
func (ISystemConfigFeatureSet) ContractEventName() string {
	return ISystemConfigFeatureSetEventName
}

// UnpackFeatureSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeatureSet(bytes32 indexed feature, bool indexed enabled)
func (iSystemConfig *ISystemConfig) UnpackFeatureSetEvent(log *types.Log) (*ISystemConfigFeatureSet, error) {
	event := "FeatureSet"
	if len(log.Topics) == 0 || log.Topics[0] != iSystemConfig.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ISystemConfigFeatureSet)
	if len(log.Data) > 0 {
		if err := iSystemConfig.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iSystemConfig.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
