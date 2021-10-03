package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	c := &Calculator{map[string]Calculate{}}
	c.Save("add", &Adder{})
	fmt.Println(c.CalculateRes("add", 1, 2))
}

type Adder struct {
}

func (a *Adder) GetResult(k int, b int) int {
	return k + b
}
