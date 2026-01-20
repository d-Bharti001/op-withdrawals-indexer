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

// IL1CrossDomainMessengerMetaData contains all meta data concerning the IL1CrossDomainMessenger contract.
var IL1CrossDomainMessengerMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SentMessageExtension1\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ENCODING_OVERHEAD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLOOR_CALLDATA_OVERHEAD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSAGE_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_GAS_CALLDATA_OVERHEAD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_CALL_OVERHEAD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_CONSTANT_OVERHEAD\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_GAS_CHECK_BUFFER\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAY_RESERVED_GAS\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TX_BASE_GAS\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"}],\"name\":\"baseGas\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"failedMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherMessenger\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"relayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"successfulMessages\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "IL1CrossDomainMessenger",
}

// IL1CrossDomainMessenger is an auto generated Go binding around an Ethereum contract.
type IL1CrossDomainMessenger struct {
	abi abi.ABI
}

// NewIL1CrossDomainMessenger creates a new instance of IL1CrossDomainMessenger.
func NewIL1CrossDomainMessenger() *IL1CrossDomainMessenger {
	parsed, err := IL1CrossDomainMessengerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IL1CrossDomainMessenger{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IL1CrossDomainMessenger) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackENCODINGOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddd5a40f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ENCODING_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackENCODINGOVERHEAD() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("ENCODING_OVERHEAD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackENCODINGOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xddd5a40f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ENCODING_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackENCODINGOVERHEAD() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("ENCODING_OVERHEAD")
}

// UnpackENCODINGOVERHEAD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xddd5a40f.
//
// Solidity: function ENCODING_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackENCODINGOVERHEAD(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("ENCODING_OVERHEAD", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackFLOORCALLDATAOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe46e245a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function FLOOR_CALLDATA_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackFLOORCALLDATAOVERHEAD() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("FLOOR_CALLDATA_OVERHEAD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFLOORCALLDATAOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe46e245a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function FLOOR_CALLDATA_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackFLOORCALLDATAOVERHEAD() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("FLOOR_CALLDATA_OVERHEAD")
}

// UnpackFLOORCALLDATAOVERHEAD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe46e245a.
//
// Solidity: function FLOOR_CALLDATA_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackFLOORCALLDATAOVERHEAD(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("FLOOR_CALLDATA_OVERHEAD", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackMESSAGEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f827a5a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackMESSAGEVERSION() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("MESSAGE_VERSION")
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
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackMESSAGEVERSION() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("MESSAGE_VERSION")
}

// UnpackMESSAGEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackMESSAGEVERSION(data []byte) (uint16, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("MESSAGE_VERSION", data)
	if err != nil {
		return *new(uint16), err
	}
	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	return out0, nil
}

// PackMINGASCALLDATAOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x028f85f7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackMINGASCALLDATAOVERHEAD() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("MIN_GAS_CALLDATA_OVERHEAD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINGASCALLDATAOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x028f85f7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackMINGASCALLDATAOVERHEAD() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("MIN_GAS_CALLDATA_OVERHEAD")
}

// UnpackMINGASCALLDATAOVERHEAD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x028f85f7.
//
// Solidity: function MIN_GAS_CALLDATA_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackMINGASCALLDATAOVERHEAD(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("MIN_GAS_CALLDATA_OVERHEAD", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackMINGASDYNAMICOVERHEADDENOMINATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c568498.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackMINGASDYNAMICOVERHEADDENOMINATOR() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINGASDYNAMICOVERHEADDENOMINATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c568498.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackMINGASDYNAMICOVERHEADDENOMINATOR() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR")
}

// UnpackMINGASDYNAMICOVERHEADDENOMINATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0c568498.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackMINGASDYNAMICOVERHEADDENOMINATOR(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("MIN_GAS_DYNAMIC_OVERHEAD_DENOMINATOR", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackMINGASDYNAMICOVERHEADNUMERATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2828d7e8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackMINGASDYNAMICOVERHEADNUMERATOR() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMINGASDYNAMICOVERHEADNUMERATOR is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2828d7e8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackMINGASDYNAMICOVERHEADNUMERATOR() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR")
}

// UnpackMINGASDYNAMICOVERHEADNUMERATOR is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2828d7e8.
//
// Solidity: function MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackMINGASDYNAMICOVERHEADNUMERATOR(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("MIN_GAS_DYNAMIC_OVERHEAD_NUMERATOR", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackOTHERMESSENGER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fce812c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackOTHERMESSENGER() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("OTHER_MESSENGER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOTHERMESSENGER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9fce812c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackOTHERMESSENGER() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("OTHER_MESSENGER")
}

// UnpackOTHERMESSENGER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9fce812c.
//
// Solidity: function OTHER_MESSENGER() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackOTHERMESSENGER(data []byte) (common.Address, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("OTHER_MESSENGER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRELAYCALLOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c1d6a69.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackRELAYCALLOVERHEAD() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("RELAY_CALL_OVERHEAD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRELAYCALLOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c1d6a69.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackRELAYCALLOVERHEAD() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("RELAY_CALL_OVERHEAD")
}

// UnpackRELAYCALLOVERHEAD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4c1d6a69.
//
// Solidity: function RELAY_CALL_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackRELAYCALLOVERHEAD(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("RELAY_CALL_OVERHEAD", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackRELAYCONSTANTOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83a74074.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackRELAYCONSTANTOVERHEAD() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("RELAY_CONSTANT_OVERHEAD")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRELAYCONSTANTOVERHEAD is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83a74074.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackRELAYCONSTANTOVERHEAD() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("RELAY_CONSTANT_OVERHEAD")
}

// UnpackRELAYCONSTANTOVERHEAD is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x83a74074.
//
// Solidity: function RELAY_CONSTANT_OVERHEAD() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackRELAYCONSTANTOVERHEAD(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("RELAY_CONSTANT_OVERHEAD", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackRELAYGASCHECKBUFFER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5644cfdf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackRELAYGASCHECKBUFFER() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("RELAY_GAS_CHECK_BUFFER")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRELAYGASCHECKBUFFER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5644cfdf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackRELAYGASCHECKBUFFER() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("RELAY_GAS_CHECK_BUFFER")
}

// UnpackRELAYGASCHECKBUFFER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5644cfdf.
//
// Solidity: function RELAY_GAS_CHECK_BUFFER() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackRELAYGASCHECKBUFFER(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("RELAY_GAS_CHECK_BUFFER", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackRELAYRESERVEDGAS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8cbeeef2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackRELAYRESERVEDGAS() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("RELAY_RESERVED_GAS")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRELAYRESERVEDGAS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8cbeeef2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackRELAYRESERVEDGAS() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("RELAY_RESERVED_GAS")
}

// UnpackRELAYRESERVEDGAS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8cbeeef2.
//
// Solidity: function RELAY_RESERVED_GAS() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackRELAYRESERVEDGAS(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("RELAY_RESERVED_GAS", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackTXBASEGAS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f7d3922.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function TX_BASE_GAS() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackTXBASEGAS() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("TX_BASE_GAS")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTXBASEGAS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f7d3922.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function TX_BASE_GAS() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackTXBASEGAS() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("TX_BASE_GAS")
}

// UnpackTXBASEGAS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2f7d3922.
//
// Solidity: function TX_BASE_GAS() view returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackTXBASEGAS(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("TX_BASE_GAS", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackBaseGas is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb28ade25.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackBaseGas(message []byte, minGasLimit uint32) []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("baseGas", message, minGasLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBaseGas is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb28ade25.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackBaseGas(message []byte, minGasLimit uint32) ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("baseGas", message, minGasLimit)
}

// UnpackBaseGas is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb28ade25.
//
// Solidity: function baseGas(bytes _message, uint32 _minGasLimit) pure returns(uint64)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackBaseGas(data []byte) (uint64, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("baseGas", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackFailedMessages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4e7f8bd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackFailedMessages(arg0 [32]byte) []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("failedMessages", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFailedMessages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4e7f8bd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackFailedMessages(arg0 [32]byte) ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("failedMessages", arg0)
}

// UnpackFailedMessages is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa4e7f8bd.
//
// Solidity: function failedMessages(bytes32 ) view returns(bool)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackFailedMessages(data []byte) (bool, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("failedMessages", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMessageNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecc70428.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function messageNonce() view returns(uint256)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackMessageNonce() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("messageNonce")
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
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackMessageNonce() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("messageNonce")
}

// UnpackMessageNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackMessageNonce(data []byte) (*big.Int, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("messageNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOtherMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb505d80.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function otherMessenger() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackOtherMessenger() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("otherMessenger")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOtherMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb505d80.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function otherMessenger() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackOtherMessenger() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("otherMessenger")
}

// UnpackOtherMessenger is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdb505d80.
//
// Solidity: function otherMessenger() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackOtherMessenger(data []byte) (common.Address, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("otherMessenger", data)
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
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackPaused() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("paused")
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
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackPaused() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("paused")
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackPaused(data []byte) (bool, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRelayMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd764ad0b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) payable returns()
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackRelayMessage(nonce *big.Int, sender common.Address, target common.Address, value *big.Int, minGasLimit *big.Int, message []byte) []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("relayMessage", nonce, sender, target, value, minGasLimit, message)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRelayMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd764ad0b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function relayMessage(uint256 _nonce, address _sender, address _target, uint256 _value, uint256 _minGasLimit, bytes _message) payable returns()
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackRelayMessage(nonce *big.Int, sender common.Address, target common.Address, value *big.Int, minGasLimit *big.Int, message []byte) ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("relayMessage", nonce, sender, target, value, minGasLimit, message)
}

// PackSendMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3dbb202b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _minGasLimit) payable returns()
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackSendMessage(target common.Address, message []byte, minGasLimit uint32) []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("sendMessage", target, message, minGasLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSendMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3dbb202b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sendMessage(address _target, bytes _message, uint32 _minGasLimit) payable returns()
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackSendMessage(target common.Address, message []byte, minGasLimit uint32) ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("sendMessage", target, message, minGasLimit)
}

// PackSuccessfulMessages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1b1b209.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackSuccessfulMessages(arg0 [32]byte) []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("successfulMessages", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSuccessfulMessages is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1b1b209.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackSuccessfulMessages(arg0 [32]byte) ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("successfulMessages", arg0)
}

// UnpackSuccessfulMessages is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb1b1b209.
//
// Solidity: function successfulMessages(bytes32 ) view returns(bool)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackSuccessfulMessages(data []byte) (bool, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("successfulMessages", data)
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
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackVersion() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("version")
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
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackVersion() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackVersion(data []byte) (string, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackXDomainMessageSender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e296e45.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) PackXDomainMessageSender() []byte {
	enc, err := iL1CrossDomainMessenger.abi.Pack("xDomainMessageSender")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackXDomainMessageSender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e296e45.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) TryPackXDomainMessageSender() ([]byte, error) {
	return iL1CrossDomainMessenger.abi.Pack("xDomainMessageSender")
}

// UnpackXDomainMessageSender is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackXDomainMessageSender(data []byte) (common.Address, error) {
	out, err := iL1CrossDomainMessenger.abi.Unpack("xDomainMessageSender", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// IL1CrossDomainMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerFailedRelayedMessage struct {
	MsgHash [32]byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const IL1CrossDomainMessengerFailedRelayedMessageEventName = "FailedRelayedMessage"

// ContractEventName returns the user-defined event name.
func (IL1CrossDomainMessengerFailedRelayedMessage) ContractEventName() string {
	return IL1CrossDomainMessengerFailedRelayedMessageEventName
}

// UnpackFailedRelayedMessageEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed msgHash)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackFailedRelayedMessageEvent(log *types.Log) (*IL1CrossDomainMessengerFailedRelayedMessage, error) {
	event := "FailedRelayedMessage"
	if len(log.Topics) == 0 || log.Topics[0] != iL1CrossDomainMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1CrossDomainMessengerFailedRelayedMessage)
	if len(log.Data) > 0 {
		if err := iL1CrossDomainMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1CrossDomainMessenger.abi.Events[event].Inputs {
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

// IL1CrossDomainMessengerInitialized represents a Initialized event raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const IL1CrossDomainMessengerInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (IL1CrossDomainMessengerInitialized) ContractEventName() string {
	return IL1CrossDomainMessengerInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackInitializedEvent(log *types.Log) (*IL1CrossDomainMessengerInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != iL1CrossDomainMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1CrossDomainMessengerInitialized)
	if len(log.Data) > 0 {
		if err := iL1CrossDomainMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1CrossDomainMessenger.abi.Events[event].Inputs {
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

// IL1CrossDomainMessengerRelayedMessage represents a RelayedMessage event raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerRelayedMessage struct {
	MsgHash [32]byte
	Raw     *types.Log // Blockchain specific contextual infos
}

const IL1CrossDomainMessengerRelayedMessageEventName = "RelayedMessage"

// ContractEventName returns the user-defined event name.
func (IL1CrossDomainMessengerRelayedMessage) ContractEventName() string {
	return IL1CrossDomainMessengerRelayedMessageEventName
}

// UnpackRelayedMessageEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RelayedMessage(bytes32 indexed msgHash)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackRelayedMessageEvent(log *types.Log) (*IL1CrossDomainMessengerRelayedMessage, error) {
	event := "RelayedMessage"
	if len(log.Topics) == 0 || log.Topics[0] != iL1CrossDomainMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1CrossDomainMessengerRelayedMessage)
	if len(log.Data) > 0 {
		if err := iL1CrossDomainMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1CrossDomainMessenger.abi.Events[event].Inputs {
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

// IL1CrossDomainMessengerSentMessage represents a SentMessage event raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerSentMessage struct {
	Target       common.Address
	Sender       common.Address
	Message      []byte
	MessageNonce *big.Int
	GasLimit     *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const IL1CrossDomainMessengerSentMessageEventName = "SentMessage"

// ContractEventName returns the user-defined event name.
func (IL1CrossDomainMessengerSentMessage) ContractEventName() string {
	return IL1CrossDomainMessengerSentMessageEventName
}

// UnpackSentMessageEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackSentMessageEvent(log *types.Log) (*IL1CrossDomainMessengerSentMessage, error) {
	event := "SentMessage"
	if len(log.Topics) == 0 || log.Topics[0] != iL1CrossDomainMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1CrossDomainMessengerSentMessage)
	if len(log.Data) > 0 {
		if err := iL1CrossDomainMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1CrossDomainMessenger.abi.Events[event].Inputs {
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

// IL1CrossDomainMessengerSentMessageExtension1 represents a SentMessageExtension1 event raised by the IL1CrossDomainMessenger contract.
type IL1CrossDomainMessengerSentMessageExtension1 struct {
	Sender common.Address
	Value  *big.Int
	Raw    *types.Log // Blockchain specific contextual infos
}

const IL1CrossDomainMessengerSentMessageExtension1EventName = "SentMessageExtension1"

// ContractEventName returns the user-defined event name.
func (IL1CrossDomainMessengerSentMessageExtension1) ContractEventName() string {
	return IL1CrossDomainMessengerSentMessageExtension1EventName
}

// UnpackSentMessageExtension1Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SentMessageExtension1(address indexed sender, uint256 value)
func (iL1CrossDomainMessenger *IL1CrossDomainMessenger) UnpackSentMessageExtension1Event(log *types.Log) (*IL1CrossDomainMessengerSentMessageExtension1, error) {
	event := "SentMessageExtension1"
	if len(log.Topics) == 0 || log.Topics[0] != iL1CrossDomainMessenger.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IL1CrossDomainMessengerSentMessageExtension1)
	if len(log.Data) > 0 {
		if err := iL1CrossDomainMessenger.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iL1CrossDomainMessenger.abi.Events[event].Inputs {
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
