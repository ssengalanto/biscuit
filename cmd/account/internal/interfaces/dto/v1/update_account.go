package v1

import "time"

// UpdateAccountRequest - update account request dto.
type UpdateAccountRequest struct {
	FirstName   *string                 `json:"firstName" example:"John"`
	LastName    *string                 `json:"lastName" example:"Doe"`
	Phone       *string                 `json:"phone" example:"09066871243"`
	DateOfBirth *time.Time              `json:"dateOfBirth" example:"2000-11-12T13:14:15Z"`
	Locations   *[]UpdateAddressRequest `json:"locations"`
} // @name UpdateAccountRequest

// UpdateAddressRequest - update address request dto.
type UpdateAddressRequest struct {
	ID         string  `json:"id" example:"3d3e36e1-9533-4408-8677-9d693a9ed8d4"`
	Street     *string `json:"street" validate:"required" example:"365 Talon I Real 1740"`
	Unit       *string `json:"unit" example:"Unit 206 Rm. 5"`
	City       *string `json:"city" example:"San Pedro"`
	District   *string `json:"district" example:"Laguna"`
	State      *string `json:"state" example:"Calabarzon"`
	Country    *string `json:"country" example:"Philippines"`
	PostalCode *string `json:"postalCode" example:"4023"`
}

// UpdateAccountResponse - update account response dto.
type UpdateAccountResponse struct {
	ID string `json:"id"`
}
