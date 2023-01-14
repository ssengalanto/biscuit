package v1

import (
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
)

// DeactivateAccountCommand contains required fields for account deactivation.
type DeactivateAccountCommand struct {
	ID string `json:"id"`
}

// NewDeactivateAccountCommand creates a new command for account deactivation.
func NewDeactivateAccountCommand(input dto.DeactivateAccountRequest) *DeactivateAccountCommand {
	return &DeactivateAccountCommand{
		ID: input.ID,
	}
}
