package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/address"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/person"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
)

// UpdateAccountCommandHandler - command handler struct for updating account, satisfies midt.RequestHandler.
type UpdateAccountCommandHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
	cache             account.Cache
}

// NewUpdateAccountCommandHandler creates a new command handler that handles account updates.
func NewUpdateAccountCommandHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
	cache account.Cache,
) *UpdateAccountCommandHandler {
	return &UpdateAccountCommandHandler{
		log:               logger,
		accountRepository: accountRepository,
		cache:             cache,
	}
}

func (u *UpdateAccountCommandHandler) Name() string {
	return fmt.Sprintf("%T", &UpdateAccountCommand{})
}

func (u *UpdateAccountCommandHandler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	empty := dto.UpdateAccountResponse{}

	command, ok := request.(*UpdateAccountCommand)
	if !ok {
		u.log.Error("invalid command", map[string]any{"command": command})
		return empty, errors.ErrInternal
	}

	id, err := uuid.Parse(command.ID)
	if err != nil {
		u.log.Error("invalid uuid", map[string]any{"command": command, "error": err})
		return empty, fmt.Errorf("%w: uuid `%s`", errors.ErrInvalid, command.ID)
	}

	acct, err := u.accountRepository.FindByID(ctx, id)
	if err != nil {
		return empty, err
	}

	err = u.updateAccount(&acct, command)
	if err != nil {
		return empty, err
	}

	err = u.accountRepository.Update(ctx, acct)
	if err != nil {
		return empty, err
	}

	err = u.cache.Delete(ctx, command.ID)
	if err != nil {
		return empty, err
	}

	response := dto.UpdateAccountResponse{ID: command.ID}

	return response, err
}

func (u *UpdateAccountCommandHandler) updateAccount(
	acct *account.Entity,
	cmd *UpdateAccountCommand,
) error {
	err := acct.UpdatePersonDetails(person.UpdateDetailsInput{
		FirstName:   cmd.FirstName,
		LastName:    cmd.LastName,
		Phone:       cmd.Phone,
		DateOfBirth: cmd.DateOfBirth,
	})
	if err != nil {
		u.log.Error("person update failed", map[string]any{"command": cmd, "error": err})
		return err
	}

	if cmd.Locations != nil {
		err = u.updateAddress(acct, *cmd.Locations)
		if err != nil {
			u.log.Error(
				"address update failed",
				map[string]any{"command": cmd, "error": err},
			)
			return err
		}
	}

	return nil
}

func (u UpdateAccountCommandHandler) updateAddress(acct *account.Entity, addrs []dto.UpdateAddressRequest) error {
	var addrInputs []account.UpdateAddressInput
	for _, addr := range addrs {
		input := account.UpdateAddressInput{
			ID: addr.ID,
			Components: address.UpdateComponentsInput{
				Street:     addr.Street,
				Unit:       addr.Unit,
				City:       addr.City,
				District:   addr.District,
				State:      addr.State,
				Country:    addr.Country,
				PostalCode: addr.PostalCode,
			},
		}
		addrInputs = append(addrInputs, input)
	}

	err := acct.UpdatePersonAddress(addrInputs)
	if err != nil {
		return err
	}

	return nil
}
