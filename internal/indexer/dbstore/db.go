package dbstore

import (
	"context"
	"database/sql"
	"op-withdrawals-indexer/internal/database/postgresql"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*PostgresStore, error) {
	postgresDb, err := postgresql.New(connStr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: postgresDb,
	}, nil
}

func (s *PostgresStore) DB() DbTx {
	return s.db
}

func (s *PostgresStore) WithTx(ctx context.Context, fn func(tx DbTx) error) error {
	tx, err := s.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *PostgresStore) CloseConnection() error {
	return s.db.Close()
}
