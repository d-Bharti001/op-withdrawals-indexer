package decoder

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IStandardBridgeFinalizeBridgeETHInputArgs struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
}

type IStandardBridgeFinalizeBridgeERC20InputArgs struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
}

func (d *EventDecoder) ParseFinalizeBridgeETHInputArgs(data []byte) (*IStandardBridgeFinalizeBridgeETHInputArgs, error) {
	args := make(map[string]any)

	err := d.IL2StandardBridgeABI.Methods["finalizeBridgeETH"].Inputs.UnpackIntoMap(args, data)
	if err != nil {
		return nil, err
	}

	result := new(IStandardBridgeFinalizeBridgeETHInputArgs)

	_from, ok := args["_from"]
	if !ok {
		return nil, errors.New("cannot find _from in finalizeBridgeETH inputs")
	}
	_fromVal, ok := _from.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _from from finalizeBridgeETH inputs")
	}

	_to, ok := args["_to"]
	if !ok {
		return nil, errors.New("cannot find _to in finalizeBridgeETH inputs")
	}
	_toVal, ok := _to.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _to from finalizeBridgeETH inputs")
	}

	_amount, ok := args["_amount"]
	if !ok {
		return nil, errors.New("cannot find _amount in finalizeBridgeETH inputs")
	}
	_amountVal, ok := _amount.(*big.Int)
	if !ok {
		return nil, errors.New("cannot decode _amount from finalizeBridgeETH inputs")
	}

	_extraData, ok := args["_extraData"]
	if !ok {
		return nil, errors.New("cannot find _extraData in finalizeBridgeETH inputs")
	}
	_extraDataVal, ok := _extraData.([]byte)
	if !ok {
		return nil, errors.New("cannot decode _extraData from finalizeBridgeETH inputs")
	}

	result.From = _fromVal
	result.To = _toVal
	result.Amount = _amountVal
	result.ExtraData = _extraDataVal

	return result, nil
}

func (d *EventDecoder) ParseFinalizeBridgeERC20InputArgs(data []byte) (*IStandardBridgeFinalizeBridgeERC20InputArgs, error) {
	args := make(map[string]any)

	err := d.IL2StandardBridgeABI.Methods["finalizeBridgeERC20"].Inputs.UnpackIntoMap(args, data)
	if err != nil {
		return nil, err
	}

	result := new(IStandardBridgeFinalizeBridgeERC20InputArgs)

	_localToken, ok := args["_localToken"]
	if !ok {
		return nil, errors.New("cannot find _localToken in finalizeBridgeERC20 inputs")
	}
	_localTokenVal, ok := _localToken.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _localToken from finalizeBridgeERC20 inputs")
	}

	_remoteToken, ok := args["_remoteToken"]
	if !ok {
		return nil, errors.New("cannot find _remoteToken in finalizeBridgeERC20 inputs")
	}
	_remoteTokenVal, ok := _remoteToken.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _remoteToken from finalizeBridgeERC20 inputs")
	}

	_from, ok := args["_from"]
	if !ok {
		return nil, errors.New("cannot find _from in finalizeBridgeERC20 inputs")
	}
	_fromVal, ok := _from.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _from from finalizeBridgeERC20 inputs")
	}

	_to, ok := args["_to"]
	if !ok {
		return nil, errors.New("cannot find _to in finalizeBridgeERC20 inputs")
	}
	_toVal, ok := _to.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _to from finalizeBridgeERC20 inputs")
	}

	_amount, ok := args["_amount"]
	if !ok {
		return nil, errors.New("cannot find _amount in finalizeBridgeERC20 inputs")
	}
	_amountVal, ok := _amount.(*big.Int)
	if !ok {
		return nil, errors.New("cannot decode _amount from finalizeBridgeERC20 inputs")
	}

	_extraData, ok := args["_extraData"]
	if !ok {
		return nil, errors.New("cannot find _extraData in finalizeBridgeERC20 inputs")
	}
	_extraDataVal, ok := _extraData.([]byte)
	if !ok {
		return nil, errors.New("cannot decode _extraData from finalizeBridgeERC20 inputs")
	}

	result.LocalToken = _localTokenVal
	result.RemoteToken = _remoteTokenVal
	result.From = _fromVal
	result.To = _toVal
	result.Amount = _amountVal
	result.ExtraData = _extraDataVal

	return result, nil
}
