package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type ConcurrentCounter struct {
	counts map[string]*int32
	mu     sync.Mutex
}

func NewConcurrentCounter() *ConcurrentCounter {
	return &ConcurrentCounter{counts: make(map[string]*int32)}
}

func (cc *ConcurrentCounter) Get(key string) int32 {
	cc.mu.Lock()
	defer cc.mu.Unlock()
	if count, exists := cc.counts[key]; exists {
		return atomic.LoadInt32(count)
	}
	return 0
}

func (cc *ConcurrentCounter) Increment(key string) {
	cc.mu.Lock() // Lock just for map access, not for atomic operations.
	defer cc.mu.Unlock()
	if count, exists := cc.counts[key]; exists {
		atomic.AddInt32(count, 1)
	} else {
		var initialValue int32
		cc.counts[key] = &initialValue
		atomic.AddInt32(cc.counts[key], 1)
	}
}

func main() {
	cc := NewConcurrentCounter()
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cc.Increment("key")
		}(i)
	}
	wg.Wait()

	fmt.Println("Final count:", cc.Get("key")) // should print 1000
}
