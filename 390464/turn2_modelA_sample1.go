package main

import (
	"fmt"
	"time"
)

// Worker function
func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d started processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate some processing
		fmt.Printf("Worker %d finished processing job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 5)

	// Create worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for processing to finish (you may use sync.WaitGroup for a real case)
	time.Sleep(6 * time.Second)
	fmt.Println("All jobs are processed.")
}
