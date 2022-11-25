package pgsql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewConnection initializes a new postgres database connection pool.
func NewConnection(dsn string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
