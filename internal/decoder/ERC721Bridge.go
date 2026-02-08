package decoder

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IERC721BridgeFinalizeBridgeERC721InputArgs struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	TokenID     *big.Int
	ExtraData   []byte
}

func (d *EventDecoder) ParseFinalizeBridgeERC721InputArgs(data []byte) (*IERC721BridgeFinalizeBridgeERC721InputArgs, error) {
	args := make(map[string]any)

	err := d.IL2ERC721BridgeABI.Methods["finalizeBridgeERC721"].Inputs.UnpackIntoMap(args, data)
	if err != nil {
		return nil, err
	}

	result := new(IERC721BridgeFinalizeBridgeERC721InputArgs)

	_localToken, ok := args["_localToken"]
	if !ok {
		return nil, errors.New("cannot find _localToken in finalizeBridgeERC721 inputs")
	}
	_localTokenVal, ok := _localToken.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _localToken from finalizeBridgeERC721 inputs")
	}

	_remoteToken, ok := args["_remoteToken"]
	if !ok {
		return nil, errors.New("cannot find _remoteToken in finalizeBridgeERC721 inputs")
	}
	_remoteTokenVal, ok := _remoteToken.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _remoteToken from finalizeBridgeERC721 inputs")
	}

	_from, ok := args["_from"]
	if !ok {
		return nil, errors.New("cannot find _from in finalizeBridgeERC721 inputs")
	}
	_fromVal, ok := _from.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _from from finalizeBridgeERC721 inputs")
	}

	_to, ok := args["_to"]
	if !ok {
		return nil, errors.New("cannot find _to in finalizeBridgeERC721 inputs")
	}
	_toVal, ok := _to.(common.Address)
	if !ok {
		return nil, errors.New("cannot decode _to from finalizeBridgeERC721 inputs")
	}

	_tokenId, ok := args["_tokenId"]
	if !ok {
		return nil, errors.New("cannot find _tokenId in finalizeBridgeERC721 inputs")
	}
	_tokenIdVal, ok := _tokenId.(*big.Int)
	if !ok {
		return nil, errors.New("cannot decode _tokenId from finalizeBridgeERC721 inputs")
	}

	_extraData, ok := args["_extraData"]
	if !ok {
		return nil, errors.New("cannot find _extraData in finalizeBridgeERC721 inputs")
	}
	_extraDataVal, ok := _extraData.([]byte)
	if !ok {
		return nil, errors.New("cannot decode _extraData from finalizeBridgeERC721 inputs")
	}

	result.LocalToken = _localTokenVal
	result.RemoteToken = _remoteTokenVal
	result.From = _fromVal
	result.To = _toVal
	result.TokenID = _tokenIdVal
	result.ExtraData = _extraDataVal

	return result, nil
}
