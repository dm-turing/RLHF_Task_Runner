package main

import (
	"fmt"
	"sync"
	"time"
)

type FIFOQueue struct {
	queue []string
	lock  sync.Mutex
}

func (q *FIFOQueue) Enqueue(task string) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, task)
}

func (q *FIFOQueue) Dequeue() (string, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if len(q.queue) == 0 {
		return "", false
	}
	task := q.queue[0]
	q.queue = q.queue[1:]
	return task, true
}

func processTask(task string) {
	fmt.Printf("Processing task: %s\n", task)
	time.Sleep(time.Second)
}

func main() {
	queue := &FIFOQueue{}

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
