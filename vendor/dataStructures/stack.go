package ds

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T

	if s.IsEmpty() {
		return zero, false
	}

	item := s.items[s.Size()-1]
	s.items = s.items[:s.Size()-1]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T

	if s.IsEmpty() {
		return zero, false
	}

	return s.items[s.Size()-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
