package mapx

import "sync"

type SafeMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

func (sm *SafeMap[K, V]) LoadOrStore(key K, newVal V) (actual V, loaded bool) {

	oldVal, ok := sm.get(key)
	if ok {
		return oldVal, true
	}
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	oldVal, ok = sm.m[key]
	if ok {
		return oldVal, true
	}
	sm.m[key] = newVal
	return newVal, false
}

func (sm *SafeMap[K, V]) get(key K) (V, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	oldVal, ok := sm.m[key]
	return oldVal, ok
}
