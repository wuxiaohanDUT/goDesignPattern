package simple_factory

import "fmt"

type DrinkAble interface {
	Drink()
}

type Tee struct {
}

func (t *Tee) Drink() {
	fmt.Println("Drink Tee")
}

type Coffee struct {
}

func (c *Coffee) Drink() {
	fmt.Println("Drink Coffee")
}

func CreatDrinking(s string) DrinkAble {
	switch s {
	case "Tee":
		return &Tee{}
	case "Coffee":
		return &Coffee{}
	}
	return nil
}
