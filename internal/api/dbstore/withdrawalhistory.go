package dbstore

import (
	"context"
	"op-withdrawals-indexer/internal/api/dbmodels"
	"op-withdrawals-indexer/internal/database/dbtypes"

	"github.com/ethereum/go-ethereum/common"
)

func (s *PostgresStore) WithdrawalHistory(ctx context.Context, addr common.Address, chainID, sinceTimestamp, limit, offset uint64) ([]*dbmodels.WithdrawalDetails, uint64, error) {
	query := `
		WITH
			proven_flags AS (
				SELECT
					wpt.withdrawal_chain_id,
					wpt.withdrawal_hash,
					COUNT(*) > 0 AS proven,
					bool_or(wpt.proof_submitter = $1) AS proved_by_user
				FROM withdrawal_proven_txs wpt
				WHERE wpt.withdrawal_chain_id = $2
				GROUP BY wpt.withdrawal_chain_id, wpt.withdrawal_hash
			),
			finalization_flags AS (
				SELECT
					wft.withdrawal_chain_id,
					wft.withdrawal_hash,
					bool_or(wft.success) AS finalization_success
				FROM withdrawal_finalized_txs wft
				WHERE wft.withdrawal_chain_id = $2
				GROUP BY wft.withdrawal_chain_id, wft.withdrawal_hash
			)

		SELECT
			w.withdrawal_hash,
			w.chain_id,
			w.withdrawal_nonce,
			w.withdrawal_sender,
			w.withdrawal_target,
			w.withdrawal_value,
			w.withdrawal_gas_limit,
			w.withdrawal_data,
			w.tx_hash,
			w.block_number,
			w.block_timestamp,
			wi.decoded_action AS info_decoded_action,
			wi.token_address AS info_token_address,
			wi.from_address AS info_from_address,
			wi.to_address AS info_to_address,
			wi.token_id AS info_token_id,
			wi.amount AS info_amount,
			t.chain_id as token_chain_id,
			t.token_class,
			t.token_name,
			t.token_symbol,
			t.token_decimals,
			COALESCE(pf.proven, FALSE) AS proven,
			ff.finalization_success,
			COUNT(*) OVER() AS total_count

		FROM withdrawals w

		LEFT JOIN withdrawal_info wi
			ON
				wi.chain_id = w.chain_id
				AND wi.withdrawal_hash = w.withdrawal_hash

		LEFT JOIN tokens t
			ON
				t.chain_id = wi.chain_id
				AND t.token_address = wi.token_address

		LEFT JOIN proven_flags pf
			ON
				pf.withdrawal_chain_id = w.chain_id
				AND pf.withdrawal_hash = w.withdrawal_hash

		LEFT JOIN finalization_flags ff
			ON
				ff.withdrawal_chain_id = w.chain_id
				AND ff.withdrawal_hash = w.withdrawal_hash

		WHERE
			w.chain_id = $2
			AND (
				w.withdrawal_sender = $1
				OR wi.from_address = $1
				OR pf.proved_by_user
			)
			AND (
				w.block_timestamp >= $3
				OR ff.finalization_success IS NULL
			)

		ORDER BY w.block_timestamp DESC
		LIMIT $4 OFFSET $5;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	rows, err := s.db.QueryContext(
		ctx,
		query,
		dbtypes.Address(addr),
		chainID,
		sinceTimestamp,
		limit,
		offset,
	)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var totalCount uint64
	result := make([]*dbmodels.WithdrawalDetails, 0)

	for rows.Next() {
		var row dbmodels.WithdrawalDetailsDBRow
		err = rows.Scan(
			&row.WithdrawalHash,
			&row.ChainID,
			&row.Nonce,
			&row.Sender,
			&row.Target,
			&row.Value,
			&row.GasLimit,
			&row.Data,
			&row.TxHash,
			&row.BlockNumber,
			&row.BlockTimestamp,
			&row.DecodedAction,
			&row.TokenAddress,
			&row.FromAddress,
			&row.ToAddress,
			&row.TokenID,
			&row.Amount,
			&row.TokenChainID,
			&row.TokenClass,
			&row.TokenName,
			&row.TokenSymbol,
			&row.TokenDecimals,
			&row.Proven,
			&row.FinalizationSuccess,
			&totalCount,
		)
		if err != nil {
			return nil, 0, err
		}

		domainModel, err := row.ToDomainModel()
		if err != nil {
			return nil, 0, err
		}

		result = append(result, domainModel)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return result, totalCount, nil
}
