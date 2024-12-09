package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
	done <- true
}

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, done)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for w := 1; w <= 3; w++ {
		<-done
	}

	fmt.Println("All jobs are processed.")
}
