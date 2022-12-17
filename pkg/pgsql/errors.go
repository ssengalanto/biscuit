package pgsql

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

// Errors used by the pgsql package.

// ErrConnectionFailed is returned when postgres database connection failed.
var ErrConnectionFailed = fmt.Errorf("pgsql database connection failed")

func ErrorCode(err error) string {
	var pgErr *pgconn.PgError

	ok := errors.As(err, &pgErr)
	if !ok {
		return ""
	}

	return pgErr.Code
}
