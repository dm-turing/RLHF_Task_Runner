package main

import (
	"fmt"
	"unsafe"
)

type Node interface {
	GetData() int
}

type structNode struct {
	data int
}

func (s *structNode) GetData() int {
	return s.data
}

func (s *structNode) setData(newData int) {
	s.data = newData
}

type pointerNode struct {
	*structNode
}

func (p *pointerNode) GetData() int {
	return p.structNode.GetData()
}

func (p *pointerNode) setData(newData int) {
	p.structNode.setData(newData)
}

type Tree struct {
	root Node
}

func newTreeWithStructNode(data int) *Tree {
	return &Tree{root: &structNode{data: data}}
}

func newTreeWithPointerNode(data int) *Tree {
	return &Tree{root: &pointerNode{&structNode{data: data}}}
}

func main() {
	// Memory Model for structNode:
	//         +-----+
	//     root |     |
	//         +-----+
	//         |
	// +-----+-----+
	// | data |     |
	// +-----+-----+

	// Memory Model for pointerNode:
	//         +-----+
	//     root |     |
	//         +-----+
	//         |
	// +-----+-----+-----+
	// | structNode |     |
	// +-----+-----+-----+
	//         |
	// +-----+-----+
	// | data |     |
	// +-----+-----+

	t1 := newTreeWithStructNode(42)
	t2 := newTreeWithPointerNode(42)

	// Let's see how passing interface values behaves:
	var traverseTree func(Node)
	traverseTree = func(node Node) {
		if node != nil {
			fmt.Println(node.GetData())
			// traverseTree(node)
		}
	}
	traverseTree(t1.root) // Output: 42
	traverseTree(t2.root) // Output: 42

	// Now, let's compare the memory sizes:
	// NOTE: The actual memory sizes might vary depending on the platform and optimizations.
	// But the difference in model illustrates the point.
	var s *structNode
	var p *pointerNode

	fmt.Printf("Size of structNode: %v bytes\n", unsafe.Sizeof(s))
	fmt.Printf("Size of pointerNode: %v bytes\n", unsafe.Sizeof(p))
}
