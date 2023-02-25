package v1

import (
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
)

// DeactivateAccountCommand contains required fields for account deactivation.
type DeactivateAccountCommand struct {
	ID string `json:"id"`
}

// NewDeactivateAccountCommand creates a new command for account deactivation.
func NewDeactivateAccountCommand(input dtov1.DeactivateAccountRequest) *DeactivateAccountCommand {
	return &DeactivateAccountCommand{
		ID: input.ID,
	}
}
