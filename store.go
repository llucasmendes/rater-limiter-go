package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimiterStore interface {
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, duration time.Duration) error
}

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client: client}
}

func (r *RedisStore) Incr(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

func (r *RedisStore) Expire(ctx context.Context, key string, duration time.Duration) error {
	return r.client.Expire(ctx, key, duration).Err()
}
