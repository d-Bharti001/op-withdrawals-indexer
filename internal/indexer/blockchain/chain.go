package blockchain

import (
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	ID                     uint64
	Name                   string
	RPCClient              *ethclient.Client
	BlockConfirmationDepth uint64
	PollingInterval        time.Duration
}

type L2ChainOnL1 struct {
	L2Chain
	StartBlock  uint64
	L1Contracts L1Contracts
}
