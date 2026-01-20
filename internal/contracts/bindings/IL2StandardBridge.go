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

// IL2StandardBridgeMetaData contains all meta data concerning the IL2StandardBridge contract.
var IL2StandardBridgeMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"DepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_BRIDGE\",\"outputs\":[{\"internalType\":\"contractIStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherBridge\",\"outputs\":[{\"internalType\":\"contractIStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "IL2StandardBridge",
}

// IL2StandardBridge is an auto generated Go binding around an Ethereum contract.
type IL2StandardBridge struct {
	abi abi.ABI
}

// NewIL2StandardBridge creates a new instance of IL2StandardBridge.
func NewIL2StandardBridge() *IL2StandardBridge {
	parsed, err := IL2StandardBridgeMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IL2StandardBridge{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IL2StandardBridge) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackMESSENGER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x927ede2d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MESSENGER() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) PackMESSENGER() []byte {
	enc, err := iL2StandardBridge.abi.Pack("MESSENGER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMESSENGER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x927ede2d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MESSENGER() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) TryPackMESSENGER() ([]byte, error) {
	return iL2StandardBridge.abi.Pack("MESSENGER")
}

// UnpackMESSENGER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) UnpackMESSENGER(data []byte) (common.Address, error) {
	out, err := iL2StandardBridge.abi.Unpack("MESSENGER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOTHERBRIDGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7f46ddb2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) PackOTHERBRIDGE() []byte {
	enc, err := iL2StandardBridge.abi.Pack("OTHER_BRIDGE")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOTHERBRIDGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7f46ddb2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) TryPackOTHERBRIDGE() ([]byte, error) {
	return iL2StandardBridge.abi.Pack("OTHER_BRIDGE")
}

// UnpackOTHERBRIDGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) UnpackOTHERBRIDGE(data []byte) (common.Address, error) {
	out, err := iL2StandardBridge.abi.Unpack("OTHER_BRIDGE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackBridgeERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87087623.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL2StandardBridge *IL2StandardBridge) PackBridgeERC20(localToken common.Address, remoteToken common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL2StandardBridge.abi.Pack("bridgeERC20", localToken, remoteToken, amount, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBridgeERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87087623.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL2StandardBridge *IL2StandardBridge) TryPackBridgeERC20(localToken common.Address, remoteToken common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("bridgeERC20", localToken, remoteToken, amount, minGasLimit, extraData)
}

// PackBridgeERC20To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x540abf73.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL2StandardBridge *IL2StandardBridge) PackBridgeERC20To(localToken common.Address, remoteToken common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL2StandardBridge.abi.Pack("bridgeERC20To", localToken, remoteToken, to, amount, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBridgeERC20To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x540abf73.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL2StandardBridge *IL2StandardBridge) TryPackBridgeERC20To(localToken common.Address, remoteToken common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("bridgeERC20To", localToken, remoteToken, to, amount, minGasLimit, extraData)
}

// PackBridgeETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09fc8843.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) PackBridgeETH(minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL2StandardBridge.abi.Pack("bridgeETH", minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBridgeETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09fc8843.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) TryPackBridgeETH(minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("bridgeETH", minGasLimit, extraData)
}

// PackBridgeETHTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe11013dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) PackBridgeETHTo(to common.Address, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL2StandardBridge.abi.Pack("bridgeETHTo", to, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBridgeETHTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe11013dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) TryPackBridgeETHTo(to common.Address, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("bridgeETHTo", to, minGasLimit, extraData)
}

// PackDeposits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f601f66.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (iL2StandardBridge *IL2StandardBridge) PackDeposits(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := iL2StandardBridge.abi.Pack("deposits", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeposits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f601f66.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (iL2StandardBridge *IL2StandardBridge) TryPackDeposits(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("deposits", arg0, arg1)
}

// UnpackDeposits is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (iL2StandardBridge *IL2StandardBridge) UnpackDeposits(data []byte) (*big.Int, error) {
	out, err := iL2StandardBridge.abi.Unpack("deposits", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackFinalizeBridgeERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0166a07a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (iL2StandardBridge *IL2StandardBridge) PackFinalizeBridgeERC20(localToken common.Address, remoteToken common.Address, from common.Address, to common.Address, amount *big.Int, extraData []byte) []byte {
	enc, err := iL2StandardBridge.abi.Pack("finalizeBridgeERC20", localToken, remoteToken, from, to, amount, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinalizeBridgeERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0166a07a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (iL2StandardBridge *IL2StandardBridge) TryPackFinalizeBridgeERC20(localToken common.Address, remoteToken common.Address, from common.Address, to common.Address, amount *big.Int, extraData []byte) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("finalizeBridgeERC20", localToken, remoteToken, from, to, amount, extraData)
}

// PackFinalizeBridgeETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1635f5fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) PackFinalizeBridgeETH(from common.Address, to common.Address, amount *big.Int, extraData []byte) []byte {
	enc, err := iL2StandardBridge.abi.Pack("finalizeBridgeETH", from, to, amount, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinalizeBridgeETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1635f5fd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) TryPackFinalizeBridgeETH(from common.Address, to common.Address, amount *big.Int, extraData []byte) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("finalizeBridgeETH", from, to, amount, extraData)
}

// PackL1TokenBridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36c717c1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function l1TokenBridge() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) PackL1TokenBridge() []byte {
	enc, err := iL2StandardBridge.abi.Pack("l1TokenBridge")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackL1TokenBridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36c717c1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function l1TokenBridge() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) TryPackL1TokenBridge() ([]byte, error) {
	return iL2StandardBridge.abi.Pack("l1TokenBridge")
}

// UnpackL1TokenBridge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x36c717c1.
//
// Solidity: function l1TokenBridge() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) UnpackL1TokenBridge(data []byte) (common.Address, error) {
	out, err := iL2StandardBridge.abi.Unpack("l1TokenBridge", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3cb747bf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function messenger() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) PackMessenger() []byte {
	enc, err := iL2StandardBridge.abi.Pack("messenger")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3cb747bf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function messenger() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) TryPackMessenger() ([]byte, error) {
	return iL2StandardBridge.abi.Pack("messenger")
}

// UnpackMessenger is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) UnpackMessenger(data []byte) (common.Address, error) {
	out, err := iL2StandardBridge.abi.Unpack("messenger", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackOtherBridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc89701a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function otherBridge() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) PackOtherBridge() []byte {
	enc, err := iL2StandardBridge.abi.Pack("otherBridge")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOtherBridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc89701a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function otherBridge() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) TryPackOtherBridge() ([]byte, error) {
	return iL2StandardBridge.abi.Pack("otherBridge")
}

// UnpackOtherBridge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (iL2StandardBridge *IL2StandardBridge) UnpackOtherBridge(data []byte) (common.Address, error) {
	out, err := iL2StandardBridge.abi.Unpack("otherBridge", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() pure returns(string)
func (iL2StandardBridge *IL2StandardBridge) PackVersion() []byte {
	enc, err := iL2StandardBridge.abi.Pack("version")
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
func (iL2StandardBridge *IL2StandardBridge) TryPackVersion() ([]byte, error) {
	return iL2StandardBridge.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (iL2StandardBridge *IL2StandardBridge) UnpackVersion(data []byte) (string, error) {
	out, err := iL2StandardBridge.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32b7006d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function withdraw(address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) PackWithdraw(l2Token common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL2StandardBridge.abi.Pack("withdraw", l2Token, amount, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32b7006d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function withdraw(address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) TryPackWithdraw(l2Token common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("withdraw", l2Token, amount, minGasLimit, extraData)
}

// PackWithdrawTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3a79548.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) PackWithdrawTo(l2Token common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL2StandardBridge.abi.Pack("withdrawTo", l2Token, to, amount, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackWithdrawTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3a79548.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function withdrawTo(address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL2StandardBridge *IL2StandardBridge) TryPackWithdrawTo(l2Token common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL2StandardBridge.abi.Pack("withdrawTo", l2Token, to, amount, minGasLimit, extraData)
}

// IL2StandardBridgeDepositFinalized represents a DepositFinalized event raised by the IL2StandardBridge contract.
type IL2StandardBridgeDepositFinalized struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL2StandardBridgeDepositFinalizedEventName = "DepositFinalized"

// ContractEventName returns the user-defined event name.
func (IL2StandardBridgeDepositFinalized) ContractEventName() string {
	return IL2StandardBridgeDepositFinalizedEventName
}

// UnpackDepositFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DepositFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (iL2StandardBridge *IL2StandardBridge) UnpackDepositFinalizedEvent(log *types.Log) (*IL2StandardBridgeDepositFinalized, error) {
	event := "DepositFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iL2StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL2StandardBridgeDepositFinalized)
	if len(log.Data) > 0 {
		if err := iL2StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL2StandardBridge.abi.Events[event].Inputs {
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

// IL2StandardBridgeERC20BridgeFinalized represents a ERC20BridgeFinalized event raised by the IL2StandardBridge contract.
type IL2StandardBridgeERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const IL2StandardBridgeERC20BridgeFinalizedEventName = "ERC20BridgeFinalized"

// ContractEventName returns the user-defined event name.
func (IL2StandardBridgeERC20BridgeFinalized) ContractEventName() string {
	return IL2StandardBridgeERC20BridgeFinalizedEventName
}

// UnpackERC20BridgeFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (iL2StandardBridge *IL2StandardBridge) UnpackERC20BridgeFinalizedEvent(log *types.Log) (*IL2StandardBridgeERC20BridgeFinalized, error) {
	event := "ERC20BridgeFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iL2StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL2StandardBridgeERC20BridgeFinalized)
	if len(log.Data) > 0 {
		if err := iL2StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL2StandardBridge.abi.Events[event].Inputs {
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

// IL2StandardBridgeERC20BridgeInitiated represents a ERC20BridgeInitiated event raised by the IL2StandardBridge contract.
type IL2StandardBridgeERC20BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const IL2StandardBridgeERC20BridgeInitiatedEventName = "ERC20BridgeInitiated"

// ContractEventName returns the user-defined event name.
func (IL2StandardBridgeERC20BridgeInitiated) ContractEventName() string {
	return IL2StandardBridgeERC20BridgeInitiatedEventName
}

// UnpackERC20BridgeInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (iL2StandardBridge *IL2StandardBridge) UnpackERC20BridgeInitiatedEvent(log *types.Log) (*IL2StandardBridgeERC20BridgeInitiated, error) {
	event := "ERC20BridgeInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iL2StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL2StandardBridgeERC20BridgeInitiated)
	if len(log.Data) > 0 {
		if err := iL2StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL2StandardBridge.abi.Events[event].Inputs {
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

// IL2StandardBridgeETHBridgeFinalized represents a ETHBridgeFinalized event raised by the IL2StandardBridge contract.
type IL2StandardBridgeETHBridgeFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL2StandardBridgeETHBridgeFinalizedEventName = "ETHBridgeFinalized"

// ContractEventName returns the user-defined event name.
func (IL2StandardBridgeETHBridgeFinalized) ContractEventName() string {
	return IL2StandardBridgeETHBridgeFinalizedEventName
}

// UnpackETHBridgeFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (iL2StandardBridge *IL2StandardBridge) UnpackETHBridgeFinalizedEvent(log *types.Log) (*IL2StandardBridgeETHBridgeFinalized, error) {
	event := "ETHBridgeFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iL2StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL2StandardBridgeETHBridgeFinalized)
	if len(log.Data) > 0 {
		if err := iL2StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL2StandardBridge.abi.Events[event].Inputs {
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

// IL2StandardBridgeETHBridgeInitiated represents a ETHBridgeInitiated event raised by the IL2StandardBridge contract.
type IL2StandardBridgeETHBridgeInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL2StandardBridgeETHBridgeInitiatedEventName = "ETHBridgeInitiated"

// ContractEventName returns the user-defined event name.
func (IL2StandardBridgeETHBridgeInitiated) ContractEventName() string {
	return IL2StandardBridgeETHBridgeInitiatedEventName
}

// UnpackETHBridgeInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (iL2StandardBridge *IL2StandardBridge) UnpackETHBridgeInitiatedEvent(log *types.Log) (*IL2StandardBridgeETHBridgeInitiated, error) {
	event := "ETHBridgeInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iL2StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL2StandardBridgeETHBridgeInitiated)
	if len(log.Data) > 0 {
		if err := iL2StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL2StandardBridge.abi.Events[event].Inputs {
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

// IL2StandardBridgeWithdrawalInitiated represents a WithdrawalInitiated event raised by the IL2StandardBridge contract.
type IL2StandardBridgeWithdrawalInitiated struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL2StandardBridgeWithdrawalInitiatedEventName = "WithdrawalInitiated"

// ContractEventName returns the user-defined event name.
func (IL2StandardBridgeWithdrawalInitiated) ContractEventName() string {
	return IL2StandardBridgeWithdrawalInitiatedEventName
}

// UnpackWithdrawalInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (iL2StandardBridge *IL2StandardBridge) UnpackWithdrawalInitiatedEvent(log *types.Log) (*IL2StandardBridgeWithdrawalInitiated, error) {
	event := "WithdrawalInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iL2StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL2StandardBridgeWithdrawalInitiated)
	if len(log.Data) > 0 {
		if err := iL2StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL2StandardBridge.abi.Events[event].Inputs {
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
