package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
)

func (s *PostgresStore) SaveToken(ctx context.Context, token *models.Token) error {
	return nil
}

func (s *PostgresStore) GetToken(ctx context.Context, chainId uint64, tokenAddr common.Address) (*models.Token, error) {
	return nil, nil
}
