package slice

import "github.com/xyb7910/gogp/mapx"

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
