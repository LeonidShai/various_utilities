package simpleexample

import (
	"fmt"
	"time"
)

func selectFunc1(done <-chan struct{}, finish <-chan struct{}) {
	select {
	case <-done:
		fmt.Println("selectFunc1: case done")
	case <-finish:
		fmt.Println("selectFunc1: case finish")
	}
}

func selectFunc2(done <-chan struct{}, finish <-chan struct{}) {
	select {
	case <-done:
		fmt.Println("selectFunc2: case done")
	case <-finish:
		fmt.Println("selectFunc2: case finish")
	}
}

func forSelectFunc(done <-chan struct{}, finish <-chan struct{}) {
	var i int
	for {
		i++
		fmt.Printf("forSelectFunc: loop for: %d\n", i)
		select {
		case <-done:
			fmt.Println("forSelectFunc: case done")

			return
		case <-finish:
			fmt.Println("forSelectFunc: case finish")

			return
		default:
			fmt.Printf("forSelectFunc: default: %d\n", i)
		}
	}
}

func SelectForSelectWorker() {
	fmt.Println("----- Select and For-Select Example -----")
	done := make(chan struct{})
	finish := make(chan struct{})

	go selectFunc1(done, finish)
	go selectFunc2(done, finish)
	go forSelectFunc(done, finish)
	time.Sleep(1 * time.Millisecond)
	for i := 0; i <= 10; i++ {
		fmt.Printf("Main: %d\n", i)
	}

	select {
	case done <- struct{}{}:
		fmt.Println("Main: send done to Func")
	case finish <- struct{}{}:
		fmt.Println("Main: send finish to Func")
	default:
		fmt.Println("Main: default: has been already finished")
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println("Main: close channels")

	close(done)
	close(finish)
	time.Sleep(1 * time.Millisecond)
	fmt.Println("--------------------------")
}
