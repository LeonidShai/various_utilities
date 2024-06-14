package simpleexample

import (
    "fmt"
)

func printer(done <-chan struct{}, finished chan<- struct{}) {
    for i:=0; i<=7; i++ {
        select {
        case <-done:
            fmt.Println("case done")
            return
        default:
            fmt.Printf("Printer: %d\n", i)
        }
    }
    close(finished)
    fmt.Println("printer has finished")
}

func SimpleExampleWorker() {
    fmt.Println("----- Simple Example -----")
    done := make(chan struct{})
    finished := make(chan struct{})

    go printer(done, finished)
    for i:=0; i<=45; i++ {
        fmt.Printf("Main: %d\n", i)
    }

    select {
    case <-finished:
        fmt.Println("Printer has been already finished!")
    default:
        done <- struct{}{}
        fmt.Println("send done to printer")
        fmt.Println("main has been already finished")
    }

    close(done)
    fmt.Println("--------------------------")
}

