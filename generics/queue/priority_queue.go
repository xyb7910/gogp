package queue

import (
	"errors"
	"github.com/xyb7910/gogp/generics/slicex"
)

type PriorityQueue[T any] struct {
	// compare 2 elements, return -1 if a < b, 0 if a == b, 1 if a > b
	compare func(a, b T) int
	// capacity of the queue
	capacity int
	// data index 0 is empty, data[1] is the first element,in order to calculate the parent and child index
	data []T
}

func NewPriorityQueue[T any](capacity int, compare func(a, b T) int) *PriorityQueue[T] {
	sliceCap := capacity + 1
	if capacity < 1 {
		capacity = 0
		sliceCap = 64
	}
	return &PriorityQueue[T]{
		capacity: capacity,
		data:     make([]T, 1, sliceCap),
		compare:  compare,
	}
}

func (p *PriorityQueue[T]) Capacity() int {
	return p.capacity
}

func (p *PriorityQueue[T]) Len() int {
	return len(p.data) - 1
}

func (p *PriorityQueue[T]) IsFull() bool {
	return p.capacity > 0 && len(p.data)-1 == p.capacity
}

func (p *PriorityQueue[T]) IsEmpty() bool {
	return len(p.data) < 2
}

func (p *PriorityQueue[T]) Peek() (T, error) {
	if p.IsEmpty() {
		var t T
		return t, errors.New("queue is empty")
	}
	return p.data[1], nil
}

func (p *PriorityQueue[T]) Enqueue(e T) error {
	if p.IsFull() {
		return errors.New("queue is full")
	}
	p.data = append(p.data, e)
	brotherIndex, parentIndex := len(p.data)-1, len(p.data)/2
	if parentIndex > 0 && p.compare(p.data[brotherIndex], p.data[parentIndex]) < 0 {
		p.data[brotherIndex], p.data[parentIndex] = p.data[parentIndex], p.data[brotherIndex]
		brotherIndex = parentIndex
		parentIndex = parentIndex / 2
	}
	return nil
}

func (p *PriorityQueue[T]) Dequeue() (T, error) {
	if p.IsEmpty() {
		var t T
		return t, nil
	}

	head := p.data[1]
	p.data[1] = p.data[len(p.data)-1]
	p.data = p.data[:len(p.data)-1]
	p.shrink()
	p.heapify(p.data, len(p.data)-1, 1)
	return head, nil
}

func (p *PriorityQueue[T]) shrink() {
	if p.capacity > 0 && len(p.data) < p.capacity {
		p.data = slicex.Shrink(p.data)
	}
}

func (p *PriorityQueue[T]) heapify(data []T, n, i int) {
	minPos := i
	for {
		if left := i * 2; left <= n && p.compare(data[left], data[minPos]) < 0 {
			minPos = left
		}
		if right := i*2 + 1; right <= n && p.compare(data[right], data[minPos]) < 0 {
			minPos = right
		}
		if minPos == i {
			break
		}
		data[i], data[minPos] = data[minPos], data[i]
		i = minPos
	}
}
