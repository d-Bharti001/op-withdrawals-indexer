package decoder

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ICrossDomainMessengerRelayMessageInputArgs struct {
	Nonce       *big.Int
	Sender      common.Address
	Target      common.Address
	Value       *big.Int
	MinGasLimit *big.Int
	Message     []byte
}

func (d *EventDecoder) ParseRelayMessageInputArgs(data []byte) (*ICrossDomainMessengerRelayMessageInputArgs, error) {
	args := make(map[string]any)

	err := d.IL2CrossDomainMessengerABI.Methods["relayMessage"].Inputs.UnpackIntoMap(args, data)
	if err != nil {
		return nil, err
	}

	result := new(ICrossDomainMessengerRelayMessageInputArgs)

	_nonce, ok := args["_nonce"]
	if !ok {
		return nil, errors.New("cannot find _nonce in relayMessage inputs")
	}
	_nonceVal, ok := _nonce.(*big.Int)
	if !ok {
		return nil, errors.New("cannot decode _nonce from relayMessage inputs")
	}

	_sender, ok := args["_sender"]
	if !ok {
		return nil, errors.New("cannot find _sender in relayMessage inputs")
	}
	_senderVal, ok := _sender.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _sender from relayMessage inputs")
	}

	_target, ok := args["_target"]
	if !ok {
		return nil, errors.New("cannot find _target in relayMessage inputs")
	}
	_targetVal, ok := _target.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _target from relayMessage inputs")
	}

	_value, ok := args["_value"]
	if !ok {
		return nil, errors.New("cannot find _value in relayMessage inputs")
	}
	_valueVal, ok := _value.(*big.Int)
	if !ok {
		return nil, errors.New("cannot decode _value from relayMessage inputs")
	}

	_minGasLimit, ok := args["_minGasLimit"]
	if !ok {
		return nil, errors.New("cannot find _minGasLimit in relayMessage inputs")
	}
	_minGasLimitVal, ok := _minGasLimit.(*big.Int)
	if !ok {
		return nil, errors.New("cannot decode _minGasLimit from relayMessage inputs")
	}

	_message, ok := args["_message"]
	if !ok {
		return nil, errors.New("cannot find _message in relayMessage inputs")
	}
	_messageVal, ok := _message.([]byte)
	if !ok {
		return nil, errors.New("cannot decode _message from relayMessage inputs")
	}

	result.Nonce = _nonceVal
	result.Sender = _senderVal
	result.Target = _targetVal
	result.Value = _valueVal
	result.MinGasLimit = _minGasLimitVal
	result.Message = _messageVal

	return result, nil
}
