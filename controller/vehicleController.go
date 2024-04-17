package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asibulhasanshanto/restapi/model"
	"github.com/asibulhasanshanto/restapi/services"
)

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	// get json data from request
	vehicle := &model.Vehicle{}

	

	// decode json data
	err := json.NewDecoder(r.Body).Decode(vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// create vehicle
	vehicle = services.CreateVehicle(vehicle)

	// return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicle)

}
