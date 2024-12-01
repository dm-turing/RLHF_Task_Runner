package main

import (
	"testing"
)

func findKey(d map[int]int, key int) bool {
	_, found := d[key]
	return found
}

func BenchmarkFindKeyMap(b *testing.B) {
	data := make(map[int]int)
	for i := 0; i < 10000; i++ {
		data[i] = i
	}
	key := 1000
	for i := 0; i < b.N; i++ {
		findKey(data, key)
	}
}

func findKeySlice(arr []int, key int) bool {
	for _, val := range arr {
		if val == key {
			return true
		}
	}
	return false
}

func BenchmarkFindKeySlice(b *testing.B) {
	data := make([]int, 10000)
	for i := 0; i < len(data); i++ {
		data[i] = i
	}
	key := 1000
	for i := 0; i < b.N; i++ {
		findKeySlice(data, key)
	}
}

// func main() {
// 	testing.Main()
// }
