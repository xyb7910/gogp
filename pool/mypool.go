package pool

import (
	"sync"
	"unsafe"
)

type MyPool struct {
	p      sync.Pool
	maxCnt int32
	cnt    int32
}

func (p *MyPool) Get() any {
	return p.p.Get()
}

func (p *MyPool) Put(v any) {
	// 控制大对象
	if unsafe.Sizeof(v) > 1024 {
		return
	}
	p.p.Put(v)
}
