package decorator

import "testing"

func TestDecorator(t *testing.T) {
	water := &Water{}
	milk := &Milk{Water: water}
	milkTee := &Tee{Water: milk}
	milkTee.DrinkIt()
}
