package decorator

type HouseBlend struct {
	description string
}

func (e *HouseBlend) Cost() float64 {
	return 0.89
}

func (e *HouseBlend) GetDescription() string {
	return e.description
}

func NewHouseBlend() *HouseBlend {
	houseBlend := new(HouseBlend)
	houseBlend.description = "House Blend Coffee"
	return houseBlend
}
