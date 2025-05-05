package stack

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

const testStackCapacity = 4

func TestMainApi(t *testing.T) {
	type tests struct {
		name   string
		verify func(*testing.T, *stack)
	}

	checkers := []tests{
		{
			name:   "Check Default Capacity",
			verify: verifyDefaultCapacity,
		},
		{
			name:   "Push 1 elem",
			verify: verifyPush,
		},
		{
			name:   "Check Capacity",
			verify: verifyCapacity,
		},
		{
			name:   "Pop elem",
			verify: verifyPop,
		},
		{
			name:   "Pop elem from empty stack",
			verify: verifyPopEmptyStack,
		},
	}

	testStack := NewStack(testStackCapacity)
	for _, test := range checkers {
		// fmt.Printf("== Test: %s\n", test.name)
		// test.verify(t, &testStack)
		// fmt.Println("== Success end of test")
		// fmt.Println()
		t.Run(test.name, func(t *testing.T) {
			test.verify(t, &testStack)
		})
	}
}

func BenchmarkMainApi(b *testing.B) {
	var testCapacity int = 10000
	type benchs struct {
		name  string
		bench func(*testing.B, *stack)
	}

	benchers := []benchs{
		{
			name:  "push elem",
			bench: benchPushElem,
		},
	}

	testStack := NewStack(testCapacity)
	for _, test := range benchers {
		// b.ResetTimer()
		b.Run(test.name, func(b *testing.B) {
			test.bench(b, &testStack)
		})
	}
}

func benchPushElem(b *testing.B, testStack *stack) {
	var testElem int = 1
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		if i%testStack.capacity == 0 {
			testStack.Clear()
		}
		b.StartTimer()
		_ = testStack.Push(testElem)
	}
}

func verifyPush(t *testing.T, testStack *stack) {
	var (
		testElem     int = 1
		expectedSize int = 1
	)

	err := testStack.Push(testElem)
	assert.NoError(t, err)
	assert.EqualValues(t, expectedSize, testStack.Size())
}

func verifyDefaultCapacity(t *testing.T, _ *stack) {
	testStack := NewStack(0)
	assert.EqualValues(t, defaultCapacity, testStack.Capacity())
	assert.EqualValues(t, 0, testStack.Size())
}

func verifyCapacity(t *testing.T, testStack *stack) {
	var expectedCapacity int = 4
	assert.EqualValues(t, expectedCapacity, testStack.Capacity())
}

func verifyPop(t *testing.T, testStack *stack) {
	var expectedElem int = 1
	elem, err := testStack.Pop()
	assert.NoError(t, err)
	assert.EqualValues(t, expectedElem, elem)
	assert.EqualValues(t, 0, testStack.Size())
}

func verifyPopEmptyStack(t *testing.T, testStack *stack) {
	elem, err := testStack.Pop()
	assert.ErrorIs(t, errEmptyStack, err)
	assert.EqualValues(t, 0, elem)
	assert.EqualValues(t, 0, testStack.Size())
}
