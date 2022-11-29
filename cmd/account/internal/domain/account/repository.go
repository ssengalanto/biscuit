package account

import "context"

type Repository interface {
	Save(ctx context.Context, entity Entity) (Entity, error)
	FindOneById(ctx context.Context, id string) (Entity, error)
	FindOneByEmail(ctx context.Context, email string) (Entity, error)
	DeactivateAccount(ctx context.Context, id string) (Entity, error)
	ActivateAccount(ctx context.Context, id string) (Entity, error)
	DeleteOneById(ctx context.Context, id string) (Entity, error)
}
