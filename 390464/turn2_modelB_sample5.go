package main

import (
	"fmt"
	"sync"
)

type LIFOQueue struct {
	queue []string
	lock  sync.Mutex
}

func (q *LIFOQueue) Enqueue(task string) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, task)
}

func (q *LIFOQueue) Dequeue() (string, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) == 0 {
		return "", false
	}
	task := q.queue[len(q.queue)-1]
	q.queue = q.queue[:len(q.queue)-1]
	return task, true
}

func processTask(task string) {
	fmt.Printf("Processing task: %s\n", task)
}

func main() {
	queue := &LIFOQueue{}

	for _, task := range []string{"A", "B", "C", "D", "E"} {
		go func(t string) {
			queue.Enqueue(t)
		}(task)
	}

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				task, ok := queue.Dequeue()
				if !ok {
					break
				}
				processTask(task)
			}
		}()
	}

	wg.Wait()
}
