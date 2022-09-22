package decorator

type Car interface {
	GetClass() string
	GetDescription() string
	Cost() float64
}
