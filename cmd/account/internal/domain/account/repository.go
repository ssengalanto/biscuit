package account

import "context"

type Repository interface {
	Save(ctx context.Context, entity Entity) (Entity, error)
	FindByID(ctx context.Context, id string) (Entity, error)
	FindByEmail(ctx context.Context, email string) (Entity, error)
	DeactivateAccount(ctx context.Context, id string) (Entity, error)
	ActivateAccount(ctx context.Context, id string) (Entity, error)
	DeleteById(ctx context.Context, id string) (Entity, error)
}
