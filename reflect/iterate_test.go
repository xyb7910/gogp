package reflect

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterate(t *testing.T) {
	testcases := []struct {
		name    string
		input   any
		wantRes []any
		wantErr error
	}{
		{
			name:    "slice",
			input:   []int{1, 2, 3},
			wantRes: []any{1, 2, 3},
		},
		{
			name:    "array",
			input:   [3]int{1, 2, 3},
			wantRes: []any{1, 2, 3},
		},
		{
			name:    "string",
			input:   "123456",
			wantRes: []any{1, 2, 3, 4, 5, 6},
		},
		{
			name:    "invalid",
			input:   map[string]int{"a": 1},
			wantErr: errors.New("input must be array, slice or string"),
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Iterate(tt.input)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, tt.wantRes, res)
		})
	}
}

func TestIterateMap(t *testing.T) {
	testcases := []struct {
		name       string
		input      any
		wantKeys   []any
		wantValues []any
		wantErr    error
	}{
		{
			name:    "nil",
			input:   nil,
			wantErr: errors.New("input must be map"),
		},
		{
			name:       "map",
			input:      map[string]int{"a": 1, "b": 2, "c": 3},
			wantKeys:   []any{"a", "b", "c"},
			wantValues: []any{1, 2, 3},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			keys, values, err := IterateMapV2(tt.input)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, tt.wantKeys, keys)
			assert.Equal(t, tt.wantValues, values)
		})
	}
}
