package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter int32 = 0

const threshold = 10000

func incrementAndCheck() {
	for {
		old := atomic.LoadInt32(&counter)
		if old >= threshold {
			return
		}
		atomic.AddInt32(&counter, 1)
		if atomic.LoadInt32(&counter) == threshold {
			fmt.Println("Reached threshold by:", old+1)
			return
		}
	}
}

func main() {
	go incrementAndCheck()
	go incrementAndCheck()
	time.Sleep(2 * time.Second)
	fmt.Println("Final counter:", atomic.LoadInt32(&counter))
}
