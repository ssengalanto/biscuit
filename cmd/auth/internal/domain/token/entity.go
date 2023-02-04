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
	ExpiredAt    time.Time `json:"expiredAt" validate:"required"`
}

// Payload contains required fields for token claims.
type Payload struct {
	AccountID string
	ClientID  string
	Issuer    string
	ExpiresIn time.Duration
}

// New creates a new token entity.
func New(accountID uuid.UUID, clientID string) Entity {
	return Entity{
		ID:        uuid.New(),
		AccountID: accountID,
		ClientID:  clientID,
	}
}

// GenerateAccessToken generates a JWT access token and sets the access token state of the token entity.
func (e *Entity) GenerateAccessToken(iss, pk string, ei time.Duration) error {
	p := Payload{
		AccountID: e.AccountID.String(),
		ClientID:  e.ClientID,
		Issuer:    iss,
		ExpiresIn: ei,
	}

	at, err := NewJWT(p, Base64RSAPrivateKey(pk))
	if err != nil {
		return err
	}

	e.AccessToken = at

	return nil
}

// GenerateRefreshToken generates a JWT refresh token and sets the access token state of the token entity.
func (e *Entity) GenerateRefreshToken(iss, pk string, ei time.Duration) error {
	p := Payload{
		AccountID: e.AccountID.String(),
		ClientID:  e.ClientID,
		Issuer:    iss,
		ExpiresIn: ei,
	}

	rt, err := NewJWT(p, Base64RSAPrivateKey(pk))
	if err != nil {
		return err
	}

	e.RefreshToken = rt

	return nil
}
