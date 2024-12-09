package main

import (
	"fmt"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	stack := &Stack{}

	// Push tasks onto the stack
	stack.Push("Task 1")
	stack.Push("Task 2")
	stack.Push("Task 3")

	// Pop and process tasks
	for len(stack.items) > 0 {
		task := stack.Pop()
		fmt.Printf("Processing %s\n", task) // Will process in LIFO order
	}
}
