package address

import (
	"github.com/ssengalanto/potato-project/pkg/validator"
)

// Components - address component value object.
type Components struct {
	PlaceID          string `json:"placeID" validate:"required"`
	AddressLine1     Names  `json:"addressLine1" validate:"required"`
	AddressLine2     Names  `json:"addressLine2"`
	City             Names  `json:"city" validate:"required"`
	State            Names  `json:"state" validate:"required"`
	Country          Names  `json:"country" validate:"required"`
	PostalCode       Names  `json:"postalCode" validate:"required"`
	FormattedAddress string `json:"formattedAddress" validate:"required"`
}

// IsValid checks the validity of the person details.
func (c Components) IsValid() (bool, error) {
	err := validator.Struct(c)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Geometry - address lat lng value object.
type Geometry struct {
	Lat float64 `json:"lat" validate:"required"`
	Lng float64 `json:"lng" validate:"required"`
}

// IsValid checks the validity of the person details.
func (g Geometry) IsValid() (bool, error) {
	err := validator.Struct(g)
	if err != nil {
		return false, err
	}

	return true, nil
}
