package main

import (
	"fmt"
	"golang-practice/builder"
)

func main() {
	cakeMaker := builder.CakeMaker{}
	chocolateCake := cakeMaker.MakeCake(&builder.ChocolateCakeBuilder{})
	vanillaCake := cakeMaker.MakeCake(&builder.VanillaCakeBuilder{})
	fmt.Println(chocolateCake)
	fmt.Println(vanillaCake)
}
