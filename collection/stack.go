package collection

type StackIsEmptyError struct{}

func (e StackIsEmptyError) Error() string {
	return "stack is empty"
}

type Stack struct {
	elems *ListNode
}

type ListNode struct {
	elem interface{}
	prev *ListNode
}

// InitStack create new empty stack
func InitStack() *Stack {
	return &Stack{}
}

// IsEmpty return true when stack really empty
// otherwise return false
func (s *Stack) IsEmpty() bool {
	return s.elems == nil
}

// Push append new unknown element to stack
func (s *Stack) Push(el interface{}) {
	s.elems = &ListNode{el, s.elems}
}

// Top return top element in stack, can return error
// when stack is empty
func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, StackIsEmptyError{}
	}
	return s.elems.elem, nil
}

// Pop return and remove top element from stack, can
// return error when stack is empty
func (s *Stack) Pop() (interface{}, error) {
	top, err := s.Top()
	if err != nil {
		return nil, err
	}
	s.elems = s.elems.prev
	return top, nil
}

// IsNonEmpty is opposite to empty, this return
// true when stack has elements otherwise return false
func (s *Stack) IsNonEmpty() bool {
	return !s.IsEmpty()
}
