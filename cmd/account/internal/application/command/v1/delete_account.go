package v1

import (
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
)

// DeleteAccountCommand contains required fields for account deletion.
type DeleteAccountCommand struct {
	ID string `json:"id"`
}

// NewDeleteAccountCommand creates a new command for account deletion.
func NewDeleteAccountCommand(input dtov1.DeleteAccountRequest) *DeleteAccountCommand {
	return &DeleteAccountCommand{
		ID: input.ID,
	}
}
