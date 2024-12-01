package main

import (
	"testing"
)

// A simple linked list node
type node struct {
	data int
	next *node
}

func appendToLinkedList(head **node, data int) {
	newNode := &node{data: data}
	if *head == nil {
		*head = newNode
	} else {
		current := *head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func BenchmarkAppendToLinkedList(b *testing.B) {
	head := (*node)(nil)
	for i := 0; i < b.N; i++ {
		appendToLinkedList(&head, i)
	}
}
