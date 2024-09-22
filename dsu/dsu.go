// Package dsu provides an implementation of a Disjoint Set Union data structure
// that is set of disjoint sets with efficient union and membership check.
package dsu

// DSU implements a Disjoint Set Union.
type DSU[T comparable] struct {
	parent map[T]T
}

// New returns an empty DSU.
func New[T comparable](values ...T) *DSU[T] {
	m := make(map[T]T, len(values))
	for _, v := range values {
		m[v] = v
	}
	return &DSU[T]{
		parent: m,
	}
}

func (s *DSU[T]) Push(value T) {
	if _, ok := s.parent[value]; ok {
		return
	}

	s.parent[value] = value
}

func (s *DSU[T]) root(value T) T {
	if value == s.parent[value] {
		return value
	}

	s.parent[value] = s.root(s.parent[value])
	return s.parent[value]
}

func (s *DSU[T]) AreInSameSet(a, b T) bool {
	a, b = s.root(a), s.root(b)
	return a == b
}

func (s *DSU[T]) Union(a, b T) bool {
	a, b = s.root(a), s.root(b)
	if a == b {
		return false
	}

	s.parent[a] = b
	return true
}

// Copy returns a copy of this stack.
func (s *DSU[T]) Iter(yield func(T) bool) bool {
	for value := range s.parent {
		if !yield(value) {
			return false
		}
	}

	return true
}
