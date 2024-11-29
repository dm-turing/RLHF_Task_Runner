package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter int32 = 0

func incrementCounter(n int) {
	for i := 0; i < n; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func main() {
	const nThreads = 10
	const iterations = 1_000_000

	go func() {
		defer time.Sleep(time.Second)
		fmt.Println("Final Counter:", atomic.LoadInt32(&counter))
	}()

	for i := 0; i < nThreads; i++ {
		go incrementCounter(iterations)
	}

	time.Sleep(2 * time.Second)
}
