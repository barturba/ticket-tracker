package api

import (
	"log/slog"
	"net/http"

	cihandlers "github.com/barturba/ticket-tracker/internal/api/handlers/cis"
	companyhandlers "github.com/barturba/ticket-tracker/internal/api/handlers/companies"
	"github.com/barturba/ticket-tracker/internal/api/handlers/healthcheck"
	incidenthandlers "github.com/barturba/ticket-tracker/internal/api/handlers/incidents"
	userhandlers "github.com/barturba/ticket-tracker/internal/api/handlers/users"
	"github.com/barturba/ticket-tracker/internal/database"
)

func AddRouteHealthcheck(mux *http.ServeMux, logger *slog.Logger) {
	mux.Handle("GET /v1/healthcheck", healthcheck.Get(logger))
}

func AddRoutesIncidents(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	mux.Handle("GET /v1/incidents", incidenthandlers.Get(logger, db))
	mux.Handle("GET /v1/incidents_all", incidenthandlers.GetAll(logger, db))
	mux.Handle("POST /v1/incidents", incidenthandlers.Post(logger, db))
	mux.Handle("GET /v1/incidents/{id}", incidenthandlers.GetByID(logger, db))
	mux.Handle("GET /v1/incidents_latest", incidenthandlers.GetLatest(logger, db))
	mux.Handle("PUT /v1/incidents/{id}", incidenthandlers.Put(logger, db))
	mux.Handle("DELETE /v1/incidents/{id}", incidenthandlers.Delete(logger, db))
}

func AddRoutesCompanies(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	mux.Handle("GET /v1/companies", companyhandlers.Get(logger, db))
	mux.Handle("GET /v1/companies_all", companyhandlers.GetAll(logger, db))
	mux.Handle("POST /v1/companies", companyhandlers.Post(logger, db))
	mux.Handle("GET /v1/companies/{id}", companyhandlers.GetByID(logger, db))
	mux.Handle("GET /v1/companies_latest", companyhandlers.GetLatest(logger, db))
	mux.Handle("PUT /v1/companies/{id}", companyhandlers.Put(logger, db))
	mux.Handle("DELETE /v1/companies/{id}", companyhandlers.Delete(logger, db))
}

func AddRoutesUsers(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	mux.Handle("GET /v1/users", userhandlers.Get(logger, db))
	mux.Handle("GET /v1/users_all", userhandlers.GetAll(logger, db))
	mux.Handle("POST /v1/users", userhandlers.Post(logger, db))
	mux.Handle("GET /v1/users/{id}", userhandlers.GetByID(logger, db))
	mux.Handle("GET /v1/users_latest", userhandlers.GetLatest(logger, db))
	mux.Handle("PUT /v1/users/{id}", userhandlers.Put(logger, db))
	mux.Handle("DELETE /v1/users/{id}", userhandlers.Delete(logger, db))
}

func AddRoutesConfigurationItems(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	mux.Handle("GET /v1/cis", cihandlers.Get(logger, db))
	mux.Handle("GET /v1/cis_all", cihandlers.GetAll(logger, db))
	mux.Handle("POST /v1/cis", cihandlers.Post(logger, db))
	mux.Handle("GET /v1/cis/{id}", cihandlers.GetByID(logger, db))
	mux.Handle("GET /v1/cis_latest", cihandlers.GetLatest(logger, db))
	mux.Handle("PUT /v1/cis/{id}", cihandlers.Put(logger, db))
	mux.Handle("DELETE /v1/cis/{id}", cihandlers.Delete(logger, db))
}
