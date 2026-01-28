package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveWithdrawalFinalizedTx(ctx context.Context, tx *models.WithdrawalFinalizedTx) (bool, error) {
	return false, nil
}
