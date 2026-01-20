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

// IL1StandardBridgeMetaData contains all meta data concerning the IL1StandardBridge contract.
var IL1StandardBridgeMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHWithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_BRIDGE\",\"outputs\":[{\"internalType\":\"contractIStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeETHWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherBridge\",\"outputs\":[{\"internalType\":\"contractIStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "IL1StandardBridge",
}

// IL1StandardBridge is an auto generated Go binding around an Ethereum contract.
type IL1StandardBridge struct {
	abi abi.ABI
}

// NewIL1StandardBridge creates a new instance of IL1StandardBridge.
func NewIL1StandardBridge() *IL1StandardBridge {
	parsed, err := IL1StandardBridgeMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IL1StandardBridge{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IL1StandardBridge) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackMESSENGER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x927ede2d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MESSENGER() view returns(address)
func (iL1StandardBridge *IL1StandardBridge) PackMESSENGER() []byte {
	enc, err := iL1StandardBridge.abi.Pack("MESSENGER")
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
func (iL1StandardBridge *IL1StandardBridge) TryPackMESSENGER() ([]byte, error) {
	return iL1StandardBridge.abi.Pack("MESSENGER")
}

// UnpackMESSENGER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (iL1StandardBridge *IL1StandardBridge) UnpackMESSENGER(data []byte) (common.Address, error) {
	out, err := iL1StandardBridge.abi.Unpack("MESSENGER", data)
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
func (iL1StandardBridge *IL1StandardBridge) PackOTHERBRIDGE() []byte {
	enc, err := iL1StandardBridge.abi.Pack("OTHER_BRIDGE")
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
func (iL1StandardBridge *IL1StandardBridge) TryPackOTHERBRIDGE() ([]byte, error) {
	return iL1StandardBridge.abi.Pack("OTHER_BRIDGE")
}

// UnpackOTHERBRIDGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (iL1StandardBridge *IL1StandardBridge) UnpackOTHERBRIDGE(data []byte) (common.Address, error) {
	out, err := iL1StandardBridge.abi.Unpack("OTHER_BRIDGE", data)
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
func (iL1StandardBridge *IL1StandardBridge) PackBridgeERC20(localToken common.Address, remoteToken common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("bridgeERC20", localToken, remoteToken, amount, minGasLimit, extraData)
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
func (iL1StandardBridge *IL1StandardBridge) TryPackBridgeERC20(localToken common.Address, remoteToken common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("bridgeERC20", localToken, remoteToken, amount, minGasLimit, extraData)
}

// PackBridgeERC20To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x540abf73.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL1StandardBridge *IL1StandardBridge) PackBridgeERC20To(localToken common.Address, remoteToken common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("bridgeERC20To", localToken, remoteToken, to, amount, minGasLimit, extraData)
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
func (iL1StandardBridge *IL1StandardBridge) TryPackBridgeERC20To(localToken common.Address, remoteToken common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("bridgeERC20To", localToken, remoteToken, to, amount, minGasLimit, extraData)
}

// PackBridgeETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09fc8843.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) PackBridgeETH(minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("bridgeETH", minGasLimit, extraData)
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
func (iL1StandardBridge *IL1StandardBridge) TryPackBridgeETH(minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("bridgeETH", minGasLimit, extraData)
}

// PackBridgeETHTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe11013dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) PackBridgeETHTo(to common.Address, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("bridgeETHTo", to, minGasLimit, extraData)
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
func (iL1StandardBridge *IL1StandardBridge) TryPackBridgeETHTo(to common.Address, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("bridgeETHTo", to, minGasLimit, extraData)
}

// PackDepositERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x58a997f6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL1StandardBridge *IL1StandardBridge) PackDepositERC20(l1Token common.Address, l2Token common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("depositERC20", l1Token, l2Token, amount, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDepositERC20 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x58a997f6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL1StandardBridge *IL1StandardBridge) TryPackDepositERC20(l1Token common.Address, l2Token common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("depositERC20", l1Token, l2Token, amount, minGasLimit, extraData)
}

// PackDepositERC20To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x838b2520.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL1StandardBridge *IL1StandardBridge) PackDepositERC20To(l1Token common.Address, l2Token common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("depositERC20To", l1Token, l2Token, to, amount, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDepositERC20To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x838b2520.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iL1StandardBridge *IL1StandardBridge) TryPackDepositERC20To(l1Token common.Address, l2Token common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("depositERC20To", l1Token, l2Token, to, amount, minGasLimit, extraData)
}

// PackDepositETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1a1a882.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function depositETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) PackDepositETH(minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("depositETH", minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDepositETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1a1a882.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function depositETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) TryPackDepositETH(minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("depositETH", minGasLimit, extraData)
}

// PackDepositETHTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a2ac6d5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function depositETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) PackDepositETHTo(to common.Address, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("depositETHTo", to, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDepositETHTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a2ac6d5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function depositETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) TryPackDepositETHTo(to common.Address, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("depositETHTo", to, minGasLimit, extraData)
}

// PackDeposits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f601f66.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (iL1StandardBridge *IL1StandardBridge) PackDeposits(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := iL1StandardBridge.abi.Pack("deposits", arg0, arg1)
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
func (iL1StandardBridge *IL1StandardBridge) TryPackDeposits(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("deposits", arg0, arg1)
}

// UnpackDeposits is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (iL1StandardBridge *IL1StandardBridge) UnpackDeposits(data []byte) (*big.Int, error) {
	out, err := iL1StandardBridge.abi.Unpack("deposits", data)
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
func (iL1StandardBridge *IL1StandardBridge) PackFinalizeBridgeERC20(localToken common.Address, remoteToken common.Address, from common.Address, to common.Address, amount *big.Int, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("finalizeBridgeERC20", localToken, remoteToken, from, to, amount, extraData)
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
func (iL1StandardBridge *IL1StandardBridge) TryPackFinalizeBridgeERC20(localToken common.Address, remoteToken common.Address, from common.Address, to common.Address, amount *big.Int, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("finalizeBridgeERC20", localToken, remoteToken, from, to, amount, extraData)
}

// PackFinalizeBridgeETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1635f5fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) PackFinalizeBridgeETH(from common.Address, to common.Address, amount *big.Int, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("finalizeBridgeETH", from, to, amount, extraData)
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
func (iL1StandardBridge *IL1StandardBridge) TryPackFinalizeBridgeETH(from common.Address, to common.Address, amount *big.Int, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("finalizeBridgeETH", from, to, amount, extraData)
}

// PackFinalizeERC20Withdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9f9e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (iL1StandardBridge *IL1StandardBridge) PackFinalizeERC20Withdrawal(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("finalizeERC20Withdrawal", l1Token, l2Token, from, to, amount, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinalizeERC20Withdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9f9e675.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (iL1StandardBridge *IL1StandardBridge) TryPackFinalizeERC20Withdrawal(l1Token common.Address, l2Token common.Address, from common.Address, to common.Address, amount *big.Int, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("finalizeERC20Withdrawal", l1Token, l2Token, from, to, amount, extraData)
}

// PackFinalizeETHWithdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1532ec34.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) PackFinalizeETHWithdrawal(from common.Address, to common.Address, amount *big.Int, extraData []byte) []byte {
	enc, err := iL1StandardBridge.abi.Pack("finalizeETHWithdrawal", from, to, amount, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinalizeETHWithdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1532ec34.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (iL1StandardBridge *IL1StandardBridge) TryPackFinalizeETHWithdrawal(from common.Address, to common.Address, amount *big.Int, extraData []byte) ([]byte, error) {
	return iL1StandardBridge.abi.Pack("finalizeETHWithdrawal", from, to, amount, extraData)
}

// PackL2TokenBridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91c49bf8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function l2TokenBridge() view returns(address)
func (iL1StandardBridge *IL1StandardBridge) PackL2TokenBridge() []byte {
	enc, err := iL1StandardBridge.abi.Pack("l2TokenBridge")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackL2TokenBridge is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91c49bf8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function l2TokenBridge() view returns(address)
func (iL1StandardBridge *IL1StandardBridge) TryPackL2TokenBridge() ([]byte, error) {
	return iL1StandardBridge.abi.Pack("l2TokenBridge")
}

// UnpackL2TokenBridge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (iL1StandardBridge *IL1StandardBridge) UnpackL2TokenBridge(data []byte) (common.Address, error) {
	out, err := iL1StandardBridge.abi.Unpack("l2TokenBridge", data)
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
func (iL1StandardBridge *IL1StandardBridge) PackMessenger() []byte {
	enc, err := iL1StandardBridge.abi.Pack("messenger")
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
func (iL1StandardBridge *IL1StandardBridge) TryPackMessenger() ([]byte, error) {
	return iL1StandardBridge.abi.Pack("messenger")
}

// UnpackMessenger is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (iL1StandardBridge *IL1StandardBridge) UnpackMessenger(data []byte) (common.Address, error) {
	out, err := iL1StandardBridge.abi.Unpack("messenger", data)
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
func (iL1StandardBridge *IL1StandardBridge) PackOtherBridge() []byte {
	enc, err := iL1StandardBridge.abi.Pack("otherBridge")
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
func (iL1StandardBridge *IL1StandardBridge) TryPackOtherBridge() ([]byte, error) {
	return iL1StandardBridge.abi.Pack("otherBridge")
}

// UnpackOtherBridge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (iL1StandardBridge *IL1StandardBridge) UnpackOtherBridge(data []byte) (common.Address, error) {
	out, err := iL1StandardBridge.abi.Unpack("otherBridge", data)
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
// Solidity: function version() view returns(string)
func (iL1StandardBridge *IL1StandardBridge) PackVersion() []byte {
	enc, err := iL1StandardBridge.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function version() view returns(string)
func (iL1StandardBridge *IL1StandardBridge) TryPackVersion() ([]byte, error) {
	return iL1StandardBridge.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (iL1StandardBridge *IL1StandardBridge) UnpackVersion(data []byte) (string, error) {
	out, err := iL1StandardBridge.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// IL1StandardBridgeERC20BridgeFinalized represents a ERC20BridgeFinalized event raised by the IL1StandardBridge contract.
type IL1StandardBridgeERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const IL1StandardBridgeERC20BridgeFinalizedEventName = "ERC20BridgeFinalized"

// ContractEventName returns the user-defined event name.
func (IL1StandardBridgeERC20BridgeFinalized) ContractEventName() string {
	return IL1StandardBridgeERC20BridgeFinalizedEventName
}

// UnpackERC20BridgeFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (iL1StandardBridge *IL1StandardBridge) UnpackERC20BridgeFinalizedEvent(log *types.Log) (*IL1StandardBridgeERC20BridgeFinalized, error) {
	event := "ERC20BridgeFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iL1StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1StandardBridgeERC20BridgeFinalized)
	if len(log.Data) > 0 {
		if err := iL1StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1StandardBridge.abi.Events[event].Inputs {
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

// IL1StandardBridgeERC20BridgeInitiated represents a ERC20BridgeInitiated event raised by the IL1StandardBridge contract.
type IL1StandardBridgeERC20BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const IL1StandardBridgeERC20BridgeInitiatedEventName = "ERC20BridgeInitiated"

// ContractEventName returns the user-defined event name.
func (IL1StandardBridgeERC20BridgeInitiated) ContractEventName() string {
	return IL1StandardBridgeERC20BridgeInitiatedEventName
}

// UnpackERC20BridgeInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (iL1StandardBridge *IL1StandardBridge) UnpackERC20BridgeInitiatedEvent(log *types.Log) (*IL1StandardBridgeERC20BridgeInitiated, error) {
	event := "ERC20BridgeInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iL1StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1StandardBridgeERC20BridgeInitiated)
	if len(log.Data) > 0 {
		if err := iL1StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1StandardBridge.abi.Events[event].Inputs {
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

// IL1StandardBridgeERC20DepositInitiated represents a ERC20DepositInitiated event raised by the IL1StandardBridge contract.
type IL1StandardBridgeERC20DepositInitiated struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL1StandardBridgeERC20DepositInitiatedEventName = "ERC20DepositInitiated"

// ContractEventName returns the user-defined event name.
func (IL1StandardBridgeERC20DepositInitiated) ContractEventName() string {
	return IL1StandardBridgeERC20DepositInitiatedEventName
}

// UnpackERC20DepositInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (iL1StandardBridge *IL1StandardBridge) UnpackERC20DepositInitiatedEvent(log *types.Log) (*IL1StandardBridgeERC20DepositInitiated, error) {
	event := "ERC20DepositInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iL1StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1StandardBridgeERC20DepositInitiated)
	if len(log.Data) > 0 {
		if err := iL1StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1StandardBridge.abi.Events[event].Inputs {
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

// IL1StandardBridgeERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the IL1StandardBridge contract.
type IL1StandardBridgeERC20WithdrawalFinalized struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL1StandardBridgeERC20WithdrawalFinalizedEventName = "ERC20WithdrawalFinalized"

// ContractEventName returns the user-defined event name.
func (IL1StandardBridgeERC20WithdrawalFinalized) ContractEventName() string {
	return IL1StandardBridgeERC20WithdrawalFinalizedEventName
}

// UnpackERC20WithdrawalFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (iL1StandardBridge *IL1StandardBridge) UnpackERC20WithdrawalFinalizedEvent(log *types.Log) (*IL1StandardBridgeERC20WithdrawalFinalized, error) {
	event := "ERC20WithdrawalFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iL1StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1StandardBridgeERC20WithdrawalFinalized)
	if len(log.Data) > 0 {
		if err := iL1StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1StandardBridge.abi.Events[event].Inputs {
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

// IL1StandardBridgeETHBridgeFinalized represents a ETHBridgeFinalized event raised by the IL1StandardBridge contract.
type IL1StandardBridgeETHBridgeFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL1StandardBridgeETHBridgeFinalizedEventName = "ETHBridgeFinalized"

// ContractEventName returns the user-defined event name.
func (IL1StandardBridgeETHBridgeFinalized) ContractEventName() string {
	return IL1StandardBridgeETHBridgeFinalizedEventName
}

// UnpackETHBridgeFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (iL1StandardBridge *IL1StandardBridge) UnpackETHBridgeFinalizedEvent(log *types.Log) (*IL1StandardBridgeETHBridgeFinalized, error) {
	event := "ETHBridgeFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iL1StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1StandardBridgeETHBridgeFinalized)
	if len(log.Data) > 0 {
		if err := iL1StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1StandardBridge.abi.Events[event].Inputs {
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

// IL1StandardBridgeETHBridgeInitiated represents a ETHBridgeInitiated event raised by the IL1StandardBridge contract.
type IL1StandardBridgeETHBridgeInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL1StandardBridgeETHBridgeInitiatedEventName = "ETHBridgeInitiated"

// ContractEventName returns the user-defined event name.
func (IL1StandardBridgeETHBridgeInitiated) ContractEventName() string {
	return IL1StandardBridgeETHBridgeInitiatedEventName
}

// UnpackETHBridgeInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (iL1StandardBridge *IL1StandardBridge) UnpackETHBridgeInitiatedEvent(log *types.Log) (*IL1StandardBridgeETHBridgeInitiated, error) {
	event := "ETHBridgeInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iL1StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1StandardBridgeETHBridgeInitiated)
	if len(log.Data) > 0 {
		if err := iL1StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1StandardBridge.abi.Events[event].Inputs {
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

// IL1StandardBridgeETHDepositInitiated represents a ETHDepositInitiated event raised by the IL1StandardBridge contract.
type IL1StandardBridgeETHDepositInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL1StandardBridgeETHDepositInitiatedEventName = "ETHDepositInitiated"

// ContractEventName returns the user-defined event name.
func (IL1StandardBridgeETHDepositInitiated) ContractEventName() string {
	return IL1StandardBridgeETHDepositInitiatedEventName
}

// UnpackETHDepositInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (iL1StandardBridge *IL1StandardBridge) UnpackETHDepositInitiatedEvent(log *types.Log) (*IL1StandardBridgeETHDepositInitiated, error) {
	event := "ETHDepositInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iL1StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1StandardBridgeETHDepositInitiated)
	if len(log.Data) > 0 {
		if err := iL1StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1StandardBridge.abi.Events[event].Inputs {
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

// IL1StandardBridgeETHWithdrawalFinalized represents a ETHWithdrawalFinalized event raised by the IL1StandardBridge contract.
type IL1StandardBridgeETHWithdrawalFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IL1StandardBridgeETHWithdrawalFinalizedEventName = "ETHWithdrawalFinalized"

// ContractEventName returns the user-defined event name.
func (IL1StandardBridgeETHWithdrawalFinalized) ContractEventName() string {
	return IL1StandardBridgeETHWithdrawalFinalizedEventName
}

// UnpackETHWithdrawalFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (iL1StandardBridge *IL1StandardBridge) UnpackETHWithdrawalFinalizedEvent(log *types.Log) (*IL1StandardBridgeETHWithdrawalFinalized, error) {
	event := "ETHWithdrawalFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iL1StandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1StandardBridgeETHWithdrawalFinalized)
	if len(log.Data) > 0 {
		if err := iL1StandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1StandardBridge.abi.Events[event].Inputs {
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
