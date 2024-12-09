package main

import (
	"fmt"
	"time"
)

func worker(id int, sem chan struct{}) {
	defer func() { <-sem }()
	fmt.Printf("Worker %d is working\n", id)
	time.Sleep(2 * time.Second) // Simulating work
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	const maxConcurrentWorkers = 3
	sem := make(chan struct{}, maxConcurrentWorkers)

	for i := 1; i <= 10; i++ {
		sem <- struct{}{} // Acquire a slot
		go worker(i, sem)
	}

	// Wait for all goroutines to finish
	time.Sleep(10 * time.Second)
	fmt.Println("All workers done.")
}
