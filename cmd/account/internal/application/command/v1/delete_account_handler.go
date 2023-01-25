package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
)

// DeleteAccountCommandHandler - command handler struct for account deletion, satisfies midt.RequestHandler.
type DeleteAccountCommandHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
	cache             account.Cache
}

// NewDeleteAccountCommandHandler creates a new command handler that handles account deletion.
func NewDeleteAccountCommandHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
	cache account.Cache,
) *DeleteAccountCommandHandler {
	return &DeleteAccountCommandHandler{
		log:               logger,
		accountRepository: accountRepository,
		cache:             cache,
	}
}

func (d *DeleteAccountCommandHandler) Name() string {
	return fmt.Sprintf("%T", &DeleteAccountCommand{})
}

func (d *DeleteAccountCommandHandler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	empty := dto.DeleteAccountResponse{}

	command, ok := request.(*DeleteAccountCommand)
	if !ok {
		d.log.Error("invalid command", map[string]any{"command": command})
		return empty, errors.ErrInternal
	}

	id, err := uuid.Parse(command.ID)
	if err != nil {
		d.log.Error("invalid uuid", map[string]any{"command": command, "error": err})
		return empty, fmt.Errorf("%w: uuid `%s`", errors.ErrInvalid, command.ID)
	}

	err = d.accountRepository.DeleteByID(ctx, id)
	if err != nil {
		return empty, err
	}

	d.cache.Delete(ctx, command.ID) //nolint:errcheck //unnecessary

	response := dto.DeleteAccountResponse{ID: command.ID}

	return response, err
}
