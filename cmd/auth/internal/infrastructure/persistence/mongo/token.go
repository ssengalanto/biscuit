package mongo

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/auth/internal/domain/token"
)

// Token mongo model.
type Token struct {
	ID           uuid.UUID `json:"id" bson:"_id"`
	AccountID    uuid.UUID `json:"accountId" bson:"accountId"`
	ClientID     string    `json:"clientId" bson:"clientId"`
	AccessToken  string    `json:"accessToken" bson:"accessToken"`
	RefreshToken string    `json:"refreshToken" bson:"accessToken"`
	ExpiredAt    time.Time `json:"expiredAt" bson:"expiredAt"`
	UpdatedAt    time.Time `json:"updatedAt" bson:"updatedAt"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
}

// ToEntity transforms the token model to token entity.
func (t Token) ToEntity() token.Entity {
	return token.Entity{
		ID:           t.ID,
		AccountID:    t.AccountID,
		ClientID:     t.ClientID,
		AccessToken:  token.JWT(t.AccessToken),
		RefreshToken: token.JWT(t.RefreshToken),
		ExpiredAt:    t.ExpiredAt,
	}
}
