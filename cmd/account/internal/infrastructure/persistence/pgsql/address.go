package pgsql

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
)

// Address pgsql model.
type Address struct {
	ID         uuid.UUID `json:"id" db:"id"`
	PersonID   uuid.UUID `json:"personId" db:"person_id"`
	Street     string    `json:"street" db:"street"`
	Unit       string    `json:"unit" db:"unit"`
	City       string    `json:"city" db:"city"`
	District   string    `json:"district" db:"district"`
	State      string    `json:"state" db:"state"`
	Country    string    `json:"country" db:"country"`
	PostalCode string    `json:"postalCode" db:"postal_code"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}

// ToEntity transforms the address model to account entity.
func (a Address) ToEntity() address.Entity {
	return address.Entity{
		ID:       a.ID,
		PersonID: a.PersonID,
		Components: address.Components{
			Street:     a.Street,
			Unit:       a.Unit,
			City:       a.City,
			District:   a.District,
			State:      a.State,
			Country:    a.Country,
			PostalCode: a.PostalCode,
		},
	}
}
