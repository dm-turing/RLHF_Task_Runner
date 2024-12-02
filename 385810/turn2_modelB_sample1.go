package main

import (
	"time"

	"github.com/patrickmn/go-cache"
)

func tokenBucketRateLimiter(key string, rate float64) bool {
	cache := cache.New(time.Minute, time.Minute)

	// Calculate tokens per second
	tokensPerSecond := rate
	item, _ := cache.Get(key)
	tokens := item.(int)
	// Generate tokens based on the current time and rate
	now := time.Now()
	item, _ = cache.Get(key + "_last_update")
	lastUpdate := item.(time.Time)
	duration := now.Sub(lastUpdate)
	tokens += int(duration.Seconds() * tokensPerSecond)

	// Limit tokens to a maximum capacity
	maxTokens := int(rate * 10) // Assuming 10 seconds of token buffer
	if tokens > maxTokens {
		tokens = maxTokens
	}

	// Check if there are enough tokens for the request
	if tokens > 0 {
		tokens--
		cache.Set(key, tokens, time.Second)
		cache.Set(key+"_last_update", now, time.Second)
		return true
	}
	return false
}
