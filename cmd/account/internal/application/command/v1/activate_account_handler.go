package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
)

// ActivateAccountCommandHandler - command handler struct for account activation, satisfies midt.RequestHandler.
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
	empty := dtov1.ActivateAccountResponse{}

	cmd := request.(*ActivateAccountCommand) //nolint:errcheck //intentional panic

	id, err := uuid.Parse(cmd.ID)
	if err != nil {
		a.log.Error("invalid uuid", map[string]any{"command": cmd, "error": err})
		return empty, fmt.Errorf("%w: uuid `%s`", errors.ErrInvalid, cmd.ID)
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

	a.cache.Delete(ctx, cmd.ID) //nolint:errcheck //unnecessary

	response := dtov1.ActivateAccountResponse{ID: cmd.ID}

	return response, err
}
