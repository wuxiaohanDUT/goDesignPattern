package abstract_factory

import "testing"

func TestFactory(t *testing.T) {
	factoryA := &FactoryA{}
	factoryB := &FactoryB{}
	d1 := factoryA.CreateDrink()
	d2 := factoryB.CreateDrink()
	e1 := factoryA.CreateEat()
	e2 := factoryB.CreateEat()
	d1.Drink()
	d2.Drink()
	e1.Eat()
	e2.Eat()
}
