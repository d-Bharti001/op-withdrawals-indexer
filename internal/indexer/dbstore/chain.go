package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveChain(ctx context.Context, chain *models.Chain) error {
	query := `
		INSERT INTO chains (chain_id, chain_name, source_chain_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (chain_id)
		DO UPDATE SET
			chain_name      = EXCLUDED.chain_name,
			source_chain_id = EXCLUDED.source_chain_id
		WHERE
			chains.chain_name      IS DISTINCT FROM EXCLUDED.chain_name OR
			chains.source_chain_id IS DISTINCT FROM EXCLUDED.source_chain_id;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	chainRow := chain.ToDBRow()

	_, err := s.db.ExecContext(
		ctx,
		query,
		chainRow.ID,
		chainRow.Name,
		chainRow.SourceChainID,
	)

	return err
}
