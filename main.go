package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nerd500/clinic-appointment/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	// env config
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT string not found in the env")
	}
	dbURL := os.Getenv("dbURL")
	if dbURL == "" {
		log.Fatal("dbURL string not found in the env")
	}

	// DB Config
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	dbQuery := database.New(conn)

	apiCfg := apiConfig{
		DB: dbQuery,
	}

	// router cinfig
	router := GetRouter()
	router.Mount("/v1", V1Router(apiCfg))

	// server config
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// Server starting and error handling
	log.Println("Starting server at port: ", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("port: ", portString)

}
