package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveWithdrawalFinalizedTx(ctx context.Context, tx *models.WithdrawalFinalizedTx) error {
	query := `
		INSERT INTO withdrawal_finalized_txs (
			withdrawal_hash,
			withdrawal_chain_id,
			success,
			tx_hash,
			tx_chain_id,
			tx_caller,
			block_number,
			block_hash,
			block_timestamp
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (
			withdrawal_chain_id,
			withdrawal_hash,
			tx_chain_id,
			tx_hash
		)
		DO UPDATE SET
			success			= EXCLUDED.success,
			tx_caller       = EXCLUDED.tx_caller,
			block_number    = EXCLUDED.block_number,
			block_hash      = EXCLUDED.block_hash,
			block_timestamp = EXCLUDED.block_timestamp
		WHERE
			withdrawal_finalized_txs.success         IS DISTINCT FROM EXCLUDED.success OR
			withdrawal_finalized_txs.tx_caller       IS DISTINCT FROM EXCLUDED.tx_caller OR
			withdrawal_finalized_txs.block_number    IS DISTINCT FROM EXCLUDED.block_number OR
			withdrawal_finalized_txs.block_hash      IS DISTINCT FROM EXCLUDED.block_hash OR
			withdrawal_finalized_txs.block_timestamp IS DISTINCT FROM EXCLUDED.block_timestamp;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	wdFinalizedTxRow := tx.ToDBRow()

	_, err := s.db.ExecContext(
		ctx,
		query,
		wdFinalizedTxRow.WithdrawalHash,
		wdFinalizedTxRow.WithdrawalChainID,
		wdFinalizedTxRow.Success,
		wdFinalizedTxRow.TxHash,
		wdFinalizedTxRow.TxChainID,
		wdFinalizedTxRow.TxCaller,
		wdFinalizedTxRow.BlockNumber,
		wdFinalizedTxRow.BlockHash,
		wdFinalizedTxRow.BlockTimestamp,
	)

	return err
}
