package api

import (
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/api/handlers"
	"github.com/barturba/ticket-tracker/internal/database"
)

func SetupRoutes(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	// Healthcheck
	mux.Handle("GET /v1/healthcheck", handlers.GetHealthcheck(logger))

	// Incidents
	mux.Handle("GET /v1/incidents", handlers.ListIncidents(logger, db))
	mux.Handle("GET /v1/incidents_all", handlers.ListAllIncidents(logger, db))
	mux.Handle("POST /v1/incidents", handlers.UpdateIncident(logger, db))
	mux.Handle("GET /v1/incidents/{id}", handlers.GetIncidentByID(logger, db))
	mux.Handle("GET /v1/incidents_latest", handlers.ListRecentIncidents(logger, db))
	mux.Handle("PUT /v1/incidents/{id}", handlers.UpdateIncident(logger, db))
	mux.Handle("DELETE /v1/incidents/{id}", handlers.DeleteIncident(logger, db))

	// Companies
	mux.Handle("GET /v1/companies", handlers.ListCompanies(logger, db))
	mux.Handle("GET /v1/companies_all", handlers.ListAllCompanies(logger, db))
	mux.Handle("POST /v1/companies", handlers.CreateCompany(logger, db))
	mux.Handle("GET /v1/companies/{id}", handlers.GetCompanyByID(logger, db))
	mux.Handle("GET /v1/companies_latest", handlers.ListRecentCompanies(logger, db))
	mux.Handle("PUT /v1/companies/{id}", handlers.UpdateCompany(logger, db))
	mux.Handle("DELETE /v1/companies/{id}", handlers.DeleteCompany(logger, db))

	// Users
	mux.Handle("GET /v1/users", handlers.ListUsers(logger, db))
	mux.Handle("GET /v1/users_all", handlers.ListAllUsers(logger, db))
	mux.Handle("POST /v1/users", handlers.CreateUser(logger, db))
	mux.Handle("GET /v1/users/{id}", handlers.GetUser(logger, db))
	mux.Handle("GET /v1/users_latest", handlers.ListRecentUsers(logger, db))
	mux.Handle("PUT /v1/users/{id}", handlers.UpdateUser(logger, db))
	mux.Handle("DELETE /v1/users/{id}", handlers.DeleteUser(logger, db))

	// Configuration Items
	mux.Handle("GET /v1/cis", handlers.ListCIs(logger, db))
	mux.Handle("GET /v1/cis_all", handlers.ListAllCIs(logger, db))
	mux.Handle("POST /v1/cis", handlers.CreateCI(logger, db))
	mux.Handle("GET /v1/cis/{id}", handlers.GetCI(logger, db))
	mux.Handle("GET /v1/cis_latest", handlers.ListRecentCIs(logger, db))
	mux.Handle("PUT /v1/cis/{id}", handlers.UpdateCI(logger, db))
	mux.Handle("DELETE /v1/cis/{id}", handlers.DeleteCI(logger, db))
}
