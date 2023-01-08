package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/account"
	"github.com/ssengalanto/hex/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/hex/pkg/errors"
	"github.com/ssengalanto/hex/pkg/interfaces"
)

// DeactivateAccountCommandHandler - command handler struct for account deactivation, satisfies mediatr.RequestHandler.
type DeactivateAccountCommandHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
	cache             account.Cache
}

// NewDeactivateAccountCommandHandler creates a new command handler that handles account deactivation.
func NewDeactivateAccountCommandHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
	cache account.Cache,
) *DeactivateAccountCommandHandler {
	return &DeactivateAccountCommandHandler{
		log:               logger,
		accountRepository: accountRepository,
		cache:             cache,
	}
}

func (d *DeactivateAccountCommandHandler) Name() string {
	return fmt.Sprintf("%T", &DeactivateAccountCommand{})
}

func (d *DeactivateAccountCommandHandler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	empty := dto.DeactivateAccountResponseDto{}

	command, ok := request.(*DeactivateAccountCommand)
	if !ok {
		d.log.Error("invalid command", map[string]any{"command": command})
		return empty, fmt.Errorf("%w: command", errors.ErrInvalid)
	}
	d.log.Info(fmt.Sprintf("executing %s", d.Name()), nil)

	id, err := uuid.Parse(command.ID)
	if err != nil {
		d.log.Error("invalid uuid", map[string]any{"command": command, "error": err})
		return empty, fmt.Errorf("%w: uuid %s", errors.ErrInvalid, command.ID)
	}

	acct, err := d.accountRepository.FindByID(ctx, id)
	if err != nil {
		return empty, err
	}

	acct.Deactivate()

	err = d.accountRepository.Update(ctx, acct)
	if err != nil {
		return empty, err
	}

	err = d.cache.Delete(ctx, command.ID)
	if err != nil {
		return empty, err
	}

	response := dto.DeactivateAccountResponseDto{ID: command.ID}

	return response, err
}
