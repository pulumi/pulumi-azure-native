package collections

import (
	"cmp"
	"slices"
)

type OrderableSet[T cmp.Ordered] struct {
	m map[T]struct{}
}

func NewOrderableSet[T cmp.Ordered](values ...T) *OrderableSet[T] {
	newSet := &OrderableSet[T]{m: make(map[T]struct{})}
	for _, value := range values {
		newSet.Add(value)
	}
	return newSet
}

func (s *OrderableSet[T]) Add(value T) {
	s.m[value] = struct{}{}
}

func (s *OrderableSet[T]) Has(value T) bool {
	_, ok := s.m[value]
	return ok
}

func (s *OrderableSet[T]) Remove(value T) {
	delete(s.m, value)
}

func (s *OrderableSet[T]) Count() int {
	return len(s.m)
}

func (s *OrderableSet[T]) SortedValues() []T {
	values := make([]T, 0, len(s.m))
	for value := range s.m {
		values = append(values, value)
	}
	slices.Sort(values)
	return values
}
