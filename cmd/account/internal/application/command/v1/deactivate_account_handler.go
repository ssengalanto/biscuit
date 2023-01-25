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

// DeactivateAccountCommandHandler - command handler struct for account deactivation, satisfies midt.RequestHandler.
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
	empty := dto.DeactivateAccountResponse{}

	cmd := request.(*DeactivateAccountCommand) //nolint:errcheck //intentional panic

	id, err := uuid.Parse(cmd.ID)
	if err != nil {
		d.log.Error("invalid uuid", map[string]any{"command": cmd, "error": err})
		return empty, fmt.Errorf("%w: uuid `%s`", errors.ErrInvalid, cmd.ID)
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

	d.cache.Delete(ctx, cmd.ID) //nolint:errcheck //unnecessary

	res := dto.DeactivateAccountResponse{ID: cmd.ID}

	return res, err
}
