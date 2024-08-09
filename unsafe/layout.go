package unsafe

import (
	"fmt"
	"reflect"
)

// PrintFieldOffset  print field offset
// Go 使用字长对齐，即每个字段的偏移量都是字长的整数倍
func PrintFieldOffset(entity any) {
	typ := reflect.TypeOf(entity)
	for i := 0; i < typ.NumField(); i++ {
		fd := typ.Field(i)
		fmt.Printf("field %s offset %d\n", fd.Name, fd.Offset)
	}
}
