package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	id        uint64
	name      string
	rpcClient *ethclient.Client
}

func (c *Chain) ChainID() uint64 {
	return c.id
}

func (c *Chain) ChainName() string {
	return c.name
}

func (c *Chain) LatestBlockNumber(ctx context.Context) (uint64, error) {
	return c.rpcClient.BlockNumber(ctx)
}

func (c *Chain) CloseConnection() {
	c.rpcClient.Close()
}
