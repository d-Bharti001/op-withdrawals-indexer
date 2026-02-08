package dbstore

import (
	"context"

	"op-withdrawals-indexer/internal/database/models"
)

func (s *PostgresStore) SaveWithdrawalInfo(ctx context.Context, db DbTx, withdrawalInfo *models.WithdrawalInformation) error {
	query := `
		INSERT INTO withdrawal_info (
			withdrawal_hash,
			chain_id,
			decoded_action,
			token_address,
			from_address,
			to_address,
			token_id,
			amount
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (chain_id, withdrawal_hash)
		DO NOTHING;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	wdInfoRow := withdrawalInfo.ToDBRow()

	_, err := db.ExecContext(
		ctx,
		query,
		wdInfoRow.WithdrawalHash,
		wdInfoRow.ChainID,
		wdInfoRow.DecodedAction,
		wdInfoRow.TokenAddress,
		wdInfoRow.FromAddress,
		wdInfoRow.ToAddress,
		wdInfoRow.TokenID,
		wdInfoRow.Amount,
	)

	return err
}
