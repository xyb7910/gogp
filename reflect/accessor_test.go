package reflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReflectAccessor_Field(t *testing.T) {
	testcases := []struct {
		name string

		// input
		entity interface{}
		field  string

		// output
		wantErr error
		wantVal int
	}{
		{
			name:    "success",
			entity:  &User{Age: 10},
			field:   "Age",
			wantVal: 10,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			// 新建一个ReflectAccessor
			accessor, err := NewReflectAccessor(tt.entity)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			val, err := accessor.Field(tt.field)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantVal, val)
		})
	}
}

func TestNewReflectAccessor_SetField(t *testing.T) {
	testcases := []struct {
		name    string
		entity  *User
		field   string
		newVal  int
		wantErr error
	}{
		{
			name:   "set age",
			entity: &User{},
			field:  "Age",
			newVal: 18,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			accessor, err := NewReflectAccessor(tt.entity)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
				return
			}
			err = accessor.SetField(tt.field, tt.newVal)
			if err != nil {
				return
			}
			assert.Equal(t, tt.newVal, tt.entity.Age)
		})
	}

}
