package dto

import "time"

// UpdateAccountRequest - update account request dto.
type UpdateAccountRequest struct {
	FirstName   *string                 `json:"firstName"`
	LastName    *string                 `json:"lastName"`
	Phone       *string                 `json:"phone"`
	DateOfBirth *time.Time              `json:"dateOfBirth"`
	Locations   *[]UpdateAddressRequest `json:"locations"`
}

// UpdateAddressRequest - update address request dto.
type UpdateAddressRequest struct {
	ID         string  `json:"id"`
	Street     *string `json:"street"`
	Unit       *string `json:"unit"`
	City       *string `json:"city"`
	District   *string `json:"district"`
	State      *string `json:"state"`
	Country    *string `json:"country"`
	PostalCode *string `json:"postalCode"`
}

// UpdateAccountResponse - update account response dto.
type UpdateAccountResponse struct {
	ID string `json:"id"`
}
