package queue

import (
	"context"
	"golang.org/x/sync/semaphore"
	"sync"
)

type ConcurrentArrayBlockingQueue[T any] struct {
	data  []T
	mutex *sync.RWMutex
	head  int

	tail int

	count int

	enqueueCap *semaphore.Weighted
	dequeueCap *semaphore.Weighted

	zero T
}

func NewConcurrentArrayBlockingQueue[T any](capacity int, zero T) *ConcurrentArrayBlockingQueue[T] {
	mutex := &sync.RWMutex{}
	semaForEnqueue := semaphore.NewWeighted(int64(capacity))
	semaForDequeue := semaphore.NewWeighted(int64(capacity))

	_ = semaForEnqueue.Acquire(context.TODO(), int64(capacity))

	res := &ConcurrentArrayBlockingQueue[T]{
		data:       make([]T, capacity),
		mutex:      mutex,
		enqueueCap: semaForEnqueue,
		dequeueCap: semaForDequeue,
	}
	return res
}

func (c *ConcurrentArrayBlockingQueue[T]) Enqueue(ctx context.Context, value T) error {
	// TODO: implement me
	panic("implement me")
}
