package v1

import (
	"time"

	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
)

// CreateAccountCommand contains required fields for account creation.
type CreateAccountCommand struct {
	Email       string                       `json:"email"`
	Password    string                       `json:"password"`
	Active      bool                         `json:"active"`
	FirstName   string                       `json:"firstName"`
	LastName    string                       `json:"lastName"`
	Phone       string                       `json:"phone"`
	DateOfBirth time.Time                    `json:"dateOfBirth"`
	Locations   []dtov1.CreateAddressRequest `json:"locations"`
}

// NewCreateAccountCommand creates a new command for account creation.
func NewCreateAccountCommand(input dtov1.CreateAccountRequest) *CreateAccountCommand {
	return &CreateAccountCommand{
		Email:       input.Email,
		Password:    input.Password,
		Active:      input.Active,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Phone:       input.Phone,
		DateOfBirth: input.DateOfBirth,
		Locations:   input.Locations,
	}
}
