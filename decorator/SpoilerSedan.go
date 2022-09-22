package decorator

type SpoilerSedan struct {
	car Car
}

func (s *SpoilerSedan) GetClass() string {
	return s.car.GetClass()
}

func (s *SpoilerSedan) GetDescription() string {

	if s.car.GetClass() == "sedan" {
		return s.car.GetDescription() + ", Spoiler"
	}

	return s.car.GetDescription()
}

func (s *SpoilerSedan) Cost() float64 {

	if s.car.GetClass() == "sedan" {
		return s.car.Cost() + 1000
	}

	return s.car.Cost()
}

func AddSpoilerSedan(product AutoProduct) *SpoilerSedan {
	switch product.(type) {
	case Car:
		s := new(SpoilerSedan)
		s.car = product.(Car)
		return s
	}
	return nil
}
