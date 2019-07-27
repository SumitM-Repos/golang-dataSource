package dataSource

const (
	AddCarsInDB = "INSERT INTO Cars (Name,ManufacturingYear,Colour,CarNumber) VALUES (?,?,?,?)"

	GetCarsByName   = "Select c.Name, c.ManufacturingYear, c.Colour, c.CarNumber from Cars c where c.Name = ?"
	GetCars         = "Select c.Name, c.ManufacturingYear, c.Colour, c.CarNumber from Cars c"
	GetCarsByColour = "Select c.Name, c.ManufacturingYear, c.Colour, c.CarNumber from Cars c where c.Colour = ?"

	GetCarsByManufacturingYear = "Select c.Name, c.ManufacturingYear, c.Colour, c.CarNumber from Cars c where c.ManufacturingYear = ?"
)
