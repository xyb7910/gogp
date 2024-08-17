package slicex

import "github.com/xyb7910/gogp"

// Max 返回 slicex 中最大的元素及其元素下标
func Max[T gogp.RealNumber](src []T) (T, int) {
	maxVal := src[0]
	index := 0
	for i := 0; i < len(src); i++ {
		if src[i] > maxVal {
			maxVal = src[i]
			index = i
		}
	}
	return maxVal, index
}

// Min 返回 slicex 中最小的元素及其元素下标
func Min[T gogp.RealNumber](src []T) (minVal T, index int) {
	minVal = src[0]
	index = 0
	for i := 0; i < len(src); i++ {
		if src[i] < minVal {
			minVal = src[i]
			index = i
		}
	}
	return minVal, index
}

// Sum 返回 slicex 中所有元素的和
func Sum[T gogp.RealNumber](src []T) T {
	var sum T
	for _, v := range src {
		sum += v
	}
	return sum
}
