package dbstore

import (
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

func (s *PostgresStore) CloseConnection() error {
	return s.db.Close()
}
