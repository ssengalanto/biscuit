package v1

import (
	"github.com/ssengalanto/hex/cmd/account/internal/interfaces/dto"
)

// DeleteAccountCommand contains required fields for account deletion.
type DeleteAccountCommand struct {
	ID string `json:"id"`
}

// NewDeleteAccountCommand creates a new command for account deletion.
func NewDeleteAccountCommand(input dto.DeleteAccountRequestDto) *DeleteAccountCommand {
	return &DeleteAccountCommand{
		ID: input.ID,
	}
}
