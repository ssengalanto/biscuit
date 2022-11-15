package person

import (
	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
)

// Entity - Person Entity.
type Entity struct {
	ID        uuid.UUID       `json:"id"`
	AccountID uuid.UUID       `json:"accountId"`
	Details   Details         `json:"details"`
	Avatar    string          `json:"avatar"`
	Address   *address.Entity `json:"address"`
}

// New creates a new person entity.
func New() Entity {
	return Entity{
		ID: uuid.New(),
	}
}
