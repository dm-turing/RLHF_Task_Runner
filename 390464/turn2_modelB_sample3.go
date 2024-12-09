package main

import (
	"container/heap"
	"fmt"
	"time"
)

type Item struct {
	value    string // The value of the item
	priority int    // The priority of the item
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority // Max-Heap for higher priority first
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Push(x interface{}) {
	heap.Push(pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	return heap.Pop(pq)
}

func processTask(task string) {
	fmt.Printf("Processing task: %s\n", task)
	time.Sleep(time.Second)
}

func main() {
	queue := make(PriorityQueue, 0)
	heap.Init(&queue)

	heap.Push(&queue, &Item{
		value:    "Task 3",
		priority: 3,
	})
	heap.Push(&queue, &Item{
		value:    "Task 1",
		priority: 1,
	})
	heap.Push(&queue, &Item{
		value:    "Task 2",
		priority: 2,
	})

	for queue.Len() > 0 {
		item := heap.Pop(&queue).(*Item)
		fmt.Printf("Processing high priority task: %s\n", item.value)
		processTask(item.value)
	}
}
