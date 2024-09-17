package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/joho/godotenv"
)

const (
	JWT_EXPIRES_IN_SECONDS = 3600
)

type ApiConfig struct {
	DB        *database.Queries
	JWTSecret string
}

func NewServer() *http.Server {

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

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := ApiConfig{
		DB:        dbQueries,
		JWTSecret: jwtSecret,
	}

	mux := http.NewServeMux()

	// API Endpoints

	mux.HandleFunc("POST /v1/users", apiCfg.handleUsers)
	mux.HandleFunc("GET /v1/organizations", apiCfg.middlewareAuth(apiCfg.getOrganizations))
	mux.HandleFunc("PUT /v1/organizations", apiCfg.middlewareAuth(apiCfg.updateOrganizations))
	mux.HandleFunc("POST /v1/configuration-items", apiCfg.middlewareAuth(apiCfg.handleConfigurationItems))
	mux.HandleFunc("GET /v1/configuration-items", apiCfg.middlewareAuth(apiCfg.getConfigurationItems))
	mux.HandleFunc("POST /v1/companies", apiCfg.middlewareAuth(apiCfg.handleCompanies))
	mux.HandleFunc("POST /v1/incidents", apiCfg.middlewareAuth(apiCfg.handleIncidents))
	mux.HandleFunc("GET /v1/incidents", apiCfg.middlewareAuth(apiCfg.getIncidents))

	// Page Endpoints

	mux.HandleFunc("GET /companies", apiCfg.middlewareAuthPage(apiCfg.handleCompaniesPage))
	mux.HandleFunc("GET /configuration-items", apiCfg.middlewareAuthPage(apiCfg.handleConfigurationItemsPage))

	mux.HandleFunc("GET /incidents", apiCfg.middlewareAuthPage(apiCfg.handleIncidentsPage))
	mux.HandleFunc("GET /incidents/new", apiCfg.middlewareAuthPage(apiCfg.handleIncidentsNewPage))
	mux.HandleFunc("POST /incidents", apiCfg.middlewareAuthPage(apiCfg.handleIncidentsPostPage))
	mux.HandleFunc("GET /incidents/{id}/edit", apiCfg.middlewareAuthPage(apiCfg.handleIncidentsEditPage))
	mux.HandleFunc("GET /incidents/{id}", apiCfg.middlewareAuthPage(apiCfg.handleIncidentsGetPage))
	mux.HandleFunc("PUT /incidents/{id}", apiCfg.middlewareAuthPage(apiCfg.handleIncidentsUpdatePage))

	// Login Endpoints

	mux.HandleFunc("POST /v1/login", apiCfg.handleLogin)
	mux.HandleFunc("GET /login", apiCfg.handleLoginPage)
	mux.HandleFunc("GET /get", apiCfg.getCookieHandler)

	srv := &http.Server{
		Handler:      mux,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	return srv
}
