//nolint:revive,stylecheck //unnecessary for this package
package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/biscuit/pkg/constants"
)

// NewSqlDb returns a sqlx stub.
func NewSqlDb() (*sqlx.DB, sqlmock.Sqlmock, error) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, ErrDatabaseStubConnection
	}

	db := sqlx.NewDb(mockDb, constants.SlqMockDriver)

	return db, mock, nil
}
