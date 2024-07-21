package slice

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	arr := []int{1, 2, 3}
	back := Add(arr, 4, 2)
	fmt.Println(back)
}

func TestDiffSet(t *testing.T) {
	src := []int{1, 2, 3, 4, 5}
	//srcMap := mapx.ToMap(src)
	//fmt.Println(srcMap)
	dist := []int{2, 3, 4, 5, 6}
	res := UnionSet(src, dist)
	fmt.Println(res)
}
