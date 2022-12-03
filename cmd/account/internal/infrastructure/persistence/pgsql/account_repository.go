package pgsql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
)

type AccountRepository struct {
	db *sqlx.DB
}

// NewAccountRepository creates a new account repository.
func NewAccountRepository(db *sqlx.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Exists checks if an account record with specific ID exists in the database.
func (a *AccountRepository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
	var count int

	stmt, err := a.db.PreparexContext(ctx, AccountQueries["exists"])
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

// Create inserts a new account record in the database.
func (a *AccountRepository) Create(ctx context.Context, entity account.Entity) (account.Entity, error) {
	acc := Account{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	stmt, err := tx.PreparexContext(ctx, AccountQueries["create"])
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

	err = tx.Commit()
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// FindByID gets an account record with specific ID in the database.
func (a *AccountRepository) FindByID(ctx context.Context, id uuid.UUID) (account.Entity, error) {
	acc := Account{}

	stmt, err := a.db.PreparexContext(ctx, AccountQueries["findByID"])
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

// FindByEmail gets an account record with specific email in the database.
func (a *AccountRepository) FindByEmail(ctx context.Context, email string) (account.Entity, error) {
	acc := Account{}

	stmt, err := a.db.PreparexContext(ctx, AccountQueries["findByEmail"])
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

// UpdateByID updates an account record with specific id in the database.
func (a *AccountRepository) UpdateByID(ctx context.Context, entity account.Entity) (account.Entity, error) {
	acc := Account{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	stmt, err := tx.PreparexContext(ctx, AccountQueries["updateByID"])
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

	err = tx.Commit()
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

// DeleteByID deletes an account record with specific ID in the database.
func (a *AccountRepository) DeleteByID(ctx context.Context, id uuid.UUID) (account.Entity, error) {
	acc := Account{}

	stmt, err := a.db.PreparexContext(ctx, AccountQueries["deleteByID"])
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

// AccountQueries is a map holds all queries for account table.
var AccountQueries = map[string]string{ //nolint:gochecknoglobals //intended
	"exists":      accountExistsQuery,
	"create":      createAccountQuery,
	"findByID":    findByIDQuery,
	"findByEmail": findByEmailQuery,
	"updateByID":  updateByIDQuery,
	"deleteByID":  deleteByIDQuery,
}

const accountExistsQuery = `
	SELECT COUNT(1)
	FROM account
	WHERE id = $1`

const createAccountQuery = `
	INSERT INTO account (id, email, password, active, last_login_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *`

const findByIDQuery = `
	SELECT id, email, password, active, last_login_at
	FROM account
	WHERE id = $1`

const findByEmailQuery = `
	SELECT id, email, password, active, last_login_at
	FROM account
	WHERE email = $1`

const updateByIDQuery = `
	UPDATE account
	SET email = $2, password = $3, active = $4, last_login_at = $5, updated_at = NOW()
	FROM account
	WHERE id = $1
	RETURNING *`

const deleteByIDQuery = `
	DELETE FROM account
	WHERE id = $1
	RETURNING *`
