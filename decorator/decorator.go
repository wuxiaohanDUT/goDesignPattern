package decorator

import "fmt"

type Drink interface {
	DrinkIt()
}

type Water struct {
}

func (w *Water) DrinkIt() {
	fmt.Println("Water")
}

type Coffee struct {
	Water Drink
}

func (c *Coffee) DrinkIt() {
	c.Water.DrinkIt()
	fmt.Println("Coffee")
}
