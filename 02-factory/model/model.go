package model

// Car 產品 interface
type Car interface {
	Drive() string
}

// Toyota 產品實例
type Toyota struct{}

func (c *Toyota) Drive() string {
	return "Toyota is running"
}

type Factory interface {
	CreateCar() Car
}

// ToyotaFactory ConcreteCreator 實現工廠接口
type ToyotaFactory struct{}

func (c *ToyotaFactory) CreateCar() Car {
	return &Toyota{}
}
