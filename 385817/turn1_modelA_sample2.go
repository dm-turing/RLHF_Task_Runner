package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedCounter = 0
var increments = 1000000
var mu sync.Mutex

func incrementCounter() {
	for i := 0; i < increments; i++ {
		mu.Lock()
		sharedCounter++
		mu.Unlock()
	}
}

func main() {
	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		go incrementCounter()
	}

	time.Sleep(2 * time.Second) // Wait for goroutines to complete
	fmt.Println("Final shared counter value:", sharedCounter)
}
