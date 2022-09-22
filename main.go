package main

import (
	"decorator/decorator"
	"decorator/decoratorCoffe"
	"fmt"
)

func main() {
	var beverage1 decoratorCoffe.Beverage = decoratorCoffe.NewEspresso()
	beverage1 = decoratorCoffe.NewMocha(beverage1)
	beverage1 = decoratorCoffe.NewSoy(beverage1)
	beverage1 = decoratorCoffe.NewMocha(beverage1)
	fmt.Printf("%v $%v \n", beverage1.GetDescription(), beverage1.Cost())

	var beverage2 decoratorCoffe.Beverage = decoratorCoffe.NewHouseBlend()
	beverage2 = decoratorCoffe.NewSoy(beverage2)
	beverage2 = decoratorCoffe.NewSoy(beverage2)
	beverage2 = decoratorCoffe.NewMocha(beverage2)
	fmt.Printf("%v $%v \n", beverage2.GetDescription(), beverage2.Cost())

	var camry decorator.AutoProduct = decorator.NewToyoutaCamry()
	camry = decorator.AddSpoilerSedan(camry)
	camry = decorator.AddSpoilerSedan(camry)
	fmt.Printf("%v $%v \n", camry.GetDescription(), camry.Cost())

	var cruiser decorator.AutoProduct = decorator.NewToyoutaLandCruiser()
	cruiser = decorator.AddSpoilerCrossover(cruiser)
	fmt.Printf("%v $%v \n", cruiser.GetDescription(), cruiser.Cost())
}
