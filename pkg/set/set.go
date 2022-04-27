package set

// Set is a non-thread safe generic type to hold a set of un-ordered unique
// `comparable` items.
type Set[T comparable] struct {
	items map[T]struct{}
}

// New builds a new Set of type `T`.
func New[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

// Add appends a new item to the Set if it doesn't already exist.
func (s *Set[T]) Add(v T) {
	s.items[v] = struct{}{}
}

// Del removes an item by value.
func (s *Set[T]) Del(v T) {
	delete(s.items, v)
}

// In returns a bool indicating the existence of an item by value.
func (s *Set[T]) In(v T) bool {
	_, found := s.items[v]
	return found
}

// Len returns the length of the Set.
func (s *Set[T]) Len() int {
	return len(s.items)
}
