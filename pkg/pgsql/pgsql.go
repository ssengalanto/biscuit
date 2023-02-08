package pgsql

import (
	"context"
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/biscuit/pkg/constants"
)

// NewConnection initializes a new postgres database connection pool.
func NewConnection(user, pwd, host, port, dbn, ssl string) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ResourceTimeout)
	defer cancel()

	hp := net.JoinHostPort(host, port)
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", user, pwd, hp, dbn, ssl)
	db, err := sqlx.ConnectContext(ctx, constants.PgsqlDriver, dsn)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnectionFailed, err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnectionFailed, err)
	}

	return db, nil
}
