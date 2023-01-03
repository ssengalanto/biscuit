package dto

import "time"

// UpdateAccountRequestDto - update account request dto.
type UpdateAccountRequestDto struct {
	FirstName   *string                    `json:"firstName"`
	LastName    *string                    `json:"lastName"`
	Phone       *string                    `json:"phone"`
	DateOfBirth *time.Time                 `json:"dateOfBirth"`
	Locations   *[]UpdateAddressRequestDto `json:"locations"`
}

// UpdateAddressRequestDto - update address request dto.
type UpdateAddressRequestDto struct {
	ID         string  `json:"id"`
	Street     *string `json:"street"`
	Unit       *string `json:"unit"`
	City       *string `json:"city"`
	District   *string `json:"district"`
	State      *string `json:"state"`
	Country    *string `json:"country"`
	PostalCode *string `json:"postalCode"`
}

// UpdateAccountResponseDto - update account response dto.
type UpdateAccountResponseDto struct {
	ID string `json:"id"`
}
