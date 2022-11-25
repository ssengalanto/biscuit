package pgsql

import (
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
)

// Address pgsql model.
type Address struct {
	ID               uuid.UUID     `json:"id"`
	PersonID         uuid.UUID     `json:"personId"`
	PlaceID          string        `json:"placeId"`
	AddressLine1     address.Names `json:"addressLine1"`
	AddressLine2     address.Names `json:"addressLine2"`
	City             address.Names `json:"city"`
	State            address.Names `json:"state"`
	Country          address.Names `json:"country"`
	PostalCode       address.Names `json:"postalCode"`
	FormattedAddress string        `json:"formattedAddress"`
	Lat              float64       `json:"lat"`
	Lng              float64       `json:"lng"`
	CreatedAt        time.Time     `json:"createdAt"`
	UpdatedAt        time.Time     `json:"updatedAt"`
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
