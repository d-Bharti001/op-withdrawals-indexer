package indexer

import (
	"context"
	"fmt"
	"log"
	"sync"

	"op-withdrawals-indexer/internal/database/models"
	"op-withdrawals-indexer/internal/decoder"
	"op-withdrawals-indexer/internal/indexer/blockchain"
	"op-withdrawals-indexer/internal/indexer/dbstore"

	"github.com/ethereum/go-ethereum/common"
)

type IndexerApp struct {
	ctx    context.Context
	cancel context.CancelFunc

	wg sync.WaitGroup

	stopOnce sync.Once

	l1Scanner L1Scanner
	l2Scanner L2Scanner

	dbProvider dbstore.DBStoreProvider
}

func NewIndexer(ctx context.Context, cfg IndexerInitConfig) (*IndexerApp, error) {
	ctx, cancel := context.WithCancel(ctx)

	l1Chain, err := blockchain.NewL1Chain(ctx, blockchain.L1ChainInitConfig{
		ChainInitConfig: blockchain.ChainInitConfig{
			ID:     cfg.L1ChainID,
			Name:   cfg.L1ChainName,
			RPCUrl: cfg.L1RPCUrl,
		},
	})
	if err != nil {
		cancel()
		return nil, err
	}

	l2Chain, err := blockchain.NewL2Chain(ctx, blockchain.L2ChainInitConfig{
		ChainInitConfig: blockchain.ChainInitConfig{
			ID:     cfg.L2ChainID,
			Name:   cfg.L2ChainName,
			RPCUrl: cfg.L2RPCUrl,
		},
		L1RPCUrl:         cfg.L1RPCUrl,
		L1ChainID:        cfg.L1ChainID,
		SystemConfigAddr: cfg.SystemConfigAddr,
	})
	if err != nil {
		cancel()
		l1Chain.CloseConnection()
		return nil, err
	}

	db, err := dbstore.NewPostgresStore(
		cfg.DBConnStr,
		cfg.DBMaxOpenConn,
		cfg.DBMaxIdleConn,
		cfg.DBMaxIdleTime,
	)
	if err != nil {
		cancel()
		l1Chain.CloseConnection()
		l2Chain.CloseConnection()
		return nil, err
	}

	evtDecoder, err := decoder.NewEventDecoder()
	if err != nil {
		cancel()
		l1Chain.CloseConnection()
		l2Chain.CloseConnection()
		db.CloseConnection()
		return nil, err
	}

	portalAddresses := []common.Address{
		l2Chain.OptimismPortalAddr(),
	}
	portalAddressToChainID := map[common.Address]uint64{
		l2Chain.OptimismPortalAddr(): l2Chain.ChainID(),
	}
	minL1BlockNumber := l2Chain.L1StartBlockNumber()

	indexer := IndexerApp{
		ctx:        ctx,
		cancel:     cancel,
		dbProvider: db,
		l1Scanner: L1Scanner{
			l1Provider:                l1Chain,
			portalAddresses:           portalAddresses,
			portalAddressToChainID:    portalAddressToChainID,
			minBlockToConsider:        minL1BlockNumber,
			blockScanBatchSizeLimit:   cfg.L1BlockScanBatchSizeLimit,
			unstableBlocksDepth:       cfg.L1UnstableBlocksDepth,
			pollingInterval:           DefaultL1PollingInterval,
			chainIndexerStateTableKey: fmt.Sprintf("%d:%d", cfg.L2ChainID, cfg.L1ChainID),
		},
		l2Scanner: L2Scanner{
			l2Provider:                l2Chain,
			evtDecoder:                *evtDecoder,
			blockScanBatchSizeLimit:   cfg.L2BlockScanBatchSizeLimit,
			unstableBlocksDepth:       cfg.L2UnstableBlocksDepth,
			pollingInterval:           DefaultL2PollingInterval,
			chainIndexerStateTableKey: fmt.Sprintf("%d", cfg.L2ChainID),
		},
	}

	return &indexer, nil
}

func (app *IndexerApp) Start() error {
	log.Println("Starting indexer...")

	// Save chains into database

	if err := app.dbProvider.SaveL1Chain(
		app.ctx,
		&models.Chain{
			ID:            app.l1Scanner.l1Provider.ChainID(),
			Name:          app.l1Scanner.l1Provider.ChainName(),
			SourceChainID: nil,
		},
	); err != nil {
		return err
	}

	l2SourceChainID := app.l2Scanner.l2Provider.L1ChainID()

	if err := app.dbProvider.SaveL2Chain(
		app.ctx,
		&models.Chain{
			ID:            app.l2Scanner.l2Provider.ChainID(),
			Name:          app.l2Scanner.l2Provider.ChainName(),
			SourceChainID: &l2SourceChainID,
		},
	); err != nil {
		return err
	}

	// Start scanners

	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		log.Println("Starting L1 Scanner...")
		app.l1Scanner.Start(app.ctx, app.dbProvider)
	}()

	app.wg.Add(1)
	go func() {
		defer app.wg.Done()
		log.Println("Starting L2 Scanner...")
		app.l2Scanner.Start(app.ctx, app.dbProvider)
	}()

	return nil
}

func (app *IndexerApp) Stop() {
	app.stopOnce.Do(func() {
		log.Println("Stopping indexer...")

		ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
		defer cancel()

		// Signal shutdown
		if app.cancel != nil {
			app.cancel()
		}

		log.Println("Stopping scanner services...")

		// Wait for scanners or timeout
		done := make(chan struct{})

		go func() {
			app.wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			log.Println("Scanners stopped cleanly")
		case <-ctx.Done():
			log.Println("Shutdown timeout exceeded:", ctx.Err())
		}

		log.Println("Closing connections...")

		app.l1Scanner.l1Provider.CloseConnection()
		app.l2Scanner.l2Provider.CloseConnection()
		app.dbProvider.CloseConnection()

		log.Println("Indexer stopped.")
	})
}
