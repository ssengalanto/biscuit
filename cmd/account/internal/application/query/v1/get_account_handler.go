package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/account"
	"github.com/ssengalanto/potato-project/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/potato-project/pkg/errors"
	"github.com/ssengalanto/potato-project/pkg/interfaces"
)

// GetAccountQueryHandler - query handler struct for account creation, satisfies mediatr.RequestHandler.
type GetAccountQueryHandler struct {
	log               interfaces.Logger
	accountRepository account.Repository
	cache             account.Cache
}

// NewGetAccountQueryHandler creates a new query handler that handles account retrieval.
func NewGetAccountQueryHandler(
	logger interfaces.Logger,
	accountRepository account.Repository,
	cache account.Cache,
) *GetAccountQueryHandler {
	return &GetAccountQueryHandler{
		log:               logger,
		accountRepository: accountRepository,
		cache:             cache,
	}
}

func (c *GetAccountQueryHandler) Name() string {
	return fmt.Sprintf("%T", &GetAccountQuery{})
}

func (c *GetAccountQueryHandler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	var result account.Entity
	empty := dto.GetAccountResponseDto{}

	query, ok := request.(*GetAccountQuery)
	if !ok {
		c.log.Error("invalid query", map[string]any{"query": query})
		return nil, fmt.Errorf("%w: query", errors.ErrInvalid)
	}

	cachedAccount, err := c.cache.Get(ctx, query.ID)
	if err != nil {
		return empty, err
	}

	if cachedAccount != nil {
		return transformResponse(*cachedAccount), err
	}

	id, err := uuid.Parse(query.ID)
	if err != nil {
		c.log.Error("invalid uuid", map[string]any{"query": query, "error": err})
		return nil, fmt.Errorf("%w: uuid %s", errors.ErrInvalid, query.ID)
	}

	result, err = c.accountRepository.FindByID(ctx, id)
	if err != nil {
		return empty, err
	}

	c.cache.Set(ctx, result.ID.String(), result)

	return transformResponse(result), err
}

func transformResponse(entity account.Entity) dto.GetAccountResponseDto {
	var locations []dto.LocationResponseDto

	for _, addr := range *entity.Person.Address {
		location := dto.LocationResponseDto{
			ID:         addr.ID.String(),
			Street:     addr.Components.Street,
			Unit:       addr.Components.Unit,
			City:       addr.Components.City,
			District:   addr.Components.District,
			State:      addr.Components.State,
			Country:    addr.Components.Country,
			PostalCode: addr.Components.PostalCode,
		}
		locations = append(locations, location)
	}

	return dto.GetAccountResponseDto{
		ID:     entity.ID.String(),
		Email:  entity.Email.String(),
		Active: entity.Active,
		Person: dto.PersonResponseDto{
			ID:          entity.Person.ID.String(),
			FirstName:   entity.Person.Details.FirstName,
			LastName:    entity.Person.Details.LastName,
			Email:       entity.Person.Details.Email,
			Phone:       entity.Person.Details.Phone,
			DateOfBirth: entity.Person.Details.DateOfBirth,
		},
		Locations: locations,
	}
}
