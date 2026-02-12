package indexer

import (
	"context"
	"log"
	"time"

	"op-withdrawals-indexer/internal/database/models"
	"op-withdrawals-indexer/internal/indexer/blockchain"
	"op-withdrawals-indexer/internal/indexer/dbstore"

	"github.com/ethereum/go-ethereum/common"
)

type L1Scanner struct {
	l1Provider blockchain.L1ChainProvider

	portalAddresses        []common.Address
	portalAddressToChainID map[common.Address]uint64

	minBlockToConsider      uint64
	blockScanBatchSizeLimit uint64
	unstableBlocksDepth     uint64
	pollingInterval         time.Duration

	chainIndexerStateTableKey string
}

func (s *L1Scanner) Start(ctx context.Context, db dbstore.DBStoreProvider) error {
	timerPivot := time.Now()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		scanBlocksErr := s.scanBlocks(ctx, db)
		if scanBlocksErr != nil {
			log.Println("L1 Scanner error :::", scanBlocksErr)
		}

		timeElapsed := time.Since(timerPivot)

		if timeElapsed > s.pollingInterval {
			fullIntervals := timeElapsed / s.pollingInterval
			timerPivot = timerPivot.Add(fullIntervals * s.pollingInterval)
		} else if scanBlocksErr == nil {
			timerPivot = timerPivot.Add(s.pollingInterval)
		}

		sleep := time.Until(timerPivot)
		if scanBlocksErr != nil {
			sleep = 200 * time.Millisecond
		}

		if sleep > 0 {
			timer := time.NewTimer(sleep)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
			}
		}
	}
}

func (s *L1Scanner) scanBlocks(ctx context.Context, db dbstore.DBStoreProvider) error {
	noScannedBlock := false
	latestScannedBlock, err := db.GetChainLatestScannedBlockNumber(ctx, s.chainIndexerStateTableKey)
	if err != nil {
		if err == dbstore.ErrNotFound {
			noScannedBlock = true
		} else {
			return err
		}
	}

	startBlock := s.minBlockToConsider

	if !noScannedBlock && latestScannedBlock >= s.unstableBlocksDepth {
		lowerBound := latestScannedBlock - s.unstableBlocksDepth
		startBlock = max(startBlock, lowerBound)
	}

	latestChainBlock, err := s.l1Provider.LatestBlockNumber(ctx)
	if err != nil {
		return err
	}

	for {
		if startBlock > latestChainBlock {
			break
		}

		endBlock := min(startBlock+s.blockScanBatchSizeLimit-1, latestChainBlock)
		if endBlock < startBlock {
			endBlock = startBlock
		}

		logs, err := s.l1Provider.PortalWithdrawalLogs(ctx, s.portalAddresses, startBlock, endBlock)
		if err != nil {
			log.Println("L1 Scanner error while fetching portal withdrawal logs :::", err)

			// Re-iterate after some time
			timer := time.NewTimer(50 * time.Millisecond)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
				continue
			}
		}

		// Process fetched logs
		var logRecordError error
		for _, log := range logs {
			logRecordError = s.savePortalEvent(ctx, db, log)
			if logRecordError != nil {
				break
			}
		}
		if logRecordError != nil {
			log.Println("L1 Scanner error while saving portal withdrawal event :::", logRecordError)

			// Re-iterate after some time
			timer := time.NewTimer(50 * time.Millisecond)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
				continue
			}
		}

		err = db.SaveChainLatestScannedBlockNumber(ctx, s.chainIndexerStateTableKey, endBlock)
		if err != nil {
			log.Println("L1 Scanner error while saving latest scanned block number :::", err)

			// Re-iterate after some time
			timer := time.NewTimer(20 * time.Millisecond)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
				continue
			}
		}

		log.Printf("L1 Scanner [chain id %d]: scanned blocks [%d..%d]\n", s.l1Provider.ChainID(), startBlock, endBlock)

		startBlock = endBlock + 1
	}

	return nil
}

func (s *L1Scanner) savePortalEvent(
	ctx context.Context,
	db dbstore.DBStoreProvider,
	log blockchain.OptimismPortalRelevantWithdrawalEvent,
) error {
	switch event := log.(type) {
	case *blockchain.WithdrawalProvenEvent:
		l2ChainID, ok := s.portalAddressToChainID[event.Raw.Address]
		if !ok {
			return nil
		}

		return db.SaveWithdrawalProvenTx(ctx,
			&models.WithdrawalProvenTx{
				WithdrawalHash:    event.WithdrawalHash,
				WithdrawalChainID: l2ChainID,
				ProofSubmitter:    event.ProofSubmitter,
				TxHash:            event.Raw.TxHash,
				TxChainID:         s.l1Provider.ChainID(),
				BlockNumber:       event.Raw.BlockNumber,
				BlockHash:         event.Raw.BlockHash,
				BlockTimestamp:    event.Raw.BlockTimestamp,
			})

	case *blockchain.WithdrawalFinalizedEvent:
		l2ChainID, ok := s.portalAddressToChainID[event.Raw.Address]
		if !ok {
			return nil
		}

		return db.SaveWithdrawalFinalizedTx(ctx,
			&models.WithdrawalFinalizedTx{
				WithdrawalHash:    event.WithdrawalHash,
				WithdrawalChainID: l2ChainID,
				Success:           event.Success,
				TxHash:            event.Raw.TxHash,
				TxChainID:         s.l1Provider.ChainID(),
				BlockNumber:       event.Raw.BlockNumber,
				BlockHash:         event.Raw.BlockHash,
				BlockTimestamp:    event.Raw.BlockTimestamp,
			})

	default:
		return nil
	}
}
