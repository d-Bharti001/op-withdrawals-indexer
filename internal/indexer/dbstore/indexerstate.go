package dbstore

import "context"

func (s *PostgresStore) SaveChainLatestSyncedBlockNumber(ctx context.Context, chainKey string, blockNumber uint64) error {
	return nil
}

func (s *PostgresStore) GetChainLatestSyncedBlockNumber(ctx context.Context, chainKey string) (uint64, error) {
	return 0, nil
}
