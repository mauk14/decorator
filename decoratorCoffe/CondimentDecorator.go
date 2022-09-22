package decoratorCoffe

type CondimentDecorator interface {
	GetDescription() string
	Cost() float64
}
