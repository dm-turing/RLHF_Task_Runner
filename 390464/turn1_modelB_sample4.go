package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	queue := make(chan int, 2)

	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			queue <- i
			fmt.Println("Produced:", i)
			time.Sleep(time.Millisecond * 100)
		}
		close(queue)
	}()

	// Consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range queue {
			fmt.Println("Consumed:", num)
		}
	}()

	wg.Wait()
}
