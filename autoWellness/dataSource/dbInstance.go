package dataSource

import (
	"database/sql"
	"log"

	"github.com/ContinuumLLC/autoWellness/model"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDataSource struct {
	*sql.DB
}

func NewMySqlDataStore(connectionString string) (*MySQLDataSource, error) {
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
	}
	return &MySQLDataSource{DB: db}, nil
}

func (mds *MySQLDataSource) Close() {
	mds.Close()
}

func (mds *MySQLDataSource) AddCar(car *model.Car) error {

	tx, err := mds.Begin()
	if err != nil {
		log.Print(err)
		return err
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare(AddCarsInDB)
	if err != nil {
		log.Print(err)
		return err
	}

	_, err = stmt.Exec(car.Name, car.ManufacturingYear, car.Colour, car.CarNumber)
	if err != nil {
		log.Print(err)
		return err
	}

	defer stmt.Close()

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (mds *MySQLDataSource) GetCars() (*[]model.Car, error) {

	rows, err := mds.Query(GetCars)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	cars := []model.Car{}
	for rows.Next() {

		car := model.Car{}
		err := rows.Scan(&car.Name, &car.ManufacturingYear, &car.Colour, &car.CarNumber)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		cars = append(cars, car)

	}
	return &cars, nil
}
