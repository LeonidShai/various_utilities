package stack

import "fmt"

func StackWorker() {
	fmt.Println("--- stack worker ---")
	var stackCapacity int = 4
	st := NewStack(stackCapacity)
	for i := 0; i < stackCapacity; i++ {
		_ = st.Push(i)
	}
	st.Print()

	_ = st.Push(5)
	_, _ = st.Pop()
	st.Print()

	_ = st.Push(5)
	st.Print()

	st.Clear()
	st.Print()
	fmt.Println("--- ---")
}

func Stack2Worker() {
	fmt.Println("--- stack 2 worker ---")
	var stackCapacity int = 4
	st := CreateNewStack(stackCapacity)
	for i := 0; i < stackCapacity; i++ {
		_ = st.Push(i)
	}
	st.Print()

	_ = st.Push(5)
	_, _ = st.Pop()
	st.Print()

	_ = st.Push(5)
	st.Print()

	st.Clear()
	st.Print()

	_ = st.Push(6)
	st.Print()
	_, _ = st.Pop()
	st.Print()
	_, _ = st.Pop()
	fmt.Println("--- ---")
}
