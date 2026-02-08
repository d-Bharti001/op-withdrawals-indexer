package decoder

import (
	"bytes"
	"errors"
	"op-withdrawals-indexer/internal/contracts/bindings"
	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type EventDecoder struct {
	IL2ToL1MessagePasserABI    abi.ABI
	IL2CrossDomainMessengerABI abi.ABI
	IL2StandardBridgeABI       abi.ABI
	IL2ERC721BridgeABI         abi.ABI
}

var ErrDecodeFailed = errors.New("could not decode event action")

func NewEventDecoder() (*EventDecoder, error) {
	IL2ToL1MessagePasserABI, err := bindings.IL2ToL1MessagePasserMetaData.ParseABI()
	if err != nil {
		return nil, err
	}
	IL2CrossDomainMessengerABI, err := bindings.IL2CrossDomainMessengerMetaData.ParseABI()
	if err != nil {
		return nil, err
	}
	IL2StandardBridgeABI, err := bindings.IL2StandardBridgeMetaData.ParseABI()
	if err != nil {
		return nil, err
	}
	IL2ERC721BridgeABI, err := bindings.IL2ERC721BridgeMetaData.ParseABI()
	if err != nil {
		return nil, err
	}

	return &EventDecoder{
		IL2ToL1MessagePasserABI:    *IL2ToL1MessagePasserABI,
		IL2CrossDomainMessengerABI: *IL2CrossDomainMessengerABI,
		IL2StandardBridgeABI:       *IL2StandardBridgeABI,
		IL2ERC721BridgeABI:         *IL2ERC721BridgeABI,
	}, nil
}

func (d *EventDecoder) DecodeL2MessagePassedEvent(event *bindings.IL2ToL1MessagePasserMessagePassed, chainID uint64) (*models.WithdrawalInformation, error) {
	result := new(models.WithdrawalInformation)

	result.WithdrawalHash = event.WithdrawalHash
	result.ChainID = chainID

	fnCalldata := event.Data

	if len(fnCalldata) == 0 {
		result.DecodedAction = string(models.EthTransferAction)
		result.FromAddress = &event.Sender
		result.ToAddress = &event.Target
		result.Amount = event.Value

		return result, nil
	}

	if len(fnCalldata) < 4 {
		return nil, ErrDecodeFailed
	}

	if bytes.Equal(fnCalldata[:4], d.IL2CrossDomainMessengerABI.Methods["relayMessage"].ID) {
		args, err := d.ParseRelayMessageInputArgs(fnCalldata[4:])
		if err != nil {
			return nil, ErrDecodeFailed
		}

		if len(args.Message) == 0 {
			result.DecodedAction = string(models.EthTransferAction)
			result.FromAddress = &args.Sender
			result.ToAddress = &args.Target
			result.Amount = args.Value

			return result, nil
		}

		if len(args.Message) < 4 {
			return nil, ErrDecodeFailed
		}

		fnCalldata = args.Message
	}

	// Here, fnCalldata can be == event.Data, which can be a direct call to finalizeBridge...(...)
	// Although operationally it won't be a case with the standard OP Stack withdrawal flow,
	// it is being allowed here for ABI flexibility.

	if bytes.Equal(fnCalldata[:4], d.IL2StandardBridgeABI.Methods["finalizeBridgeETH"].ID) {
		args, err := d.ParseFinalizeBridgeETHInputArgs(fnCalldata[4:])
		if err != nil {
			return nil, ErrDecodeFailed
		}

		result.DecodedAction = string(models.EthTransferAction)
		result.FromAddress = &args.From
		result.ToAddress = &args.To
		result.Amount = args.Amount

		return result, nil
	}

	if bytes.Equal(fnCalldata[:4], d.IL2StandardBridgeABI.Methods["finalizeBridgeERC20"].ID) {
		args, err := d.ParseFinalizeBridgeERC20InputArgs(fnCalldata[4:])
		if err != nil {
			return nil, ErrDecodeFailed
		}

		result.DecodedAction = string(models.ERC20TransferAction)
		result.TokenAddress = &args.RemoteToken // this is _localToken on the current chain
		result.FromAddress = &args.From
		result.ToAddress = &args.To
		result.Amount = args.Amount

		return result, nil
	}

	if bytes.Equal(fnCalldata[:4], d.IL2ERC721BridgeABI.Methods["finalizeBridgeERC721"].ID) {
		args, err := d.ParseFinalizeBridgeERC721InputArgs(fnCalldata[4:])
		if err != nil {
			return nil, ErrDecodeFailed
		}

		result.DecodedAction = string(models.ERC721TransferAction)
		result.TokenAddress = &args.RemoteToken // this is _localToken on the current chain
		result.FromAddress = &args.From
		result.ToAddress = &args.To
		result.TokenID = args.TokenID

		return result, nil
	}

	return nil, ErrDecodeFailed
}
