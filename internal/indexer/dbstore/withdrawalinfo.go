package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveWithdrawalInfo(ctx context.Context, withdrawalInfo *models.WithdrawalInformation) error {
	return nil
}
