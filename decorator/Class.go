package decorator

type Sedan struct {
}

func (s *Sedan) GetClass() string {
	return "sedan"
}

type Crossover struct {
}

func (crossover *Crossover) GetClass() string {
	return "crossover"
}
