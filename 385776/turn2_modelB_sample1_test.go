// Example: Benchmarking a new sorting algorithm during development
package main

import (
	"testing"
)

func newSortingAlgorithm(arr []int) []int {
	// Implementation of the new sorting algorithm
	return arr
}

func BenchmarkNewSortingAlgorithm(b *testing.B) {
	arr := make([]int, 1000)
	for i := 0; i < len(arr); i++ {
		arr[i] = len(arr) - i - 1
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		newSortingAlgorithm(arr)
	}
}
