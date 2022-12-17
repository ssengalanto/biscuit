package account

import (
	"context"
	"github.com/ssengalanto/potato-project/cmd/account/internal/domain/address"

	"github.com/google/uuid"
)

type Repository interface {
	Exists(ctx context.Context, id uuid.UUID) (bool, error)
	Create(ctx context.Context, entity Entity) (Entity, error)
	CreatePersonAddresses(ctx context.Context, entities []address.Entity) ([]address.Entity, error)
	FindByID(ctx context.Context, id uuid.UUID) (Entity, error)
	FindByEmail(ctx context.Context, email string) (Entity, error)
	Update(ctx context.Context, entity Entity) (Entity, error)
	DeleteByID(ctx context.Context, id uuid.UUID) (Entity, error)
}
