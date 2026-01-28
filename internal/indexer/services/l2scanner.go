package indexer

import (
	"op-withdrawals-indexer/internal/indexer/blockchain"
	"time"
)

type L2Scanner struct {
	l2PollingInterval time.Duration
	l2Provider        blockchain.L2ChainProvider

	// TODO:
	// dbStore (interface for Postgresql)
	// memcache (interface for Redis)
}
