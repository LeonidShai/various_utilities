package stack

import "fmt"

func StackWorker() {
	fmt.Println("--- stack worker ---")
	st := NewStack(4)
	for i := 0; i < 4; i++ {
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
