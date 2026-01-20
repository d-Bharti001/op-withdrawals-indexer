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

// IERC1155MetaData contains all meta data concerning the IERC1155 contract.
var IERC1155MetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "IERC1155",
}

// IERC1155 is an auto generated Go binding around an Ethereum contract.
type IERC1155 struct {
	abi abi.ABI
}

// NewIERC1155 creates a new instance of IERC1155.
func NewIERC1155() *IERC1155 {
	parsed, err := IERC1155MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC1155{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC1155) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00fdd58e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (iERC1155 *IERC1155) PackBalanceOf(account common.Address, id *big.Int) []byte {
	enc, err := iERC1155.abi.Pack("balanceOf", account, id)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00fdd58e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (iERC1155 *IERC1155) TryPackBalanceOf(account common.Address, id *big.Int) ([]byte, error) {
	return iERC1155.abi.Pack("balanceOf", account, id)
}

// UnpackBalanceOf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (iERC1155 *IERC1155) UnpackBalanceOf(data []byte) (*big.Int, error) {
	out, err := iERC1155.abi.Unpack("balanceOf", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBalanceOfBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e1273f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (iERC1155 *IERC1155) PackBalanceOfBatch(accounts []common.Address, ids []*big.Int) []byte {
	enc, err := iERC1155.abi.Pack("balanceOfBatch", accounts, ids)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBalanceOfBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4e1273f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (iERC1155 *IERC1155) TryPackBalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]byte, error) {
	return iERC1155.abi.Pack("balanceOfBatch", accounts, ids)
}

// UnpackBalanceOfBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (iERC1155 *IERC1155) UnpackBalanceOfBatch(data []byte) ([]*big.Int, error) {
	out, err := iERC1155.abi.Unpack("balanceOfBatch", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (iERC1155 *IERC1155) PackIsApprovedForAll(account common.Address, operator common.Address) []byte {
	enc, err := iERC1155.abi.Pack("isApprovedForAll", account, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsApprovedForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe985e9c5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (iERC1155 *IERC1155) TryPackIsApprovedForAll(account common.Address, operator common.Address) ([]byte, error) {
	return iERC1155.abi.Pack("isApprovedForAll", account, operator)
}

// UnpackIsApprovedForAll is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (iERC1155 *IERC1155) UnpackIsApprovedForAll(data []byte) (bool, error) {
	out, err := iERC1155.abi.Unpack("isApprovedForAll", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSafeBatchTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2eb2c2d6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (iERC1155 *IERC1155) PackSafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) []byte {
	enc, err := iERC1155.abi.Pack("safeBatchTransferFrom", from, to, ids, values, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeBatchTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2eb2c2d6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (iERC1155 *IERC1155) TryPackSafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) ([]byte, error) {
	return iERC1155.abi.Pack("safeBatchTransferFrom", from, to, ids, values, data)
}

// PackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf242432a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (iERC1155 *IERC1155) PackSafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) []byte {
	enc, err := iERC1155.abi.Pack("safeTransferFrom", from, to, id, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSafeTransferFrom is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf242432a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (iERC1155 *IERC1155) TryPackSafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) ([]byte, error) {
	return iERC1155.abi.Pack("safeTransferFrom", from, to, id, value, data)
}

// PackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (iERC1155 *IERC1155) PackSetApprovalForAll(operator common.Address, approved bool) []byte {
	enc, err := iERC1155.abi.Pack("setApprovalForAll", operator, approved)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetApprovalForAll is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa22cb465.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (iERC1155 *IERC1155) TryPackSetApprovalForAll(operator common.Address, approved bool) ([]byte, error) {
	return iERC1155.abi.Pack("setApprovalForAll", operator, approved)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC1155 *IERC1155) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := iERC1155.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC1155 *IERC1155) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return iERC1155.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (iERC1155 *IERC1155) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := iERC1155.abi.Unpack("supportsInterface", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// IERC1155ApprovalForAll represents a ApprovalForAll event raised by the IERC1155 contract.
type IERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC1155ApprovalForAllEventName = "ApprovalForAll"

// ContractEventName returns the user-defined event name.
func (IERC1155ApprovalForAll) ContractEventName() string {
	return IERC1155ApprovalForAllEventName
}

// UnpackApprovalForAllEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (iERC1155 *IERC1155) UnpackApprovalForAllEvent(log *types.Log) (*IERC1155ApprovalForAll, error) {
	event := "ApprovalForAll"
	if len(log.Topics) == 0 || log.Topics[0] != iERC1155.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC1155ApprovalForAll)
	if len(log.Data) > 0 {
		if err := iERC1155.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC1155.abi.Events[event].Inputs {
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

// IERC1155TransferBatch represents a TransferBatch event raised by the IERC1155 contract.
type IERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC1155TransferBatchEventName = "TransferBatch"

// ContractEventName returns the user-defined event name.
func (IERC1155TransferBatch) ContractEventName() string {
	return IERC1155TransferBatchEventName
}

// UnpackTransferBatchEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (iERC1155 *IERC1155) UnpackTransferBatchEvent(log *types.Log) (*IERC1155TransferBatch, error) {
	event := "TransferBatch"
	if len(log.Topics) == 0 || log.Topics[0] != iERC1155.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC1155TransferBatch)
	if len(log.Data) > 0 {
		if err := iERC1155.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC1155.abi.Events[event].Inputs {
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

// IERC1155TransferSingle represents a TransferSingle event raised by the IERC1155 contract.
type IERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const IERC1155TransferSingleEventName = "TransferSingle"

// ContractEventName returns the user-defined event name.
func (IERC1155TransferSingle) ContractEventName() string {
	return IERC1155TransferSingleEventName
}

// UnpackTransferSingleEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (iERC1155 *IERC1155) UnpackTransferSingleEvent(log *types.Log) (*IERC1155TransferSingle, error) {
	event := "TransferSingle"
	if len(log.Topics) == 0 || log.Topics[0] != iERC1155.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC1155TransferSingle)
	if len(log.Data) > 0 {
		if err := iERC1155.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC1155.abi.Events[event].Inputs {
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

// IERC1155URI represents a URI event raised by the IERC1155 contract.
type IERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const IERC1155URIEventName = "URI"

// ContractEventName returns the user-defined event name.
func (IERC1155URI) ContractEventName() string {
	return IERC1155URIEventName
}

// UnpackURIEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event URI(string value, uint256 indexed id)
func (iERC1155 *IERC1155) UnpackURIEvent(log *types.Log) (*IERC1155URI, error) {
	event := "URI"
	if len(log.Topics) == 0 || log.Topics[0] != iERC1155.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC1155URI)
	if len(log.Data) > 0 {
		if err := iERC1155.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC1155.abi.Events[event].Inputs {
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
