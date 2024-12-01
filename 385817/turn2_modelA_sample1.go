package main

import (
	"math/rand"
	"sync"
	"testing"
)

func init() {
	rand.Seed(int64(42)) // Random seed for consistent tests
}

var shareVal = 0
var mtx = &sync.Mutex{}
var chanSize = 10000
var ch = make(chan int, chanSize)

// BenchmarkProducerChannel : Test concurrent writes with channel.
func BenchmarkProducerChannel(b *testing.B) {
	for i := 0; i < 1000; i++ {
		ch <- rand.Int()
	}
}

// BenchmarkConsumerChannel : Test concurrent reads with channel.
func BenchmarkConsumerChannel(b *testing.B) {
	for i := 0; i < 1000; i++ {
		<-ch
	}
}

// BenchmarkProducerLock : Test concurrent writes with mutex.
func BenchmarkProducerLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mtx.Lock()
		shareVal = rand.Int()
		mtx.Unlock()
	}
}

// BenchmarkConsumerLock : Test concurrent reads with mutex.
func BenchmarkConsumerLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mtx.Lock()
		_ = shareVal
		mtx.Unlock()
	}
}
