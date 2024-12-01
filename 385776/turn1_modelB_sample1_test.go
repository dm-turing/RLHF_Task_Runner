package main

import (
	"testing"
)

func slowFunction(n int) []int {
	result := make([]int, n)
	for i := range result {
		result[i] = i * i
	}
	return result
}

func BenchmarkSlowFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slowFunction(1000)
	}
}

func fastFunction(n int) []int {
	return make([]int, n)
}

func BenchmarkFastFunction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastFunction(1000)
	}
}

// func main() {
// 	testing.Main()
// }
