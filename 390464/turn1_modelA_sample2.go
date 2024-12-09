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

func main() {
	q := &FIFOQueue{}
	for i := 1; i <= 5; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 5; i++ {
		if value, ok := q.Dequeue(); ok {
			fmt.Println(value)
		}
	}
}
