package server

import (
	"database/sql"
	"embed"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/joho/godotenv"
)

//go:generate ./tailwindcss -o static/css/tailwind.css

//go:embed static
var static embed.FS

const (
	JWT_EXPIRES_IN_SECONDS = 3600
)

type ApiConfig struct {
	DB                    *database.Queries
	JWTSecret             string
	MenuItems             models.MenuItems
	ProfileItems          models.MenuItems
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

	menuItems := models.MenuItems{
		models.MenuItem{
			Name: "Incidents List",
			Link: "/incidents",
		},
		models.MenuItem{
			Name: "Configuration Items List",
			Link: "/configuration-items",
		},
		models.MenuItem{
			Name: "Companies List",
			Link: "/companies",
		},
		models.MenuItem{
			Name: "Users List",

			Link: "/users",
		}}

	profileItems := models.MenuItems{
		models.MenuItem{
			Name: "Settings",
			Link: "/settings",
		},
		models.MenuItem{
			Name: "Log Out",
			Link: "/logout",
		}}

	logo := "/static/images/logo.png"
	ProfilePicPlaceholder := "/static/images/profile_placeholder.webp"

	apiCfg := ApiConfig{
		DB:                    dbQueries,
		JWTSecret:             jwtSecret,
		MenuItems:             menuItems,
		ProfileItems:          profileItems,
		Logo:                  logo,
		ProfilePicPlaceholder: ProfilePicPlaceholder,
	}

	srv := &http.Server{
		Handler:      apiCfg.Routes(),
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	return srv
}
