package main

import (
	"testing"
)

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		result = append(result, left[i])
	}

	for ; j < len(right); j++ {
		result = append(result, right[j])
	}

	return result
}

func BenchmarkBubbleSort(b *testing.B) {
	arr := make([]int, 10000)
	for i := 0; i < len(arr); i++ {
		arr[i] = len(arr) - i - 1
	}
	for i := 0; i < b.N; i++ {
		bubbleSort(arr)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	arr := make([]int, 10000)
	for i := 0; i < len(arr); i++ {
		arr[i] = len(arr) - i - 1
	}
	for i := 0; i < b.N; i++ {
		mergeSort(arr)
	}
}

// func main() {
// 	testing.Main()
// }
