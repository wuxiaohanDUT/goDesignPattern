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

type Milk struct {
	Water Drink
}

func (m *Milk) DrinkIt() {
	m.Water.DrinkIt()
	fmt.Println("Milk")
}

type Tee struct {
	Water Drink
}

func (t *Tee) DrinkIt() {
	t.Water.DrinkIt()
	fmt.Println("Tee")
}
