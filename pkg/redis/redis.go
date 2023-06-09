package redis

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ssengalanto/biscuit/pkg/constants"
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
func NewUniversalClient(host, port, pwd string, db int) (redis.UniversalClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ResourceTimeout)
	defer cancel()

	hp := net.JoinHostPort(host, port)
	c := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:           []string{hp},
		Password:        pwd,
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

	_, err := c.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnectionFailed, err)
	}

	return c, nil
}
