package dbstore

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
)

func (s *PostgresStore) SaveWithdrawal(ctx context.Context, withdrawal *models.Withdrawal) error {
	query := `
		INSERT INTO withdrawals (
			withdrawal_hash,
			chain_id,
			withdrawal_sender,
			withdrawal_target,
			withdrawal_data,
			withdrawal_value,
			tx_hash,
			tx_caller,
			block_number,
			block_hash,
			block_timestamp
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT (chain_id, withdrawal_hash)
		DO UPDATE SET
			tx_hash         = EXCLUDED.tx_hash,
			tx_caller       = EXCLUDED.tx_caller,
			block_number    = EXCLUDED.block_number,
			block_hash      = EXCLUDED.block_hash,
			block_timestamp = EXCLUDED.block_timestamp
		WHERE
			withdrawals.tx_hash         IS DISTINCT FROM EXCLUDED.tx_hash OR
			withdrawals.tx_caller       IS DISTINCT FROM EXCLUDED.tx_caller OR
			withdrawals.block_number    IS DISTINCT FROM EXCLUDED.block_number OR
			withdrawals.block_hash      IS DISTINCT FROM EXCLUDED.block_hash OR
			withdrawals.block_timestamp IS DISTINCT FROM EXCLUDED.block_timestamp;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	withdrawalRow := withdrawal.ToDBRow()

	_, err := s.db.ExecContext(
		ctx,
		query,
		withdrawalRow.Hash,
		withdrawalRow.ChainID,
		withdrawalRow.Sender,
		withdrawalRow.Target,
		withdrawalRow.Data,
		withdrawalRow.Value,
		withdrawalRow.TxHash,
		withdrawalRow.TxCaller,
		withdrawalRow.BlockNumber,
		withdrawalRow.BlockHash,
		withdrawalRow.BlockTimestamp,
	)

	return err
}

func (s *PostgresStore) GetWithdrawal(ctx context.Context, chainID uint64, withdrawalHash common.Hash) (*models.Withdrawal, error) {
	query := `
		SELECT
			withdrawal_hash,
			chain_id,
			withdrawal_sender,
			withdrawal_target,
			withdrawal_data,
			withdrawal_value,
			tx_hash,
			tx_caller,
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

	err := s.db.QueryRowContext(
		ctx,
		query,
		chainID,
		strings.ToLower(withdrawalHash.Hex()),
	).Scan(
		&withdrawalRow.Hash,
		&withdrawalRow.ChainID,
		&withdrawalRow.Sender,
		&withdrawalRow.Target,
		&withdrawalRow.Data,
		&withdrawalRow.Value,
		&withdrawalRow.TxHash,
		&withdrawalRow.TxCaller,
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
