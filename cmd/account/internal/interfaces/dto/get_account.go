package dto

import "time"

// GetAccountRequest - get account request dto.
type GetAccountRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

// PersonResponse - person field response dto.
type PersonResponse struct {
	ID          string    `json:"id" example:"63237c24-c6f3-49bd-808b-e7764e75ebd1"`
	FirstName   string    `json:"firstName" example:"John"`
	LastName    string    `json:"lastName" example:"Doe"`
	Email       string    `json:"email" example:"johndoe@example.com"`
	Phone       string    `json:"phone" example:"09066871243"`
	DateOfBirth time.Time `json:"dateOfBirth" example:"2000-2-20"`
} // @name GetAccountResponse

// LocationResponse - address field response dto.
type LocationResponse struct {
	ID         string `json:"id" example:"3d3e36e1-9533-4408-8677-9d693a9ed8d4"`
	Street     string `json:"street" validate:"required" example:"365 Talon I Real 1740"`
	Unit       string `json:"unit,omitempty" example:"Unit 206 Rm. 5"`
	City       string `json:"city" validate:"required" example:"San Pedro"`
	District   string `json:"district" validate:"required" example:"Laguna"`
	State      string `json:"state" validate:"required" example:"Calabarzon"`
	Country    string `json:"country" validate:"required" example:"Philippines"`
	PostalCode string `json:"postalCode" validate:"required" example:"4023"`
} // @name LocationResponse

// GetAccountResponse - get account response dto.
type GetAccountResponse struct {
	ID        string             `json:"id" example:"b058dc29-546a-41e2-8e9c-f8142b9d194c"`
	Email     string             `json:"email" example:"johndoe@example.com"`
	Active    bool               `json:"active" example:"true"`
	Person    PersonResponse     `json:"person"`
	Locations []LocationResponse `json:"locations"`
} // @name GetAccountResponse
