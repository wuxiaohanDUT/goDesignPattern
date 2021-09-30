package decorator

import "testing"

func TestDecorator(t *testing.T) {
	water := &Water{}
	coffe := Coffee{Water: water}
	coffe.DrinkIt()
}
