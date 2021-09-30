package factory_method

import "testing"

func TestFactory(t *testing.T) {
	coffeFactory := &CoffeeFactory{}
	milkFactory := &MilkFactory{}
	coffee := coffeFactory.CreateDrinking()
	milk := milkFactory.CreateDrinking()
	coffee.Drink()
	milk.Drink()
}
