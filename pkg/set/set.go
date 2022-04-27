package set

type Set[T comparable] struct {
	items []T
}

func New[T comparable]() *Set[T] {
	return &Set[T]{}
}

func (s *Set[T]) Add(v T) {
	if _, found := s.In(v); found {
		return
	}
	s.items = append(s.items, v)
}

func (s *Set[T]) Del(v T) {
	if i, found := s.In(v); found {
		s.items[i] = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
	}
}

func (s *Set[T]) In(v T) (int, bool) {
	for i, k := range s.items {
		if k == v {
			return i, true
		}
	}
	return -1, false
}

func (s *Set[T]) Len() int {
	return len(s.items)
}
