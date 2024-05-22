package main

import (
	"context"
	"testing"
	"time"
)

type MockStore struct {
	data     map[string]int64
	expiries map[string]time.Duration
}

func NewMockStore() *MockStore {
	return &MockStore{
		data:     make(map[string]int64),
		expiries: make(map[string]time.Duration),
	}
}

func (m *MockStore) Incr(ctx context.Context, key string) (int64, error) {
	m.data[key]++
	return m.data[key], nil
}

func (m *MockStore) Expire(ctx context.Context, key string, duration time.Duration) error {
	m.expiries[key] = duration
	return nil
}

func TestRateLimiter(t *testing.T) {
	config := Config{
		RateLimitIP:    5,
		RateLimitToken: 10,
		BlockDuration:  60,
	}

	store := NewMockStore()
	limiter := NewRateLimiter(store, config)
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
