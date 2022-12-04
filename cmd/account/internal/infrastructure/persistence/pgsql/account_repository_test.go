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

func TestAccountRepository_Exists(t *testing.T) {
	testCases := []struct {
		name    string
		payload uuid.UUID
		rows    *sqlmock.Rows
		assert  func(t *testing.T, result bool, err error)
	}{
		{
			name:    "record exists",
			payload: uuid.New(),
			rows:    sqlmock.NewRows([]string{"count"}).AddRow(1),
			assert: func(t *testing.T, result bool, err error) {
				require.NoError(t, err)
				require.True(t, result)
			},
		},
		{
			name:    "record doesn't exists",
			payload: uuid.New(),
			rows:    sqlmock.NewRows([]string{"count"}).AddRow(0),
			assert: func(t *testing.T, result bool, err error) {
				require.NoError(t, err)
				require.False(t, result)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, dbmock, err := mock.NewSqlDb()
			require.NoError(t, err)
			defer db.Close()

			repo := pgsql.NewAccountRepository(db)

			query, ok := pgsql.AccountQueries["exists"]
			require.True(t, ok)

			stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))

			stmt.ExpectQuery().WithArgs(tc.payload).WillReturnRows(tc.rows)

			result, err := repo.Exists(context.Background(), tc.payload)
			tc.assert(t, result, err)
		})
	}
}

func TestAccountRepository_Create(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	query, ok := pgsql.AccountQueries["create"]
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

	result, err := repo.Create(context.Background(), entity)
	require.NoError(t, err)
	require.Equal(t, result, entity)
}

func TestAccountRepository_FindByID(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	query, ok := pgsql.AccountQueries["findByID"]
	require.True(t, ok)

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))

	stmt.ExpectQuery().WithArgs(entity.ID).WillReturnRows(rows)

	result, err := repo.FindByID(context.Background(), entity.ID)
	require.NoError(t, err)
	require.Equal(t, result, entity)
}

func TestAccountRepository_FindByEmail(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	query, ok := pgsql.AccountQueries["findByEmail"]
	require.True(t, ok)

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))

	stmt.ExpectQuery().WithArgs(entity.Email.String()).WillReturnRows(rows)

	result, err := repo.FindByEmail(context.Background(), entity.Email.String())
	require.NoError(t, err)
	require.Equal(t, result, entity)
}

func TestAccountRepository_UpdateByID(t *testing.T) {
	entity := newAccountEntity()
	update := newAccountEntity()
	update.ID = entity.ID

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	query, ok := pgsql.AccountQueries["updateByID"]
	require.True(t, ok)

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	dbmock.ExpectBegin()
	stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))
	stmt.ExpectQuery().WithArgs(
		update.ID,
		update.Email,
		update.Password,
		update.Active,
		update.LastLoginAt).WillReturnRows(rows)
	dbmock.ExpectCommit()

	result, err := repo.UpdateByID(context.Background(), update)
	require.NoError(t, err)
	require.Equal(t, result, entity)
}

func TestAccountRepository_DeleteByID(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	query, ok := pgsql.AccountQueries["deleteByID"]
	require.True(t, ok)

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))

	stmt.ExpectQuery().WithArgs(entity.ID).WillReturnRows(rows)

	result, err := repo.DeleteByID(context.Background(), entity.ID)
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
