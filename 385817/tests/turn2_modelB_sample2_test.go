package main

import (
	"sync"
	"testing"
)

var sharedData = 0
var mu sync.Mutex

func BenchmarkExplicitLocking(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		sharedData = i
		mu.Unlock()
	}
}
