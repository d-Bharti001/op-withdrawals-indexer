package blockchain

import (
	"context"

	"op-withdrawals-indexer/internal/contracts/bindings"
	"op-withdrawals-indexer/internal/database/models"

	bind2 "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Chain) TokenInfo(ctx context.Context, tokenAddr common.Address, typeHint models.TokenClass) (*models.Token, error) {
	ctx, cancel := context.WithTimeout(ctx, DefaultRpcTimeout)
	defer cancel()

	result := new(models.Token)

	result.Address = tokenAddr
	result.ChainID = c.ChainID()
	result.Class = string(typeHint)

	contract := bindings.NewIERC20Metadata()
	instance := contract.Instance(c.rpcClient, tokenAddr)

	name, err := bind2.Call(instance, nil, contract.PackName(), contract.UnpackName)
	if err != nil {
		return nil, err
	}

	symbol, err := bind2.Call(instance, nil, contract.PackSymbol(), contract.UnpackSymbol)
	if err != nil {
		return nil, err
	}

	switch typeHint {
	case models.ERC20Token:
		decimals, err := bind2.Call(instance, nil, contract.PackDecimals(), contract.UnpackDecimals)
		if err != nil {
			return nil, err
		}
		result.Decimals = uint64(decimals)
	}

	result.Name = name
	result.Symbol = symbol

	return result, nil
}
