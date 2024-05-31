package main

import "fmt"

type common struct {
	pparam *uint32
}

func main() {
	var com common
	var temp uint32 = 5
	// *com.pparam = temp  // panic: runtime error: invalid memory address or nil pointer dereference
	com.pparam = &temp
	fmt.Println(*com.pparam)
}
