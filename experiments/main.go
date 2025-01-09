package main

import (
	"fmt"

	yamlparser "experiments/yaml_parser"
)

type common struct {
	pparam *uint32
}

func experimentInvalidMemoryAddress() {
	var com common
	var temp uint32 = 5
	// *com.pparam = temp  // panic: runtime error: invalid memory address or nil pointer dereference
	com.pparam = &temp
	fmt.Println(*com.pparam)
}

func experimentSliceWithNil() {
	type simple struct {
		indx int
	}

	var s *simple

	var list []*simple = []*simple{s}
	for _, elem := range list {
		if elem == nil {
			fmt.Println("Elem is nil")
		} else {
			fmt.Println(elem)
		}
	}
}

func experimentWithSlices() {
	list := make([]int, 0, 2)
	fmt.Println(list)
	if len(list) != 0 {
		fmt.Println("List isn't equal 0")
	}
}

func experimentCloseOfClosedChannel() {
	ch := make(chan struct{})

	close(ch)
	close(ch) // panic: close of closed channel
}

func main() {
	// experimentInvalidMemoryAddress()
	// experimentSliceWithNil()
	// experimentCloseOfClosedChannel()
	// experimentWithSlices()

	// Experiments with parsing yaml file (test.yaml)
	yamlparser.ParseYamlFile(yamlparser.TestFileName)
}
