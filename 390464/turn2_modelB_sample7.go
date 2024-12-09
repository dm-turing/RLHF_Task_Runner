package main

import (
	"fmt"
	"time"
)

func listener(broadcast <-chan string, id int) {
	for event := range broadcast {
		fmt.Printf("Listener %d received event: %s\n", id, event)
	}
}

func main() {
	broadcast := make(chan string)

	go listener(broadcast, 1)
	go listener(broadcast, 2)
	go listener(broadcast, 3)

	time.Sleep(time.Second)
	broadcast <- "Event 1"
	time.Sleep(time.Second)
	broadcast <- "Event 2"
	time.Sleep(time.Second)
	close(broadcast)
}
