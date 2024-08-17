package slicex

// matchFunc 判断元素是否满足条件
type matchFunc[T any] func(src T) bool
