package api

import (
	"log/slog"
	"net/http"

	ciHandler "github.com/barturba/ticket-tracker/internal/api/handlers/cihandler"
	companyHandler "github.com/barturba/ticket-tracker/internal/api/handlers/companyhandler"
	"github.com/barturba/ticket-tracker/internal/api/handlers/healthcheck"
	incidentHandler "github.com/barturba/ticket-tracker/internal/api/handlers/incidenthandler"
	userHandler "github.com/barturba/ticket-tracker/internal/api/handlers/userhandler"
	"github.com/barturba/ticket-tracker/internal/database"
)

func SetupRoutes(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	// Healthcheck
	mux.Handle("GET /v1/healthcheck", healthcheck.Get(logger))

	// Incidents
	mux.Handle("GET /v1/incidents", incidentHandler.Get(logger, db))
	mux.Handle("GET /v1/incidents_all", incidentHandler.GetAll(logger, db))
	mux.Handle("POST /v1/incidents", incidentHandler.Post(logger, db))
	mux.Handle("GET /v1/incidents/{id}", incidentHandler.GetByID(logger, db))
	mux.Handle("GET /v1/incidents_latest", incidentHandler.GetLatest(logger, db))
	mux.Handle("PUT /v1/incidents/{id}", incidentHandler.Put(logger, db))
	mux.Handle("DELETE /v1/incidents/{id}", incidentHandler.Delete(logger, db))

	// Companies
	mux.Handle("GET /v1/companies", companyHandler.Get(logger, db))
	mux.Handle("GET /v1/companies_all", companyHandler.GetAll(logger, db))
	mux.Handle("POST /v1/companies", companyHandler.Post(logger, db))
	mux.Handle("GET /v1/companies/{id}", companyHandler.GetByID(logger, db))
	mux.Handle("GET /v1/companies_latest", companyHandler.GetLatest(logger, db))
	mux.Handle("PUT /v1/companies/{id}", companyHandler.Put(logger, db))
	mux.Handle("DELETE /v1/companies/{id}", companyHandler.Delete(logger, db))

	// Users
	mux.Handle("GET /v1/users", userHandler.Get(logger, db))
	mux.Handle("GET /v1/users_all", userHandler.GetAll(logger, db))
	mux.Handle("POST /v1/users", userHandler.Post(logger, db))
	mux.Handle("GET /v1/users/{id}", userHandler.GetByID(logger, db))
	mux.Handle("GET /v1/users_latest", userHandler.GetLatest(logger, db))
	mux.Handle("PUT /v1/users/{id}", userHandler.Put(logger, db))
	mux.Handle("DELETE /v1/users/{id}", userHandler.Delete(logger, db))

	// Configuration Items
	mux.Handle("GET /v1/cis", ciHandler.Get(logger, db))
	mux.Handle("GET /v1/cis_all", ciHandler.GetAll(logger, db))
	mux.Handle("POST /v1/cis", ciHandler.Post(logger, db))
	mux.Handle("GET /v1/cis/{id}", ciHandler.GetByID(logger, db))
	mux.Handle("GET /v1/cis_latest", ciHandler.GetLatest(logger, db))
	mux.Handle("PUT /v1/cis/{id}", ciHandler.Put(logger, db))
	mux.Handle("DELETE /v1/cis/{id}", ciHandler.Delete(logger, db))
}
