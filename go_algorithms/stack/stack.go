package stack

import (
	"errors"
	"fmt"
)

var (
	errEmptyStack = errors.New("stack is empty")
	errFullStack  = errors.New("stack is full")
)

type stack struct {
	size     int
	capacity int
	array    []int
}

// Create stack with capacity.
// If capacity = 0, stack will have a capacity of 10.
func NewStack(capacity int) stack {
	return stack{
		size:     0,
		capacity: capacity,
		array:    make([]int, 0, capacity),
	}
}

func (s *stack) Push(elem int) error {
	fmt.Println("Add element: ", elem)
	if s.size == s.capacity {
		fmt.Println("Stack is full")

		return errFullStack
	}

	s.array = append(s.array, elem)
	s.size++
	fmt.Println("Add element success")

	return nil
}

func (s *stack) Pop() (int, error) {
	if s.size == 0 {
		fmt.Println("Empty stak")

		return 0, errEmptyStack
	}

	s.size--
	elem := s.array[s.size]
	s.array = s.array[:s.size]
	fmt.Println("Pop element: ", elem)
	return elem, nil
}

func (s *stack) Size() int {
	return s.size
}

func (s *stack) Capacity() int {
	return s.capacity
}

func (s *stack) Clear() {
	s.size = 0
	s.array = s.array[:0]

	fmt.Println("Clear stak")
}

func (s *stack) Print() {
	fmt.Println("Stack size: ", s.size)
	if s.size == 0 {
		return
	}

	for i := 0; i < s.size; i++ {
		fmt.Printf("%d ", s.array[i])
	}
	fmt.Println()
}
