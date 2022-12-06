package pgsql

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
)

// Address pgsql model.
type Address struct {
	ID               uuid.UUID     `json:"id" db:"id"`
	PersonID         uuid.UUID     `json:"personId" db:"person_id"`
	PlaceID          string        `json:"placeId" db:"place_id"`
	AddressLine1     address.Names `json:"addressLine1" db:"address_line1"`
	AddressLine2     address.Names `json:"addressLine2" db:"address_line2"`
	City             address.Names `json:"city" db:"city"`
	State            address.Names `json:"state" db:"state"`
	Country          address.Names `json:"country" db:"country"`
	PostalCode       address.Names `json:"postalCode" db:"postal_code"`
	FormattedAddress string        `json:"formattedAddress" db:"formatted_address"`
	Lat              float64       `json:"lat" db:"lat"`
	Lng              float64       `json:"lng" db:"lng"`
	CreatedAt        time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time     `json:"updatedAt" db:"updated_at"`
}

// ToEntity transforms the address model to account entity.
func (a Address) ToEntity() address.Entity {
	return address.Entity{
		ID:       a.ID,
		PersonID: a.PersonID,
		Components: address.Components{
			PlaceID:          a.PlaceID,
			AddressLine1:     a.AddressLine1,
			AddressLine2:     a.AddressLine2,
			City:             a.City,
			State:            a.State,
			Country:          a.Country,
			PostalCode:       a.PostalCode,
			FormattedAddress: a.FormattedAddress,
		},
		Geometry: address.Geometry{
			Lat: a.Lat,
			Lng: a.Lng,
		},
	}
}
