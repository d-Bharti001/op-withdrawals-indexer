package indexer

import (
	"context"
	"time"

	"op-withdrawals-indexer/internal/database/models"
	"op-withdrawals-indexer/internal/indexer/blockchain"
	"op-withdrawals-indexer/internal/indexer/dbstore"
)

type IndexerApp struct {
	globalCtx context.Context

	dbProvider dbstore.DBStoreProvider

	l1Provider blockchain.L1ChainProvider
	l2Provider blockchain.L2ChainProvider

	l1PollingInterval time.Duration
	l2PollingInterval time.Duration
}

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

	db, err := dbstore.NewPostgresStore(
		cfg.DBConnStr,
		cfg.DBMaxOpenConn,
		cfg.DBMaxIdleConn,
		cfg.DBMaxIdleTime,
	)
	if err != nil {
		return nil, err
	}

	indexer := IndexerApp{
		globalCtx:         ctx,
		dbProvider:        db,
		l1Provider:        l1Chain,
		l2Provider:        l2Chain,
		l1PollingInterval: DefaultL1PollingInterval,
		l2PollingInterval: DefaultL2PollingInterval,
	}

	return &indexer, nil
}

func (app *IndexerApp) Start() error {
	if err := app.dbProvider.SaveL1Chain(
		app.globalCtx,
		&models.Chain{
			ID:            app.l1Provider.ID(),
			Name:          app.l1Provider.Name(),
			SourceChainID: nil,
		},
	); err != nil {
		return err
	}

	l2SourceChainID := app.l2Provider.L1ChainID()

	if err := app.dbProvider.SaveL2Chain(
		app.globalCtx,
		&models.Chain{
			ID:            app.l2Provider.ID(),
			Name:          app.l2Provider.Name(),
			SourceChainID: &l2SourceChainID,
		},
	); err != nil {
		return err
	}

	if err := app.StartL1Scan(); err != nil {
		return err
	}

	if err := app.StartL2Scan(); err != nil {
		return err
	}

	return nil
}
