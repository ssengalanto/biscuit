package token

import (
	"time"

	"github.com/google/uuid"
)

const (
	tokenExpiry = 15 * time.Minute
)

type Entity struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	AccountID string    `json:"account_id" validate:"required"`
	ClientID  string    `json:"client_id"`
	ExpiredAt time.Time `json:"expired_at"`
}

// New creates a new token entity.
func New(accountID, clientID string) Entity {
	now := time.Now()
	return Entity{
		ID:        uuid.New(),
		AccountID: accountID,
		ClientID:  clientID,
		ExpiredAt: now.Add(tokenExpiry),
	}
}
