package queue

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Compare() func(a, b int) int {
	return func(a, b int) int {
		if a > b {
			return 1
		}
		if a < b {
			return -1
		}
		return 0
	}
}

func TestNewPriorityQueue(t *testing.T) {
	data := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	testCases := []struct {
		name       string
		queue      *PriorityQueue[int]
		capacity   int
		data       []int
		expectData []int
	}{
		{
			name:       "无边际",
			queue:      NewPriorityQueue(0, Compare()),
			capacity:   0,
			data:       data,
			expectData: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:       "有边际",
			queue:      NewPriorityQueue(len(data), Compare()),
			capacity:   len(data),
			data:       data,
			expectData: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, 0, tc.queue.Len())
			for _, v := range data {
				err := tc.queue.Enqueue(v)
				assert.NoError(t, err)
				if err != nil {
					return
				}
			}
			//assert.Equal(t, tc.capacity, tc.queue.Capacity())
			//assert.Equal(t, len(data), tc.queue.Len())
			res := make([]int, 0, len(data))
			if tc.queue.Len() > 0 {
				ele, err := tc.queue.Dequeue()
				assert.NoError(t, err)
				if err != nil {
					return
				}
				res = append(res, ele)
			}
			//assert.Equal(t, tc.expectData, res)
		})
	}
}

func TestPriorityQueue_Peek(t *testing.T) {
	testCases := []struct {
		name     string
		capacity int
		data     []int
		wantErr  error
	}{
		{
			name:     "有数据",
			capacity: 0,
			data:     []int{6, 5, 4, 3, 2, 1},
			wantErr:  errors.New("queue is empty"),
		},
		{
			name:     "无数据",
			capacity: 0,
			data:     []int{},
			wantErr:  errors.New("queue is empty"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			q := NewPriorityQueue[int](tc.capacity, Compare())
			for _, el := range tc.data {
				err := q.Enqueue(el)
				require.NoError(t, err)
			}
			for q.Len() > 0 {
				peek, err := q.Peek()
				assert.NoError(t, err)
				el, _ := q.Dequeue()
				assert.Equal(t, el, peek)
			}
			_, err := q.Peek()
			assert.Equal(t, tc.wantErr, err)
		})

	}
}
