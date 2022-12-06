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

// Create inserts a new account record in the database.
func (a *AccountRepository) Create(ctx context.Context, entity account.Entity) (account.Entity, error) {
	acc := account.Entity{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	query := MustBeValidAccountQuery(QueryCreateAccount)
	stmt, err := tx.PreparexContext(ctx, query)
	if err != nil {
		return acc, err
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
		return acc, err
	}

	err = tx.Commit()
	if err != nil {
		return acc, err
	}

	return acc, nil
}

// FindByID gets an account record with specific ID in the database.
func (a *AccountRepository) FindByID(ctx context.Context, id uuid.UUID) (account.Entity, error) {
	acc := Account{}

	query := MustBeValidAccountQuery(QueryFindByID)
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

// FindByEmail gets an account record with specific email in the database.
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

// UpdateByID updates an account record with specific id in the database.
func (a *AccountRepository) UpdateByID(ctx context.Context, entity account.Entity) (account.Entity, error) {
	acc := Account{}

	tx := a.db.MustBeginTx(ctx, nil)
	defer tx.Rollback() //nolint:errcheck //unnecessary

	query := MustBeValidAccountQuery(QueryUpdateByID)
	stmt, err := tx.PreparexContext(ctx, query)
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
