package atomicx

import "sync/atomic"

type Value[T any] struct {
	val atomic.Value
}

func NewValue[T any]() *Value[T] {
	var t T
	return NewValueOf[T](t)
}

func NewValueOf[T any](t T) *Value[T] {
	val := atomic.Value{}
	val.Store(t)
	return &Value[T]{
		val: val,
	}
}

func (v *Value[T]) Load() (val T) {
	data := v.val.Load()
	val = data.(T)
	return
}

func (v *Value[T]) Store(val T) {
	v.val.Store(val)
}

func (v *Value[T]) Swap(newVal T) (oldVal T) {
	data := v.val.Swap(newVal)
	oldVal = data.(T)
	return
}

func (v *Value[T]) CompareAndSwap(oldVal, newVal T) (swapped bool) {
	return v.val.CompareAndSwap(oldVal, newVal)
}
