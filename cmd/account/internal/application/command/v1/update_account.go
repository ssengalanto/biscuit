package v1

import (
	"time"

	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
)

// UpdateAccountCommand contains required fields for updating account.
type UpdateAccountCommand struct {
	ID          string                         `json:"id"`
	FirstName   *string                        `json:"firstName"`
	LastName    *string                        `json:"lastName"`
	Phone       *string                        `json:"phone"`
	DateOfBirth *time.Time                     `json:"dateOfBirth"`
	Locations   *[]dto.UpdateAddressRequestDto `json:"locations"`
}

// NewUpdateAccountCommand creates a new command for updating account.
func NewUpdateAccountCommand(accountID string, input dto.UpdateAccountRequestDto) *UpdateAccountCommand {
	return &UpdateAccountCommand{
		ID:          accountID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Phone:       input.Phone,
		DateOfBirth: input.DateOfBirth,
		Locations:   input.Locations,
	}
}
