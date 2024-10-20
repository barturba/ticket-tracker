package config

import (
	"log"
	"os"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/joho/godotenv"
)

func Config() data.Config {

	// Load ENV variables from .env file.
	_ = godotenv.Load()

	// Get the environment type.
	env := os.Getenv("ENV")
	if env != "development" && env != "production" {
		env = "production" // Default to production if not set.
	}

	// Get the host name.
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		log.Fatal("SERVER_HOST environment variable is not set")
	}

	// Get the port number.
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT environment variable is not set")
	}

	// Set the database URL based on the environment.
	dbURL := ""
	if env == "development" {
		dbURL = os.Getenv("DATABASE_URL_DEV")
		if dbURL == "" {
			log.Fatal("DATABASE_URL_DEV environment variable is not set")
		}
	} else {
		dbURL = os.Getenv("DATABASE_URL_PROD")
		if dbURL == "" {
			log.Fatal("DATABASE_URL_PROD environment variable is not set")
		}
	}

	// Create a config structure to hold the configuration values.
	config := data.Config{
		Env:   env,
		Host:  host,
		Port:  port,
		DBURL: dbURL,
	}
	return config
}
