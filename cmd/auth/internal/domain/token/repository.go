package token

import (
	"context"

	"github.com/google/uuid"
)

// Repository is the token entity contract for infrastructure (persistence) layer.
type Repository interface {
	Create(ctx context.Context, entity Entity) error
	FindByID(ctx context.Context, id uuid.UUID) (Entity, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}
