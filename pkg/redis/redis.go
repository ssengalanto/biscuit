package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	maxRetries      = 5
	minRetryBackoff = 300 * time.Millisecond
	maxRetryBackoff = 500 * time.Millisecond
	dialTimeout     = 5 * time.Second
	readTimeout     = 5 * time.Second
	writeTimeout    = 3 * time.Second
	minIdleConn     = 20
	poolTimeout     = 6 * time.Second
	idleTimeout     = 12 * time.Second
	poolSize        = 300
)

// NewUniversalClient creates a new redis universal client.
func NewUniversalClient(host string, db int, password string) (redis.UniversalClient, error) {
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:           []string{host},
		Password:        password,
		DB:              db,
		MaxRetries:      maxRetries,
		MinRetryBackoff: minRetryBackoff,
		MaxRetryBackoff: maxRetryBackoff,
		DialTimeout:     dialTimeout,
		ReadTimeout:     readTimeout,
		WriteTimeout:    writeTimeout,
		PoolSize:        poolSize,
		MinIdleConns:    minIdleConn,
		PoolTimeout:     poolTimeout,
		IdleTimeout:     idleTimeout,
	})

	_, err := c.Ping(context.Background()).Result()
	if err != nil {
		return nil, ErrConnectionFailed
	}

	return c, nil
}
