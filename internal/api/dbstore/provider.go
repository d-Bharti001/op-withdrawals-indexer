package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/api/dbmodels"

	"github.com/ethereum/go-ethereum/common"
)

type DBStoreProvider interface {
	WithdrawalHistory(
		ctx context.Context,
		addr common.Address,
		chainID,
		sinceTimestamp,
		limit,
		offset uint64,
	) (
		list []*dbmodels.WithdrawalDetails,
		totalCount uint64,
		err error,
	)

	CloseConnection() error
}
