package slicex

// Remove 从 slicex 中移除某个元素, 返回新的 slicex
func Remove[T comparable](src []T, element T) []T {
	existed := Contains(src, element)
	if !existed {
		panic("element not existed")
	}
	index := IndexOf(src, element)
	return append(src[:index], src[index+1:]...)
}

// RemoveIndex 从 slicex 中移除某个下标的元素, 返回新的 slicex
func RemoveIndex[T any](src []T, index int) []T {
	if index < 0 || index > len(src) {
		panic("index out of range")
	}
	return append(src[:index], src[index+1:]...)
}

// FilterRemove 过滤 slicex 中的元素, 返回新的 slicex
func FilterRemove[T any](src []T, filter func(idx int, src T) bool) []T {
	emptyPos := 0
	for idx := range src {
		if filter(idx, src[idx]) {
			continue
		}
		// 移动元素
		src[emptyPos] = src[idx]
		emptyPos++
	}
	return src[:emptyPos]
}
