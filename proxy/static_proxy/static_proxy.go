package static_proxy

import "fmt"

type RentHouse interface {
	rentHouse()
}

type HouseRenter struct {
}

func (h *HouseRenter) rentHouse() {
	fmt.Println("rent house")
}

type Proxy struct {
	RentHouse
}

func (p *Proxy) rentHouse() {
	fmt.Println("do something")
	p.RentHouse.rentHouse()
	fmt.Println("do something")
}
