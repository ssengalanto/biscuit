package pgsql_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
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

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	dbmock.ExpectBegin()

	accountRow := createAccountRow(entity)
	createAccountQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryCreateAccount)
	accountStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAccountQuery))
	accountStmt.ExpectQuery().WithArgs(
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt).WillReturnRows(accountRow)

	personRow := createPersonRow(entity)
	createPersonQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryCreatePerson)
	personStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createPersonQuery))
	personStmt.ExpectQuery().WithArgs(
		entity.Person.ID,
		entity.Person.AccountID,
		entity.Person.Details.FirstName,
		entity.Person.Details.LastName,
		entity.Person.Details.Email,
		entity.Person.Details.Phone,
		entity.Person.Details.DateOfBirth,
		entity.Person.Avatar,
	).WillReturnRows(personRow)

	createAddressQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryCreateAddress)
	addressStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAddressQuery))

	for _, addr := range *entity.Person.Address {
		addressRows := createAddressRow(addr)
		addressStmt.ExpectQuery().WithArgs(
			addr.ID,
			addr.PersonID,
			addr.Components.PlaceID,
			addr.Components.AddressLine1.MustEncodeJSON(),
			addr.Components.AddressLine2.MustEncodeJSON(),
			addr.Components.City.MustEncodeJSON(),
			addr.Components.State.MustEncodeJSON(),
			addr.Components.Country.MustEncodeJSON(),
			addr.Components.PostalCode.MustEncodeJSON(),
			addr.Components.FormattedAddress,
			addr.Geometry.Lat,
			addr.Geometry.Lng,
		).WillReturnRows(addressRows)
	}

	dbmock.ExpectCommit()

	result, err := repo.Create(context.Background(), entity)

	require.NoError(t, err)
	require.Equal(t, entity, result)
}

func TestAccountRepository_FindByID(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

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
	require.Equal(t, result, entity)
}

func TestAccountRepository_FindByEmail(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	dbmock.ExpectBegin()

	accountRow := createAccountRow(entity)
	findAccountByEmailQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryFindAccountByEmail)
	accountStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(findAccountByEmailQuery))
	accountStmt.ExpectQuery().WithArgs(entity.Email).WillReturnRows(accountRow)

	personRow := createPersonRow(entity)
	createPersonQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryFindPersonByAccountID)
	personStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createPersonQuery))
	personStmt.ExpectQuery().WithArgs(entity.ID).WillReturnRows(personRow)

	createAddressQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryFindAddressByPersonID)
	addressStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAddressQuery))
	addressRows := createAddressRows(entity)
	addressStmt.ExpectQuery().WithArgs(entity.Person.ID).WillReturnRows(addressRows)

	dbmock.ExpectCommit()

	result, err := repo.FindByEmail(context.Background(), entity.Email.String())
	require.NoError(t, err)
	require.Equal(t, result, entity)
}

func TestAccountRepository_Update(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

	dbmock.ExpectBegin()

	accountRow := createAccountRow(entity)
	createAccountQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryUpdateAccountByID)
	accountStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAccountQuery))
	accountStmt.ExpectQuery().WithArgs(
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt).WillReturnRows(accountRow)

	personRow := createPersonRow(entity)
	createPersonQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryUpdatePersonByID)
	personStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createPersonQuery))
	personStmt.ExpectQuery().WithArgs(
		entity.Person.ID,
		entity.Person.Details.FirstName,
		entity.Person.Details.LastName,
		entity.Person.Details.Email,
		entity.Person.Details.Phone,
		entity.Person.Details.DateOfBirth,
		entity.Person.Avatar,
	).WillReturnRows(personRow)

	createAddressQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryUpdateAddressByID)
	addressStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAddressQuery))

	for _, addr := range *entity.Person.Address {
		addressRows := createAddressRow(addr)
		addressStmt.ExpectQuery().WithArgs(
			addr.ID,
			addr.Components.PlaceID,
			addr.Components.AddressLine1.MustEncodeJSON(),
			addr.Components.AddressLine2.MustEncodeJSON(),
			addr.Components.City.MustEncodeJSON(),
			addr.Components.State.MustEncodeJSON(),
			addr.Components.Country.MustEncodeJSON(),
			addr.Components.PostalCode.MustEncodeJSON(),
			addr.Components.FormattedAddress,
			addr.Geometry.Lat,
			addr.Geometry.Lng,
		).WillReturnRows(addressRows)
	}

	dbmock.ExpectCommit()

	result, err := repo.Update(context.Background(), entity)

	require.NoError(t, err)
	require.Equal(t, entity, result)
}

func TestAccountRepository_DeleteByID(t *testing.T) {
	entity := newAccountEntity()

	db, dbmock, err := mock.NewSqlDb()
	require.NoError(t, err)
	defer db.Close()

	repo := pgsql.NewAccountRepository(db)

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

	deleteAccountByIDQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryDeleteAccountByID)
	deleteStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(deleteAccountByIDQuery))
	deleteStmt.ExpectExec().WithArgs(entity.ID).WillReturnResult(sqlmock.NewResult(0, 3))

	result, err := repo.DeleteByID(context.Background(), entity.ID)
	require.NoError(t, err)
	require.Equal(t, result, entity)
}
