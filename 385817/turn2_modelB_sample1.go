package main

import (
	"testing"
)

func BenchmarkChannelCommunication(b *testing.B) {
	chanSize := 1000
	msgChannel := make(chan int, chanSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgChannel <- i
		<-msgChannel
	}
}
