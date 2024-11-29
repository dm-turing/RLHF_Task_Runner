package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counterAtomic int32 = 0
	counterMutex  int32 = 0
	mutex         sync.Mutex
)

func incrementAtomic() {
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&counterAtomic, 1)
	}
}

func incrementMutex() {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		counterMutex++
		mutex.Unlock()
	}
}

func main() {
	numGoroutines := runtime.GOMAXPROCS(0)

	fmt.Println("Starting test with", numGoroutines, "goroutines")

	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		go incrementAtomic()
		go incrementMutex()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Final atomic counter:", atomic.LoadInt32(&counterAtomic))
	fmt.Println("Final mutex counter:", counterMutex)
	fmt.Println("Elapsed time:", time.Since(start))
}
