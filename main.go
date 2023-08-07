package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nerd500/clinic-appointment/routes"
)

func main() {

	// env config
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT string not found in the env")
	}

	// router cinfig
	router := routes.GetRouter()
	router.Mount("/v1", routes.V1Router())

	// server config
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// Server starting and error handling
	log.Println("Starting server at port: ", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("port: ", portString)

}
