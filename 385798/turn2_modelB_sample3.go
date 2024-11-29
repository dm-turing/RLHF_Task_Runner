package main

import (
	"fmt"
	"sync/atomic"
)

var (
	flag   int32 = 0
	shared       = 42
)

func setShared() {
	for atomic.LoadInt32(&flag) == 1 {
	}
	shared = 100
	atomic.StoreInt32(&flag, 1)
}

func getShared() {
	for atomic.LoadInt32(&flag) == 1 {
	}
	fmt.Println("Shared value:", shared)
	atomic.StoreInt32(&flag, 1)
}
