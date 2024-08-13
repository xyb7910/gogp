package slice

// Replace 将 slice 中某个元素替换为另一个元素, 返回新的 slice
func Replace[T comparable](src []T, old, new T) []T {
	existed := Contains(src, old)
	if !existed {
		panic("element not existed")
	}
	index := IndexOf(src, old)
	src[index] = new
	return src
}

// ReplaceIndex 将 slice 中某个下标的元素替换为另一个元素, 返回新的 slice
func ReplaceIndex[T any](src []T, index int, element T) []T {
	if index < 0 || index > len(src) {
		panic("index out of range")
	}
	src[index] = element
	return src
}
