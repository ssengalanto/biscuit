package v1

import (
	"context"
	"fmt"

	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
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
	entity := account.Entity{}

	command, ok := request.(*CreateAccountCommand)
	if !ok {
		c.log.Error("invalid command", map[string]any{"command": command})
		return nil, fmt.Errorf("%w: command", errors.ErrInvalid)
	}

	acct := account.New(command.Email, command.Password, command.Active)
	err := acct.IsValid()
	if err != nil {
		c.log.Error("account is invalid", map[string]any{"account": acct, "error": err})
		return nil, fmt.Errorf("%w: account", errors.ErrInvalid)
	}

	err = acct.HashPassword()
	if err != nil {
		c.log.Error("hashing password failed", map[string]any{"account": acct, "error": err})
		return nil, err
	}

	pers := person.New(acct.ID, command.FirstName, command.LastName, command.Email, command.Phone, command.DateOfBirth)
	err = pers.IsValid()
	if err != nil {
		c.log.Error("person is invalid", map[string]any{"person": pers, "error": err})
		return nil, fmt.Errorf("%w: person", errors.ErrInvalid)
	}

	entity = acct
	entity.Person = &pers

	err = c.accountRepository.Create(ctx, entity)
	if err != nil {
		return nil, err
	}

	response := dto.CreateAccountResponseDto{ID: acct.ID.String()}

	return response, err
}
