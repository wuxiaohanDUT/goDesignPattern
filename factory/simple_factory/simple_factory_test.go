package simple_factory

import "testing"

func TestFactory(t *testing.T) {
	tee := CreatDrinking("Tee")
	coffee := CreatDrinking("Coffee")
	tee.Drink()
	coffee.Drink()
}
