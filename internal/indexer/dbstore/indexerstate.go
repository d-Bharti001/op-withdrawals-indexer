package dbstore

import (
	"context"
	"database/sql"
	"errors"
)

func (s *PostgresStore) SaveChainLatestSyncedBlockNumber(ctx context.Context, chainKey string, blockNumber uint64) error {
	query := `
		INSERT INTO chain_indexer_states (chain_key, latest_synced_block)
		VALUES ($1, $2)
		ON CONFLICT (chain_key)
		DO UPDATE SET
			latest_synced_block = EXCLUDED.latest_synced_block
		WHERE
			chain_indexer_states.latest_synced_block < EXCLUDED.latest_synced_block;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	_, err := s.db.ExecContext(
		ctx,
		query,
		chainKey,
		blockNumber,
	)

	return err
}

func (s *PostgresStore) GetChainLatestSyncedBlockNumber(ctx context.Context, chainKey string) (uint64, error) {
	query := `
		SELECT latest_synced_block
		FROM chain_indexer_states
		WHERE chain_key = $1;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	var resultBlockNumber uint64

	err := s.db.QueryRowContext(
		ctx,
		query,
		chainKey,
	).Scan(&resultBlockNumber)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNotFound
		}
		return 0, err
	}

	return resultBlockNumber, nil
}
