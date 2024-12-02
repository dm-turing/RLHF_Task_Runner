package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
)

const (
	maxRequestsPerMinute = 60
	maxRequestsPerHour   = 3600
)

type RateLimit struct {
	gorm.Model
	IP        string `json:"ip"`
	Limit     int64  `json:"limit"`
	Remaining int64  `json:"remaining"`
	ResetTime int64  `json:"reset_time"`
}

type rateLimiter struct {
	db *gorm.DB
}

func newRateLimiter(db *gorm.DB) *rateLimiter {
	return &rateLimiter{db}
}

func (rl *rateLimiter) limit(r *http.Request, limit int64) bool {
	ip := r.RemoteAddr
	rl.updateLimit(ip, limit)

	limitC := rl.getLimit(ip)
	if limitC.Remaining == 0 {
		time.Sleep(time.Until(time.Unix(limitC.ResetTime, 0)))
		rl.updateLimit(ip, limit)
	}

	if limitC.Remaining <= 0 {
		w := http.ResponseWriter
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		http.ServeJSON(w, nil, map[string]interface{}{
			"error":     "too many requests",
			"retry":     strconv.Itoa(int(time.Until(time.Unix(limitC.ResetTime, 0)).Seconds())),
			"limit":     limit,
			"remaining": 0,
			"reset":     limitC.ResetTime,
		})
		return false
	}

	return true
}

func (rl *rateLimiter) getLimit(ip string) *RateLimit {
	var limit RateLimit
	rl.db.Where("ip = ?", ip).First(&limit)
	return &limit
}

func (rl *rateLimiter) updateLimit(ip string, limit int64) {
	now := time.Now().Unix()
	limitC := rl.getLimit(ip)
	remaining := limitC.Remaining
	resetTime := limitC.ResetTime

	if remaining == 0 {
		remaining = limit
		resetTime = now + (limit/maxRequestsPerMinute)*60
	} else {
		remaining--
	}

	rl.db.Save(&RateLimit{
		Model:     limitC.Model,
		IP:        ip,
		Limit:     limit,
		Remaining: remaining,
		ResetTime: resetTime,
	})
}

func (rl *rateLimiter) cleanup() {
	c := cron.New()
	c.Schedule("@midnight", func() {
		rl.db.Delete(&RateLimit{})
	})
	c.Start()
}

func main() {
	// Initialize your database here
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create rate limiter
	limiter := newRateLimiter(db)

	// Register routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if limiter.limit(r, maxRequestsPerMinute) {
			fmt.Fprintf(w, "Hello, world!")
		}
	})

	// Start rate limiter cleanup cron job
	limiter.cleanup()

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
