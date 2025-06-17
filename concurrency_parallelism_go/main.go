package main

import (
	"fmt"
	"runtime"

	// simpleexample "concurrency_parallelism/simple_example"
	errgrp "concurrency_parallelism/err_grp"
)

func maxParallelism() {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	fmt.Printf("MaxProcs: %d, CPU: %d\n", maxProcs, numCPU)
}

func main() {
	// fmt.Println("Concurrency and Parallelism in Go")
	// maxParallelism()

	// Simple example
	// simpleexample.SimpleExampleWorker()

	// Select and For-Select
	// simpleexample.SelectForSelectWorker()
	if err := errgrp.Run(); err != nil {
		fmt.Printf("error from errgrp: %v\n", err)
	}
	fmt.Println("finished func main")
}
