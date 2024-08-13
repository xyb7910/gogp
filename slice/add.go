package slice

// Add 向 slice 中添加元素, 返回新的 slice; 如果 index 超出范围, 则 panic
func Add[T any](target []T, element T, index int) []T {
	if index < 0 || index > len(target) {
		panic("index out of range")
	}
	return append(target[:index], append([]T{element}, target[index:]...)...)
}
