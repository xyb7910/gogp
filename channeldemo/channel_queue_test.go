package channeldemo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBroker1(t *testing.T) {
	b := NewBroker1()
	str1 := ""
	b.Subscribe(func(s string) {
		str1 = str1 + s
	})

	str2 := ""
	b.Subscribe(func(s string) {
		str2 = str2 + s
	})

	b.Producer("hello")
	b.Producer(" ")
	b.Producer("world")

	time.Sleep(time.Second)
	assert.Equal(t, "hello world", str1)
	assert.Equal(t, "hello world", str2)
}
