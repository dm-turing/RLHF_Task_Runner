package main

import (
	"context"
	"time"

	"github.com/go-redis/redis"
	"github.com/opencontrol/limiter"
)

type redisKeyResolver struct {
	rd *redis.Client
}

func (r redisKeyResolver) Resolve(ctx context.Context) (string, error) {
	ip := r.getRemoteIP(ctx)
	return ip, nil
}

func (r redisKeyResolver) getRemoteIP(ctx context.Context) string {
	// Implement your logic to get the IP address (e.g., from X-Forwarded-For header)
	return "your_ip_address_here"
}

func main() {
	rd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rd.Ping().Result()
	if err != nil {
		panic(err)
	}

	keyResolver := &redisKeyResolver{rd}
	l := limiter.NewLimiter(
		limiter.Every(time.Second),
		10,
		limiter.WithKeyResolver(keyResolver),
		limiter.WithStorage(limiter.NewRedisStorage(rd)),
	)
	// ... (rest of the code remains the same)
}
