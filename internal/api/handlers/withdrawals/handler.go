package withdrawals

import "op-withdrawals-indexer/internal/api/dbstore"

type WithdrawalsHandler struct {
	db dbstore.DBStoreProvider
}

func NewWithdrawalsHandler(db dbstore.DBStoreProvider) *WithdrawalsHandler {
	return &WithdrawalsHandler{
		db,
	}
}
