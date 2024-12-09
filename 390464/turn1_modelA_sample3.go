package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // The value of the item
	priority int    // The priority of the item
	index    int    // The index of the item in the heap
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority // Min-Heap
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func main() {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Item{
		value:    "task1",
		priority: 2,
	})
	heap.Push(&pq, &Item{
		value:    "task2",
		priority: 1,
	})
	heap.Push(&pq, &Item{
		value:    "task3",
		priority: 3,
	})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Processed %s with priority %d\n", item.value, item.priority)
	}
}
