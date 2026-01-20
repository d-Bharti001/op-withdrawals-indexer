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

// IERC721BridgeMetaData contains all meta data concerning the IERC721Bridge contract.
var IERC721BridgeMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC721BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC721BridgeInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_BRIDGE\",\"outputs\":[{\"internalType\":\"contractIERC721Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC721To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherBridge\",\"outputs\":[{\"internalType\":\"contractIERC721Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "IERC721Bridge",
}

// IERC721Bridge is an auto generated Go binding around an Ethereum contract.
type IERC721Bridge struct {
	abi abi.ABI
}

// NewIERC721Bridge creates a new instance of IERC721Bridge.
func NewIERC721Bridge() *IERC721Bridge {
	parsed, err := IERC721BridgeMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IERC721Bridge{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IERC721Bridge) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackMESSENGER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x927ede2d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MESSENGER() view returns(address)
func (iERC721Bridge *IERC721Bridge) PackMESSENGER() []byte {
	enc, err := iERC721Bridge.abi.Pack("MESSENGER")
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
func (iERC721Bridge *IERC721Bridge) TryPackMESSENGER() ([]byte, error) {
	return iERC721Bridge.abi.Pack("MESSENGER")
}

// UnpackMESSENGER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (iERC721Bridge *IERC721Bridge) UnpackMESSENGER(data []byte) (common.Address, error) {
	out, err := iERC721Bridge.abi.Unpack("MESSENGER", data)
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
func (iERC721Bridge *IERC721Bridge) PackOTHERBRIDGE() []byte {
	enc, err := iERC721Bridge.abi.Pack("OTHER_BRIDGE")
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
func (iERC721Bridge *IERC721Bridge) TryPackOTHERBRIDGE() ([]byte, error) {
	return iERC721Bridge.abi.Pack("OTHER_BRIDGE")
}

// UnpackOTHERBRIDGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (iERC721Bridge *IERC721Bridge) UnpackOTHERBRIDGE(data []byte) (common.Address, error) {
	out, err := iERC721Bridge.abi.Unpack("OTHER_BRIDGE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackBridgeERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3687011a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (iERC721Bridge *IERC721Bridge) PackBridgeERC721(localToken common.Address, remoteToken common.Address, tokenId *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iERC721Bridge.abi.Pack("bridgeERC721", localToken, remoteToken, tokenId, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBridgeERC721 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3687011a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bridgeERC721(address _localToken, address _remoteToken, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (iERC721Bridge *IERC721Bridge) TryPackBridgeERC721(localToken common.Address, remoteToken common.Address, tokenId *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iERC721Bridge.abi.Pack("bridgeERC721", localToken, remoteToken, tokenId, minGasLimit, extraData)
}

// PackBridgeERC721To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa557452.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (iERC721Bridge *IERC721Bridge) PackBridgeERC721To(localToken common.Address, remoteToken common.Address, to common.Address, tokenId *big.Int, minGasLimit uint32, extraData []byte) []byte {
	enc, err := iERC721Bridge.abi.Pack("bridgeERC721To", localToken, remoteToken, to, tokenId, minGasLimit, extraData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBridgeERC721To is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa557452.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function bridgeERC721To(address _localToken, address _remoteToken, address _to, uint256 _tokenId, uint32 _minGasLimit, bytes _extraData) returns()
func (iERC721Bridge *IERC721Bridge) TryPackBridgeERC721To(localToken common.Address, remoteToken common.Address, to common.Address, tokenId *big.Int, minGasLimit uint32, extraData []byte) ([]byte, error) {
	return iERC721Bridge.abi.Pack("bridgeERC721To", localToken, remoteToken, to, tokenId, minGasLimit, extraData)
}

// PackMessenger is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3cb747bf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function messenger() view returns(address)
func (iERC721Bridge *IERC721Bridge) PackMessenger() []byte {
	enc, err := iERC721Bridge.abi.Pack("messenger")
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
func (iERC721Bridge *IERC721Bridge) TryPackMessenger() ([]byte, error) {
	return iERC721Bridge.abi.Pack("messenger")
}

// UnpackMessenger is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (iERC721Bridge *IERC721Bridge) UnpackMessenger(data []byte) (common.Address, error) {
	out, err := iERC721Bridge.abi.Unpack("messenger", data)
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
func (iERC721Bridge *IERC721Bridge) PackOtherBridge() []byte {
	enc, err := iERC721Bridge.abi.Pack("otherBridge")
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
func (iERC721Bridge *IERC721Bridge) TryPackOtherBridge() ([]byte, error) {
	return iERC721Bridge.abi.Pack("otherBridge")
}

// UnpackOtherBridge is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (iERC721Bridge *IERC721Bridge) UnpackOtherBridge(data []byte) (common.Address, error) {
	out, err := iERC721Bridge.abi.Unpack("otherBridge", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// IERC721BridgeERC721BridgeFinalized represents a ERC721BridgeFinalized event raised by the IERC721Bridge contract.
type IERC721BridgeERC721BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	TokenId     *big.Int
	ExtraData   []byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const IERC721BridgeERC721BridgeFinalizedEventName = "ERC721BridgeFinalized"

// ContractEventName returns the user-defined event name.
func (IERC721BridgeERC721BridgeFinalized) ContractEventName() string {
	return IERC721BridgeERC721BridgeFinalizedEventName
}

// UnpackERC721BridgeFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (iERC721Bridge *IERC721Bridge) UnpackERC721BridgeFinalizedEvent(log *types.Log) (*IERC721BridgeERC721BridgeFinalized, error) {
	event := "ERC721BridgeFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iERC721Bridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC721BridgeERC721BridgeFinalized)
	if len(log.Data) > 0 {
		if err := iERC721Bridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC721Bridge.abi.Events[event].Inputs {
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

// IERC721BridgeERC721BridgeInitiated represents a ERC721BridgeInitiated event raised by the IERC721Bridge contract.
type IERC721BridgeERC721BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	TokenId     *big.Int
	ExtraData   []byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const IERC721BridgeERC721BridgeInitiatedEventName = "ERC721BridgeInitiated"

// ContractEventName returns the user-defined event name.
func (IERC721BridgeERC721BridgeInitiated) ContractEventName() string {
	return IERC721BridgeERC721BridgeInitiatedEventName
}

// UnpackERC721BridgeInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ERC721BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 tokenId, bytes extraData)
func (iERC721Bridge *IERC721Bridge) UnpackERC721BridgeInitiatedEvent(log *types.Log) (*IERC721BridgeERC721BridgeInitiated, error) {
	event := "ERC721BridgeInitiated"
	if len(log.Topics) == 0 || log.Topics[0] != iERC721Bridge.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IERC721BridgeERC721BridgeInitiated)
	if len(log.Data) > 0 {
		if err := iERC721Bridge.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iERC721Bridge.abi.Events[event].Inputs {
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
