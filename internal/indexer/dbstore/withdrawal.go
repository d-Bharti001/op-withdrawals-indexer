package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveWithdrawal(ctx context.Context, withdrawal *models.Withdrawal) error {
	return nil
}
