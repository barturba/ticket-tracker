package main

import (
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/server/cis"
	"github.com/barturba/ticket-tracker/internal/server/companies"
	"github.com/barturba/ticket-tracker/internal/server/incidents"
	"github.com/barturba/ticket-tracker/internal/server/users"
)

func addRoutesIncidents(
	mux *http.ServeMux,
	logger *slog.Logger,
	config config,
	db *database.Queries) {
	mux.Handle("GET /v1/incidents", incidents.Get(logger, db))
	mux.Handle("POST /v1/incidents", incidents.Post(logger, db))
	mux.Handle("GET /v1/incidents/{id}", incidents.GetByID(logger, db))
	mux.Handle("GET /v1/incidents_count", incidents.GetCount(logger, db))
	mux.Handle("GET /v1/incidents_latest", incidents.GetLatest(logger, db))
	mux.Handle("PUT /v1/incidents/{id}", incidents.Put(logger, db))
	mux.Handle("DELETE /v1/incidents/{id}", incidents.Delete(logger, db))
}

func addRoutesCompanies(
	mux *http.ServeMux,
	logger *slog.Logger,
	config config,
	db *database.Queries) {
	mux.Handle("GET /v1/companies", companies.Get(logger, db))
	mux.Handle("POST /v1/companies", companies.Post(logger, db))
	mux.Handle("GET /v1/companies/{id}", companies.GetByID(logger, db))
	mux.Handle("GET /v1/companies_count", companies.GetCount(logger, db))
	mux.Handle("GET /v1/companies_latest", companies.GetLatest(logger, db))
	mux.Handle("PUT /v1/companies/{id}", companies.Put(logger, db))
	mux.Handle("DELETE /v1/companies/{id}", companies.Delete(logger, db))
}

func addRoutesUsers(
	mux *http.ServeMux,
	logger *slog.Logger,
	config config,
	db *database.Queries) {
	mux.Handle("GET /v1/users", users.Get(logger, db))
	mux.Handle("POST /v1/users", users.Post(logger, db))
	mux.Handle("GET /v1/users/{id}", users.GetByID(logger, db))
	mux.Handle("GET /v1/users_count", users.GetCount(logger, db))
	mux.Handle("GET /v1/users_latest", users.GetLatest(logger, db))
	mux.Handle("PUT /v1/users/{id}", users.Put(logger, db))
	mux.Handle("DELETE /v1/users/{id}", users.Delete(logger, db))
}

func addRoutesConfigurationItems(
	mux *http.ServeMux,
	logger *slog.Logger,
	config config,
	db *database.Queries) {
	mux.Handle("GET /v1/cis", cis.Get(logger, db))
	mux.Handle("POST /v1/cis", cis.Post(logger, db))
	mux.Handle("GET /v1/cis/{id}", cis.GetByID(logger, db))
	mux.Handle("GET /v1/cis_count", cis.GetCount(logger, db))
	mux.Handle("GET /v1/cis_latest", cis.GetLatest(logger, db))
	mux.Handle("PUT /v1/cis/{id}", cis.Put(logger, db))
	mux.Handle("DELETE /v1/cis/{id}", cis.Delete(logger, db))
}
