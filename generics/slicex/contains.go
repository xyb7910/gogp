package slicex

import (
	"github.com/xyb7910/gogp/generics/mapx"
)

// Contains 判断 slicex 中是否包含某个元素
func Contains[T comparable](src []T, element T) bool {
	for _, v := range src {
		if equal(v, element) {
			return true
		}
	}
	return false
}

// ContainsAll 判断 slicex 中是否包含所有元素
func ContainsAll[T comparable](src, dist []T) bool {
	distMap := mapx.ToMap(dist)
	for _, v := range src {
		if _, existed := distMap[v]; !existed {
			return false
		}
	}
	return true
}

// ContainsAny 判断 slicex 中是否包含任意一个元素
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
