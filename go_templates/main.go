package main

import (
	"fmt"
)

func main() {
	fmt.Println("go templates")

	// template Singleton
	singletonWorker()

	// template Adapter
	adapterWorker()
}
