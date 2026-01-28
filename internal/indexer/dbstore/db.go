package dbstore

import (
	"database/sql"
	"op-withdrawals-indexer/internal/database/postgresql"
	"time"
)

const DBQueryTimeout = 3 * time.Second

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(addr string, maxOpenConns, maxIdleConns int, maxIdleTime string) (*PostgresStore, error) {
	postgresDb, err := postgresql.New(addr, maxOpenConns, maxIdleConns, maxIdleTime)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: postgresDb,
	}, nil
}
