package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
)

type DBStoreProvider interface {
	// SaveL1Chain should not update source_chain_id
	SaveL1Chain(context.Context, *models.Chain) error

	// SaveL2Chain should update source_chain_id if the stored value is different
	SaveL2Chain(context.Context, *models.Chain) error

	GetWithdrawal(ctx context.Context, chainID uint64, withdrawalHash common.Hash) (*models.Withdrawal, error)
	SaveWithdrawal(context.Context, *models.Withdrawal) error

	SaveWithdrawalProvenTx(context.Context, *models.WithdrawalProvenTx) error
	SaveWithdrawalFinalizedTx(context.Context, *models.WithdrawalFinalizedTx) error

	SaveWithdrawalInfo(context.Context, *models.WithdrawalInformation) error

	GetToken(
		ctx context.Context,
		chainId uint64,
		tokenAddr common.Address,
	) (*models.Token, error)
	SaveToken(context.Context, *models.Token) error

	SaveChainLatestSyncedBlockNumber(ctx context.Context, chainKey string, blockNumber uint64) error
	GetChainLatestSyncedBlockNumber(ctx context.Context, chainKey string) (blockNumber uint64, err error)

	CloseConnection() error
}
