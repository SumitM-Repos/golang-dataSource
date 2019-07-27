package dataSource

import "github.com/ContinuumLLC/autoWellness/model"

type DBS struct {
	DBs []DB `json:"DBs"`
}
type DB struct {
	UserID   string `json:"USERID"`
	Password string `json:"Password"`
	Server   string `json:"Server"`
	Database string `json:"Database"`
}

type DataSource interface {
	GetCars() (*[]model.Car, error)
	//GetCarsByName(name string) (*[]model.Car, error)
	//GetCarsByManufacturingYear(year int64) (*[]model.Car, error)
	Close()
	AddCar(car *model.Car) error
}

func GetDB(connectionString string) (DataSource, error) {
	return NewMySqlDataStore(connectionString)

}
