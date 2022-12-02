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

// Save insert a new account record in the database.
func (a *AccountRepository) Save(ctx context.Context, entity account.Entity) (account.Entity, error) {
	acc := Account{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	stmt, err := tx.PreparexContext(ctx, AccountQueries["save"])
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

// FindByID get an account record with specific ID in the database.
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

// AccountQueries is a map holds all queries for account table.
var AccountQueries = map[string]string{ //nolint:gochecknoglobals //intended
	"save":     saveAccountQuery,
	"findByID": findByIDQuery,
}

const saveAccountQuery = `
	INSERT INTO account (id, email, password, active, last_login_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *`

const findByIDQuery = `
	SELECT id, email, password, active, last_login_at
	FROM account
	WHERE id = $1`
