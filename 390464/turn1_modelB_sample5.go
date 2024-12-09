package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	queue := make(chan int)

	// Producer
	go func() {
		for i := 1; i <= 3; i++ {
			queue <- i
			fmt.Println("Produced:", i)
			time.Sleep(time.Millisecond * 100)
		}
		close(queue)
	}()

	// Consumers
	wg.Add(2)
	go func() {
		defer wg.Done()
		for num := range queue {
			fmt.Println("Consumer 1:", num)
		}
	}()

	go func() {
		defer wg.Done()
		for num := range queue {
			fmt.Println("Consumer 2:", num)
		}
	}()

	wg.Wait()
}
