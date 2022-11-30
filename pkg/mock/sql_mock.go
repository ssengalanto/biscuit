//nolint:revive,stylecheck //unnecessary for this package
package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/potato-project/pkg/constants"
)

func NewSqlDb() (*sqlx.DB, sqlmock.Sqlmock, error) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, ErrDatabaseStubConnection
	}

	db := sqlx.NewDb(mockDb, constants.SlqMockDriver)

	return db, mock, nil
}
