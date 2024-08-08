package pool

import (
	"fmt"
	"sync"
	"testing"
)

type User struct {
	Id   int64
	Name string
}

func (u *User) Reset() {
	u.Id = 0
	u.Name = ""
}

func (u *User) changeName(newName string) {
	u.Name = newName
}

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return &User{}
		},
	}
	// pool 里面没有的时候就会创建一个
	u1 := pool.Get().(*User)
	u1.Id = 1
	u1.Name = "Tom"
	// 重置
	u1.Reset()
	// 必须把对象放回池子，不然还会创建一个新的对象
	pool.Put(u1)
	u2 := pool.Get().(*User)
	fmt.Println(u2)
}
