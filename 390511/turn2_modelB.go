package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ConcurrentHashMap struct {
	n       int
	stripes []*sync.RWMutex
	maps    []map[string]int
}

func NewConcurrentHashMap(n int) *ConcurrentHashMap {
	return &ConcurrentHashMap{
		n:       n,
		stripes: make([]*sync.RWMutex, n),
		maps:    make([]map[string]int, n),
	}
}

func (chm *ConcurrentHashMap) hash(key string) int {
	return hash(key) % chm.n
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = 31*h + int(c)
	}
	return h
}

func (chm *ConcurrentHashMap) Get(key string) int {
	idx := chm.hash(key)
	chm.stripes[idx].RLock()
	defer chm.stripes[idx].RUnlock()
	return chm.maps[idx][key]
}

func (chm *ConcurrentHashMap) Put(key string, value int) {
	idx := chm.hash(key)
	chm.stripes[idx].Lock()
	defer chm.stripes[idx].Unlock()
	chm.maps[idx][key] = value
}

func (chm *ConcurrentHashMap) Increment(key string) {
	idx := chm.hash(key)
	chm.stripes[idx].Lock()
	defer chm.stripes[idx].Unlock()
	chm.maps[idx][key]++
}

func (chm *ConcurrentHashMap) Decrement(key string) {
	idx := chm.hash(key)
	chm.stripes[idx].Lock()
	defer chm.stripes[idx].Unlock()
	if chm.maps[idx][key] > 0 {
		chm.maps[idx][key]--
	}
}

func main() {
	chm := NewConcurrentHashMap(16) // Number of stripes
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
	for i := range chm.maps {
		chm.stripes[i].RLock()
		defer chm.stripes[i].RUnlock()
		for url, count := range chm.maps[i] {
			fmt.Printf("%s: %d\n", url, count)
		}
	}
}
