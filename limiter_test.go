package main

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRateLimiter(t *testing.T) {
	config := Config{
		RedisAddr:      "localhost:6379",
		RedisPassword:  "",
		RateLimitIP:    5,
		RateLimitToken: 10,
		BlockDuration:  60,
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPassword,
		DB:       0,
	})

	limiter := NewRateLimiter(rdb, config)
	ctx := context.Background()
	key := "test-ip"

	for i := 0; i < 5; i++ {
		allowed, err := limiter.AllowRequest(ctx, key, config.RateLimitIP)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if !allowed {
			t.Fatalf("Expected allowed, got %v", allowed)
		}
	}

	allowed, err := limiter.AllowRequest(ctx, key, config.RateLimitIP)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if allowed {
		t.Fatalf("Expected not allowed, got %v", allowed)
	}
}
