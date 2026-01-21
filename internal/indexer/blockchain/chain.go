package blockchain

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	id                     uint64
	name                   string
	rpcClient              *ethclient.Client
	blockConfirmationDepth uint64
}

func (c *Chain) ID() uint64 {
	return c.id
}

func (c *Chain) Name() string {
	return c.name
}

func (c *Chain) BlockConfirmationDepth() uint64 {
	return c.blockConfirmationDepth
}

func (c *Chain) LatestBlockNumber(ctx context.Context) (uint64, error) {
	return c.rpcClient.BlockNumber(ctx)
}
