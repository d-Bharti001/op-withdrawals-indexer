package dbstore

import (
	"context"
	"database/sql"
	"errors"

	"op-withdrawals-indexer/internal/database/dbtypes"
	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
)

func (s *PostgresStore) SaveWithdrawal(ctx context.Context, db DbTx, withdrawal *models.Withdrawal) error {
	query := `
		INSERT INTO withdrawals (
			withdrawal_hash,
			chain_id,
			withdrawal_nonce,
			withdrawal_sender,
			withdrawal_target,
			withdrawal_value,
			withdrawal_gas_limit,
			withdrawal_data,
			tx_hash,
			block_number,
			block_hash,
			block_timestamp
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		ON CONFLICT (chain_id, withdrawal_hash)
		DO UPDATE SET
			tx_hash         = EXCLUDED.tx_hash,
			block_number    = EXCLUDED.block_number,
			block_hash      = EXCLUDED.block_hash,
			block_timestamp = EXCLUDED.block_timestamp
		WHERE
			withdrawals.tx_hash         IS DISTINCT FROM EXCLUDED.tx_hash OR
			withdrawals.block_number    IS DISTINCT FROM EXCLUDED.block_number OR
			withdrawals.block_hash      IS DISTINCT FROM EXCLUDED.block_hash OR
			withdrawals.block_timestamp IS DISTINCT FROM EXCLUDED.block_timestamp;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	withdrawalRow := withdrawal.ToDBRow()

	_, err := db.ExecContext(
		ctx,
		query,
		withdrawalRow.Hash,
		withdrawalRow.ChainID,
		withdrawalRow.Nonce,
		withdrawalRow.Sender,
		withdrawalRow.Target,
		withdrawalRow.Value,
		withdrawalRow.GasLimit,
		withdrawalRow.Data,
		withdrawalRow.TxHash,
		withdrawalRow.BlockNumber,
		withdrawalRow.BlockHash,
		withdrawalRow.BlockTimestamp,
	)

	return err
}

func (s *PostgresStore) GetWithdrawal(ctx context.Context, db DbTx, chainID uint64, withdrawalHash common.Hash) (*models.Withdrawal, error) {
	query := `
		SELECT
			withdrawal_hash,
			chain_id,
			withdrawal_nonce,
			withdrawal_sender,
			withdrawal_target,
			withdrawal_value,
			withdrawal_gas_limit,
			withdrawal_data,
			tx_hash,
			block_number,
			block_hash,
			block_timestamp
		FROM
			withdrawals
		WHERE
			chain_id = $1 AND withdrawal_hash = $2;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	var withdrawalRow models.WithdrawalDBRow

	err := db.QueryRowContext(
		ctx,
		query,
		chainID,
		dbtypes.Hash(withdrawalHash),
	).Scan(
		&withdrawalRow.Hash,
		&withdrawalRow.ChainID,
		&withdrawalRow.Nonce,
		&withdrawalRow.Sender,
		&withdrawalRow.Target,
		&withdrawalRow.Value,
		&withdrawalRow.GasLimit,
		&withdrawalRow.Data,
		&withdrawalRow.TxHash,
		&withdrawalRow.BlockNumber,
		&withdrawalRow.BlockHash,
		&withdrawalRow.BlockTimestamp,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return withdrawalRow.ToDomainModel(), nil
}
