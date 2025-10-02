package main

import (
	"errors"
	"fmt"
	// getid "experiments/get_id"
	// yamlparser "experiments/yaml_parser"
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

func expDeferWithIf() {
	var err error
	var a int = 5
	defer func() {
		if err != nil {
			fmt.Printf("we have error: %v\n", a)
		} else {
			fmt.Printf("we don't have error: %v. Uhhu!!!\n", a)
		}
	}()

	a = 2 + 3
	err = errors.New("first error")
}

func main() {
	// experimentInvalidMemoryAddress()
	// experimentSliceWithNil()
	// experimentCloseOfClosedChannel()
	// experimentWithSlices()

	// Experiments with parsing yaml file (test.yaml)
	// yamlparser.ParseYamlFile(yamlparser.TestFileName)

	// GetID
	// getid.GetIdWorker()
	expDeferWithIf()
	fmt.Println("end of main")
}
