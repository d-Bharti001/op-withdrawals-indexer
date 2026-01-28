package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveChain(ctx context.Context, chain *models.Chain) (bool, error) {
	chainDBRow := chain.ToDBRow()

	query := `
		WITH existing AS (
			SELECT 1
			FROM chains
			WHERE chain_id = $1
		),
		upsert AS (
			INSERT INTO chains (chain_id, chain_name, source_chain_id)
			VALUES ($1, $2, $3)
			ON CONFLICT (chain_id)
			DO UPDATE SET
				chain_name = EXCLUDED.chain_name,
				source_chain_id = EXCLUDED.source_chain_id
			WHERE
				chains.chain_name IS DISTINCT FROM EXCLUDED.chain_name OR
				chains.source_chain_id IS DISTINCT FROM EXCLUDED.source_chain_id
			RETURNING 1
		)
		SELECT
			EXISTS (SELECT 1 FROM existing) AS already_present;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	var alreadyPresent bool

	err := s.db.QueryRowContext(
		ctx,
		query,
		chainDBRow.ID,
		chainDBRow.Name,
		chainDBRow.SourceChainID,
	).Scan(&alreadyPresent)

	return alreadyPresent, err
}
