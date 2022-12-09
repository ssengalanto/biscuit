package pgsql_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
)

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

func createAccountRow(entity account.Entity) *sqlmock.Rows {
	return sqlmock.NewRows([]string{"id", "email", "password", "active", "last_login_at"}).
		AddRow(entity.ID, entity.Email, entity.Password, entity.Active, entity.LastLoginAt)
}

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

func createAddressRow(addr address.Entity) *sqlmock.Rows {
	return sqlmock.NewRows(
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
}

func createAddressRows(entity account.Entity) *sqlmock.Rows {
	rows := sqlmock.NewRows(
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
	)

	for _, addr := range *entity.Person.Address {
		rows.
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
	}

	return rows
}
