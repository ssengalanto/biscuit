package mock

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/address"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/person"
)

// NewAccount create a new mock account entity.
func NewAccount() account.Entity {
	return account.New(
		gofakeit.Email(),
		gofakeit.Password(
			true,
			true,
			true,
			true,
			false,
			10, //nolint:gomnd // unnecessary
		),
		gofakeit.Bool(),
	)
}

// NewPerson create a new mock person entity.
func NewPerson() person.Entity {
	return person.New(
		uuid.New(),
		gofakeit.FirstName(),
		gofakeit.LastName(),
		gofakeit.Email(),
		gofakeit.Phone(),
		gofakeit.Date(),
	)
}

// NewAddress create a new mock address entity.
func NewAddress() address.Entity {
	a := gofakeit.Address()
	addr := address.New(uuid.New(), address.Components{
		Street:     a.Street,
		Unit:       a.Street,
		City:       a.City,
		District:   a.City,
		State:      a.State,
		Country:    a.Country,
		PostalCode: a.Zip,
	})

	return addr
}

// NewAddresses create a new list of mock address entity.
func NewAddresses() []address.Entity {
	a := gofakeit.Address()

	var addrs []address.Entity
	addr := address.New(uuid.New(), address.Components{
		Street:     a.Street,
		Unit:       a.Street,
		City:       a.City,
		District:   a.City,
		State:      a.State,
		Country:    a.Country,
		PostalCode: a.Zip,
	})
	addrs = append(addrs, addr)

	return addrs
}

// NewAccountEntity create a new account entity aggregate.
func NewAccountEntity() account.Entity {
	acct := NewAccount()
	pers := NewPerson()
	addrs := NewAddresses()
	return account.AggregateAccount(acct, pers, addrs)
}
