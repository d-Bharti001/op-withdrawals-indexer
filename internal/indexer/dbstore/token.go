package dbstore

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"op-withdrawals-indexer/internal/database/models"

	"github.com/ethereum/go-ethereum/common"
)

func (s *PostgresStore) SaveToken(ctx context.Context, db DbTx, token *models.Token) error {
	query := `
		INSERT INTO tokens (
			token_address,
			chain_id,
			token_class,
			token_name,
			token_symbol,
			token_decimals
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (chain_id, token_address)
		DO NOTHING;
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	tokenRow := token.ToDBRow()

	_, err := db.ExecContext(
		ctx,
		query,
		tokenRow.Address,
		tokenRow.ChainID,
		tokenRow.Class,
		tokenRow.Name,
		tokenRow.Symbol,
		tokenRow.Decimals,
	)

	return err
}

func (s *PostgresStore) GetToken(ctx context.Context, db DbTx, chainID uint64, tokenAddr common.Address) (*models.Token, error) {
	query := `
		SELECT
			token_address,
			chain_id,
			token_class,
			token_name,
			token_symbol,
			token_decimals
		FROM
			tokens
		WHERE
			chain_id = $1 AND token_address = $2
	`

	ctx, cancel := context.WithTimeout(ctx, DBQueryTimeout)
	defer cancel()

	var tokenRow models.TokenDBRow

	err := db.QueryRowContext(
		ctx,
		query,
		chainID,
		strings.ToLower(tokenAddr.Hex()),
	).Scan(
		&tokenRow.Address,
		&tokenRow.ChainID,
		&tokenRow.Class,
		&tokenRow.Name,
		&tokenRow.Symbol,
		&tokenRow.Decimals,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return tokenRow.ToDomainModel(), nil
}
