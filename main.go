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
	"golang.org/x/crypto/bcrypt"
)

const (
	JWT_EXPIRES_IN_SECONDS = 3600
)

type apiConfig struct {
	DB        *database.Queries
	JWTSecret string
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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	adminPasswordTest := os.Getenv("ADMIN_PASSWORD")
	if adminPasswordTest == "" {
		log.Fatal("ADMIN_PASSWORD environment variable is not set")
	}
	dat, err := bcrypt.GenerateFromPassword([]byte(adminPasswordTest), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("couldn't generate admin password")
	}
	fmt.Printf("%s\n", string(dat))

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB:        dbQueries,
		JWTSecret: jwtSecret,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/users", apiCfg.handleUsers)
	mux.HandleFunc("GET /v1/organizations", apiCfg.middlewareAuth(apiCfg.getOrganizations))
	mux.HandleFunc("PUT /v1/organizations", apiCfg.middlewareAuth(apiCfg.updateOrganizations))
	mux.HandleFunc("POST /v1/configuration-items", apiCfg.middlewareAuth(apiCfg.handleConfigurationItems))
	mux.HandleFunc("GET /v1/configuration-items", apiCfg.middlewareAuth(apiCfg.getConfigurationItems))
	mux.HandleFunc("POST /v1/companies", apiCfg.middlewareAuth(apiCfg.handleCompanies))
	mux.HandleFunc("POST /v1/incidents", apiCfg.middlewareAuth(apiCfg.handleIncidents))
	mux.HandleFunc("GET /v1/incidents", apiCfg.middlewareAuth(apiCfg.getIncidents))
	mux.HandleFunc("GET /incidents", apiCfg.middlewareAuth(apiCfg.handleIncidentsPage))
	mux.HandleFunc("GET /incidents/{id}/edit", apiCfg.middlewareAuth(apiCfg.handleIncidentsEditPage))
	mux.HandleFunc("POST /v1/login", apiCfg.handleLogin)
	mux.HandleFunc("GET /login", apiCfg.handleLoginPage)
	mux.HandleFunc("GET /get", apiCfg.getCookieHandler)

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
