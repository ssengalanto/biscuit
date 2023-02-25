package v1

import (
	"time"

	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
)

// UpdateAccountCommand contains required fields for updating account.
type UpdateAccountCommand struct {
	ID          string                        `json:"id"`
	FirstName   *string                       `json:"firstName"`
	LastName    *string                       `json:"lastName"`
	Phone       *string                       `json:"phone"`
	DateOfBirth *time.Time                    `json:"dateOfBirth"`
	Locations   *[]dtov1.UpdateAddressRequest `json:"locations"`
}

// NewUpdateAccountCommand creates a new command for updating account.
func NewUpdateAccountCommand(accountID string, input dtov1.UpdateAccountRequest) *UpdateAccountCommand {
	return &UpdateAccountCommand{
		ID:          accountID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Phone:       input.Phone,
		DateOfBirth: input.DateOfBirth,
		Locations:   input.Locations,
	}
}
