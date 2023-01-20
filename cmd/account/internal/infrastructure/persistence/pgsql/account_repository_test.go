package pgsql_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/infrastructure/persistence/pgsql"
	"github.com/ssengalanto/biscuit/pkg/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAccountRepository(t *testing.T) {
	db, _, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock.NewMockLogger(ctrl)

	repo := pgsql.NewAccountRepository(logger, db)
	assert.NotNil(t, repo)
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
				assert.True(t, result)
			},
		},
		{
			name:    "record doesn't exists",
			payload: uuid.New(),
			rows:    sqlmock.NewRows([]string{"count"}).AddRow(0),
			assert: func(t *testing.T, result bool, err error) {
				require.NoError(t, err)
				assert.False(t, result)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, dbmock, err := mock.NewSqlDb()
			require.NoError(t, err)
			defer db.Close()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			logger := mock.NewMockLogger(ctrl)
			repo := pgsql.NewAccountRepository(logger, db)

			query := pgsql.MustBeValidAccountQuery(pgsql.QueryExists)
			stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))
			stmt.ExpectQuery().WithArgs(tc.payload).WillReturnRows(tc.rows)

			result, err := repo.Exists(context.Background(), tc.payload)
			tc.assert(t, result, err)
		})
	}
}

func TestAccountRepository_Create(t *testing.T) {
	entity := newAccountEntity()
	*entity.Person.Address = nil

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock.NewMockLogger(ctrl)
	repo := pgsql.NewAccountRepository(logger, db)

	dbmock.ExpectBegin()

	createAccountQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryCreateAccount)
	accountStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAccountQuery))
	accountStmt.ExpectExec().WithArgs(
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt).WillReturnResult(sqlmock.NewResult(0, 1))

	createPersonQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryCreatePerson)
	personStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createPersonQuery))
	personStmt.ExpectExec().WithArgs(
		entity.Person.ID,
		entity.Person.AccountID,
		entity.Person.Details.FirstName,
		entity.Person.Details.LastName,
		entity.Person.Details.Email,
		entity.Person.Details.Phone,
		entity.Person.Details.DateOfBirth,
		entity.Person.Avatar,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	createAddressQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryCreateAddress)
	addressStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAddressQuery))

	for _, addr := range *entity.Person.Address {
		addressStmt.ExpectExec().WithArgs(
			addr.ID,
			addr.PersonID,
			addr.Components.Street,
			addr.Components.Unit,
			addr.Components.City,
			addr.Components.District,
			addr.Components.State,
			addr.Components.Country,
			addr.Components.PostalCode,
		).WillReturnResult(sqlmock.NewResult(0, int64(len(*entity.Person.Address))))
	}

	dbmock.ExpectCommit()

	err = repo.Create(context.Background(), entity)
	require.NoError(t, err)
}

func TestAccountRepository_FindByID(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock.NewMockLogger(ctrl)
	repo := pgsql.NewAccountRepository(logger, db)

	dbmock.ExpectBegin()

	accountRow := createAccountRow(entity)
	findAccountByIDQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryFindAccountByID)
	accountStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(findAccountByIDQuery))
	accountStmt.ExpectQuery().WithArgs(entity.ID).WillReturnRows(accountRow)

	personRow := createPersonRow(entity)
	createPersonQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryFindPersonByAccountID)
	personStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createPersonQuery))
	personStmt.ExpectQuery().WithArgs(entity.ID).WillReturnRows(personRow)

	createAddressQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryFindAddressByPersonID)
	addressStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAddressQuery))
	addressRows := createAddressRows(entity)
	addressStmt.ExpectQuery().WithArgs(entity.Person.ID).WillReturnRows(addressRows)

	dbmock.ExpectCommit()

	result, err := repo.FindByID(context.Background(), entity.ID)
	require.NoError(t, err)
	assert.Equal(t, result, entity)
}

func TestAccountRepository_Update(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock.NewMockLogger(ctrl)
	repo := pgsql.NewAccountRepository(logger, db)

	dbmock.ExpectBegin()

	createAccountQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryUpdateAccountByID)
	accountStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAccountQuery))
	accountStmt.ExpectExec().WithArgs(
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt).WillReturnResult(sqlmock.NewResult(0, 1))

	createPersonQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryUpdatePersonByID)
	personStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createPersonQuery))
	personStmt.ExpectExec().WithArgs(
		entity.Person.ID,
		entity.Person.Details.FirstName,
		entity.Person.Details.LastName,
		entity.Person.Details.Email,
		entity.Person.Details.Phone,
		entity.Person.Details.DateOfBirth,
		entity.Person.Avatar,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	createAddressQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryUpdateAddressByID)
	addressStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAddressQuery))

	for _, addr := range *entity.Person.Address {
		addressStmt.ExpectExec().WithArgs(
			addr.ID,
			addr.PersonID,
			addr.Components.Street,
			addr.Components.Unit,
			addr.Components.City,
			addr.Components.District,
			addr.Components.State,
			addr.Components.Country,
			addr.Components.PostalCode,
		).WillReturnResult(sqlmock.NewResult(0, 1))
	}

	dbmock.ExpectCommit()

	err = repo.Update(context.Background(), entity)
	require.NoError(t, err)
}

func TestAccountRepository_DeleteByID(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := mock.NewMockLogger(ctrl)
	repo := pgsql.NewAccountRepository(logger, db)

	deleteAccountByIDQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryDeleteAccountByID)
	deleteStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(deleteAccountByIDQuery))
	deleteStmt.ExpectExec().WithArgs(entity.ID).WillReturnResult(sqlmock.NewResult(0, 3))

	err = repo.DeleteByID(context.Background(), entity.ID)
	require.NoError(t, err)
}
