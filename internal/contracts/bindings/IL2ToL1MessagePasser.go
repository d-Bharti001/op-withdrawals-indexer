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

// IL2ToL1MessagePasserMetaData contains all meta data concerning the IL2ToL1MessagePasser contract.
var IL2ToL1MessagePasserMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"}],\"name\":\"MessagePassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawerBalanceBurnt\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSAGE_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initiateWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sentMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "IL2ToL1MessagePasser",
}

// IL2ToL1MessagePasser is an auto generated Go binding around an Ethereum contract.
type IL2ToL1MessagePasser struct {
	abi abi.ABI
}

// NewIL2ToL1MessagePasser creates a new instance of IL2ToL1MessagePasser.
func NewIL2ToL1MessagePasser() *IL2ToL1MessagePasser {
	parsed, err := IL2ToL1MessagePasserMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IL2ToL1MessagePasser{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IL2ToL1MessagePasser) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackMESSAGEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f827a5a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) PackMESSAGEVERSION() []byte {
	enc, err := iL2ToL1MessagePasser.abi.Pack("MESSAGE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMESSAGEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f827a5a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) TryPackMESSAGEVERSION() ([]byte, error) {
	return iL2ToL1MessagePasser.abi.Pack("MESSAGE_VERSION")
}

// UnpackMESSAGEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) UnpackMESSAGEVERSION(data []byte) (uint16, error) {
	out, err := iL2ToL1MessagePasser.abi.Unpack("MESSAGE_VERSION", data)
	if err != nil {
		return *new(uint16), err
	}
	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	return out0, nil
}

// PackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44df8e70.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function burn() returns()
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) PackBurn() []byte {
	enc, err := iL2ToL1MessagePasser.abi.Pack("burn")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBurn is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x44df8e70.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function burn() returns()
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) TryPackBurn() ([]byte, error) {
	return iL2ToL1MessagePasser.abi.Pack("burn")
}

// PackInitiateWithdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2b3e5ac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) PackInitiateWithdrawal(target common.Address, gasLimit *big.Int, data []byte) []byte {
	enc, err := iL2ToL1MessagePasser.abi.Pack("initiateWithdrawal", target, gasLimit, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitiateWithdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2b3e5ac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) TryPackInitiateWithdrawal(target common.Address, gasLimit *big.Int, data []byte) ([]byte, error) {
	return iL2ToL1MessagePasser.abi.Pack("initiateWithdrawal", target, gasLimit, data)
}

// PackMessageNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecc70428.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function messageNonce() view returns(uint256)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) PackMessageNonce() []byte {
	enc, err := iL2ToL1MessagePasser.abi.Pack("messageNonce")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMessageNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecc70428.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function messageNonce() view returns(uint256)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) TryPackMessageNonce() ([]byte, error) {
	return iL2ToL1MessagePasser.abi.Pack("messageNonce")
}

// UnpackMessageNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) UnpackMessageNonce(data []byte) (*big.Int, error) {
	out, err := iL2ToL1MessagePasser.abi.Unpack("messageNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSentMessages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82e3702d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) PackSentMessages(arg0 [32]byte) []byte {
	enc, err := iL2ToL1MessagePasser.abi.Pack("sentMessages", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSentMessages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82e3702d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) TryPackSentMessages(arg0 [32]byte) ([]byte, error) {
	return iL2ToL1MessagePasser.abi.Pack("sentMessages", arg0)
}

// UnpackSentMessages is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) UnpackSentMessages(data []byte) (bool, error) {
	out, err := iL2ToL1MessagePasser.abi.Unpack("sentMessages", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function version() view returns(string)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) PackVersion() []byte {
	enc, err := iL2ToL1MessagePasser.abi.Pack("version")
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
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) TryPackVersion() ([]byte, error) {
	return iL2ToL1MessagePasser.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) UnpackVersion(data []byte) (string, error) {
	out, err := iL2ToL1MessagePasser.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// IL2ToL1MessagePasserMessagePassed represents a MessagePassed event raised by the IL2ToL1MessagePasser contract.
type IL2ToL1MessagePasserMessagePassed struct {
	Nonce          *big.Int
	Sender         common.Address
	Target         common.Address
	Value          *big.Int
	GasLimit       *big.Int
	Data           []byte
	WithdrawalHash [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const IL2ToL1MessagePasserMessagePassedEventName = "MessagePassed"

// ContractEventName returns the user-defined event name.
func (IL2ToL1MessagePasserMessagePassed) ContractEventName() string {
	return IL2ToL1MessagePasserMessagePassedEventName
}

// UnpackMessagePassedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) UnpackMessagePassedEvent(log *types.Log) (*IL2ToL1MessagePasserMessagePassed, error) {
	event := "MessagePassed"
	if len(log.Topics) == 0 || log.Topics[0] != iL2ToL1MessagePasser.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL2ToL1MessagePasserMessagePassed)
	if len(log.Data) > 0 {
		if err := iL2ToL1MessagePasser.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL2ToL1MessagePasser.abi.Events[event].Inputs {
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

// IL2ToL1MessagePasserWithdrawerBalanceBurnt represents a WithdrawerBalanceBurnt event raised by the IL2ToL1MessagePasser contract.
type IL2ToL1MessagePasserWithdrawerBalanceBurnt struct {
	Amount *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const IL2ToL1MessagePasserWithdrawerBalanceBurntEventName = "WithdrawerBalanceBurnt"

// ContractEventName returns the user-defined event name.
func (IL2ToL1MessagePasserWithdrawerBalanceBurnt) ContractEventName() string {
	return IL2ToL1MessagePasserWithdrawerBalanceBurntEventName
}

// UnpackWithdrawerBalanceBurntEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (iL2ToL1MessagePasser *IL2ToL1MessagePasser) UnpackWithdrawerBalanceBurntEvent(log *types.Log) (*IL2ToL1MessagePasserWithdrawerBalanceBurnt, error) {
	event := "WithdrawerBalanceBurnt"
	if len(log.Topics) == 0 || log.Topics[0] != iL2ToL1MessagePasser.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL2ToL1MessagePasserWithdrawerBalanceBurnt)
	if len(log.Data) > 0 {
		if err := iL2ToL1MessagePasser.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL2ToL1MessagePasser.abi.Events[event].Inputs {
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
