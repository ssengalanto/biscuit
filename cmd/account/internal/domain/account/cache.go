package account

import (
	"context"
)

// Cache - account entity contract for infrastructure (cache) layer.
type Cache interface {
	Set(ctx context.Context, key string, value Entity)
	Get(ctx context.Context, key string) (*Entity, error)
	Delete(ctx context.Context, key string) error
}
