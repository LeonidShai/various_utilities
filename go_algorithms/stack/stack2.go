package stack

import "fmt"

const defaultCapacity = 10

type node struct {
	data   int
	bottom *node
}

type stack2 struct {
	size     int
	capacity int
	top      *node
}

func CreateNewStack(capacity int) stack2 {
	if capacity == 0 {
		capacity = defaultCapacity
	}

	return stack2{
		size:     0,
		capacity: capacity,
		top:      nil,
	}
}

func (s *stack2) Push(elem int) error {
	fmt.Println("Push elem: ", elem)
	if s.size == s.capacity {
		fmt.Println("Stack is full")
		return errFullStack
	}

	nd := node{
		data: elem,
	}

	if s.top != nil {
		nd.bottom = s.top
	}

	s.top = &nd
	s.size++

	fmt.Println("Pushing is success")
	return nil
}

func (s *stack2) Pop() (int, error) {
	if s.size == 0 {
		fmt.Println("Stack is empty")
		return 0, errEmptyStack
	}

	elem := s.top.data
	s.top = s.top.bottom
	s.size--
	fmt.Println("Pop elem: ", elem)

	return elem, nil
}

func (s *stack2) Size() int {
	return s.size
}

func (s *stack2) Capacity() int {
	return s.capacity
}

func (s *stack2) Clear() {
	fmt.Println("Clear stack")

	s.size = 0
	s.top = nil
}

func (s *stack2) Print() {
	fmt.Println("Stack size: ", s.size)
	tmp := s.top
	for tmp != nil {
		fmt.Printf("%d ", tmp.data)
		tmp = tmp.bottom
	}
	fmt.Println()
}
