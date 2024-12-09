package main

import (
	"fmt"
	"sync"
	"time"
)

type BoundedQueue struct {
	queue []int
	limit int
	lock  sync.Mutex
	cond  *sync.Cond
}

func NewBoundedQueue(limit int) *BoundedQueue {
	bq := &BoundedQueue{
		queue: make([]int, 0),
		limit: limit,
	}
	bq.cond = sync.NewCond(&bq.lock)
	return bq
}

func (bq *BoundedQueue) Enqueue(value int) {
	bq.lock.Lock()
	defer bq.lock.Unlock()

	for len(bq.queue) == bq.limit {
		bq.cond.Wait() // Wait if the queue is full
	}
	bq.queue = append(bq.queue, value)
	bq.cond.Signal()
}

func (bq *BoundedQueue) Dequeue() (int, bool) {
	bq.lock.Lock()
	defer bq.lock.Unlock()

	for len(bq.queue) == 0 {
		bq.cond.Wait() // Wait if the queue is empty
	}
	value := bq.queue[0]
	bq.queue = bq.queue[1:]
	bq.cond.Signal()
	return value, true
}

func worker(id int, q *BoundedQueue) {
	for i := 0; i < 5; i++ {
		task := i + id*5 // Unique task identifier
		q.Enqueue(task)
		fmt.Printf("Worker %d enqueued task %d\n", id, task)
		time.Sleep(500 * time.Millisecond) // Throttle to simulate work
	}
}

func main() {
	bq := NewBoundedQueue(5)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, bq)
		}(i)
	}

	wg.Wait() // Wait for all workers to finish

	// Dequeue all tasks
	for i := 1; i <= 15; i++ {
		task, ok := bq.Dequeue()
		if !ok {
			break
		}
		fmt.Printf("Processed task %d\n", task)
	}
}
