package session

import (
	"context"
	"errors"
)

// MemorySession use memory to store session, it's not safe for concurrent, only for test
type MemorySession struct {
	data   map[string]any
	claims Claims
}

func (m *MemorySession) Set(ctx *context.Context, key string, val any) error {
	m.data[key] = val
	return nil
}

func (m *MemorySession) Get(ctx *context.Context, key string) (any, error) {
	val, ok := m.data[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return val, nil
}

func (m *MemorySession) Del(ctx *context.Context, key string) error {
	delete(m.data, key)
	return nil
}

func (m *MemorySession) Destroy(ctx *context.Context) error {
	return nil
}

func (m *MemorySession) Claims() Claims {
	return m.claims
}

func NewMemorySession() *MemorySession {
	return &MemorySession{
		data:   map[string]any{},
		claims: Claims{},
	}
}

var _ Session = &MemorySession{}
