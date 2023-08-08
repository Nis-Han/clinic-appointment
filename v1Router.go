package main

import (
	"github.com/go-chi/chi"
)

// Function to create v1 router
//All the paths defined will be accessed as https://baseurl/v1/...

func V1Router(apiCfg apiConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", HealthCheck)
	router.Post("/user", apiCfg.CreateUser)
	return router
}
