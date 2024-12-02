package main

import (
	"testing"
	"testing/iotest"
)

func BenchmarkMyReader(b *testing.B) {
	yourReader := yourCustomReader{}
	r := iotest.TestReader(yourReader, make([]byte, 1024*1024)) // 1 MB data
	for i := 0; i < b.N; i++ {
		if _, err := yourReader.Read(r.Bytes()); err != nil {
			b.Fatal(err)
		}
	}
}
