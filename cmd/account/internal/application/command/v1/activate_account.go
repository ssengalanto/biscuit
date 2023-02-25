package v1

import (
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
)

// ActivateAccountCommand contains required fields for account activation.
type ActivateAccountCommand struct {
	ID string `json:"id"`
}

// NewActivateAccountCommand creates a new command for account activation.
func NewActivateAccountCommand(input dtov1.ActivateAccountRequest) *ActivateAccountCommand {
	return &ActivateAccountCommand{
		ID: input.ID,
	}
}
