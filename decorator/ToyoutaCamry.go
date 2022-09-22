package decorator

type ToyoutaCamry struct {
}

func (t *ToyoutaCamry) GetClass() string {
	return "sedan"
}

func (t *ToyoutaCamry) GetDescription() string {
	return "Toyouta Camry 3.5б (2022 year)"
}

func (t *ToyoutaCamry) Cost() float64 {
	return 30000
}

func NewToyoutaCamry() *ToyoutaCamry {
	return new(ToyoutaCamry)
}
