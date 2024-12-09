package main

import (
	"fmt"
	"sync"
)

type StripedMap struct {
	maps       []map[string]int
	mutexes    []sync.RWMutex
	numStripes int
}

func NewStripedMap(numStripes int) *StripedMap {
	return &StripedMap{
		maps:       make([]map[string]int, numStripes),
		mutexes:    make([]sync.RWMutex, numStripes),
		numStripes: numStripes,
	}
}

func (sm *StripedMap) getStripe(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % sm.numStripes
}

func (sm *StripedMap) Get(key string) int {
	stripe := sm.getStripe(key)
	sm.mutexes[stripe].RLock()
	defer sm.mutexes[stripe].RUnlock()
	return sm.maps[stripe][key]
}

func (sm *StripedMap) Put(key string, value int) {
	stripe := sm.getStripe(key)
	sm.mutexes[stripe].Lock()
	defer sm.mutexes[stripe].Unlock()
	if sm.maps[stripe] == nil {
		sm.maps[stripe] = make(map[string]int)
	}
	sm.maps[stripe][key] = value
}

func main() {
	sm := NewStripedMap(8) // 8 stripes for concurrency
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Put(fmt.Sprintf("key%d", i), i)
		}(i)
	}
	wg.Wait()
}
