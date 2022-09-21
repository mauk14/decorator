package decorator

type CondimentDecorator interface {
	GetDescription() string
	Cost() float64
}
