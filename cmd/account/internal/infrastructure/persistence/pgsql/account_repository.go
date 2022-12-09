package pgsql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
)

type AccountRepository struct {
	db *sqlx.DB
}

// NewAccountRepository creates a new account repository.
func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Exists checks if an account record with the specified ID exists in the database.
func (a *AccountRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int

	query := MustBeValidAccountQuery(QueryExists)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return false, err
	}

	row := stmt.QueryRowxContext(ctx, id)

	err = row.Scan(&count)
	if err != nil {
		return false, err
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}

// Create begins a new transaction to process and insert a new account record together with its associated
// person and address records. If transaction fails it will roll back all the changes it made,
// otherwise it will commit the changes to the database.
func (a *AccountRepository) Create(ctx context.Context, entity account.Entity) (account.Entity, error) {
	e := account.Entity{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	acc, err := createAccount(ctx, tx, entity)
	if err != nil {
		return e, err
	}

	p, err := createPerson(ctx, tx, *entity.Person)
	if err != nil {
		return e, err
	}

	addrs, err := createAddress(ctx, tx, *entity.Person.Address)
	if err != nil {
		return e, err
	}

	tx.Commit() //nolint:errcheck //unnecessary

	e = buildAccountEntity(acc, p, addrs)
	return e, nil
}

// FindByID gets an account record with the specified ID in the database
// together with its associated person and address records.
func (a *AccountRepository) FindByID(ctx context.Context, id uuid.UUID) (account.Entity, error) {
	entity := account.Entity{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	acc, err := findAccountByID(ctx, tx, id)
	if err != nil {
		return entity, err
	}

	p, err := findPersonByAccountID(ctx, tx, acc.ID)
	if err != nil {
		return entity, err
	}

	addrs, err := findAddressByPersonID(ctx, tx, p.ID)
	if err != nil {
		return entity, err
	}

	tx.Commit() //nolint:errcheck //unnecessary

	entity = buildAccountEntity(acc, p, addrs)
	return entity, nil
}

// FindByEmail gets an account record with the specified email in the database.
func (a *AccountRepository) FindByEmail(ctx context.Context, email string) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryFindByEmail)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return account.Entity{}, err
	}

	row := stmt.QueryRowxContext(ctx, email)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// UpdateByID updates an account record with the specified ID in the database.
func (a *AccountRepository) UpdateByID(ctx context.Context, entity account.Entity) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryUpdateByID)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return account.Entity{}, err
	}

	row := stmt.QueryRowxContext(
		ctx,
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt,
	)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// DeleteByID deletes an account record with the specified ID in the database.
func (a *AccountRepository) DeleteByID(ctx context.Context, id uuid.UUID) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryDeleteByID)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return account.Entity{}, err
	}

	row := stmt.QueryRowxContext(ctx, id)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// createAccount inserts a new account record in the database.
func createAccount(ctx context.Context, tx *sqlx.Tx, entity account.Entity) (account.Entity, error) {
	a := Account{}

	query := MustBeValidAccountQuery(QueryCreateAccount)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return a.ToEntity(), err
	}

	row := stmt.QueryRowxContext(
		ctx,
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt,
	)

	err = row.StructScan(&a)
	if err != nil {
		return a.ToEntity(), err
	}

	return a.ToEntity(), nil
}

// createPerson inserts a new person record associated with account in the database.
func createPerson(ctx context.Context, tx *sqlx.Tx, entity person.Entity) (person.Entity, error) {
	p := Person{}

	query := MustBeValidAccountQuery(QueryCreatePerson)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return p.ToEntity(), err
	}

	row := stmt.QueryRowxContext(
		ctx,
		entity.ID,
		entity.AccountID,
		entity.Details.FirstName,
		entity.Details.LastName,
		entity.Details.Email,
		entity.Details.Phone,
		entity.Details.DateOfBirth,
		entity.Avatar,
	)

	err = row.StructScan(&p)
	if err != nil {
		return p.ToEntity(), err
	}

	return p.ToEntity(), nil
}

// createAddress inserts a new address record associated with person in the database.
func createAddress(ctx context.Context, tx *sqlx.Tx, entities []address.Entity) ([]address.Entity, error) {
	var addrs []address.Entity

	query := MustBeValidAccountQuery(QueryCreateAddress)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}

	for _, entity := range entities {
		a := Address{}

		row := stmt.QueryRowxContext(
			ctx,
			entity.ID,
			entity.PersonID,
			entity.Components.PlaceID,
			entity.Components.AddressLine1.MustEncodeJSON(),
			entity.Components.AddressLine2.MustEncodeJSON(),
			entity.Components.City.MustEncodeJSON(),
			entity.Components.State.MustEncodeJSON(),
			entity.Components.Country.MustEncodeJSON(),
			entity.Components.PostalCode.MustEncodeJSON(),
			entity.Components.FormattedAddress,
			entity.Geometry.Lat,
			entity.Geometry.Lng,
		)

		err = row.StructScan(&a)
		if err != nil {
			return nil, err
		}

		addrs = append(addrs, a.ToEntity())
	}

	return addrs, nil
}

// findAccountByID gets the account record with the specified ID in the database.
func findAccountByID(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryFindAccountByID)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return account.Entity{}, err
	}

	row := stmt.QueryRowxContext(ctx, id)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// findPersonByAccountID gets the person record associated with account in the database.
func findPersonByAccountID(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (person.Entity, error) {
	p := Person{}

	query := MustBeValidAccountQuery(QueryFindPersonByAccountID)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return p.ToEntity(), err
	}

	row := stmt.QueryRowxContext(ctx, id)

	err = row.StructScan(&p)
	if err != nil {
		return p.ToEntity(), err
	}

	return p.ToEntity(), nil
}

// findAddressByPersonID gets the list of address records associated with person in the database.
func findAddressByPersonID(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) ([]address.Entity, error) {
	var addrs []address.Entity

	query := MustBeValidAccountQuery(QueryFindAddressByPersonID)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryxContext(ctx, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		addr := Address{}
		err = rows.StructScan(&addr)
		if err != nil {
			return nil, err
		}

		addrs = append(addrs, addr.ToEntity())
	}
	defer rows.Close()

	return addrs, nil
}

// buildAccountEntity takes account, person and address entities as parameters
// and builds the account entity.
func buildAccountEntity(account account.Entity, person person.Entity, address []address.Entity) account.Entity {
	entity := account
	entity.Person = &person
	entity.Person.Address = &address

	return entity
}
