package main

import (
	"fmt"
	"sync"
	"time"
)

type Broadcast struct {
	subscribers []chan string
	lock        sync.Mutex
}

func (b *Broadcast) Subscribe() chan string {
	b.lock.Lock()
	defer b.lock.Unlock()
	ch := make(chan string)
	b.subscribers = append(b.subscribers, ch)
	return ch
}

func (b *Broadcast) Publish(message string) {
	b.lock.Lock()
	defer b.lock.Unlock()
	for _, subscriber := range b.subscribers {
		subscriber <- message
	}
}

func main() {
	broadcaster := &Broadcast{}

	// Clients subscribing to message
	client1 := broadcaster.Subscribe()
	client2 := broadcaster.Subscribe()

	go func() {
		for msg := range client1 {
			fmt.Printf("Client 1 received: %s\n", msg)
		}
	}()

	go func() {
		for msg := range client2 {
			fmt.Printf("Client 2 received: %s\n", msg)
		}
	}()

	// Publish messages
	broadcaster.Publish("Hello, subscribers!")
	broadcaster.Publish("Another event occurred!")

	// Exit after a delay to allow time for processing
	time.Sleep(time.Second)
}
