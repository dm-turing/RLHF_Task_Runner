package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	fmt.Println("Starting trace")
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	trace.Stop()
	// if err := trace.Stop(); err != nil {
	// log.Fatal(err)
	// }

	fmt.Println("Trace written to trace.out")
}
