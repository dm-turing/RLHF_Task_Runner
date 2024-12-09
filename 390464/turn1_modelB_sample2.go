package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stack := make(chan int, 10)

	// Push onto stack
	go func() {
		for i := 1; i <= 10; i++ {
			stack <- i
			fmt.Println("Pushed:", i)
			time.Sleep(time.Millisecond * 100)
		}
		close(stack)
	}()

	// Pop from stack
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range stack {
			fmt.Println("Popped:", num)
		}
	}()

	wg.Wait()
}
