package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
)

// Account postgres model.
type Account struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Active      bool      `json:"active"`
	LastLoginAt time.Time `json:"lastLoginAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ToEntity transforms the account model to account entity.
func (a Account) ToEntity() account.Entity {
	return account.Entity{
		ID:          a.ID,
		Email:       account.Email(a.Email),
		Password:    account.Password(a.Password),
		Active:      a.Active,
		LastLoginAt: a.LastLoginAt,
	}
}
