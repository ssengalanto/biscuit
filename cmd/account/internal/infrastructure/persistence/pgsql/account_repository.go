package pgsql

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/gg"
	"github.com/ssengalanto/potato-project/pkg/pgsql"
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
// person record. If transaction fails it will roll back all the changes it made,
// otherwise it will commit the changes to the database.
func (a *AccountRepository) Create(ctx context.Context, entity account.Entity) error {
	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	err := createAccount(ctx, tx, entity)
	if err != nil {
		return err
	}

	err = createPerson(ctx, tx, *entity.Person)
	if err != nil {
		return err
	}

	tx.Commit() //nolint:errcheck //unnecessary

	return nil
}

// CreatePersonAddresses begins a new transaction to process and insert a
// single or multiple address associated with person record.
// If transaction fails it will roll back all the changes it made,
// otherwise it will commit the changes to the database.
func (a *AccountRepository) CreatePersonAddresses(
	ctx context.Context,
	entities []address.Entity,
) error {
	tx := a.db.MustBeginTx(ctx, nil)

	err := createAddress(ctx, tx, entities)
	if err != nil {
		return err
	}

	tx.Commit() //nolint:errcheck //unnecessary

	return nil
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
	entity := account.Entity{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	acc, err := findAccountByEmail(ctx, tx, email)
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

// Update updates an account record in the database.
func (a *AccountRepository) Update(ctx context.Context, entity account.Entity) error {
	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	err := updateAccount(ctx, tx, entity)
	if err != nil {
		return err
	}

	err = updatePerson(ctx, tx, *entity.Person)
	if err != nil {
		return err
	}

	err = updateAddress(ctx, tx, *entity.Person.Address)
	if err != nil {
		return err
	}

	tx.Commit() //nolint:errcheck //unnecessary

	return nil
}

// DeleteByID deletes an account record with the specified ID in the database.
func (a *AccountRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	query := MustBeValidAccountQuery(QueryDeleteAccountByID)
	stmt, err := a.db.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	err = handleRowsAffected(result.RowsAffected())
	if err != nil {
		return err
	}

	return nil
}

// createAccount inserts a new account record in the database.
func createAccount(ctx context.Context, tx *sqlx.Tx, entity account.Entity) error {
	query := MustBeValidAccountQuery(QueryCreateAccount)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(
		ctx,
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt,
	)
	if err != nil {
		code := pgsql.ErrorCode(err)

		if code == pgerrcode.UniqueViolation {
			return fmt.Errorf("%w: duplicate email key value", errors.ErrInvalid)
		}

		return fmt.Errorf("%w: %s", errors.ErrInternal, err.Error())
	}

	err = handleRowsAffected(result.RowsAffected())
	if err != nil {
		return err
	}

	return nil
}

// createPerson inserts a new person record associated with account in the database.
func createPerson(ctx context.Context, tx *sqlx.Tx, entity person.Entity) error {
	query := MustBeValidAccountQuery(QueryCreatePerson)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(
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
	if err != nil {
		return err
	}

	err = handleRowsAffected(result.RowsAffected())
	if err != nil {
		return err
	}

	return nil
}

// createAddress inserts a new address record associated with person in the database.
func createAddress(ctx context.Context, tx *sqlx.Tx, entities []address.Entity) error {
	query := MustBeValidAccountQuery(QueryCreateAddress)
	stmt, preperr := tx.PreparexContext(ctx, query)
	if preperr != nil {
		return preperr
	}

	for _, entity := range entities {
		result, err := stmt.ExecContext(
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
		if err != nil {
			return err
		}

		err = handleRowsAffected(result.RowsAffected())
		if err != nil {
			return err
		}
	}

	return nil
}

// findAccountByID gets the account record with the specified ID in the database.
func findAccountByID(ctx context.Context, tx *sqlx.Tx, id uuid.UUID) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryFindAccountByID)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return acc.ToEntity(), err
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

// findAccountByEmail gets the account record with the specified email address in the database.
func findAccountByEmail(ctx context.Context, tx *sqlx.Tx, email string) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryFindAccountByEmail)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return acc.ToEntity(), err
	}

	row := stmt.QueryRowxContext(ctx, email)

	err = row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// updateAccount updates an account record in the database.
func updateAccount(ctx context.Context, tx *sqlx.Tx, entity account.Entity) error {
	query := MustBeValidAccountQuery(QueryUpdateAccountByID)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(
		ctx,
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt,
	)
	if err != nil {
		return err
	}

	err = handleRowsAffected(result.RowsAffected())
	if err != nil {
		return err
	}

	return nil
}

// updatePerson updates a person record associated with account in the database.
func updatePerson(ctx context.Context, tx *sqlx.Tx, entity person.Entity) error {
	query := MustBeValidAccountQuery(QueryUpdatePersonByID)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return err
	}

	result, err := stmt.ExecContext(
		ctx,
		entity.ID,
		entity.Details.FirstName,
		entity.Details.LastName,
		entity.Details.Email,
		entity.Details.Phone,
		entity.Details.DateOfBirth,
		entity.Avatar,
	)
	if err != nil {
		return err
	}

	err = handleRowsAffected(result.RowsAffected())
	if err != nil {
		return err
	}

	return nil
}

// updateAddress updates an address record associated with person in the database.
func updateAddress(ctx context.Context, tx *sqlx.Tx, entities []address.Entity) error {
	query := MustBeValidAccountQuery(QueryUpdateAddressByID)
	stmt, preperr := tx.PreparexContext(ctx, query)
	if preperr != nil {
		return preperr
	}

	for _, entity := range entities {
		result, err := stmt.ExecContext(
			ctx,
			entity.ID,
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
		if err != nil {
			return err
		}

		err = handleRowsAffected(result.RowsAffected())
		if err != nil {
			return err
		}
	}

	return nil
}

func handleRowsAffected(n int64, err error) error {
	if err != nil {
		return err
	}

	if !gg.Itob(n) {
		return ErrExecFailed
	}

	return nil
}

// buildAccountEntity takes account, person and address entities as parameters
// and builds the account entity.
func buildAccountEntity(account account.Entity, person person.Entity, address []address.Entity) account.Entity {
	entity := account
	entity.Person = &person
	entity.Person.Address = &address

	return entity
}
