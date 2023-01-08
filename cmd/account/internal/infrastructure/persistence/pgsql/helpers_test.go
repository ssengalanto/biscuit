package pgsql_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/account"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/address"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/person"
)

// newAccountEntity creates a new account.Entity with mock values.
func newAccountEntity() account.Entity {
	accountID := uuid.New()
	personID := uuid.New()
	email := gofakeit.Email()

	return account.Entity{
		ID:          accountID,
		Email:       account.Email(email),
		Password:    account.Password(gofakeit.Password(true, true, true, true, false, 10)),
		Active:      gofakeit.Bool(),
		LastLoginAt: gofakeit.Date(),
		Person: &person.Entity{
			ID:        personID,
			AccountID: accountID,
			Details: person.Details{
				FirstName:   gofakeit.FirstName(),
				LastName:    gofakeit.LastName(),
				Email:       email,
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

// newAccountEntity creates a new address.Entity with mock values.
func newAddressEntity(personID uuid.UUID) address.Entity {
	addr := gofakeit.Address()

	return address.Entity{
		ID:       uuid.New(),
		PersonID: personID,
		Components: address.Components{
			Street:     addr.Address,
			Unit:       addr.Address,
			City:       addr.City,
			District:   addr.City,
			State:      addr.State,
			Country:    addr.Country,
			PostalCode: addr.Zip,
		},
	}
}

// createAccountRow creates a new mock Account record.
func createAccountRow(entity account.Entity) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)
}

// createPersonRow creates a new mock Person record.
func createPersonRow(entity account.Entity) *sqlmock.Rows {
	return sqlmock.NewRows(
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
}

// createAddressRows creates a new mock Address record.
func createAddressRows(entity account.Entity) *sqlmock.Rows {
	rows := sqlmock.NewRows(
		[]string{
			"id",
			"person_id",
			"street",
			"unit",
			"city",
			"district",
			"state",
			"country",
			"postal_code",
		},
	)

	for _, addr := range *entity.Person.Address {
		rows.
			AddRow(
				addr.ID,
				addr.PersonID,
				addr.Components.Street,
				addr.Components.Unit,
				addr.Components.City,
				addr.Components.District,
				addr.Components.State,
				addr.Components.Country,
				addr.Components.PostalCode,
			)
	}

	return rows
}
