package decorator

type ToyoutaLandCruiser struct {
	Crossover
}

func (t *ToyoutaLandCruiser) GetDescription() string {
	return "Toyouta Land Cruiser 300 (2022 year)"
}

func (t *ToyoutaLandCruiser) Cost() float64 {
	return 50000
}

func NewToyoutaLandCruiser() *ToyoutaLandCruiser {
	return new(ToyoutaLandCruiser)
}
