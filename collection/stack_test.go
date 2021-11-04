package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushElements(t *testing.T) {
	stack := InitStack()

	stack.Push("d")
	el, _ := stack.Top()

	assert.Equal(t, "d", el)
}

func TestPopElements(t *testing.T) {
	stack := InitStack()

	stack.Push("a")
	stack.Push("b")

	el, _ := stack.Top()
	assert.Equal(t, "b", el)

	el, _ = stack.Pop()
	assert.Equal(t, "b", el)

	el, _ = stack.Top()
	assert.Equal(t, "a", el)
}

func TestIsEmptyWhenStackEmpty(t *testing.T) {
	stack := InitStack()

	assert.True(t, stack.IsEmpty())
}

func TestIsEmptyWhenStackNotEmpty(t *testing.T) {
	stack := InitStack()
	stack.Push("a")

	assert.False(t, stack.IsEmpty())
}

func TestErrorWhenPopInEmptyStack(t *testing.T) {
	stack := InitStack()
	_, err := stack.Pop()
	assert.Error(t, err)
}

func TestErrorWhenTopInEmptyStack(t *testing.T) {
	stack := InitStack()
	_, err := stack.Top()
	assert.Error(t, err)
}

func TestIsEmptyWorksWithState(t *testing.T) {
	stack := InitStack()
	assert.True(t, stack.IsEmpty())

	stack.Push("one")
	assert.False(t, stack.IsEmpty())

	stack.Pop()
	assert.True(t, stack.IsEmpty())
}
