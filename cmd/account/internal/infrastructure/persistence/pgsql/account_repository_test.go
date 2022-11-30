package pgsql_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/infrastructure/persistence/pgsql"
	"github.com/ssengalanto/potato-project/pkg/mock"
	"github.com/stretchr/testify/require"
)

func TestNewAccountRepository(t *testing.T) {
	db, _, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)
	require.NotNil(t, repo)
}

func TestAccountRepository_Save(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	query, ok := pgsql.AccountQueries["save"]
	require.True(t, ok)

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	dbmock.ExpectBegin()
	stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))
	stmt.ExpectQuery().WithArgs(
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt).WillReturnRows(rows)
	dbmock.ExpectCommit()

	result, err := repo.Save(context.Background(), entity)
	require.NoError(t, err)
	require.Equal(t, result, entity)
}

func newAccountEntity() account.Entity {
	return account.Entity{
		ID:          uuid.New(),
		Email:       account.Email(gofakeit.Email()),
		Password:    account.Password(gofakeit.Password(true, true, true, true, false, 10)),
		Active:      gofakeit.Bool(),
		LastLoginAt: gofakeit.Date(),
	}
}
