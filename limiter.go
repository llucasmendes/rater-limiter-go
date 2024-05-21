package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimiter struct {
	client         *redis.Client
	rateLimitIP    int
	rateLimitToken int
	blockDuration  int
}

func NewRateLimiter(client *redis.Client, config Config) *RateLimiter {
	return &RateLimiter{
		client:         client,
		rateLimitIP:    config.RateLimitIP,
		rateLimitToken: config.RateLimitToken,
		blockDuration:  config.BlockDuration,
	}
}

func (rl *RateLimiter) AllowRequest(ctx context.Context, key string, limit int) (bool, error) {
	pipe := rl.client.TxPipeline()

	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, time.Duration(rl.blockDuration)*time.Second)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return false, err
	}

	return incr.Val() <= int64(limit), nil
}

func (rl *RateLimiter) GetLimit(token string) int {
	if token != "" {
		return rl.rateLimitToken
	}
	return rl.rateLimitIP
}
