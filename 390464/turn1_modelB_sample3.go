package main

import (
	"container/heap"
	"fmt"
	"time"
)

type PriorityItem struct {
	Priority int
	Value    int
}

func (p PriorityItem) Less(q interface{}) bool {
	return p.Priority < q.(PriorityItem).Priority
}

func main() {
	var pq []PriorityItem
	heap.Init(&pq)

	// Add items to the priority queue
	heap.Push(&pq, PriorityItem{Priority: 3, Value: 1})
	heap.Push(&pq, PriorityItem{Priority: 1, Value: 2})
	heap.Push(&pq, PriorityItem{Priority: 2, Value: 3})

	// Remove items from the priority queue
	for len(pq) > 0 {
		item := heap.Pop(&pq).(PriorityItem)
		fmt.Println("Popped:", item.Value)
		time.Sleep(time.Millisecond * 100)
	}
}
