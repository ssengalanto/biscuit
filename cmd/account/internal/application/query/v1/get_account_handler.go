package v1

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/account/internal/domain/account"
	"github.com/ssengalanto/biscuit/cmd/account/internal/interfaces/dto"
	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
)

// GetAccountQueryHandler - query handler struct for account creation, satisfies midt.RequestHandler.
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

func (g *GetAccountQueryHandler) Name() string {
	return fmt.Sprintf("%T", &GetAccountQuery{})
}

func (g *GetAccountQueryHandler) Handle(
	ctx context.Context,
	request any,
) (any, error) {
	var result account.Entity
	empty := dto.GetAccountResponse{}

	query, ok := request.(*GetAccountQuery)
	if !ok {
		g.log.Error("invalid query", map[string]any{"query": query})
		return nil, fmt.Errorf("%w: query", errors.ErrInvalid)
	}
	g.log.Info(fmt.Sprintf("executing %s", g.Name()), nil)

	cachedAccount, err := g.cache.Get(ctx, query.ID)
	if err != nil {
		return empty, err
	}

	if cachedAccount != nil {
		return transformResponse(*cachedAccount), err
	}

	id, err := uuid.Parse(query.ID)
	if err != nil {
		g.log.Error("invalid uuid", map[string]any{"query": query, "error": err})
		return nil, fmt.Errorf("%w: uuid %s", errors.ErrInvalid, query.ID)
	}

	result, err = g.accountRepository.FindByID(ctx, id)
	if err != nil {
		return empty, err
	}

	g.cache.Set(ctx, result.ID.String(), result)

	return transformResponse(result), err
}

func transformResponse(entity account.Entity) dto.GetAccountResponse {
	var locations []dto.LocationResponse

	for _, addr := range *entity.Person.Address {
		location := dto.LocationResponse{
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

	return dto.GetAccountResponse{
		ID:     entity.ID.String(),
		Email:  entity.Email.String(),
		Active: entity.Active,
		Person: dto.PersonResponse{
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
