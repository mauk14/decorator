package main

import (
	"decorator/decorator"
	"fmt"
)

func main() {
	var beverage1 decorator.Beverage = decorator.NewEspresso()
	beverage1 = decorator.NewMocha(beverage1)
	beverage1 = decorator.NewSoy(beverage1)
	beverage1 = decorator.NewMocha(beverage1)
	fmt.Printf("%v $%v \n", beverage1.GetDescription(), beverage1.Cost())

	var beverage2 decorator.Beverage = decorator.NewHouseBlend()
	beverage2 = decorator.NewSoy(beverage2)
	beverage2 = decorator.NewSoy(beverage2)
	beverage2 = decorator.NewMocha(beverage2)
	fmt.Printf("%v $%v \n", beverage2.GetDescription(), beverage2.Cost())
}
