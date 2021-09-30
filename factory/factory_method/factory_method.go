package factory_method

import "fmt"

type DrinkAble interface {
	Drink()
}

type Milk struct {
}

func (m *Milk) Drink() {
	fmt.Println("Drink Milk")
}

type Coffee struct {
}

func (c *Coffee) Drink() {
	fmt.Println("Drink Coffee")
}

type CreateDrinkingFactory interface {
	CreateDrinking() DrinkAble
}

type MilkFactory struct {
}

func (m *MilkFactory) CreateDrinking() DrinkAble {
	return &Milk{}
}

type CoffeeFactory struct {
}

func (c *CoffeeFactory) CreateDrinking() DrinkAble {
	return &Coffee{}
}
