package main

import (
	"fmt"
	"sync"
)

type FIFOQueue struct {
	queue []int
	lock  sync.Mutex
}

func (q *FIFOQueue) Enqueue(value int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, value)
}

func (q *FIFOQueue) Dequeue() (int, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) == 0 {
		return 0, false
	}
	value := q.queue[0]
	q.queue = q.queue[1:]
	return value, true
}

func worker(id int, q *FIFOQueue, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := q.Dequeue()
		if !ok {
			break // Exit if no tasks are left
		}
		fmt.Printf("Worker %d processing task %d\n", id, task)
	}
}

func main() {
	q := &FIFOQueue{}
	wg := &sync.WaitGroup{}

	// Enqueue tasks
	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
	}

	// Start workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, q, wg)
	}

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All tasks processed.")
}
