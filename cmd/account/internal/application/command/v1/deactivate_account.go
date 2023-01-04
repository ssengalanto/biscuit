package v1

import (
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
)

// DeactivateAccountCommand contains required fields for account deactivation.
type DeactivateAccountCommand struct {
	ID string `json:"id"`
}

// NewDeactivateAccountCommand creates a new command for account deactivation.
func NewDeactivateAccountCommand(input dto.DeactivateAccountRequestDto) *DeactivateAccountCommand {
	return &DeactivateAccountCommand{
		ID: input.ID,
	}
}
