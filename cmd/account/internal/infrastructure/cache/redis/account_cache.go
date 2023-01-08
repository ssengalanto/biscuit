package redis

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/ssengalanto/hex/cmd/account/internal/domain/account"
	"github.com/ssengalanto/hex/pkg/interfaces"
)

type AccountCache struct {
	log    interfaces.Logger
	client redis.UniversalClient
}

// New creates a new AccountCache instance.
func New(log interfaces.Logger, client redis.UniversalClient) *AccountCache {
	return &AccountCache{log: log, client: client}
}

const accountCacheKeyPrefix = "account_cache"

func (a *AccountCache) Set(ctx context.Context, key string, value account.Entity) {
	data, err := json.Marshal(value)
	if err != nil {
		a.log.Error("json marshal failed", map[string]any{"key": key, "value": value, "error": err})
		panic(err)
	}

	a.client.HSetNX(ctx, a.keyPrefix(), key, data)
	a.log.Info("cache set", map[string]any{"key": key, "value": value})
}

func (a *AccountCache) Get(ctx context.Context, key string) (*account.Entity, error) {
	data, err := a.client.HGet(ctx, a.keyPrefix(), key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil //nolint:nilnil //necessary
		}
		a.log.Error("cache retrieval failed", map[string]any{"key": key, "error": err})
		return nil, err
	}

	var account account.Entity
	if err = json.Unmarshal(data, &account); err != nil {
		a.log.Error("json unmarshal failed", map[string]any{"key": key, "error": err})
		return nil, err
	}

	a.log.Info("cache retrieved", map[string]any{"key": key, "value": account})
	return &account, nil
}

func (a *AccountCache) Delete(ctx context.Context, key string) error {
	err := a.client.HDel(ctx, a.keyPrefix(), key).Err()
	if err != nil {
		a.log.Error("cache deletion failed", map[string]any{"key": key, "error": err})
		return err
	}

	a.log.Info("cache deleted", map[string]any{"key": key})
	return nil
}

func (a *AccountCache) keyPrefix() string {
	return accountCacheKeyPrefix
}
