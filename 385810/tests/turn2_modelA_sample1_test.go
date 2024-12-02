package main

import (
	"net/http"
	"testing"
	"time"
)

func TestRateLimiterSuccess(t *testing.T) {
	// Set up your rate limiter here
	// ...

	var resp *http.Response

	// Simulate multiple successful requests within the limit
	for i := 0; i < 9; i++ {
		resp, _ = http.Get("https://httpbin.org/get")
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status code: %d", resp.StatusCode)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
