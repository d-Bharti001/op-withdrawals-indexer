package indexer

import (
	"context"
	"sync"
	"time"

	"op-withdrawals-indexer/internal/indexer/blockchain"
	"op-withdrawals-indexer/internal/indexer/dbstore"
)

type L2Scanner struct {
	l2Provider      blockchain.L2ChainProvider
	pollingInterval time.Duration

	chainIndexerStateTableKey string
}

func (s *L2Scanner) Start(ctx context.Context, wg *sync.WaitGroup, db dbstore.DBStoreProvider) {
	defer wg.Done()

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
}
