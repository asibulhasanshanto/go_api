package controller

import (
	"encoding/json"
	"net/http"

	"github.com/asibulhasanshanto/restapi/model"
	"github.com/asibulhasanshanto/restapi/services"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// get data from request
	user := &model.User{}

	// decode json data
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create user
	user = services.CreateUser(user)

	// return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
	
}
