package dto

import "time"

// GetAccountRequest - get account request dto.
type GetAccountRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

// PersonResponse - person field response dto.
type PersonResponse struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

// LocationResponse - address field response dto.
type LocationResponse struct {
	ID         string `json:"id"`
	Street     string `json:"street" validate:"required"`
	Unit       string `json:"unit,omitempty"`
	City       string `json:"city" validate:"required"`
	District   string `json:"district" validate:"required"`
	State      string `json:"state" validate:"required"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postalCode" validate:"required"`
}

// GetAccountResponse - get account response dto.
type GetAccountResponse struct {
	ID        string             `json:"id"`
	Email     string             `json:"email"`
	Active    bool               `json:"active"`
	Person    PersonResponse     `json:"person"`
	Locations []LocationResponse `json:"locations"`
}
