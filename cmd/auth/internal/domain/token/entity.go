package token

import (
	"time"

	"github.com/google/uuid"
)

const (
	tokenExpiry = 15 * time.Minute
)

// Entity - token entity struct.
type Entity struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	AccountID string    `json:"account_id" validate:"required"`
	ClientID  string    `json:"client_id" validate:"required"`
	Token     JWT       `json:"token" validate:"required"`
	ExpiresIn time.Time `json:"expires_in" validate:"required"`
}

// Subject contains required fields for token claims.
type Subject struct {
	AccountID string
	Email     string
	ClientID  string
	Issuer    string
}

// New creates a new token entity.
func New(accountID, clientID string) Entity {
	return Entity{
		ID:        uuid.New(),
		AccountID: accountID,
		ClientID:  clientID,
	}
}
