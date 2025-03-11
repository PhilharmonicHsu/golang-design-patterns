package main

import "fmt"
import "02-factory/model"

func main() {
	car := createCar(&model.ToyotaFactory{})
	fmt.Println(car.Drive())
}

func createCar(factory model.Factory) model.Car {
	return factory.CreateCar()
}
