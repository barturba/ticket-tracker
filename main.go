package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", apiCfg.handleUsers)
	mux.HandleFunc("GET /v1/organizations", apiCfg.middlewareAuth(apiCfg.getOrganizations))
	mux.HandleFunc("PUT /v1/organizations", apiCfg.middlewareAuth(apiCfg.updateOrganizations))
	mux.HandleFunc("POST /v1/configuration-items", apiCfg.middlewareAuth(apiCfg.handleConfigurationItems))
	mux.HandleFunc("GET /v1/configuration-items", apiCfg.middlewareAuth(apiCfg.getConfigurationItems))
	mux.HandleFunc("POST /v1/companies", apiCfg.middlewareAuth(apiCfg.handleCompanies))

	fmt.Printf("ticket-tracker\n")
	srv := http.Server{
		Handler:      mux,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on ", port)
	err = srv.ListenAndServe()
	log.Fatal(err)

	fmt.Printf("the ticket-tracker has started\n")
}
