package pgsql_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
)

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
