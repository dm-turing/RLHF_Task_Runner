package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		channel <- i
		fmt.Printf("Producer sent task %d\n", i)
		time.Sleep(time.Second)
	}
}

func consumer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-channel
		if !ok {
			break
		}
		fmt.Printf("Consumer received task %d\n", task)
		time.Sleep(time.Second / 2)
	}
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan int)

	wg.Add(1)
	go producer(channel, &wg)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go consumer(channel, &wg)
	}

	wg.Wait()
	close(channel)
}
