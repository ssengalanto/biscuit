package v1

import "time"

// CreateAccountRequest - account creation request dto.
type CreateAccountRequest struct {
	Email       string                 `json:"email" validate:"required,email" example:"johndoe@example.com"`
	Password    string                 `json:"password" validate:"min=10,required" example:"t5eC9E6ldLmaf"`
	Active      bool                   `json:"active" validate:"boolean" example:"true"`
	FirstName   string                 `json:"firstName" validate:"required" example:"John"`
	LastName    string                 `json:"lastName" validate:"required" example:"Doe"`
	Phone       string                 `json:"phone" validate:"required,numeric" example:"09066871243"`
	DateOfBirth time.Time              `json:"dateOfBirth" validate:"required" example:"2000-11-12T13:14:15Z"`
	Locations   []CreateAddressRequest `json:"locations" validate:"required"`
} // @name CreateAccountRequest

// CreateAddressRequest - address creation request dto.
type CreateAddressRequest struct {
	Street     string `json:"street" validate:"required" example:"365 Talon I Real 1740"`
	Unit       string `json:"unit,omitempty" example:"Unit 206 Rm. 5"`
	City       string `json:"city" validate:"required" example:"San Pedro"`
	District   string `json:"district" validate:"required" example:"Laguna"`
	State      string `json:"state" validate:"required" example:"Calabarzon"`
	Country    string `json:"country" validate:"required" example:"Philippines"`
	PostalCode string `json:"postalCode" validate:"required" example:"4023"`
} // @name CreateAddressRequest

// CreateAccountResponse - account creation response dto.
type CreateAccountResponse struct {
	ID string `json:"id"`
}
