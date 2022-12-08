package pgsql_test

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
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

	accountRows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)
	createAccountQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryCreateAccount)
	accountStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAccountQuery))
	accountStmt.ExpectQuery().WithArgs(
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt).WillReturnRows(accountRows)

	personRows := sqlmock.NewRows(
		[]string{"id", "account_id", "first_name", "last_name", "email", "phone", "date_of_birth", "avatar"},
	).
		AddRow(
			entity.Person.ID,
			entity.Person.AccountID,
			entity.Person.Details.FirstName,
			entity.Person.Details.LastName,
			entity.Person.Details.Email,
			entity.Person.Details.Phone,
			entity.Person.Details.DateOfBirth,
			entity.Person.Avatar,
		)
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
	).WillReturnRows(personRows)

	createAddressQuery := pgsql.MustBeValidAccountQuery(pgsql.QueryCreateAddress)
	addressStmt := dbmock.ExpectPrepare(regexp.QuoteMeta(createAddressQuery))

	for _, addr := range *entity.Person.Address {
		addressRows := sqlmock.NewRows(
			[]string{
				"id",
				"person_id",
				"place_id",
				"address_line1",
				"address_line2",
				"city",
				"state",
				"country",
				"postal_code",
				"formatted_address",
				"lat",
				"lng",
			},
		).
			AddRow(
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
			)
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

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	query := pgsql.MustBeValidAccountQuery(pgsql.QueryFindByID)
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

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	query := pgsql.MustBeValidAccountQuery(pgsql.QueryFindByEmail)
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

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	query := pgsql.MustBeValidAccountQuery(pgsql.QueryUpdateByID)
	stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))
	stmt.ExpectQuery().WithArgs(
		update.ID,
		update.Email,
		update.Password,
		update.Active,
		update.LastLoginAt).WillReturnRows(rows)

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

	rows := sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)

	query := pgsql.MustBeValidAccountQuery(pgsql.QueryDeleteByID)
	stmt := dbmock.ExpectPrepare(regexp.QuoteMeta(query))
	stmt.ExpectQuery().WithArgs(entity.ID).WillReturnRows(rows)

	result, err := repo.DeleteByID(context.Background(), entity.ID)
	require.NoError(t, err)
	require.Equal(t, result, entity)
}

func newAccountEntity() account.Entity {
	accountID := uuid.New()
	personID := uuid.New()

	return account.Entity{
		ID:          accountID,
		Email:       account.Email(gofakeit.Email()),
		Password:    account.Password(gofakeit.Password(true, true, true, true, false, 10)),
		Active:      gofakeit.Bool(),
		LastLoginAt: gofakeit.Date(),
		Person: &person.Entity{
			ID:        personID,
			AccountID: accountID,
			Details: person.Details{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				Email:       gofakeit.Email(),
				Phone:       gofakeit.Phone(),
				DateOfBirth: gofakeit.Date(),
			},
			Avatar: person.Avatar(gofakeit.URL()),
			Address: &[]address.Entity{
				newAddressEntity(personID),
				newAddressEntity(personID),
			},
		},
	}
}

func newAddressEntity(personID uuid.UUID) address.Entity {
	addr := gofakeit.Address()

	return address.Entity{
		ID:       uuid.New(),
		PersonID: personID,
		Components: address.Components{
			PlaceID: gofakeit.UUID(),
			AddressLine1: address.Names{
				ShortName: addr.Street,
				LongName:  addr.Street,
			},
			AddressLine2: address.Names{
				ShortName: addr.Street,
				LongName:  addr.Street,
			},
			City: address.Names{
				ShortName: addr.City,
				LongName:  addr.City,
			},
			State: address.Names{
				ShortName: addr.State,
				LongName:  addr.State,
			},
			Country: address.Names{
				ShortName: addr.Country,
				LongName:  addr.Country,
			},
			PostalCode: address.Names{
				ShortName: addr.Zip,
				LongName:  addr.Zip,
			},
			FormattedAddress: addr.Address,
		},
		Geometry: address.Geometry{
			Lat: gofakeit.Latitude(),
			Lng: gofakeit.Longitude(),
		},
	}
}
