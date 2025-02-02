package set

type Set[T comparable] interface {
	Add(key T)
	Delete(key T)
	Exist(key T) bool
	Keys() []T
}

type MapSet[T comparable] struct {
	m map[T]struct{}
}

func (s *MapSet[T]) Delete(key T) {
	delete(s.m, key)
}

func (s *MapSet[T]) Exist(key T) bool {
	_, ok := s.m[key]
	return ok
}

func (s *MapSet[T]) Keys() []T {
	res := make([]T, 0, len(s.m))
	for key := range s.m {
		res = append(res, key)
	}
	return res
}

func NewMapSet[T comparable](size int) *MapSet[T] {
	return &MapSet[T]{
		m: make(map[T]struct{}, size),
	}
}

func (s *MapSet[T]) Add(key T) {
	s.m[key] = struct{}{}
}
