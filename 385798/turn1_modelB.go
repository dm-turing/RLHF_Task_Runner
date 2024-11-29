package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter int32 = 0

func increment() {
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func main() {
	numGoroutines := 10

	for i := 0; i < numGoroutines; i++ {
		go increment()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Final counter value:", atomic.LoadInt32(&counter))
}
