package config

import (
	"errors"
	"os"

	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/joho/godotenv"
)

func Load() (models.Config, error) {

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
		return models.Config{}, errors.New("SERVER_HOST environment variable is not set")
	}

	// Get the port number.
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return models.Config{}, errors.New("SERVER_PORT environment variable is not set")
	}

	// Set the database URL based on the environment.
	dbURL := ""
	if env == "development" {
		dbURL = os.Getenv("DATABASE_URL_DEV")
		if dbURL == "" {
			return models.Config{}, errors.New("DATABASE_URL_DEV environment variable is not set")
		}
	} else {
		dbURL = os.Getenv("DATABASE_URL_PROD")
		if dbURL == "" {
			return models.Config{}, errors.New("DATABASE_URL_PROD environment variable is not set")
		}
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return models.Config{}, errors.New("JWT_SECRET environment variable is not set")
	}

	// Create a config structure to hold the configuration values.
	config := models.Config{
		Env:       env,
		Host:      host,
		Port:      port,
		DBURL:     dbURL,
		JWTSecret: jwtSecret,
	}

	// Validate the configuration
	err := config.Validate()

	return config, err
}
