package syncx

import "sync"

type Map[K comparable, V any] struct {
	m sync.Map
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	var anyVal any
	anyVal, ok = m.m.Load(key)
	if anyVal != nil {
		value = anyVal.(V)
	}
	return
}

func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	
}
