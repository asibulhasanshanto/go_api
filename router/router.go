package router

import (
	"github.com/asibulhasanshanto/restapi/controller"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {

	r := chi.NewRouter()

	r.Route("/vehicles", func(r chi.Router) {
		r.Post("/", controller.CreateVehicle)
	})

	return r

}
