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

	// Simulate some CPU-bound work
	time.Sleep(5 * time.Second)

	// Stop the profile.
	pprof.StopCPUProfile()
	// if err := pprof.StopCPUProfile(); err != nil {
	// log.Fatal("could not stop CPU profile:", err)
	// }

	fmt.Println("CPU profile written to cpu.prof")
}
