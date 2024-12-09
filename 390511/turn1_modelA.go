package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	m sync.Map
}

// Function to add or update an item in the cache
func (c *Cache) Store(key int, value string) {
	c.m.Store(key, value)
}

// Function to retrieve an item from the cache
func (c *Cache) Load(key int) (string, bool) {
	val, ok := c.m.Load(key)
	if ok {
		return val.(string), ok
	}
	return "", ok
}

func main() {
	cache := &Cache{}

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Sample data
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Store(i, fmt.Sprintf("Value %d", i))
		}(i)
	}

	// Read from cache concurrently
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if value, ok := cache.Load(i); ok {
				fmt.Printf("Loaded: key=%d, value=%s\n", i, value)
			} else {
				fmt.Printf("Key %d not found\n", i)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	time.Sleep(1 * time.Second) // Wait for a moment before program exit
}
