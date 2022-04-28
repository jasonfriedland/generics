package stack

type Stack[T comparable] struct {
	items []T
}

// New builds a new Stack of type `T`.
func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push appends a new item to the Stack.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop returns the last value from the Stack.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	i := len(s.items) - 1
	item := s.items[i]
	s.items = s.items[:i]
	return item, true
}

// Len returns the length of the Stack.
func (s *Stack[T]) Len() int {
	return len(s.items)
}
