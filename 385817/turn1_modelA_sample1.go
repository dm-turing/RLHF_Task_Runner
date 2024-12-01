package main

import (
	"fmt"
)

var sharedCounter = 0
var increments = 1000000

func incrementCounter(ch chan int) {
	for i := 0; i < increments; i++ {
		sharedCounter++
		ch <- 1
	}
}

func main() {
	numGoroutines := 100
	done := make(chan int, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go incrementCounter(done)
	}

	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	fmt.Println("Final shared counter value:", sharedCounter)
}
