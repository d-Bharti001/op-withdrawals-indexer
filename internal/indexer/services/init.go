package indexer

import (
	"context"

	"op-withdrawals-indexer/internal/indexer/blockchain"
)

// TODO:
// insert into db if the chains are not present

func NewIndexer(ctx context.Context, cfg IndexerInitConfig) (*IndexerApp, error) {
	l1Chain, err := blockchain.NewL1Chain(ctx, blockchain.L1ChainInitConfig{
		ChainInitConfig: blockchain.ChainInitConfig{
			ID:                     cfg.L1ChainID,
			Name:                   cfg.L1ChainName,
			RPCUrl:                 cfg.L1RPCUrl,
			BlockConfirmationDepth: cfg.L1BlockConfirmationDepth,
		},
	})
	if err != nil {
		return nil, err
	}

	l2Chain, err := blockchain.NewL2Chain(ctx, blockchain.L2ChainInitConfig{
		ChainInitConfig: blockchain.ChainInitConfig{
			ID:                     cfg.L2ChainID,
			Name:                   cfg.L2ChainName,
			RPCUrl:                 cfg.L2RPCUrl,
			BlockConfirmationDepth: cfg.L2BlockConfirmationDepth,
		},
		L1RPCUrl:         cfg.L1RPCUrl,
		L1ChainID:        cfg.L1ChainID,
		SystemConfigAddr: cfg.SystemConfigAddr,
	})
	if err != nil {
		return nil, err
	}

	indexer := IndexerApp{
		l1Scanner: &L1Scanner{
			l1PollingInterval: DefaultL1PollingInterval,
			l1Provider:        l1Chain,
		},
		l2Scanner: &L2Scanner{
			l2PollingInterval: DefaultL2PollingInterval,
			l2Provider:        l2Chain,
		},
		// TODO:
		// postgres and redis
	}

	return &indexer, nil
}
