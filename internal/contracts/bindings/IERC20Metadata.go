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

// IERC20MetadataMetaData contains all meta data concerning the IERC20Metadata contract.
var IERC20MetadataMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "IERC20Metadata",
}

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	abi abi.ABI
}

// NewIERC20Metadata creates a new instance of IERC20Metadata.
func NewIERC20Metadata() *IERC20Metadata {
	parsed, err := IERC20MetadataMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC20Metadata{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC20Metadata) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iERC20Metadata *IERC20Metadata) PackAllowance(owner common.Address, spender common.Address) []byte {
	enc, err := iERC20Metadata.abi.Pack("allowance", owner, spender)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd62ed3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iERC20Metadata *IERC20Metadata) TryPackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	return iERC20Metadata.abi.Pack("allowance", owner, spender)
}

// UnpackAllowance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (iERC20Metadata *IERC20Metadata) UnpackAllowance(data []byte) (*big.Int, error) {
	out, err := iERC20Metadata.abi.Unpack("allowance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) PackApprove(spender common.Address, value *big.Int) []byte {
	enc, err := iERC20Metadata.abi.Pack("approve", spender, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApprove is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x095ea7b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) TryPackApprove(spender common.Address, value *big.Int) ([]byte, error) {
	return iERC20Metadata.abi.Pack("approve", spender, value)
}

// UnpackApprove is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) UnpackApprove(data []byte) (bool, error) {
	out, err := iERC20Metadata.abi.Unpack("approve", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iERC20Metadata *IERC20Metadata) PackBalanceOf(account common.Address) []byte {
	enc, err := iERC20Metadata.abi.Pack("balanceOf", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70a08231.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iERC20Metadata *IERC20Metadata) TryPackBalanceOf(account common.Address) ([]byte, error) {
	return iERC20Metadata.abi.Pack("balanceOf", account)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (iERC20Metadata *IERC20Metadata) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iERC20Metadata.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function decimals() view returns(uint8)
func (iERC20Metadata *IERC20Metadata) PackDecimals() []byte {
	enc, err := iERC20Metadata.abi.Pack("decimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x313ce567.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function decimals() view returns(uint8)
func (iERC20Metadata *IERC20Metadata) TryPackDecimals() ([]byte, error) {
	return iERC20Metadata.abi.Pack("decimals")
}

// UnpackDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (iERC20Metadata *IERC20Metadata) UnpackDecimals(data []byte) (uint8, error) {
	out, err := iERC20Metadata.abi.Unpack("decimals", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function name() view returns(string)
func (iERC20Metadata *IERC20Metadata) PackName() []byte {
	enc, err := iERC20Metadata.abi.Pack("name")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06fdde03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function name() view returns(string)
func (iERC20Metadata *IERC20Metadata) TryPackName() ([]byte, error) {
	return iERC20Metadata.abi.Pack("name")
}

// UnpackName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (iERC20Metadata *IERC20Metadata) UnpackName(data []byte) (string, error) {
	out, err := iERC20Metadata.abi.Unpack("name", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function symbol() view returns(string)
func (iERC20Metadata *IERC20Metadata) PackSymbol() []byte {
	enc, err := iERC20Metadata.abi.Pack("symbol")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSymbol is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x95d89b41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function symbol() view returns(string)
func (iERC20Metadata *IERC20Metadata) TryPackSymbol() ([]byte, error) {
	return iERC20Metadata.abi.Pack("symbol")
}

// UnpackSymbol is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (iERC20Metadata *IERC20Metadata) UnpackSymbol(data []byte) (string, error) {
	out, err := iERC20Metadata.abi.Unpack("symbol", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalSupply() view returns(uint256)
func (iERC20Metadata *IERC20Metadata) PackTotalSupply() []byte {
	enc, err := iERC20Metadata.abi.Pack("totalSupply")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalSupply is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18160ddd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalSupply() view returns(uint256)
func (iERC20Metadata *IERC20Metadata) TryPackTotalSupply() ([]byte, error) {
	return iERC20Metadata.abi.Pack("totalSupply")
}

// UnpackTotalSupply is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (iERC20Metadata *IERC20Metadata) UnpackTotalSupply(data []byte) (*big.Int, error) {
	out, err := iERC20Metadata.abi.Unpack("totalSupply", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) PackTransfer(to common.Address, value *big.Int) []byte {
	enc, err := iERC20Metadata.abi.Pack("transfer", to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9059cbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) TryPackTransfer(to common.Address, value *big.Int) ([]byte, error) {
	return iERC20Metadata.abi.Pack("transfer", to, value)
}

// UnpackTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) UnpackTransfer(data []byte) (bool, error) {
	out, err := iERC20Metadata.abi.Unpack("transfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) PackTransferFrom(from common.Address, to common.Address, value *big.Int) []byte {
	enc, err := iERC20Metadata.abi.Pack("transferFrom", from, to, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23b872dd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) TryPackTransferFrom(from common.Address, to common.Address, value *big.Int) ([]byte, error) {
	return iERC20Metadata.abi.Pack("transferFrom", from, to, value)
}

// UnpackTransferFrom is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (iERC20Metadata *IERC20Metadata) UnpackTransferFrom(data []byte) (bool, error) {
	out, err := iERC20Metadata.abi.Unpack("transferFrom", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     *types.Log // Blockchain specific contextual infos
}

const IERC20MetadataApprovalEventName = "Approval"

// ContractEventName returns the user-defined event name.
func (IERC20MetadataApproval) ContractEventName() string {
	return IERC20MetadataApprovalEventName
}

// UnpackApprovalEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (iERC20Metadata *IERC20Metadata) UnpackApprovalEvent(log *types.Log) (*IERC20MetadataApproval, error) {
	event := "Approval"
	if len(log.Topics) == 0 || log.Topics[0] != iERC20Metadata.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC20MetadataApproval)
	if len(log.Data) > 0 {
		if err := iERC20Metadata.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC20Metadata.abi.Events[event].Inputs {
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

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const IERC20MetadataTransferEventName = "Transfer"

// ContractEventName returns the user-defined event name.
func (IERC20MetadataTransfer) ContractEventName() string {
	return IERC20MetadataTransferEventName
}

// UnpackTransferEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (iERC20Metadata *IERC20Metadata) UnpackTransferEvent(log *types.Log) (*IERC20MetadataTransfer, error) {
	event := "Transfer"
	if len(log.Topics) == 0 || log.Topics[0] != iERC20Metadata.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC20MetadataTransfer)
	if len(log.Data) > 0 {
		if err := iERC20Metadata.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC20Metadata.abi.Events[event].Inputs {
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
