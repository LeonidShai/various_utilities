package stack

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

const testStackCapacity = 4

func Test_Push(t *testing.T) {
	var (
		testElem     int = 1
		expectedSize int = 1
	)
	testStack := NewStack(testStackCapacity)
	err := testStack.Push(testElem)
	assert.NoError(t, err)
	assert.EqualValues(t, expectedSize, testStack.Size())
}

func Test_Capacity(t *testing.T) {
	testStack := NewStack(0)
	assert.EqualValues(t, defaultCapacity, testStack.Capacity())
	assert.EqualValues(t, 0, testStack.Size())
}
