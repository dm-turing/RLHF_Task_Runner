package main

import (
	"fmt"
	"time"
)

func worker(id int, sem chan struct{}) {
	defer func() { <-sem }() // Release the semaphore when done
	// Simulate work
	fmt.Printf("Worker %d is working\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	const maxConcurrentWorkers = 3
	sem := make(chan struct{}, maxConcurrentWorkers)

	for i := 1; i <= 10; i++ {
		sem <- struct{}{} // Acquire semaphore
		go worker(i, sem)
	}

	// Wait for all workers to finish
	time.Sleep(10 * time.Second)
	fmt.Println("All workers done.")
}
