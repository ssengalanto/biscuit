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
	empty := dtov1.DeleteAccountResponse{}

	cmd := request.(*DeleteAccountCommand) //nolint:errcheck //intentional panic

	id, err := uuid.Parse(cmd.ID)
	if err != nil {
		d.log.Error("invalid uuid", map[string]any{"command": cmd, "error": err})
		return empty, fmt.Errorf("%w: uuid `%s`", errors.ErrInvalid, cmd.ID)
	}

	err = d.accountRepository.DeleteByID(ctx, id)
	if err != nil {
		return empty, err
	}

	d.cache.Delete(ctx, cmd.ID) //nolint:errcheck //unnecessary

	res := dtov1.DeleteAccountResponse{ID: cmd.ID}

	return res, err
}
