// main.go is the entry point for ticket-tracker.
package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// main is the entry point for the ticket-tracker application. It calls the run function.
func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("ticket-tracker has started\n")
}

// run initializes and starts the application.
func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// Create a structured logger for logging messages.
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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
		Env:  env,
		Host: host,
		Port: port,
	}

	// Open a database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("started database connection", "stats", db.Stats())

	// Ping the database to verify the connection.
	logger.Info("pinging db to verify the connection")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)

	// Create a new server instance.
	srv := newServer(logger, config, dbQueries)
	httpServer := &http.Server{
		Addr:         net.JoinHostPort(config.Host, config.Port),
		Handler:      srv,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the HTTP server in a new goroutine.
	go func() {
		logger.Info("starting server", "addr", httpServer.Addr, "env", config.Env, "host", config.Host, "port", config.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("error listening and serving", "error", err)
		}
	}()

	// Wait for an interrupt signal to gracefully shut down the server.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx := context.Background()
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)
		defer cancel()
		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logger.Error("error shutting down http server", "error", err)
		}
	}()
	wg.Wait()
	return nil
}

// newServer sets up the HTTP server by creating a new ServeMux and adding
// routes for incidents, companies, users, and configuration items.
//
// Returns an HTTP handler that can be used by the HTTP server.
func newServer(logger *slog.Logger, config data.Config, db *database.Queries) http.Handler {

	mux := http.NewServeMux()

	addRouteHealthcheck(mux, logger)
	addRoutesIncidents(mux, logger, db)
	addRoutesCompanies(mux, logger, db)
	addRoutesUsers(mux, logger, db)
	addRoutesConfigurationItems(mux, logger, db)

	var handler http.Handler = mux

	return handler
}
