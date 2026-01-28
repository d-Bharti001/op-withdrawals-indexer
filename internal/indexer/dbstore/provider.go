package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
)

type DBStoreProvider interface {
	SaveChain(context.Context, *models.Chain) (alreadyPresent bool, err error)

	SaveWithdrawal(context.Context, *models.Withdrawal) (alreadyPresent bool, err error)
	SaveWithdrawalProvenTx(context.Context, *models.WithdrawalProvenTx) (alreadyPresent bool, err error)
	SaveWithdrawalFinalizedTx(context.Context, *models.WithdrawalFinalizedTx) (alreadyPresent bool, err error)
	SaveWithdrawalInfo(context.Context, *models.WithdrawalInformation) (alreadyPresent bool, err error)

	SaveToken(context.Context, *models.Token) (alreadyPresent bool, err error)
	GetToken(ctx context.Context, chainId uint64, tokenAddr common.Address) (*models.Token, error)

	SaveChainLatestSyncedBlockNumber(ctx context.Context, chainKey string, blockNumber uint64) error
	GetChainLatestSyncedBlockNumber(ctx context.Context, chainKey string) (blockNumber uint64, err error)
}
