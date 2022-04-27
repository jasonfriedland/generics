package set

// Set is a non-thread safe generic type to hold a slice of un-ordered unique
// `comparable` items.
type Set[T comparable] struct {
	items []T
}

// New builds a new Set of type `T`.
func New[T comparable]() *Set[T] {
	return &Set[T]{}
}

// Add appends a new item to the Set if it doesn't already exist.
func (s *Set[T]) Add(v T) {
	if _, found := s.In(v); found {
		return
	}
	s.items = append(s.items, v)
}

// Del removes an item by value. The deletion algorithm DOES NOT retain the
// order of items.
func (s *Set[T]) Del(v T) {
	if i, found := s.In(v); found {
		// Replace the item at index `i` with the last item of the
		// slice, and discard the last item
		s.items[i] = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
	}
}

// In returns the index and a bool indicating the position and existence of an
// item by value.
func (s *Set[T]) In(v T) (int, bool) {
	for i, k := range s.items {
		if k == v {
			return i, true
		}
	}
	return -1, false
}

// Len returns the length of the Set.
func (s *Set[T]) Len() int {
	return len(s.items)
}
