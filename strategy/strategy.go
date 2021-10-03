package strategy

type Calculate interface {
	GetResult(a int, b int) int
}

type Calculator struct {
	strategyContext map[string]Calculate
}

func (c *Calculator) Save(name string, inter Calculate) {
	c.strategyContext[name] = inter
}

func (c *Calculator) Get(name string) Calculate {
	if r, ok := c.strategyContext[name]; ok {
		return r
	}
	return nil
}

func (c *Calculator) CalculateRes(strategy string, a int, b int) int {
	s := c.Get(strategy)
	if s == nil {
		panic("calculate error")
	}
	return s.GetResult(a, b)
}
