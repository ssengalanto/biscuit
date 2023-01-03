package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
)

// UpdateAccountCommandHandler - command handler struct for updating account, satisfies mediatr.RequestHandler.
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
	empty := dto.UpdateAccountResponseDto{}

	command, ok := request.(*UpdateAccountCommand)
	if !ok {
		u.log.Error("invalid command", map[string]any{"command": command})
		return empty, fmt.Errorf("%w: command", errors.ErrInvalid)
	}

	id, err := uuid.Parse(command.ID)
	if err != nil {
		u.log.Error("invalid uuid", map[string]any{"command": command, "error": err})
		return empty, fmt.Errorf("%w: uuid %s", errors.ErrInvalid, command.ID)
	}

	res, err := u.accountRepository.FindByID(ctx, id)
	if err != nil {
		return empty, err
	}

	acct, err := u.updateAccount(&res, command)
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

	response := dto.UpdateAccountResponseDto{ID: command.ID}

	return response, err
}

func (u *UpdateAccountCommandHandler) updateAccount(
	acct *account.Entity,
	cmd *UpdateAccountCommand,
) (account.Entity, error) {
	empty := account.Entity{}

	err := acct.Person.UpdateDetails(person.UpdateDetailsInput{
		FirstName:   cmd.FirstName,
		LastName:    cmd.LastName,
		Phone:       cmd.Phone,
		DateOfBirth: cmd.DateOfBirth,
	})
	if err != nil {
		u.log.Error("person update failed", map[string]any{"command": cmd, "error": err})
		return empty, err
	}

	var addrInputs []account.UpdateAddressInput
	for _, addr := range *cmd.Locations {
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

	addrerr := acct.UpdateAddress(addrInputs)
	if addrerr != nil {
		u.log.Error(
			"address update failed",
			map[string]any{"command": cmd, "input": addrInputs, "error": err},
		)
		return empty, err
	}

	return *acct, nil
}
