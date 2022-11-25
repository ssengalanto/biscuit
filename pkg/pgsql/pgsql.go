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
		return nil, ErrConnectionFailed
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		return nil, ErrConnectionFailed
	}

	return dbpool, nil
}
