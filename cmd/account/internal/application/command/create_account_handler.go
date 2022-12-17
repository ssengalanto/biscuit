package command

import (
	"context"

	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/person"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
)

type CreateAccountCommandHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
}

func NewCreateAccountCommandHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
) *CreateAccountCommandHandler {
	return &CreateAccountCommandHandler{
		log:               logger,
		accountRepository: accountRepository,
	}
}

func (c *CreateAccountCommandHandler) Handle(
	ctx context.Context,
	command *CreateAccountCommand,
) (dto.CreateAccountResponseDto, error) {
	entity := account.Entity{}
	empty := dto.CreateAccountResponseDto{}

	acct := account.New(command.Email, command.Password, command.Active)
	err := acct.IsValid()
	if err != nil {
		c.log.Error("account is invalid", map[string]any{"account": acct, "error": err})
		return empty, err
	}

	pers := person.New(acct.ID, command.FirstName, command.LastName, command.Email, command.Phone, command.DateOfBirth)
	err = pers.IsValid()
	if err != nil {
		c.log.Error("person is invalid", map[string]any{"person": pers, "error": err})
		return empty, err
	}

	entity = acct
	entity.Person = &pers

	result, err := c.accountRepository.Create(ctx, entity)
	if err != nil {
		c.log.Error("account creation failed", map[string]any{"payload": entity, "error": err})
		return empty, err
	}

	response := dto.CreateAccountResponseDto{
		ID:     result.ID.String(),
		Email:  result.Email.String(),
		Active: result.Active,
		Person: dto.PersonResponseDto{
			ID:          result.Person.ID.String(),
			FirstName:   result.Person.Details.FirstName,
			LastName:    result.Person.Details.LastName,
			Email:       result.Person.Details.Email,
			Phone:       result.Person.Details.Phone,
			DateOfBirth: result.Person.Details.DateOfBirth,
		},
	}

	return response, err
}
