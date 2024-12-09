package main

import (
	"fmt"
	"sync"
	"time"
)

func makeRequest(url string, sem chan struct{}) {
	defer func() { <-sem }()
	fmt.Printf("Making request to %s\n", url)
	time.Sleep(2 * time.Second)
	fmt.Printf("Request to %s completed\n", url)
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 5) // Limit to 5 concurrent requests
	urls := []string{
		"https://example.com/1", "https://example.com/2",
		"https://example.com/3", "https://example.com/4",
		"https://example.com/5", "https://example.com/6",
		"https://example.com/7", "https://example.com/8",
		"https://example.com/9", "https://example.com/10",
	}

	for _, url := range urls {
		sem <- struct{}{}
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			makeRequest(u, sem)
		}(url)
	}

	wg.Wait()
}
