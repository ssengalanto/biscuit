package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/account"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/address"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/person"
	"github.com/ssengalanto/hex/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/hex/pkg/errors"
	"github.com/ssengalanto/hex/pkg/interfaces"
)

// CreateAccountCommandHandler - command handler struct for account retrieval, satisfies mediatr.RequestHandler.
type CreateAccountCommandHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
	cache             account.Cache
}

// NewCreateAccountCommandHandler creates a new command handler that handles account creation.
func NewCreateAccountCommandHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
	cache account.Cache,
) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		log:               logger,
		accountRepository: accountRepository,
		cache:             cache,
	}
}

func (c *CreateAccountCommandHandler) Name() string {
	return fmt.Sprintf("%T", &CreateAccountCommand{})
}

func (c *CreateAccountCommandHandler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	empty := dto.CreateAccountResponseDto{}
	entity := account.Entity{}

	command, ok := request.(*CreateAccountCommand)
	if !ok {
		c.log.Error("invalid command", map[string]any{"command": command})
		return empty, fmt.Errorf("%w: command", errors.ErrInvalid)
	}
	c.log.Info(fmt.Sprintf("executing %s", c.Name()), nil)

	acct, err := c.createAccount(command)
	if err != nil {
		return empty, err
	}

	pers, err := c.createPerson(acct.ID, command)
	if err != nil {
		return empty, err
	}

	addrs, err := c.createAddresses(pers.ID, command)
	if err != nil {
		return empty, err
	}

	entity = acct
	entity.Person = &pers
	entity.Person.Address = &addrs

	err = c.accountRepository.Create(ctx, entity)
	if err != nil {
		return empty, err
	}

	response := dto.CreateAccountResponseDto{ID: acct.ID.String()}
	return response, err
}

func (c *CreateAccountCommandHandler) createAccount(cmd *CreateAccountCommand) (account.Entity, error) {
	empty := account.Entity{}

	acct := account.New(cmd.Email, cmd.Password, cmd.Active)
	err := acct.IsValid()
	if err != nil {
		c.log.Error("account is invalid", map[string]any{"account": acct, "error": err})
		return empty, fmt.Errorf("%w: account", errors.ErrInvalid)
	}

	err = acct.HashPassword()
	if err != nil {
		c.log.Error("hashing password failed", map[string]any{"account": acct, "error": err})
		return empty, err
	}

	return acct, nil
}

func (c *CreateAccountCommandHandler) createPerson(
	accountID uuid.UUID,
	cmd *CreateAccountCommand,
) (person.Entity, error) {
	empty := person.Entity{}

	pers := person.New(accountID, cmd.FirstName, cmd.LastName, cmd.Email, cmd.Phone, cmd.DateOfBirth)
	err := pers.IsValid()
	if err != nil {
		c.log.Error("person is invalid", map[string]any{"person": pers, "error": err})
		return empty, fmt.Errorf("%w: person", errors.ErrInvalid)
	}

	return pers, nil
}

func (c *CreateAccountCommandHandler) createAddresses(
	personID uuid.UUID,
	cmd *CreateAccountCommand,
) ([]address.Entity, error) {
	var addrs []address.Entity

	c.log.Debug("entity", map[string]any{"entity": cmd.Locations})

	for _, location := range cmd.Locations {
		component := address.Components{
			Street:     location.Street,
			Unit:       location.Unit,
			City:       location.City,
			District:   location.District,
			State:      location.State,
			Country:    location.Country,
			PostalCode: location.PostalCode,
		}

		addr := address.New(personID, component)
		err := addr.IsValid()
		if err != nil {
			c.log.Error("address is invalid", map[string]any{"address": addr, "error": err})
			return nil, fmt.Errorf("%w: person", errors.ErrInvalid)
		}

		addrs = append(addrs, addr)
	}

	return addrs, nil
}
