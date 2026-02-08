package indexer

import (
	"context"
	"log"
	"time"

	"op-withdrawals-indexer/internal/contracts/bindings"
	"op-withdrawals-indexer/internal/database/models"
	"op-withdrawals-indexer/internal/decoder"
	"op-withdrawals-indexer/internal/indexer/blockchain"
	"op-withdrawals-indexer/internal/indexer/dbstore"
)

type L2Scanner struct {
	l2Provider blockchain.L2ChainProvider

	evtDecoder decoder.EventDecoder

	minBlockToConsider      uint64
	blockScanBatchSizeLimit uint64
	unstableBlocksDepth     uint64
	pollingInterval         time.Duration

	chainIndexerStateTableKey string
}

func (s *L2Scanner) Start(ctx context.Context, db dbstore.DBStoreProvider) error {
	timerPivot := time.Now()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		scanBlocksErr := s.scanBlocks(ctx, db)
		if scanBlocksErr != nil {
			log.Println("L2 Scanner error :::", scanBlocksErr)
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

func (s *L2Scanner) scanBlocks(ctx context.Context, db dbstore.DBStoreProvider) error {
	noScannedBlock := false
	latestScannedBlock, err := db.GetChainLatestScannedBlockNumber(ctx, s.chainIndexerStateTableKey)
	if err != nil {
		if err == dbstore.ErrNotFound {
			noScannedBlock = true
		} else {
			return err
		}
	}

	startBlock := s.minBlockToConsider // i.e. 0 for L2 chain

	if !noScannedBlock && latestScannedBlock >= s.unstableBlocksDepth {
		lowerBound := latestScannedBlock - s.unstableBlocksDepth
		startBlock = max(startBlock, lowerBound)
	}

	latestChainBlock, err := s.l2Provider.LatestBlockNumber(ctx)
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

		logs, err := s.l2Provider.MessagePassedLogs(ctx, startBlock, endBlock)
		if err != nil {
			log.Println("L2 Scanner error while fetching message passed logs :::", err)

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
			logRecordError = s.saveWithdrawalEvent(ctx, db, log)
			if logRecordError != nil {
				break
			}
		}
		if logRecordError != nil {
			log.Println("L2 Scanner error while saving withdrawal event :::", logRecordError)

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
			log.Println("L2 Scanner error while saving latest scanned block number :::", err)

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

		startBlock = endBlock + 1
	}

	return nil
}

func (s *L2Scanner) saveWithdrawalEvent(ctx context.Context, db dbstore.DBStoreProvider, log *bindings.IL2ToL1MessagePasserMessagePassed) error {

	// First, get saved withdrawal
	dbWithdrawal, err := db.GetWithdrawal(ctx, db.DB(), s.l2Provider.ChainID(), log.WithdrawalHash)
	if err != nil && err != dbstore.ErrNotFound {
		return err
	}

	withdrawalModel := models.Withdrawal{
		Hash:           log.WithdrawalHash,
		ChainID:        s.l2Provider.ChainID(),
		Sender:         log.Sender,
		Target:         log.Target,
		Data:           log.Data,
		Value:          log.Value,
		TxHash:         log.Raw.TxHash,
		BlockNumber:    log.Raw.BlockNumber,
		BlockHash:      log.Raw.BlockHash,
		BlockTimestamp: log.Raw.BlockTimestamp,
	}

	// If found, and any of (tx hash, block number, block hash, block timestamp) differ, then just save (i.e. update) the withdrawal, and return.
	if dbWithdrawal != nil {

		if dbWithdrawal.TxHash != log.Raw.TxHash ||
			dbWithdrawal.BlockNumber != log.Raw.BlockNumber ||
			dbWithdrawal.BlockHash != log.Raw.BlockHash ||
			dbWithdrawal.BlockTimestamp != log.Raw.BlockTimestamp {

			return db.SaveWithdrawal(ctx, db.DB(), &withdrawalModel)
		}

		return nil
	}

	// If not found, proceed for all the steps:

	// 1. Decode the event with evtDecoder.DecodeL2MessagePassedEvent
	decodedInfo, err := s.evtDecoder.DecodeL2MessagePassedEvent(log, s.l2Provider.ChainID())
	if err != nil && err != decoder.ErrDecodeFailed {
		return err
	}

	if decodedInfo == nil {
		return db.SaveWithdrawal(ctx, db.DB(), &withdrawalModel)
	}

	// 2. If applicable, fetch token details from database. If not found, fetch details from blockchain.

	var tokenToSave *models.Token

	if decodedInfo.TokenAddress != nil {
		dbToken, err := db.GetToken(ctx, db.DB(), s.l2Provider.ChainID(), *decodedInfo.TokenAddress)
		if err != nil && err != dbstore.ErrNotFound {
			return err
		}

		if dbToken == nil {
			var tokenClass models.TokenClass
			switch decodedInfo.DecodedAction {
			case string(models.ERC20TransferAction):
				tokenClass = models.ERC20Token
			case string(models.ERC721TransferAction):
				tokenClass = models.ERC721Token
			case string(models.ERC1155TransferAction):
				tokenClass = models.ERC1155Token
			}

			// tokenToSave must be nil in case of error
			// error is ignored after the second try, skipping saving this token
			tokenToSave, err = s.l2Provider.TokenInfo(ctx, *decodedInfo.TokenAddress, tokenClass)
			if err != nil {
				tokenToSave, _ = s.l2Provider.TokenInfo(ctx, *decodedInfo.TokenAddress, tokenClass)
			}
		}
	}

	// 3. In a transaction, save withdrawal, withdrawal info (if available), token details (if available)
	return db.WithTx(ctx, func(tx dbstore.DbTx) error {

		err := db.SaveWithdrawal(ctx, tx, &withdrawalModel)
		if err != nil {
			return err
		}

		if decodedInfo != nil {
			err := db.SaveWithdrawalInfo(ctx, tx, decodedInfo)
			if err != nil {
				return err
			}
		}

		if tokenToSave != nil {
			err := db.SaveToken(ctx, tx, tokenToSave)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
