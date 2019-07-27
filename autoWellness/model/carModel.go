package model

//Car model
type Car struct {
	Name              string
	ManufacturingYear int64
	Colour            string
	CarNumber         string
}

//NewCar gives a new car instance
func NewCar(name string, yearofManufacturing int64, colour string, carNumber string) *Car {

	car := Car{Name: name, ManufacturingYear: yearofManufacturing, Colour: colour, CarNumber: carNumber}
	return &car
}
