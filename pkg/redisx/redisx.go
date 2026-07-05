// Package redisx exposes a shared Redis client initialized once per process.
package redisx

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Config holds Redis connection parameters. Password may be empty (no auth).
type Config struct {
	Addr     string
	Password string
	DB       int
}

// Client is the shared Redis connection, mirroring the postgres.DB global.
var Client *redis.Client

// Conn opens the Redis client from cfg and verifies connectivity. Callers should
// treat a non-nil error as fatal — most Tentaculum services rely on Redis for
// rate limiting or session/2FA state and cannot boot without it.
func Conn(cfg Config) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("redis ping %s: %w", cfg.Addr, err)
	}
	return nil
}

// Ping reports whether Redis is reachable (used by health checks).
func Ping(ctx context.Context) bool {
	if Client == nil {
		return false
	}
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
	}
	return Client.Ping(ctx).Err() == nil
}
