package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func RateLimitMiddleware(rl *RateLimiter) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := strings.Split(r.RemoteAddr, ":")[0]
			token := r.Header.Get("API_KEY")
			key := ip
			if token != "" {
				key = token
			}
			limit := rl.GetLimit(token)

			allowed, err := rl.AllowRequest(context.Background(), key, limit)
			if err != nil || !allowed {
				http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
