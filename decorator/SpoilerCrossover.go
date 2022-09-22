package decorator

type SpoilerCrossover struct {
	car Car
}

func (s *SpoilerCrossover) GetClass() string {
	return s.car.GetClass()
}

func (s *SpoilerCrossover) GetDescription() string {

	if s.car.GetClass() == "crossover" {
		return s.car.GetDescription() + ", Spoiler"
	}

	return s.car.GetDescription()
}

func (s *SpoilerCrossover) Cost() float64 {

	if s.car.GetClass() == "crossover" {
		return s.car.Cost() + 1500
	}

	return s.car.Cost()
}

func AddSpoilerCrossover(product AutoProduct) *SpoilerCrossover {
	switch product.(type) {
	case Car:
		s := new(SpoilerCrossover)
		s.car = product.(Car)
		return s
	}
	return nil
}
