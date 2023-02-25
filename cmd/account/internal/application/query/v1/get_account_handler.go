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
	var res account.Entity
	empty := dtov1.GetAccountResponse{}

	q := request.(*GetAccountQuery) //nolint:errcheck //intentional panic

	cachedAcct, _ := g.cache.Get(ctx, q.ID) //nolint:errcheck,nolintlint //unnecessary
	if cachedAcct != nil {
		return transformResponse(*cachedAcct), nil
	}

	id, err := uuid.Parse(q.ID)
	if err != nil {
		g.log.Error("invalid uuid", map[string]any{"q": q, "error": err})
		return empty, fmt.Errorf("%w: uuid `%s`", errors.ErrInvalid, q.ID)
	}

	res, err = g.accountRepository.FindByID(ctx, id)
	if err != nil {
		return empty, err
	}

	g.cache.Set(ctx, res.ID.String(), res)

	return transformResponse(res), err
}

func transformResponse(entity account.Entity) dtov1.GetAccountResponse {
	var locations []dtov1.LocationResponse

	if entity.Person.Address != nil {
		for _, addr := range *entity.Person.Address {
			location := dtov1.LocationResponse{
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
	}

	return dtov1.GetAccountResponse{
		ID:     entity.ID.String(),
		Email:  entity.Email.String(),
		Active: entity.Active,
		Person: dtov1.PersonResponse{
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
