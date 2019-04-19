package stack

// Stack implements a LIFO queue ('stack') using an underlying slice.
type Stack struct {
	underlying []interface{}
}

// New allocates a stack capable of holding up to `capacity` elements.
func New(capacity int) *Stack {
	return &Stack{make([]interface{}, 0, capacity)}
}

// Len returns the number of elements currently on the stack.
func (s *Stack) Len() int {
	return len(s.underlying)
}

// Cap returns the capacity of the stack. If Len() == Cap(), then Push() will fail and return false.
func (s *Stack) Cap() int {
	return cap(s.underlying)
}

// Peek shows the value that would be returned if Popped; nil, false if stack is empty.
func (s *Stack) Peek() (interface{}, bool) {
	if len(s.underlying) == 0 {
		return nil, false
	}
	return s.underlying[len(s.underlying)-1], true
}

// Pop removes the most recently pushed value; nil, false if stack is empty.
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.underlying) == 0 {
		return nil, false
	}
	last := len(s.underlying) - 1
	v := s.underlying[last]
	s.underlying = s.underlying[0:last]
	return v, true
}

// Push places the supplied value at the top of the stack; nil if stack is at capacity and cannot add value.
func (s *Stack) Push(value interface{}) bool {
	if len(s.underlying) == cap(s.underlying) {
		return false
	}
	s.underlying = append(s.underlying, value)
	return true
}
