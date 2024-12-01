package main

import (
	"fmt"
	"testing"
)

// Calculates the Fibonacci number
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func TestFib(t *testing.T) {
	fmt.Println(fib(35)) // Result: 9227465
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(25)
	}
}
