package pgsql

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
)

// namesJSON is a json string that satisfies the address.Names fields.
type namesJSON string

// mustDecodeJSON decodes the string to address.Names.
func (n namesJSON) mustDecodeJSON() address.Names {
	var names address.Names
	err := json.Unmarshal([]byte(n), &names)
	if err != nil {
		panic(err)
	}

	return names
}

// Address pgsql model.
type Address struct {
	ID               uuid.UUID `json:"id" db:"id"`
	PersonID         uuid.UUID `json:"personId" db:"person_id"`
	PlaceID          string    `json:"placeId" db:"place_id"`
	AddressLine1     namesJSON `json:"addressLine1" db:"address_line1"`
	AddressLine2     namesJSON `json:"addressLine2" db:"address_line2"`
	City             namesJSON `json:"city" db:"city"`
	State            namesJSON `json:"state" db:"state"`
	Country          namesJSON `json:"country" db:"country"`
	PostalCode       namesJSON `json:"postalCode" db:"postal_code"`
	FormattedAddress string    `json:"formattedAddress" db:"formatted_address"`
	Lat              float64   `json:"lat" db:"lat"`
	Lng              float64   `json:"lng" db:"lng"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt        time.Time `json:"updatedAt" db:"updated_at"`
}

// ToEntity transforms the address model to account entity.
func (a Address) ToEntity() address.Entity {
	return address.Entity{
		ID:       a.ID,
		PersonID: a.PersonID,
		Components: address.Components{
			PlaceID:          a.PlaceID,
			AddressLine1:     a.AddressLine1.mustDecodeJSON(),
			AddressLine2:     a.AddressLine2.mustDecodeJSON(),
			City:             a.City.mustDecodeJSON(),
			State:            a.State.mustDecodeJSON(),
			Country:          a.Country.mustDecodeJSON(),
			PostalCode:       a.PostalCode.mustDecodeJSON(),
			FormattedAddress: a.FormattedAddress,
		},
		Geometry: address.Geometry{
			Lat: a.Lat,
			Lng: a.Lng,
		},
	}
}
