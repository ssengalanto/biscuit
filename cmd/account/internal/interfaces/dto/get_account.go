package dto

import "time"

// GetAccountRequestDto - get account request dto.
type GetAccountRequestDto struct {
	ID string `json:"id" validate:"required,uuid"`
}

// PersonResponseDto - person field response dto.
type PersonResponseDto struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

// LocationResponseDto - address field response dto.
type LocationResponseDto struct {
	ID         string `json:"id"`
	Street     string `json:"street" validate:"required"`
	Unit       string `json:"unit,omitempty"`
	City       string `json:"city" validate:"required"`
	District   string `json:"district" validate:"required"`
	State      string `json:"state" validate:"required"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
}

// GetAccountResponseDto - get account response dto.
type GetAccountResponseDto struct {
	ID        string                `json:"id"`
	Email     string                `json:"email"`
	Active    bool                  `json:"active"`
	Person    PersonResponseDto     `json:"person"`
	Locations []LocationResponseDto `json:"locations"`
}
