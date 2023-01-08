package pgsql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/account"
)

// Account pgsql model.
type Account struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	Email       string       `json:"email" db:"email"`
	Password    string       `json:"password" db:"password"`
	Active      bool         `json:"active" db:"active"`
	LastLoginAt sql.NullTime `json:"lastLoginAt" db:"last_login_at"`
	CreatedAt   time.Time    `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time    `json:"updatedAt" db:"updated_at"`
}

// ToEntity transforms the account model to account entity.
func (a Account) ToEntity() account.Entity {
	return account.Entity{
		ID:          a.ID,
		Email:       account.Email(a.Email),
		Password:    account.Password(a.Password),
		Active:      a.Active,
		LastLoginAt: a.LastLoginAt.Time,
	}
}
