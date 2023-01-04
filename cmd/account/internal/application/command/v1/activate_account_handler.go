package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
)

// ActivateAccountCommandHandler - command handler struct for account activation, satisfies mediatr.RequestHandler.
type ActivateAccountCommandHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
	cache             account.Cache
}

// NewActivateAccountCommandHandler creates a new command handler that handles account activation.
func NewActivateAccountCommandHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
	cache account.Cache,
) *ActivateAccountCommandHandler {
	return &ActivateAccountCommandHandler{
		log:               logger,
		accountRepository: accountRepository,
		cache:             cache,
	}
}

func (a *ActivateAccountCommandHandler) Name() string {
	return fmt.Sprintf("%T", &ActivateAccountCommand{})
}

func (a *ActivateAccountCommandHandler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	empty := dto.ActivateAccountResponseDto{}

	command, ok := request.(*ActivateAccountCommand)
	if !ok {
		a.log.Error("invalid command", map[string]any{"command": command})
		return empty, fmt.Errorf("%w: command", errors.ErrInvalid)
	}

	id, err := uuid.Parse(command.ID)
	if err != nil {
		a.log.Error("invalid uuid", map[string]any{"command": command, "error": err})
		return empty, fmt.Errorf("%w: uuid %s", errors.ErrInvalid, command.ID)
	}

	acct, err := a.accountRepository.FindByID(ctx, id)
	if err != nil {
		return empty, err
	}

	acct.Activate()

	err = a.accountRepository.Update(ctx, acct)
	if err != nil {
		return empty, err
	}

	err = a.cache.Delete(ctx, command.ID)
	if err != nil {
		return empty, err
	}

	response := dto.ActivateAccountResponseDto{ID: command.ID}

	return response, err
}
