package queue

import "sync"

type ConcurrentPriorityQueue[T any] struct {
	priorityQueue PriorityQueue[T]
	m             sync.RWMutex
}

func NewConcurrentPriorityQueue[T any](capacity int, compare func(a, b T) int) *ConcurrentPriorityQueue[T] {
	return &ConcurrentPriorityQueue[T]{
		priorityQueue: *NewPriorityQueue[T](capacity, compare),
	}
}

func (c *ConcurrentPriorityQueue[T]) Len() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.priorityQueue.Len()
}

func (c *ConcurrentPriorityQueue[T]) Cap() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.priorityQueue.Capacity()
}

func (c *ConcurrentPriorityQueue[T]) Peek() (T, error) {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.priorityQueue.Peek()
}

func (c *ConcurrentPriorityQueue[T]) Enqueue(t T) error {
	c.m.Lock()
	defer c.m.Unlock()
	return c.priorityQueue.Enqueue(t)
}

func (c *ConcurrentPriorityQueue[T]) Dequeue() (T, error) {
	c.m.Lock()
	defer c.m.Unlock()
	return c.priorityQueue.Dequeue()
}
