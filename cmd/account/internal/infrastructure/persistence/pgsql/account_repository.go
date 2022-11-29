package pgsql

import (
	"context"

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

	row := tx.QueryRowxContext(
		ctx,
		saveAccountQuery,
		entity.ID,
		entity.Email,
		entity.Password,
		entity.Active,
		entity.LastLoginAt,
	)
	err := row.StructScan(&acc)
	if err != nil {
		return acc.ToEntity(), err
	}

	err = tx.Commit()
	if err != nil {
		return acc.ToEntity(), err
	}

	return acc.ToEntity(), nil
}

const saveAccountQuery = `
	INSERT INTO account (id, email, password, active, last_login_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING *`
