package indexer

import (
	"context"
	"time"

	"op-withdrawals-indexer/internal/indexer/blockchain"
	"op-withdrawals-indexer/internal/indexer/dbstore"
)

type L2Scanner struct {
	l2Provider blockchain.L2ChainProvider

	blockScanBatchSizeLimit uint64
	unstableBlocksDepth     uint64
	pollingInterval         time.Duration

	chainIndexerStateTableKey string
}

func (s *L2Scanner) Start(ctx context.Context, db dbstore.DBStoreProvider) error {
	// TODO: scanner logic
	// include: select { case: <-ctx.Done(): return, and then other cases }

	// ChatGPT:
	// defer wg.Done()

	// ticker := time.NewTicker(s.pollingInterval)
	// defer ticker.Stop()

	// for {
	// 	select {
	// 	case <-ctx.Done():
	// 		return
	// 	case <-ticker.C:
	// 		// scan logic here
	// 	}
	// }

	return ctx.Err()
}
