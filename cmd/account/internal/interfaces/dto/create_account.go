package dto

import "time"

// CreateAccountRequestDto - account creation request dto.
type CreateAccountRequestDto struct {
	Email       string                    `json:"email" validate:"required,email"`
	Password    string                    `json:"password" validate:"min=10,required"`
	Active      bool                      `json:"active" validate:"boolean"`
	FirstName   string                    `json:"firstName" validate:"required"`
	LastName    string                    `json:"lastName" validate:"required"`
	Phone       string                    `json:"phone" validate:"required,numeric"`
	DateOfBirth time.Time                 `json:"dateOfBirth" validate:"required"`
	Locations   []CreateAddressRequestDto `json:"locations" validate:"required"`
}

// CreateAddressRequestDto - address creation request dto.
type CreateAddressRequestDto struct {
	Street     string `json:"street" validate:"required"`
	Unit       string `json:"unit,omitempty"`
	City       string `json:"city" validate:"required"`
	District   string `json:"district" validate:"required"`
	State      string `json:"state" validate:"required"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
}

// CreateAccountResponseDto - account creation response dto.
type CreateAccountResponseDto struct {
	ID string `json:"id"`
}
