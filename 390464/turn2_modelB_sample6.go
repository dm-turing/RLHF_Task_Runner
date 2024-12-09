package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(channel chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		channel <- fmt.Sprintf("Task %d", i)
		fmt.Printf("Producer sent task %d\n", i)
		time.Sleep(time.Second)
	}
}

func consumer(channel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range channel {
		fmt.Printf("Consumer received task: %s\n", task)
		time.Sleep(time.Second / 2)
	}
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan string, 3) // Bounded queue with size 3

	wg.Add(1)
	go producer(channel, &wg)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go consumer(channel, &wg)
	}

	wg.Wait()
	close(channel)
}
