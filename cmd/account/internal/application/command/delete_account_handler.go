package command

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
)

// DeleteAccountCommandHandler - command handler struct for account deletion, satisfies mediatr.RequestHandler.
type DeleteAccountCommandHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
}

// NewDeleteAccountCommandHandler creates a new command handler that handles account deletion.
func NewDeleteAccountCommandHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
) *DeleteAccountCommandHandler {
	return &DeleteAccountCommandHandler{
		log:               logger,
		accountRepository: accountRepository,
	}
}

func (d *DeleteAccountCommandHandler) Name() string {
	return fmt.Sprintf("%T", &DeleteAccountCommand{})
}

func (d *DeleteAccountCommandHandler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	command, ok := request.(*DeleteAccountCommand)
	if !ok {
		d.log.Error("invalid command", map[string]any{"command": command})
		return nil, fmt.Errorf("%w: command", errors.ErrInvalid)
	}

	id, err := uuid.Parse(command.ID)
	if err != nil {
		d.log.Error("invalid uuid", map[string]any{"command": command})
		return nil, fmt.Errorf("%w: uuid %s", errors.ErrInvalid, command.ID)
	}

	err = d.accountRepository.DeleteByID(ctx, id)
	if err != nil {
		d.log.Error("account deletion failed", map[string]any{"id": id, "error": err})
		return nil, err
	}

	response := dto.DeleteAccountRequestDto{ID: command.ID}

	return response, err
}
