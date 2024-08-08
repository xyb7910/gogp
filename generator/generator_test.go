package generator

import (
	"fmt"
	"testing"
)

func TestGeneratorIntA(t *testing.T) {
	ch := GeneratorIntA()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
