package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

type User struct {
	Name    string
	Age     int
	Alias   []byte
	Address string
}

type UserV1 struct {
	Name    string
	Age1    int32
	Age2    int32
	Alias   byte
	Address string
}

func TestPrintFieldOffset(t *testing.T) {
	fmt.Println(unsafe.Sizeof(User{}))
	PrintFieldOffset(User{})

	fmt.Println(unsafe.Sizeof(UserV1{}))
	PrintFieldOffset(UserV1{})
}
