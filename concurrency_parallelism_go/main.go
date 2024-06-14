package main

import (
    "fmt"
    "runtime"
    simpleexample "concurrency_parallelism/simple_example"
)

func maxParallelism(){
    maxProcs := runtime.GOMAXPROCS(0)
    numCPU := runtime.NumCPU()
    fmt.Printf("MaxProcs: %d, CPU: %d\n", maxProcs, numCPU)
}

func main(){
    fmt.Println("Concurrency and Parallelism in Go")
    maxParallelism()

    // Simple example
    simpleexample.SimpleExampleWorker()
}

