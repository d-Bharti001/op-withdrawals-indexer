package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveWithdrawalProvenTx(ctx context.Context, tx *models.WithdrawalProvenTx) error {
	query := `
		INSERT INTO withdrawal_proven_txs (
			withdrawal_hash,
			withdrawal_chain_id,
			proof_submitter,
			tx_hash,
			tx_chain_id,
			block_number,
			block_hash,
			block_timestamp
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (
			withdrawal_chain_id,
			withdrawal_hash,
			tx_chain_id,
			proof_submitter
		)
		DO UPDATE SET
			tx_hash         = EXCLUDED.tx_hash,
			block_number    = EXCLUDED.block_number,
			block_hash      = EXCLUDED.block_hash,
			block_timestamp = EXCLUDED.block_timestamp
		WHERE
			withdrawal_proven_txs.tx_hash         IS DISTINCT FROM EXCLUDED.tx_hash OR
			withdrawal_proven_txs.block_number    IS DISTINCT FROM EXCLUDED.block_number OR
			withdrawal_proven_txs.block_hash      IS DISTINCT FROM EXCLUDED.block_hash OR
			withdrawal_proven_txs.block_timestamp IS DISTINCT FROM EXCLUDED.block_timestamp;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	wdProvenTxRow := tx.ToDBRow()

	_, err := s.db.ExecContext(
		ctx,
		query,
		wdProvenTxRow.WithdrawalHash,
		wdProvenTxRow.WithdrawalChainID,
		wdProvenTxRow.ProofSubmitter,
		wdProvenTxRow.TxHash,
		wdProvenTxRow.TxChainID,
		wdProvenTxRow.BlockNumber,
		wdProvenTxRow.BlockHash,
		wdProvenTxRow.BlockTimestamp,
	)

	return err
}
