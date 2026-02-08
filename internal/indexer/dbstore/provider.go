package dbstore

import (
	"context"
	"database/sql"
	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
)

type DBStoreProvider interface {
	// SaveL1Chain should not update source_chain_id
	SaveL1Chain(context.Context, *models.Chain) error

	// SaveL2Chain should update source_chain_id if the stored value is different
	SaveL2Chain(context.Context, *models.Chain) error

	GetWithdrawal(ctx context.Context, db DbTx, chainID uint64, withdrawalHash common.Hash) (*models.Withdrawal, error)
	SaveWithdrawal(context.Context, DbTx, *models.Withdrawal) error

	SaveWithdrawalProvenTx(context.Context, *models.WithdrawalProvenTx) error
	SaveWithdrawalFinalizedTx(context.Context, *models.WithdrawalFinalizedTx) error

	SaveWithdrawalInfo(context.Context, DbTx, *models.WithdrawalInformation) error

	GetToken(
		ctx context.Context,
		db DbTx,
		chainId uint64,
		tokenAddr common.Address,
	) (*models.Token, error)
	SaveToken(context.Context, DbTx, *models.Token) error

	SaveChainLatestScannedBlockNumber(ctx context.Context, chainKey string, blockNumber uint64) error
	GetChainLatestScannedBlockNumber(ctx context.Context, chainKey string) (blockNumber uint64, err error)

	DB() DbTx
	WithTx(ctx context.Context, fn func(tx DbTx) error) error

	CloseConnection() error
}

type DbTx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
