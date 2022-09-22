package decoratorCoffe

type Mocha struct {
	beverage Beverage
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() float64 {
	return m.beverage.Cost() + 0.2
}

func NewMocha(beverage Beverage) *Mocha {
	mocha := new(Mocha)
	mocha.beverage = beverage
	return mocha
}
