package abstract_factory

import "fmt"

type DrinkAble interface {
	Drink()
}

type EatAble interface {
	Eat()
}

type DrinkingA struct {
}

func (d *DrinkingA) Drink() {
	fmt.Println("DrinkingA")
}

type EatingA struct {
}

func (e *EatingA) Eat() {
	fmt.Println("EatingA")
}

type DrinkingB struct {
}

func (d *DrinkingB) Drink() {
	fmt.Println("DrinkingB")
}

type EatingB struct {
}

func (e *EatingB) Eat() {
	fmt.Println("EatingB")
}

type Factory interface {
	CreateDrink() DrinkAble
	CreateEat() EatAble
}

type FactoryA struct {
}

func (f *FactoryA) CreateDrink() DrinkAble {
	return &DrinkingA{}
}

func (f *FactoryA) CreateEat() EatAble {
	return &EatingA{}
}

type FactoryB struct {
}

func (f *FactoryB) CreateDrink() DrinkAble {
	return &DrinkingB{}
}

func (f *FactoryB) CreateEat() EatAble {
	return &EatingB{}
}
