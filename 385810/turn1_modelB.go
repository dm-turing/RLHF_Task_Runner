package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/limiter"
)

// RateLimiter is a struct that holds the limiter and its configuration.
type RateLimiter struct {
	limiter *limiter.Limiter
}

// NewRateLimiter creates a new rate limiter with the given configuration.
func NewRateLimiter(maxRequestsPerSecond int) *RateLimiter {
	l := limiter.New()
	// Create a rate limit of 10 requests per second for each IP address
	l.AddRateLimit(limiter.NewRateLimit(limiter.Every(time.Second), maxRequestsPerSecond))
	return &RateLimiter{limiter: l}
}

// LimitRequest middleware function that limits the request rate.
func (rl *RateLimiter) LimitRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "limiter", rl.limiter)
		// Check if the request is allowed based on the rate limit
		if ok := rl.limiter.Allow(ctx); !ok {
			http.Error(w, "Too many requests. Please try again later.", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	rateLimiter := NewRateLimiter(10) // Limit to 10 requests per second

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	})

	// Apply the rate limiter middleware to specific routes or the entire app
	http.Handle("/api/", rateLimiter.LimitRequest(http.StripPrefix("/api/", http.DefaultServeMux)))

	fmt.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
