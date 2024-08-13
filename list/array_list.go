package list

type ArrayList[T any] struct {
	value  []T
}

func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		value: make([]T, 0, cap),
	}
}

