package address

import (
	"github.com/ssengalanto/potato-project/pkg/validator"
)

// Components contains the core fields for address entity.
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

// Update checks the validity of the address components and updates its value.
func (c Components) Update(input Components) (Components, error) {
	_, err := input.IsValid()
	if err != nil {
		return Components{}, err
	}

	return input, nil
}

// Geometry contains lat and lng values for address entity.
type Geometry struct {
	Lat float64 `json:"lat" validate:"required"`
	Lng float64 `json:"lng" validate:"required"`
}

// IsValid checks the validity of the geometry details.
func (g Geometry) IsValid() (bool, error) {
	err := validator.Struct(g)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Update checks the validity of the address geometry and updates its value.
func (g Geometry) Update(input Geometry) (Geometry, error) {
	_, err := input.IsValid()
	if err != nil {
		return Geometry{}, err
	}

	return input, nil
}
