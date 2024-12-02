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
	fp, err := os.Create("trace.out")
	if err != nil {
		log.Fatal("could not create trace file:", err)
	}
	defer fp.Close()

	if err := trace.Start(fp); err != nil {
		log.Fatal("could not start trace:", err)
	}

	defer trace.Stop()

	// Code you want to trace goes here
	time.Sleep(10 * time.Second)

	fmt.Println("Stopping trace")
}
