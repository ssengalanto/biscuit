package query

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
	"github.com/ssengalanto/potato-project/pkg/mediatr"
)

// GetAccountQueryHandler - query handler struct for account creation, satisfies mediatr.RequestHandler.
type GetAccountQueryHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
}

// NewGetAccountQueryHandler creates a new query handler that handles account retrieval.
func NewGetAccountQueryHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
) *GetAccountQueryHandler {
	return &GetAccountQueryHandler{
		log:               logger,
		accountRepository: accountRepository,
	}
}

func (c *GetAccountQueryHandler) Name() string {
	return QueryGetAccount
}

func (c *GetAccountQueryHandler) Handle(
	ctx context.Context,
	request mediatr.Request,
) (any, error) {
	empty := dto.GetAccountResponseDto{}

	query, ok := request.(*GetAccountQuery)
	if !ok {
		c.log.Error("invalid query", map[string]any{"query": query})
		return nil, fmt.Errorf("%w: query %s", errors.ErrInvalid, query.Name())
	}

	id, err := uuid.Parse(query.ID)
	if err != nil {
		c.log.Error("invalid uuid", map[string]any{"query": query})
		return nil, fmt.Errorf("%w: uuid %s", errors.ErrInvalid, query.ID)
	}

	result, err := c.accountRepository.FindByID(ctx, id)
	if err != nil {
		c.log.Error("get account failed", map[string]any{"payload": query, "error": err})
		return empty, err
	}

	response := dto.GetAccountResponseDto{
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
