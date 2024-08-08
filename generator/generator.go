package generator

import (
	"math/rand"
)

// 实现一个最简单的带缓冲区的生成器
func GeneratorIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}
