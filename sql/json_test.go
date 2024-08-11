package sql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonColumn_Value(t *testing.T) {
	js := JsonColumn[User]{Valid: true, Val: User{LastName: "Tom"}}
	value, err := js.Value()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`{"LastName":"Tom"}`), value)
	js = JsonColumn[User]{}
	value, err = js.Value()
	assert.Nil(t, err)
	assert.Nil(t, value)
}
