package main

import (
	"context"
	"time"
)

type RateLimiter struct {
	store          RateLimiterStore
	rateLimitIP    int
	rateLimitToken int
	blockDuration  int
}

func NewRateLimiter(store RateLimiterStore, config Config) *RateLimiter {
	return &RateLimiter{
		store:          store,
		rateLimitIP:    config.RateLimitIP,
		rateLimitToken: config.RateLimitToken,
		blockDuration:  config.BlockDuration,
	}
}

func (rl *RateLimiter) AllowRequest(ctx context.Context, key string, limit int) (bool, error) {
	count, err := rl.store.Incr(ctx, key)
	if err != nil {
		return false, err
	}
	err = rl.store.Expire(ctx, key, time.Duration(rl.blockDuration)*time.Second)
	if err != nil {
		return false, err
	}
	return count <= int64(limit), nil
}

func (rl *RateLimiter) GetLimit(token string) int {
	if token != "" {
		return rl.rateLimitToken
	}
	return rl.rateLimitIP
}
