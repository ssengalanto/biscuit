package pgsql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/hex/pkg/constants"
)

// NewConnection initializes a new postgres database connection pool.
func NewConnection(dsn string) (*sqlx.DB, error) {
	ctx := context.Background()
	db, err := sqlx.ConnectContext(ctx, constants.PgsqlDriver, dsn)
	if err != nil {
		return nil, ErrConnectionFailed
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, ErrConnectionFailed
	}

	return db, nil
}
