package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/address"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/person"
	dtov1 "github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto/v1"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
)

// CreateAccountCommandHandler - command handler struct for account retrieval, satisfies midt.RequestHandler.
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
	empty := dtov1.CreateAccountResponse{}

	cmd := request.(*CreateAccountCommand) //nolint:errcheck //intentional panic

	acct, err := c.createAccount(cmd)
	if err != nil {
		return empty, err
	}

	pers, err := c.createPerson(acct.ID, cmd)
	if err != nil {
		return empty, err
	}

	addrs, err := c.createAddresses(pers.ID, cmd)
	if err != nil {
		return empty, err
	}

	entity := account.AggregateAccount(acct, pers, addrs)

	err = c.accountRepository.Create(ctx, entity)
	if err != nil {
		return empty, err
	}

	res := dtov1.CreateAccountResponse{ID: acct.ID.String()}
	return res, err
}

func (c *CreateAccountCommandHandler) createAccount(cmd *CreateAccountCommand) (account.Entity, error) {
	empty := account.Entity{}

	acct := account.New(cmd.Email, cmd.Password, cmd.Active)
	err := acct.IsValid()
	if err != nil {
		c.log.Error("account is invalid", map[string]any{"account": acct, "error": err})
		return empty, err
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
		return empty, err
	}

	return pers, nil
}

func (c *CreateAccountCommandHandler) createAddresses(
	personID uuid.UUID,
	cmd *CreateAccountCommand,
) ([]address.Entity, error) {
	var addrs []address.Entity

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
			return nil, err
		}

		addrs = append(addrs, addr)
	}

	return addrs, nil
}
