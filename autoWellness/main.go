package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/ContinuumLLC/autoWellness/dataSource"
	"github.com/ContinuumLLC/autoWellness/model"
)

var DBConnectionString string

func init() {

	file, _ := ioutil.ReadFile("DBConnectionStrings.json")
	dbs := dataSource.DBS{}
	err := json.Unmarshal([]byte(file), &dbs)
	if err != nil {
		log.Print(err)
	}
	makeDBConnectionString(dbs)
	

}
func makeDBConnectionString(dbs dataSource.DBS) {
	userName := dbs.DBs[0].UserID + ":"
	pass := dbs.DBs[0].Password + "@"
	server := "tcp(" + dbs.DBs[0].Server + ")/"
	database := dbs.DBs[0].Database

	DBConnectionString = userName + pass + server + database

}
func main() {
	db, err := dataSource.GetDB(DBConnectionString)
	if err != nil {
		log.Println(err)
	}
	fmt.Print(db)
	defer db.Close()
	err = db.AddCar(&model.Car{Name: "Maruti", ManufacturingYear: 2007, Colour: "White", CarNumber: "DL3CB6594"})
	if err != nil {
		log.Println(err)
	}
	cars, _ := db.GetCars()
	fmt.Println(cars)
	time.Sleep(time.Second * 1200)
}
