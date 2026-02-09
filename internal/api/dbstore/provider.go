package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/api/dbmodels"

	"github.com/ethereum/go-ethereum/common"
)

type DBStoreProvider interface {
	WithdrawalHistory(ctx context.Context, addr common.Address, chainID, sinceTimestamp uint64) ([]*dbmodels.WithdrawalDetails, error)
	CloseConnection() error
}
