package decoratorCoffe

type Soy struct {
	beverage    Beverage
	description string
}

func (m *Soy) GetDescription() string {
	return m.beverage.GetDescription() + ", Soy"
}

func (m *Soy) Cost() float64 {
	return m.beverage.Cost() + 0.2
}

func NewSoy(beverage Beverage) *Soy {
	soy := new(Soy)
	soy.beverage = beverage
	return soy
}
