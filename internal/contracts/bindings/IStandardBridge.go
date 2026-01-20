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

// IStandardBridgeMetaData contains all meta data concerning the IStandardBridge contract.
var IStandardBridgeMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_BRIDGE\",\"outputs\":[{\"internalType\":\"contractIStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherBridge\",\"outputs\":[{\"internalType\":\"contractIStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "IStandardBridge",
}

// IStandardBridge is an auto generated Go binding around an Ethereum contract.
type IStandardBridge struct {
	abi abi.ABI
}

// NewIStandardBridge creates a new instance of IStandardBridge.
func NewIStandardBridge() *IStandardBridge {
	parsed, err := IStandardBridgeMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IStandardBridge{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IStandardBridge) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackMESSENGER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x927ede2d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MESSENGER() view returns(address)
func (iStandardBridge *IStandardBridge) PackMESSENGER() []byte {
	enc, err := iStandardBridge.abi.Pack("MESSENGER")
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
func (iStandardBridge *IStandardBridge) TryPackMESSENGER() ([]byte, error) {
	return iStandardBridge.abi.Pack("MESSENGER")
}

// UnpackMESSENGER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (iStandardBridge *IStandardBridge) UnpackMESSENGER(data []byte) (common.Address, error) {
	out, err := iStandardBridge.abi.Unpack("MESSENGER", data)
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
func (iStandardBridge *IStandardBridge) PackOTHERBRIDGE() []byte {
	enc, err := iStandardBridge.abi.Pack("OTHER_BRIDGE")
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
func (iStandardBridge *IStandardBridge) TryPackOTHERBRIDGE() ([]byte, error) {
	return iStandardBridge.abi.Pack("OTHER_BRIDGE")
}

// UnpackOTHERBRIDGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (iStandardBridge *IStandardBridge) UnpackOTHERBRIDGE(data []byte) (common.Address, error) {
	out, err := iStandardBridge.abi.Unpack("OTHER_BRIDGE", data)
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
func (iStandardBridge *IStandardBridge) PackBridgeERC20(localToken common.Address, remoteToken common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iStandardBridge.abi.Pack("bridgeERC20", localToken, remoteToken, amount, minGasLimit, extraData)
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
func (iStandardBridge *IStandardBridge) TryPackBridgeERC20(localToken common.Address, remoteToken common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iStandardBridge.abi.Pack("bridgeERC20", localToken, remoteToken, amount, minGasLimit, extraData)
}

// PackBridgeERC20To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x540abf73.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (iStandardBridge *IStandardBridge) PackBridgeERC20To(localToken common.Address, remoteToken common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iStandardBridge.abi.Pack("bridgeERC20To", localToken, remoteToken, to, amount, minGasLimit, extraData)
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
func (iStandardBridge *IStandardBridge) TryPackBridgeERC20To(localToken common.Address, remoteToken common.Address, to common.Address, amount *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iStandardBridge.abi.Pack("bridgeERC20To", localToken, remoteToken, to, amount, minGasLimit, extraData)
}

// PackBridgeETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09fc8843.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (iStandardBridge *IStandardBridge) PackBridgeETH(minGasLimit uint32, extraData []byte) []byte {
	enc, err := iStandardBridge.abi.Pack("bridgeETH", minGasLimit, extraData)
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
func (iStandardBridge *IStandardBridge) TryPackBridgeETH(minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iStandardBridge.abi.Pack("bridgeETH", minGasLimit, extraData)
}

// PackBridgeETHTo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe11013dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (iStandardBridge *IStandardBridge) PackBridgeETHTo(to common.Address, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iStandardBridge.abi.Pack("bridgeETHTo", to, minGasLimit, extraData)
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
func (iStandardBridge *IStandardBridge) TryPackBridgeETHTo(to common.Address, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iStandardBridge.abi.Pack("bridgeETHTo", to, minGasLimit, extraData)
}

// PackDeposits is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f601f66.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (iStandardBridge *IStandardBridge) PackDeposits(arg0 common.Address, arg1 common.Address) []byte {
	enc, err := iStandardBridge.abi.Pack("deposits", arg0, arg1)
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
func (iStandardBridge *IStandardBridge) TryPackDeposits(arg0 common.Address, arg1 common.Address) ([]byte, error) {
	return iStandardBridge.abi.Pack("deposits", arg0, arg1)
}

// UnpackDeposits is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (iStandardBridge *IStandardBridge) UnpackDeposits(data []byte) (*big.Int, error) {
	out, err := iStandardBridge.abi.Unpack("deposits", data)
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
func (iStandardBridge *IStandardBridge) PackFinalizeBridgeERC20(localToken common.Address, remoteToken common.Address, from common.Address, to common.Address, amount *big.Int, extraData []byte) []byte {
	enc, err := iStandardBridge.abi.Pack("finalizeBridgeERC20", localToken, remoteToken, from, to, amount, extraData)
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
func (iStandardBridge *IStandardBridge) TryPackFinalizeBridgeERC20(localToken common.Address, remoteToken common.Address, from common.Address, to common.Address, amount *big.Int, extraData []byte) ([]byte, error) {
	return iStandardBridge.abi.Pack("finalizeBridgeERC20", localToken, remoteToken, from, to, amount, extraData)
}

// PackFinalizeBridgeETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1635f5fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (iStandardBridge *IStandardBridge) PackFinalizeBridgeETH(from common.Address, to common.Address, amount *big.Int, extraData []byte) []byte {
	enc, err := iStandardBridge.abi.Pack("finalizeBridgeETH", from, to, amount, extraData)
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
func (iStandardBridge *IStandardBridge) TryPackFinalizeBridgeETH(from common.Address, to common.Address, amount *big.Int, extraData []byte) ([]byte, error) {
	return iStandardBridge.abi.Pack("finalizeBridgeETH", from, to, amount, extraData)
}

// PackMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3cb747bf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function messenger() view returns(address)
func (iStandardBridge *IStandardBridge) PackMessenger() []byte {
	enc, err := iStandardBridge.abi.Pack("messenger")
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
func (iStandardBridge *IStandardBridge) TryPackMessenger() ([]byte, error) {
	return iStandardBridge.abi.Pack("messenger")
}

// UnpackMessenger is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (iStandardBridge *IStandardBridge) UnpackMessenger(data []byte) (common.Address, error) {
	out, err := iStandardBridge.abi.Unpack("messenger", data)
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
func (iStandardBridge *IStandardBridge) PackOtherBridge() []byte {
	enc, err := iStandardBridge.abi.Pack("otherBridge")
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
func (iStandardBridge *IStandardBridge) TryPackOtherBridge() ([]byte, error) {
	return iStandardBridge.abi.Pack("otherBridge")
}

// UnpackOtherBridge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (iStandardBridge *IStandardBridge) UnpackOtherBridge(data []byte) (common.Address, error) {
	out, err := iStandardBridge.abi.Unpack("otherBridge", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// IStandardBridgeERC20BridgeFinalized represents a ERC20BridgeFinalized event raised by the IStandardBridge contract.
type IStandardBridgeERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const IStandardBridgeERC20BridgeFinalizedEventName = "ERC20BridgeFinalized"

// ContractEventName returns the user-defined event name.
func (IStandardBridgeERC20BridgeFinalized) ContractEventName() string {
	return IStandardBridgeERC20BridgeFinalizedEventName
}

// UnpackERC20BridgeFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (iStandardBridge *IStandardBridge) UnpackERC20BridgeFinalizedEvent(log *types.Log) (*IStandardBridgeERC20BridgeFinalized, error) {
	event := "ERC20BridgeFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iStandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStandardBridgeERC20BridgeFinalized)
	if len(log.Data) > 0 {
		if err := iStandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStandardBridge.abi.Events[event].Inputs {
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

// IStandardBridgeERC20BridgeInitiated represents a ERC20BridgeInitiated event raised by the IStandardBridge contract.
type IStandardBridgeERC20BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const IStandardBridgeERC20BridgeInitiatedEventName = "ERC20BridgeInitiated"

// ContractEventName returns the user-defined event name.
func (IStandardBridgeERC20BridgeInitiated) ContractEventName() string {
	return IStandardBridgeERC20BridgeInitiatedEventName
}

// UnpackERC20BridgeInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (iStandardBridge *IStandardBridge) UnpackERC20BridgeInitiatedEvent(log *types.Log) (*IStandardBridgeERC20BridgeInitiated, error) {
	event := "ERC20BridgeInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iStandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStandardBridgeERC20BridgeInitiated)
	if len(log.Data) > 0 {
		if err := iStandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStandardBridge.abi.Events[event].Inputs {
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

// IStandardBridgeETHBridgeFinalized represents a ETHBridgeFinalized event raised by the IStandardBridge contract.
type IStandardBridgeETHBridgeFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IStandardBridgeETHBridgeFinalizedEventName = "ETHBridgeFinalized"

// ContractEventName returns the user-defined event name.
func (IStandardBridgeETHBridgeFinalized) ContractEventName() string {
	return IStandardBridgeETHBridgeFinalizedEventName
}

// UnpackETHBridgeFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (iStandardBridge *IStandardBridge) UnpackETHBridgeFinalizedEvent(log *types.Log) (*IStandardBridgeETHBridgeFinalized, error) {
	event := "ETHBridgeFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iStandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStandardBridgeETHBridgeFinalized)
	if len(log.Data) > 0 {
		if err := iStandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStandardBridge.abi.Events[event].Inputs {
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

// IStandardBridgeETHBridgeInitiated represents a ETHBridgeInitiated event raised by the IStandardBridge contract.
type IStandardBridgeETHBridgeInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const IStandardBridgeETHBridgeInitiatedEventName = "ETHBridgeInitiated"

// ContractEventName returns the user-defined event name.
func (IStandardBridgeETHBridgeInitiated) ContractEventName() string {
	return IStandardBridgeETHBridgeInitiatedEventName
}

// UnpackETHBridgeInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (iStandardBridge *IStandardBridge) UnpackETHBridgeInitiatedEvent(log *types.Log) (*IStandardBridgeETHBridgeInitiated, error) {
	event := "ETHBridgeInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iStandardBridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IStandardBridgeETHBridgeInitiated)
	if len(log.Data) > 0 {
		if err := iStandardBridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iStandardBridge.abi.Events[event].Inputs {
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
