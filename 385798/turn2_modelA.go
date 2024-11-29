package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	mutex     sync.Mutex
	atomicVar int32 = 0
	lockedVar int32 = 0
)

func atomicIncrement() {
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&atomicVar, 1)
	}
}

func lockedIncrement() {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		lockedVar++
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	numGoroutines := 10

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go atomicIncrement()
	}

	for i := 0; i < numGoroutines; i++ {
		go lockedIncrement()
	}

	for i := 0; i < numGoroutines; i++ {
		<-chan struct{}{} // Wait for goroutines to finish
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
	fmt.Println("Atomic counter value:", atomic.LoadInt32(&atomicVar))
	fmt.Println("Locked counter value:", lockedVar)
}