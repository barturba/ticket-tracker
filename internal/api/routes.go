package api

import (
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/api/handlers/cis"
	"github.com/barturba/ticket-tracker/internal/api/handlers/companies"
	"github.com/barturba/ticket-tracker/internal/api/handlers/healthcheck"
	"github.com/barturba/ticket-tracker/internal/api/handlers/incidents"
	"github.com/barturba/ticket-tracker/internal/api/handlers/users"
	"github.com/barturba/ticket-tracker/internal/database"
)

func AddRouteHealthcheck(mux *http.ServeMux, logger *slog.Logger) {
	mux.Handle("GET /v1/healthcheck", healthcheck.Get(logger))
}

func AddRoutesIncidents(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	mux.Handle("GET /v1/incidents", incidents.Get(logger, db))
	mux.Handle("GET /v1/incidents_all", incidents.GetAll(logger, db))
	mux.Handle("POST /v1/incidents", incidents.Post(logger, db))
	mux.Handle("GET /v1/incidents/{id}", incidents.GetByID(logger, db))
	mux.Handle("GET /v1/incidents_latest", incidents.GetLatest(logger, db))
	mux.Handle("PUT /v1/incidents/{id}", incidents.Put(logger, db))
	mux.Handle("DELETE /v1/incidents/{id}", incidents.Delete(logger, db))
}

func AddRoutesCompanies(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	mux.Handle("GET /v1/companies", companies.Get(logger, db))
	mux.Handle("GET /v1/companies_all", companies.GetAll(logger, db))
	mux.Handle("POST /v1/companies", companies.Post(logger, db))
	mux.Handle("GET /v1/companies/{id}", companies.GetByID(logger, db))
	mux.Handle("GET /v1/companies_latest", companies.GetLatest(logger, db))
	mux.Handle("PUT /v1/companies/{id}", companies.Put(logger, db))
	mux.Handle("DELETE /v1/companies/{id}", companies.Delete(logger, db))
}

func AddRoutesUsers(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	mux.Handle("GET /v1/users", users.Get(logger, db))
	mux.Handle("GET /v1/users_all", users.GetAll(logger, db))
	mux.Handle("POST /v1/users", users.Post(logger, db))
	mux.Handle("GET /v1/users/{id}", users.GetByID(logger, db))
	mux.Handle("GET /v1/users_latest", users.GetLatest(logger, db))
	mux.Handle("PUT /v1/users/{id}", users.Put(logger, db))
	mux.Handle("DELETE /v1/users/{id}", users.Delete(logger, db))
}

func AddRoutesConfigurationItems(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	mux.Handle("GET /v1/cis", cis.Get(logger, db))
	mux.Handle("GET /v1/cis_all", cis.GetAll(logger, db))
	mux.Handle("POST /v1/cis", cis.Post(logger, db))
	mux.Handle("GET /v1/cis/{id}", cis.GetByID(logger, db))
	mux.Handle("GET /v1/cis_latest", cis.GetLatest(logger, db))
	mux.Handle("PUT /v1/cis/{id}", cis.Put(logger, db))
	mux.Handle("DELETE /v1/cis/{id}", cis.Delete(logger, db))
}
