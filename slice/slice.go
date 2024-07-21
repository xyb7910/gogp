package slice

import (
	"github.com/xyb7910/gogp"
	"github.com/xyb7910/gogp/mapx"
)

// Add 向 slice 中添加元素, 返回新的 slice; 如果 index 超出范围, 则 panic
func Add[T any](target []T, element T, index int) []T {
	if index < 0 || index > len(target) {
		panic("index out of range")
	}
	return append(target[:index], append([]T{element}, target[index:]...)...)
}

// Max 返回 slice 中最大的元素及其元素下标
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

// Min 返回 slice 中最小的元素及其元素下标
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

// Sum 返回 slice 中所有元素的和
func Sum[T gogp.RealNumber](src []T) T {
	var sum T
	for _, v := range src {
		sum += v
	}
	return sum
}

// Contains 判断 slice 中是否包含某个元素
func Contains[T comparable](src []T, element T) bool {
	for _, v := range src {
		if equal(v, element) {
			return true
		}
	}
	return false
}

// ContainsAll 判断 slice 中是否包含所有元素
func ContainsAll[T comparable](src, dist []T) bool {
	distMap := mapx.ToMap(dist)
	for _, v := range src {
		if _, existed := distMap[v]; !existed {
			return false
		}
	}
	return true
}

// ContainsAny 判断 slice 中是否包含任意一个元素
func ContainsAny[T comparable](src, dist []T) bool {
	for _, v := range src {
		for _, d := range dist {
			if equal(v, d) {
				return true
			}
		}
	}
	return false
}

// equal 判断两个元素是否相等
func equal[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

// IndexOf 返回 slice 中某个元素的下标
func IndexOf[T comparable](src []T, element T) (index int) {
	for i, v := range src {
		if !equal(v, element) {
			panic("element not existed")
		}
		index = i
	}
	return index
}

// Remove 从 slice 中移除某个元素, 返回新的 slice
func Remove[T comparable](src []T, element T) []T {
	existed := Contains(src, element)
	if !existed {
		panic("element not existed")
	}
	index := IndexOf(src, element)
	return append(src[:index], src[index+1:]...)
}

// RemoveIndex 从 slice 中移除某个下标的元素, 返回新的 slice
func RemoveIndex[T any](src []T, index int) []T {
	if index < 0 || index > len(src) {
		panic("index out of range")
	}
	return append(src[:index], src[index+1:]...)
}

// Reverse 将 slice 反转
func Reverse[T any](src []T) []T {
	reversed := make([]T, len(src))
	for i := 0; i < len(src); i++ {
		reversed[i] = src[len(src)-i-1]
	}
	return reversed
}

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

// DiffSet return a slice, src slice existed, but dist slice not existed
/*
 collection src - collection dist = DiffSet,
*/
func DiffSet[T comparable](src, dist []T) (res []T) {
	srcMap := mapx.ToMap(src)
	for _, v := range dist {
		if _, existed := srcMap[v]; existed {
			continue
		}
		res = append(res, v)
	}
	return
}

// SymDiffSet return a slice, src slice and dist slice both not existed
func SymDiffSet[T comparable](src, dist []T) (res []T) {
	srcMap, distMap := mapx.ToMap(src), mapx.ToMap(dist)
	// Add elements in src but not in dist
	for v := range srcMap {
		if _, found := distMap[v]; !found {
			res = append(res, v)
		}
	}

	// Add elements in dist but not in src
	for v := range distMap {
		if _, found := srcMap[v]; !found {
			res = append(res, v)
		}
	}
	return
}

// IntersectSet return a slice, src slice and dist slice both existed
func IntersectSet[T comparable](src, dist []T) (res []T) {
	srcMap := mapx.ToMap(src)
	for _, v := range dist {
		if _, existed := srcMap[v]; existed {
			res = append(res, v)
		}
	}
	return
}

// UnionSet return a slice, src slice exited or dist slice existed
func UnionSet[T comparable](src, dist []T) (res []T) {
	srcMap := mapx.ToMap(src)
	for _, v := range dist {
		srcMap[v] = struct{}{}
	}
	res = ToSliceByMapKey(srcMap)
	return
}

// ToSliceByMapKey is a generic function that converts a map to a slice of its keys.
func ToSliceByMapKey[K comparable, V any](m map[K]V) []K {
	// Create a slice to hold the keys
	keys := make([]K, 0, len(m))

	// Iterate over the map and append each key to the slice
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// ToSliceByMapValue is a generic function that converts a map to a slice of its values.
func ToSliceByMapValue[K comparable, V any](m map[K]V) []V {
	// Create a slice to hold the values
	values := make([]V, 0, len(m))

	// Iterate over the map and append each value to the slice
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// Map applies a function to each element of a slice and returns a new slice with the results.
func Map[T any, U any](input []T, mapper func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}

// Reduce reduces a slice to a single value using a specified reduction function.
func Reduce[T any, U any](input []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, v := range input {
		result = reducer(result, v)
	}
	return result
}
