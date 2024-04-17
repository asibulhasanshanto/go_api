package services

import (
	"github.com/asibulhasanshanto/restapi/connection"
	"github.com/asibulhasanshanto/restapi/model"
)

var dbConnection *connection.PostgreSQL

func CreateVehicle(vehicle *model.Vehicle) *model.Vehicle {
	result := dbConnection.Db.Create(vehicle)

	if result.Error != nil {
		panic(result.Error)
	}

	return vehicle
}

func init() {
	// set the connection
	pg := connection.PostgreSQL{}
	pg.Connect()
	dbConnection = &pg

}
