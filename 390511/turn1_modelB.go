package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ConcurrentHashMap struct {
	m  map[string]int
	mu sync.RWMutex
}

func NewConcurrentHashMap() *ConcurrentHashMap {
	return &ConcurrentHashMap{m: make(map[string]int)}
}

func (chm *ConcurrentHashMap) Get(key string) int {
	chm.mu.RLock()
	defer chm.mu.RUnlock()
	return chm.m[key]
}

func (chm *ConcurrentHashMap) Put(key string, value int) {
	chm.mu.Lock()
	defer chm.mu.Unlock()
	chm.m[key] = value
}

func (chm *ConcurrentHashMap) Increment(key string) {
	chm.mu.Lock()
	defer chm.mu.Unlock()
	chm.m[key]++
}

func (chm *ConcurrentHashMap) Decrement(key string) {
	chm.mu.Lock()
	defer chm.mu.Unlock()
	if chm.m[key] > 0 {
		chm.m[key]--
	}
}

func main() {
	chm := NewConcurrentHashMap()
	urls := []string{"/", "/about", "/contact", "/", "/about", "/", "/contact", "/about"}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			chm.Increment(u)
		}(url)
	}

	wg.Wait()

	fmt.Println("URL Counts:")
	for url, count := range chm.m {
		fmt.Printf("%s: %d\n", url, count)
	}
}
