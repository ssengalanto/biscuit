package token

import (
	"time"

	"github.com/google/uuid"
)

// Entity - token entity struct.
type Entity struct {
	ID           uuid.UUID `json:"id" validate:"required"`
	AccountID    uuid.UUID `json:"accountId" validate:"required"`
	ClientID     string    `json:"clientId" validate:"required"`
	AccessToken  JWT       `json:"accessToken" validate:"required"`
	RefreshToken JWT       `json:"refreshToken" validate:"required"`
	ExpiresIn    time.Time `json:"expiresIn" validate:"required"`
}

// Payload contains required fields for token claims.
type Payload struct {
	AccountID string
	Email     string
	ClientID  string
	Issuer    string
	Expiry    time.Duration
}

// New creates a new token entity.
func New(accountID uuid.UUID, clientID string) Entity {
	return Entity{
		ID:        uuid.New(),
		AccountID: accountID,
		ClientID:  clientID,
	}
}
