package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/joho/godotenv"
)

const (
	JWT_EXPIRES_IN_SECONDS = 3600
)

type ApiConfig struct {
	DB                    *database.Queries
	JWTSecret             string
	MenuItems             models.MenuItems
	Logo                  string
	ProfilePicPlaceholder string
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

	srv := &http.Server{
		Handler:      apiCfg.Routes(),
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	return srv
}
