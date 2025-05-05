package stack

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

const testStackCapacity = 4

func Test_MainApi(t *testing.T) {
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
		fmt.Printf("== Test: %s\n", test.name)
		test.verify(t, &testStack)
		fmt.Println("== Success end of test")
		fmt.Println()
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
