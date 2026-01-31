package chainclient

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type DialChainParams struct {
	Ctx     context.Context
	URL     string
	ChainID uint64
	Name    string
}

func DialChain(params DialChainParams) (client *ethclient.Client, rpcChainId uint64, err error) {
	cl, err := ethclient.DialContext(params.Ctx, params.URL)
	if err != nil {
		return nil, 0, err
	}

	chainId, err := cl.ChainID(params.Ctx)
	if err != nil {
		cl.Close()
		return nil, 0, err
	}
	chainIdUint64 := chainId.Uint64()
	if chainIdUint64 != params.ChainID {
		cl.Close()
		return nil, 0, fmt.Errorf("chain id mismatch for chain: %v; expected %v, got %v", params.Name, params.ChainID, chainIdUint64)
	}

	return cl, chainIdUint64, nil
}
