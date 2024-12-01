package main

import (
	"sync"
	"testing"
)

var sharedData = 0
var mtx sync.Mutex

func BenchmarkExplicitLocking(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mtx.Lock()
		sharedData = i
		mtx.Unlock()
	}
}
