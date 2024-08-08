package list

import "sync"

type SafeArrayList[T any] struct {
	vals  []T
	mutex sync.RWMutex
}

func NewSafeArrayList[T any](initCap int64) *SafeArrayList[T] {
	return &SafeArrayList[T]{
		vals: make([]T, 0, initCap),
	}
}

func (sal *SafeArrayList[T]) Get(index int64) T {
	sal.mutex.RLock()
	defer sal.mutex.RUnlock()
	res := sal.vals[index]
	return res
}

func (sal *SafeArrayList[T]) Set(index int64, val T) {
	sal.mutex.Lock()
	defer sal.mutex.Unlock()
	sal.vals[index] = val
}

func (sal *SafeArrayList[T]) DeleteAt(index int64) T {
	sal.mutex.Lock()
	defer sal.mutex.Unlock()
	res := sal.vals[index]
	sal.vals = append(sal.vals[:index], sal.vals[index+1:]...)
	return res
}
