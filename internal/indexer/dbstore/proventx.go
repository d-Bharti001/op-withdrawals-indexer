package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveWithdrawalProvenTx(ctx context.Context, tx *models.WithdrawalProvenTx) error {
	return nil
}
