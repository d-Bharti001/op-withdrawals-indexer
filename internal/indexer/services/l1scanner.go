package indexer

import (
	"op-withdrawals-indexer/internal/indexer/blockchain"
	"time"
)

type L1Scanner struct {
	l1PollingInterval time.Duration
	l1Provider        blockchain.L1ChainProvider

	// TODO:
	// dbStore (interface for Postgresql)
	// memcache (interface for Redis)
}
