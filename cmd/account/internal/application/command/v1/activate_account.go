package v1

import (
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
)

// ActivateAccountCommand contains required fields for account activation.
type ActivateAccountCommand struct {
	ID string `json:"id"`
}

// NewActivateAccountCommand creates a new command for account activation.
func NewActivateAccountCommand(input dto.ActivateAccountRequest) *ActivateAccountCommand {
	return &ActivateAccountCommand{
		ID: input.ID,
	}
}
