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

// IOptimismPortal2MetaData contains all meta data concerning the IOptimismPortal2 contract.
var IOptimismPortal2MetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ContentLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRemainder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHeader\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_AlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_BadTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_CallPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_CalldataTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_GasEstimation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_GasLimitTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_ImproperDisputeGame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_InvalidDisputeGame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_InvalidLockboxState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_InvalidOutputRootProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_InvalidProofTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_InvalidRootClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_NoReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_ProofNotOldEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OptimismPortal_Unproven\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReinitializableBase_ZeroInitVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedString\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"opaqueData\",\"type\":\"bytes\"}],\"name\":\"TransactionDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"WithdrawalProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proofSubmitter\",\"type\":\"address\"}],\"name\":\"WithdrawalProvenExtension1\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_withdrawalHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_proofSubmitter\",\"type\":\"address\"}],\"name\":\"checkWithdrawal\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"finalizedWithdrawals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_withdrawalHash\",\"type\":\"bytes32\"}],\"name\":\"numProofSubmitters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofMaturityDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proofSubmitters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "IOptimismPortal2",
}

// IOptimismPortal2 is an auto generated Go binding around an Ethereum contract.
type IOptimismPortal2 struct {
	abi abi.ABI
}

// NewIOptimismPortal2 creates a new instance of IOptimismPortal2.
func NewIOptimismPortal2() *IOptimismPortal2 {
	parsed, err := IOptimismPortal2MetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &IOptimismPortal2{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *IOptimismPortal2) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCheckWithdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x71c1566e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function checkWithdrawal(bytes32 _withdrawalHash, address _proofSubmitter) view returns()
func (iOptimismPortal2 *IOptimismPortal2) PackCheckWithdrawal(withdrawalHash [32]byte, proofSubmitter common.Address) []byte {
	enc, err := iOptimismPortal2.abi.Pack("checkWithdrawal", withdrawalHash, proofSubmitter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCheckWithdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x71c1566e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function checkWithdrawal(bytes32 _withdrawalHash, address _proofSubmitter) view returns()
func (iOptimismPortal2 *IOptimismPortal2) TryPackCheckWithdrawal(withdrawalHash [32]byte, proofSubmitter common.Address) ([]byte, error) {
	return iOptimismPortal2.abi.Pack("checkWithdrawal", withdrawalHash, proofSubmitter)
}

// PackFinalizedWithdrawals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa14238e7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (iOptimismPortal2 *IOptimismPortal2) PackFinalizedWithdrawals(arg0 [32]byte) []byte {
	enc, err := iOptimismPortal2.abi.Pack("finalizedWithdrawals", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinalizedWithdrawals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa14238e7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (iOptimismPortal2 *IOptimismPortal2) TryPackFinalizedWithdrawals(arg0 [32]byte) ([]byte, error) {
	return iOptimismPortal2.abi.Pack("finalizedWithdrawals", arg0)
}

// UnpackFinalizedWithdrawals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (iOptimismPortal2 *IOptimismPortal2) UnpackFinalizedWithdrawals(data []byte) (bool, error) {
	out, err := iOptimismPortal2.abi.Unpack("finalizedWithdrawals", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackNumProofSubmitters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x513747ab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function numProofSubmitters(bytes32 _withdrawalHash) view returns(uint256)
func (iOptimismPortal2 *IOptimismPortal2) PackNumProofSubmitters(withdrawalHash [32]byte) []byte {
	enc, err := iOptimismPortal2.abi.Pack("numProofSubmitters", withdrawalHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNumProofSubmitters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x513747ab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function numProofSubmitters(bytes32 _withdrawalHash) view returns(uint256)
func (iOptimismPortal2 *IOptimismPortal2) TryPackNumProofSubmitters(withdrawalHash [32]byte) ([]byte, error) {
	return iOptimismPortal2.abi.Pack("numProofSubmitters", withdrawalHash)
}

// UnpackNumProofSubmitters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x513747ab.
//
// Solidity: function numProofSubmitters(bytes32 _withdrawalHash) view returns(uint256)
func (iOptimismPortal2 *IOptimismPortal2) UnpackNumProofSubmitters(data []byte) (*big.Int, error) {
	out, err := iOptimismPortal2.abi.Unpack("numProofSubmitters", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProofMaturityDelaySeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf653a5c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proofMaturityDelaySeconds() view returns(uint256)
func (iOptimismPortal2 *IOptimismPortal2) PackProofMaturityDelaySeconds() []byte {
	enc, err := iOptimismPortal2.abi.Pack("proofMaturityDelaySeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProofMaturityDelaySeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf653a5c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proofMaturityDelaySeconds() view returns(uint256)
func (iOptimismPortal2 *IOptimismPortal2) TryPackProofMaturityDelaySeconds() ([]byte, error) {
	return iOptimismPortal2.abi.Pack("proofMaturityDelaySeconds")
}

// UnpackProofMaturityDelaySeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf653a5c.
//
// Solidity: function proofMaturityDelaySeconds() view returns(uint256)
func (iOptimismPortal2 *IOptimismPortal2) UnpackProofMaturityDelaySeconds(data []byte) (*big.Int, error) {
	out, err := iOptimismPortal2.abi.Unpack("proofMaturityDelaySeconds", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProofSubmitters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3860f48.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proofSubmitters(bytes32 , uint256 ) view returns(address)
func (iOptimismPortal2 *IOptimismPortal2) PackProofSubmitters(arg0 [32]byte, arg1 *big.Int) []byte {
	enc, err := iOptimismPortal2.abi.Pack("proofSubmitters", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProofSubmitters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3860f48.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proofSubmitters(bytes32 , uint256 ) view returns(address)
func (iOptimismPortal2 *IOptimismPortal2) TryPackProofSubmitters(arg0 [32]byte, arg1 *big.Int) ([]byte, error) {
	return iOptimismPortal2.abi.Pack("proofSubmitters", arg0, arg1)
}

// UnpackProofSubmitters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3860f48.
//
// Solidity: function proofSubmitters(bytes32 , uint256 ) view returns(address)
func (iOptimismPortal2 *IOptimismPortal2) UnpackProofSubmitters(data []byte) (common.Address, error) {
	out, err := iOptimismPortal2.abi.Unpack("proofSubmitters", data)
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
func (iOptimismPortal2 *IOptimismPortal2) PackVersion() []byte {
	enc, err := iOptimismPortal2.abi.Pack("version")
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
func (iOptimismPortal2 *IOptimismPortal2) TryPackVersion() ([]byte, error) {
	return iOptimismPortal2.abi.Pack("version")
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (iOptimismPortal2 *IOptimismPortal2) UnpackVersion(data []byte) (string, error) {
	out, err := iOptimismPortal2.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// IOptimismPortal2TransactionDeposited represents a TransactionDeposited event raised by the IOptimismPortal2 contract.
type IOptimismPortal2TransactionDeposited struct {
	From       common.Address
	To         common.Address
	Version    *big.Int
	OpaqueData []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const IOptimismPortal2TransactionDepositedEventName = "TransactionDeposited"

// ContractEventName returns the user-defined event name.
func (IOptimismPortal2TransactionDeposited) ContractEventName() string {
	return IOptimismPortal2TransactionDepositedEventName
}

// UnpackTransactionDepositedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (iOptimismPortal2 *IOptimismPortal2) UnpackTransactionDepositedEvent(log *types.Log) (*IOptimismPortal2TransactionDeposited, error) {
	event := "TransactionDeposited"
	if len(log.Topics) == 0 || log.Topics[0] != iOptimismPortal2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IOptimismPortal2TransactionDeposited)
	if len(log.Data) > 0 {
		if err := iOptimismPortal2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iOptimismPortal2.abi.Events[event].Inputs {
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

// IOptimismPortal2WithdrawalFinalized represents a WithdrawalFinalized event raised by the IOptimismPortal2 contract.
type IOptimismPortal2WithdrawalFinalized struct {
	WithdrawalHash [32]byte
	Success        bool
	Raw            *types.Log // Blockchain specific contextual infos
}

const IOptimismPortal2WithdrawalFinalizedEventName = "WithdrawalFinalized"

// ContractEventName returns the user-defined event name.
func (IOptimismPortal2WithdrawalFinalized) ContractEventName() string {
	return IOptimismPortal2WithdrawalFinalizedEventName
}

// UnpackWithdrawalFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (iOptimismPortal2 *IOptimismPortal2) UnpackWithdrawalFinalizedEvent(log *types.Log) (*IOptimismPortal2WithdrawalFinalized, error) {
	event := "WithdrawalFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != iOptimismPortal2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IOptimismPortal2WithdrawalFinalized)
	if len(log.Data) > 0 {
		if err := iOptimismPortal2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iOptimismPortal2.abi.Events[event].Inputs {
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

// IOptimismPortal2WithdrawalProven represents a WithdrawalProven event raised by the IOptimismPortal2 contract.
type IOptimismPortal2WithdrawalProven struct {
	WithdrawalHash [32]byte
	From           common.Address
	To             common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const IOptimismPortal2WithdrawalProvenEventName = "WithdrawalProven"

// ContractEventName returns the user-defined event name.
func (IOptimismPortal2WithdrawalProven) ContractEventName() string {
	return IOptimismPortal2WithdrawalProvenEventName
}

// UnpackWithdrawalProvenEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (iOptimismPortal2 *IOptimismPortal2) UnpackWithdrawalProvenEvent(log *types.Log) (*IOptimismPortal2WithdrawalProven, error) {
	event := "WithdrawalProven"
	if len(log.Topics) == 0 || log.Topics[0] != iOptimismPortal2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IOptimismPortal2WithdrawalProven)
	if len(log.Data) > 0 {
		if err := iOptimismPortal2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iOptimismPortal2.abi.Events[event].Inputs {
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

// IOptimismPortal2WithdrawalProvenExtension1 represents a WithdrawalProvenExtension1 event raised by the IOptimismPortal2 contract.
type IOptimismPortal2WithdrawalProvenExtension1 struct {
	WithdrawalHash [32]byte
	ProofSubmitter common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const IOptimismPortal2WithdrawalProvenExtension1EventName = "WithdrawalProvenExtension1"

// ContractEventName returns the user-defined event name.
func (IOptimismPortal2WithdrawalProvenExtension1) ContractEventName() string {
	return IOptimismPortal2WithdrawalProvenExtension1EventName
}

// UnpackWithdrawalProvenExtension1Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalProvenExtension1(bytes32 indexed withdrawalHash, address indexed proofSubmitter)
func (iOptimismPortal2 *IOptimismPortal2) UnpackWithdrawalProvenExtension1Event(log *types.Log) (*IOptimismPortal2WithdrawalProvenExtension1, error) {
	event := "WithdrawalProvenExtension1"
	if len(log.Topics) == 0 || log.Topics[0] != iOptimismPortal2.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(IOptimismPortal2WithdrawalProvenExtension1)
	if len(log.Data) > 0 {
		if err := iOptimismPortal2.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range iOptimismPortal2.abi.Events[event].Inputs {
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

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (iOptimismPortal2 *IOptimismPortal2) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["ContentLengthMismatch"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackContentLengthMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["EmptyItem"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackEmptyItemError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["InvalidDataRemainder"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackInvalidDataRemainderError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["InvalidHeader"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackInvalidHeaderError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalAlreadyFinalized"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalAlreadyFinalizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalBadTarget"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalBadTargetError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalCallPaused"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalCallPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalCalldataTooLarge"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalCalldataTooLargeError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalGasEstimation"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalGasEstimationError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalGasLimitTooLow"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalGasLimitTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalImproperDisputeGame"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalImproperDisputeGameError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalInvalidDisputeGame"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalInvalidDisputeGameError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalInvalidLockboxState"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalInvalidLockboxStateError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalInvalidMerkleProof"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalInvalidMerkleProofError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalInvalidOutputRootProof"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalInvalidOutputRootProofError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalInvalidProofTimestamp"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalInvalidProofTimestampError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalInvalidRootClaim"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalInvalidRootClaimError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalNoReentrancy"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalNoReentrancyError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalProofNotOldEnough"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalProofNotOldEnoughError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OptimismPortalUnproven"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOptimismPortalUnprovenError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["OutOfGas"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackOutOfGasError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["ReinitializableBaseZeroInitVersion"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackReinitializableBaseZeroInitVersionError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["UnexpectedList"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackUnexpectedListError(raw[4:])
	}
	if bytes.Equal(raw[:4], iOptimismPortal2.abi.Errors["UnexpectedString"].ID.Bytes()[:4]) {
		return iOptimismPortal2.UnpackUnexpectedStringError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// IOptimismPortal2ContentLengthMismatch represents a ContentLengthMismatch error raised by the IOptimismPortal2 contract.
type IOptimismPortal2ContentLengthMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ContentLengthMismatch()
func IOptimismPortal2ContentLengthMismatchErrorID() common.Hash {
	return common.HexToHash("0x66c944854c74f126ddb347ede5ddccb741d5d7926bed34960c8d068de5a8f475")
}

// UnpackContentLengthMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ContentLengthMismatch()
func (iOptimismPortal2 *IOptimismPortal2) UnpackContentLengthMismatchError(raw []byte) (*IOptimismPortal2ContentLengthMismatch, error) {
	out := new(IOptimismPortal2ContentLengthMismatch)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "ContentLengthMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2EmptyItem represents a EmptyItem error raised by the IOptimismPortal2 contract.
type IOptimismPortal2EmptyItem struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyItem()
func IOptimismPortal2EmptyItemErrorID() common.Hash {
	return common.HexToHash("0x5ab458fbe9490fbea962b45039c8271fd72a0f01d48093712043b25dc6ba708f")
}

// UnpackEmptyItemError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyItem()
func (iOptimismPortal2 *IOptimismPortal2) UnpackEmptyItemError(raw []byte) (*IOptimismPortal2EmptyItem, error) {
	out := new(IOptimismPortal2EmptyItem)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "EmptyItem", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2InvalidDataRemainder represents a InvalidDataRemainder error raised by the IOptimismPortal2 contract.
type IOptimismPortal2InvalidDataRemainder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDataRemainder()
func IOptimismPortal2InvalidDataRemainderErrorID() common.Hash {
	return common.HexToHash("0x5c5537b8753217199634d3b78fb6e68bd912e435605abb57ae70f055c98106de")
}

// UnpackInvalidDataRemainderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDataRemainder()
func (iOptimismPortal2 *IOptimismPortal2) UnpackInvalidDataRemainderError(raw []byte) (*IOptimismPortal2InvalidDataRemainder, error) {
	out := new(IOptimismPortal2InvalidDataRemainder)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "InvalidDataRemainder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2InvalidHeader represents a InvalidHeader error raised by the IOptimismPortal2 contract.
type IOptimismPortal2InvalidHeader struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidHeader()
func IOptimismPortal2InvalidHeaderErrorID() common.Hash {
	return common.HexToHash("0xbabb01ddbf35fa5cb4266dbe78b38e673d97e7c3b147273565c15eb8980ded33")
}

// UnpackInvalidHeaderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidHeader()
func (iOptimismPortal2 *IOptimismPortal2) UnpackInvalidHeaderError(raw []byte) (*IOptimismPortal2InvalidHeader, error) {
	out := new(IOptimismPortal2InvalidHeader)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "InvalidHeader", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalAlreadyFinalized represents a OptimismPortal_AlreadyFinalized error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalAlreadyFinalized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_AlreadyFinalized()
func IOptimismPortal2OptimismPortalAlreadyFinalizedErrorID() common.Hash {
	return common.HexToHash("0x730a1074a9bb79dfbd475058a425676a9f7d9d7fcc58f0394a03d992edfa1e3f")
}

// UnpackOptimismPortalAlreadyFinalizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_AlreadyFinalized()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalAlreadyFinalizedError(raw []byte) (*IOptimismPortal2OptimismPortalAlreadyFinalized, error) {
	out := new(IOptimismPortal2OptimismPortalAlreadyFinalized)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalAlreadyFinalized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalBadTarget represents a OptimismPortal_BadTarget error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalBadTarget struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_BadTarget()
func IOptimismPortal2OptimismPortalBadTargetErrorID() common.Hash {
	return common.HexToHash("0xc5defbadfc82123c06d9be3cab0d0e172a0f34091e235e2f1e309fea84b84c6a")
}

// UnpackOptimismPortalBadTargetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_BadTarget()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalBadTargetError(raw []byte) (*IOptimismPortal2OptimismPortalBadTarget, error) {
	out := new(IOptimismPortal2OptimismPortalBadTarget)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalBadTarget", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalCallPaused represents a OptimismPortal_CallPaused error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalCallPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_CallPaused()
func IOptimismPortal2OptimismPortalCallPausedErrorID() common.Hash {
	return common.HexToHash("0xb9c3c2ef1a8ae62b5f397b7b328607fcb745b2d598f7e56296b79be542375399")
}

// UnpackOptimismPortalCallPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_CallPaused()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalCallPausedError(raw []byte) (*IOptimismPortal2OptimismPortalCallPaused, error) {
	out := new(IOptimismPortal2OptimismPortalCallPaused)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalCallPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalCalldataTooLarge represents a OptimismPortal_CalldataTooLarge error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalCalldataTooLarge struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_CalldataTooLarge()
func IOptimismPortal2OptimismPortalCalldataTooLargeErrorID() common.Hash {
	return common.HexToHash("0x5aa3bac9d8300dad568614c187e4b6b59476fddbf302d1be1f59e48a463eb29c")
}

// UnpackOptimismPortalCalldataTooLargeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_CalldataTooLarge()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalCalldataTooLargeError(raw []byte) (*IOptimismPortal2OptimismPortalCalldataTooLarge, error) {
	out := new(IOptimismPortal2OptimismPortalCalldataTooLarge)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalCalldataTooLarge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalGasEstimation represents a OptimismPortal_GasEstimation error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalGasEstimation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_GasEstimation()
func IOptimismPortal2OptimismPortalGasEstimationErrorID() common.Hash {
	return common.HexToHash("0xab58103661f51038625302cb475e68f8a2323b148fbf16b4c3cfc471c438b82d")
}

// UnpackOptimismPortalGasEstimationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_GasEstimation()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalGasEstimationError(raw []byte) (*IOptimismPortal2OptimismPortalGasEstimation, error) {
	out := new(IOptimismPortal2OptimismPortalGasEstimation)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalGasEstimation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalGasLimitTooLow represents a OptimismPortal_GasLimitTooLow error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalGasLimitTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_GasLimitTooLow()
func IOptimismPortal2OptimismPortalGasLimitTooLowErrorID() common.Hash {
	return common.HexToHash("0x70c8bdbdbf53af52983b9bb6c7a87e362527b8ce3768296c009421a1cd4e741b")
}

// UnpackOptimismPortalGasLimitTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_GasLimitTooLow()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalGasLimitTooLowError(raw []byte) (*IOptimismPortal2OptimismPortalGasLimitTooLow, error) {
	out := new(IOptimismPortal2OptimismPortalGasLimitTooLow)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalGasLimitTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalImproperDisputeGame represents a OptimismPortal_ImproperDisputeGame error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalImproperDisputeGame struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_ImproperDisputeGame()
func IOptimismPortal2OptimismPortalImproperDisputeGameErrorID() common.Hash {
	return common.HexToHash("0xf395240ed477d22fe64b42522f66df6f7380c0b472663779e22dc1a3ea8018e4")
}

// UnpackOptimismPortalImproperDisputeGameError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_ImproperDisputeGame()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalImproperDisputeGameError(raw []byte) (*IOptimismPortal2OptimismPortalImproperDisputeGame, error) {
	out := new(IOptimismPortal2OptimismPortalImproperDisputeGame)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalImproperDisputeGame", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalInvalidDisputeGame represents a OptimismPortal_InvalidDisputeGame error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalInvalidDisputeGame struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_InvalidDisputeGame()
func IOptimismPortal2OptimismPortalInvalidDisputeGameErrorID() common.Hash {
	return common.HexToHash("0xe29927ed709dfe3791e070296e9cf0ddd4ea4e107c19a12b81f7f78dc02de307")
}

// UnpackOptimismPortalInvalidDisputeGameError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_InvalidDisputeGame()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalInvalidDisputeGameError(raw []byte) (*IOptimismPortal2OptimismPortalInvalidDisputeGame, error) {
	out := new(IOptimismPortal2OptimismPortalInvalidDisputeGame)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalInvalidDisputeGame", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalInvalidLockboxState represents a OptimismPortal_InvalidLockboxState error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalInvalidLockboxState struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_InvalidLockboxState()
func IOptimismPortal2OptimismPortalInvalidLockboxStateErrorID() common.Hash {
	return common.HexToHash("0x9c46cd791a8a54f341ef36583cf1e0631014bd3c74f80fd9b2501b349e50a550")
}

// UnpackOptimismPortalInvalidLockboxStateError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_InvalidLockboxState()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalInvalidLockboxStateError(raw []byte) (*IOptimismPortal2OptimismPortalInvalidLockboxState, error) {
	out := new(IOptimismPortal2OptimismPortalInvalidLockboxState)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalInvalidLockboxState", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalInvalidMerkleProof represents a OptimismPortal_InvalidMerkleProof error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalInvalidMerkleProof struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_InvalidMerkleProof()
func IOptimismPortal2OptimismPortalInvalidMerkleProofErrorID() common.Hash {
	return common.HexToHash("0x2e57ef3a3ec7735847f06e2835bbd4da209a337abae5ca10ac5865078cd42af3")
}

// UnpackOptimismPortalInvalidMerkleProofError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_InvalidMerkleProof()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalInvalidMerkleProofError(raw []byte) (*IOptimismPortal2OptimismPortalInvalidMerkleProof, error) {
	out := new(IOptimismPortal2OptimismPortalInvalidMerkleProof)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalInvalidMerkleProof", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalInvalidOutputRootProof represents a OptimismPortal_InvalidOutputRootProof error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalInvalidOutputRootProof struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_InvalidOutputRootProof()
func IOptimismPortal2OptimismPortalInvalidOutputRootProofErrorID() common.Hash {
	return common.HexToHash("0x426149af805023d6931f6bb96a8f179b5c81a68d63f163b9df6d0a8b94b6cfe1")
}

// UnpackOptimismPortalInvalidOutputRootProofError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_InvalidOutputRootProof()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalInvalidOutputRootProofError(raw []byte) (*IOptimismPortal2OptimismPortalInvalidOutputRootProof, error) {
	out := new(IOptimismPortal2OptimismPortalInvalidOutputRootProof)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalInvalidOutputRootProof", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalInvalidProofTimestamp represents a OptimismPortal_InvalidProofTimestamp error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalInvalidProofTimestamp struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_InvalidProofTimestamp()
func IOptimismPortal2OptimismPortalInvalidProofTimestampErrorID() common.Hash {
	return common.HexToHash("0xb4caa4e5aecdd558fec8bc0ede27c24782a6c622bfaa90dacb73bc51642f34c5")
}

// UnpackOptimismPortalInvalidProofTimestampError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_InvalidProofTimestamp()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalInvalidProofTimestampError(raw []byte) (*IOptimismPortal2OptimismPortalInvalidProofTimestamp, error) {
	out := new(IOptimismPortal2OptimismPortalInvalidProofTimestamp)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalInvalidProofTimestamp", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalInvalidRootClaim represents a OptimismPortal_InvalidRootClaim error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalInvalidRootClaim struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_InvalidRootClaim()
func IOptimismPortal2OptimismPortalInvalidRootClaimErrorID() common.Hash {
	return common.HexToHash("0x332a57f8fc402a456ebf5639b2a0c04b7137b624adcc5844248bca5441492ed2")
}

// UnpackOptimismPortalInvalidRootClaimError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_InvalidRootClaim()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalInvalidRootClaimError(raw []byte) (*IOptimismPortal2OptimismPortalInvalidRootClaim, error) {
	out := new(IOptimismPortal2OptimismPortalInvalidRootClaim)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalInvalidRootClaim", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalNoReentrancy represents a OptimismPortal_NoReentrancy error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalNoReentrancy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_NoReentrancy()
func IOptimismPortal2OptimismPortalNoReentrancyErrorID() common.Hash {
	return common.HexToHash("0xdfeaaeb899ef9901a10a6c78c72e6bfc948041ec48fc68708fdc2b864925f591")
}

// UnpackOptimismPortalNoReentrancyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_NoReentrancy()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalNoReentrancyError(raw []byte) (*IOptimismPortal2OptimismPortalNoReentrancy, error) {
	out := new(IOptimismPortal2OptimismPortalNoReentrancy)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalNoReentrancy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalProofNotOldEnough represents a OptimismPortal_ProofNotOldEnough error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalProofNotOldEnough struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_ProofNotOldEnough()
func IOptimismPortal2OptimismPortalProofNotOldEnoughErrorID() common.Hash {
	return common.HexToHash("0xd9bc01be8ff2adf05643a6bad8780feef42d4969bceb92e68bb3dad24fcf304e")
}

// UnpackOptimismPortalProofNotOldEnoughError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_ProofNotOldEnough()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalProofNotOldEnoughError(raw []byte) (*IOptimismPortal2OptimismPortalProofNotOldEnough, error) {
	out := new(IOptimismPortal2OptimismPortalProofNotOldEnough)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalProofNotOldEnough", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OptimismPortalUnproven represents a OptimismPortal_Unproven error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OptimismPortalUnproven struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OptimismPortal_Unproven()
func IOptimismPortal2OptimismPortalUnprovenErrorID() common.Hash {
	return common.HexToHash("0xcca6afdacc11528a039e49552380a48b500084f2af2330e59bdf4365ad918b70")
}

// UnpackOptimismPortalUnprovenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OptimismPortal_Unproven()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOptimismPortalUnprovenError(raw []byte) (*IOptimismPortal2OptimismPortalUnproven, error) {
	out := new(IOptimismPortal2OptimismPortalUnproven)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OptimismPortalUnproven", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2OutOfGas represents a OutOfGas error raised by the IOptimismPortal2 contract.
type IOptimismPortal2OutOfGas struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OutOfGas()
func IOptimismPortal2OutOfGasErrorID() common.Hash {
	return common.HexToHash("0x77ebef4d460014cfc45ab055292b3a8f4a7c9e9ab42973e79fdaff6b27233461")
}

// UnpackOutOfGasError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OutOfGas()
func (iOptimismPortal2 *IOptimismPortal2) UnpackOutOfGasError(raw []byte) (*IOptimismPortal2OutOfGas, error) {
	out := new(IOptimismPortal2OutOfGas)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "OutOfGas", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2ReinitializableBaseZeroInitVersion represents a ReinitializableBase_ZeroInitVersion error raised by the IOptimismPortal2 contract.
type IOptimismPortal2ReinitializableBaseZeroInitVersion struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReinitializableBase_ZeroInitVersion()
func IOptimismPortal2ReinitializableBaseZeroInitVersionErrorID() common.Hash {
	return common.HexToHash("0x9b01afed5b67e3af3a9a91660591206df676627816e027f7470b964260da493b")
}

// UnpackReinitializableBaseZeroInitVersionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReinitializableBase_ZeroInitVersion()
func (iOptimismPortal2 *IOptimismPortal2) UnpackReinitializableBaseZeroInitVersionError(raw []byte) (*IOptimismPortal2ReinitializableBaseZeroInitVersion, error) {
	out := new(IOptimismPortal2ReinitializableBaseZeroInitVersion)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "ReinitializableBaseZeroInitVersion", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2UnexpectedList represents a UnexpectedList error raised by the IOptimismPortal2 contract.
type IOptimismPortal2UnexpectedList struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnexpectedList()
func IOptimismPortal2UnexpectedListErrorID() common.Hash {
	return common.HexToHash("0x1ff9b2e4a9eba439988cd6b8e020496ddcf6d6cefbd8ccaa0c6006330e0183fd")
}

// UnpackUnexpectedListError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnexpectedList()
func (iOptimismPortal2 *IOptimismPortal2) UnpackUnexpectedListError(raw []byte) (*IOptimismPortal2UnexpectedList, error) {
	out := new(IOptimismPortal2UnexpectedList)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "UnexpectedList", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// IOptimismPortal2UnexpectedString represents a UnexpectedString error raised by the IOptimismPortal2 contract.
type IOptimismPortal2UnexpectedString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnexpectedString()
func IOptimismPortal2UnexpectedStringErrorID() common.Hash {
	return common.HexToHash("0x4b9c6abeb08b7a3e5561c35d6edd32e38a8bf05c863166482044d76d816eaad0")
}

// UnpackUnexpectedStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnexpectedString()
func (iOptimismPortal2 *IOptimismPortal2) UnpackUnexpectedStringError(raw []byte) (*IOptimismPortal2UnexpectedString, error) {
	out := new(IOptimismPortal2UnexpectedString)
	if err := iOptimismPortal2.abi.UnpackIntoInterface(out, "UnexpectedString", raw); err != nil {
		return nil, err
	}
	return out, nil
}
