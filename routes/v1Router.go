package routes

import (
	"github.com/go-chi/chi"
	controller "github.com/nerd500/clinic-appointment/api/api_controller"
)

// Function to create v1 router
//All the paths defined will be accessed as https://baseurl/v1/...

func V1Router() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", controller.HealthCheck)
	return router
}
