package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	fmt.Println("Starting CPU profile")
	fp, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile:", err)
	}
	defer fp.Close()

	if err := pprof.StartCPUProfile(fp); err != nil {
		log.Fatal("could not start CPU profile:", err)
	}

	defer pprof.StopCPUProfile()

	// Code you want to profile goes here
	time.Sleep(10 * time.Second)

	fmt.Println("Stopping CPU profile")
}
