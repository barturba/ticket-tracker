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

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/server/companies"
	"github.com/barturba/ticket-tracker/internal/server/incidents"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	Host string
	Port string
	Env  string
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("the ticket-tracker has started\n")
}

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// Create a logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load .env file")
	}

	// Create a config
	var config config

	// Get the host name
	host := os.Getenv("HOST")
	if host == "" {
		log.Fatal("HOST environment variable is not set")
	}

	// Get the port number
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Get the environment type
	env := os.Getenv("ENV")
	if env == "" {
		log.Fatal("ENV environment variable is not set")
	}

	// Get the db url
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Open a database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	config.Env = env
	config.Host = host
	config.Port = port

	srv := NewServer2(logger, config, dbQueries)
	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}
	go func() {
		// Start the HTTP server.
		logger.Info("starting server", "addr", httpServer.Addr, "env", config.Env)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("error listening and serving", "error", err)
		}
	}()
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

func NewServer2(logger *slog.Logger, config config, db *database.Queries) http.Handler {
	mux := http.NewServeMux()
	addRoutesIncident(mux, logger, config, db)
	addRoutesCompany(mux, logger, config, db)
	var handler http.Handler = mux
	// handler = someMiddleware(handler)
	return handler
}

func addRoutesIncident(mux *http.ServeMux, logger *slog.Logger, config config, db *database.Queries) {
	mux.Handle("GET /v1/incidents", incidents.Get(logger, db))
	mux.Handle("POST /v1/incidents", incidents.Post(logger, db))
	// Get a single one
	// mux.Handle("GET /v1/incidents", incidents.Get(logger, db))
	mux.Handle("GET /v1/incidents_count", incidents.GetCount(logger, db))
	mux.Handle("GET /v1/incidents_latest", incidents.GetLatest(logger, db))
	mux.Handle("GET /v1/incidents/{id}", incidents.GetByID(logger, db))
	mux.Handle("PUT /v1/incidents/{id}", incidents.Put(logger, db))
	mux.Handle("DELETE /v1/incidents/{id}", incidents.Delete(logger, db))
}

func addRoutesCompany(mux *http.ServeMux, logger *slog.Logger, config config, db *database.Queries) {
	mux.Handle("GET /v1/companies", companies.Get(logger, db))
}
