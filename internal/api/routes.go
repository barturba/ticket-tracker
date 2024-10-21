package api

import (
	"log/slog"
	"net/http"

	cihandler "github.com/barturba/ticket-tracker/internal/api/handlers/cihandler"
	"github.com/barturba/ticket-tracker/internal/api/handlers/companyhandler"
	"github.com/barturba/ticket-tracker/internal/api/handlers/healthcheck"
	"github.com/barturba/ticket-tracker/internal/api/handlers/incidenthandler"
	"github.com/barturba/ticket-tracker/internal/api/handlers/userhandler"
	"github.com/barturba/ticket-tracker/internal/database"
)

func SetupRoutes(mux *http.ServeMux, logger *slog.Logger, db *database.Queries) {
	// Healthcheck
	mux.Handle("GET /v1/healthcheck", healthcheck.Get(logger))

	// Incidents
	mux.Handle("GET /v1/incidents", incidenthandler.ListIncidents(logger, db))
	mux.Handle("GET /v1/incidents_all", incidenthandler.ListAllIncidents(logger, db))
	mux.Handle("POST /v1/incidents", incidenthandler.UpdateIncident(logger, db))
	mux.Handle("GET /v1/incidents/{id}", incidenthandler.GetIncidentByID(logger, db))
	mux.Handle("GET /v1/incidents_latest", incidenthandler.ListRecentIncidents(logger, db))
	mux.Handle("PUT /v1/incidents/{id}", incidenthandler.UpdateIncident(logger, db))
	mux.Handle("DELETE /v1/incidents/{id}", incidenthandler.DeleteIncident(logger, db))

	// Companies
	mux.Handle("GET /v1/companies", companyhandler.ListCompanies(logger, db))
	mux.Handle("GET /v1/companies_all", companyhandler.ListAllCompanies(logger, db))
	mux.Handle("POST /v1/companies", companyhandler.CreateCompany(logger, db))
	mux.Handle("GET /v1/companies/{id}", companyhandler.GetCompanyByID(logger, db))
	mux.Handle("GET /v1/companies_latest", companyhandler.ListRecentCompanies(logger, db))
	mux.Handle("PUT /v1/companies/{id}", companyhandler.UpdateCompany(logger, db))
	mux.Handle("DELETE /v1/companies/{id}", companyhandler.DeleteCompany(logger, db))

	// Users
	mux.Handle("GET /v1/users", userhandler.ListUsers(logger, db))
	mux.Handle("GET /v1/users_all", userhandler.ListAllUsers(logger, db))
	mux.Handle("POST /v1/users", userhandler.CreateUser(logger, db))
	mux.Handle("GET /v1/users/{id}", userhandler.GetUserByID(logger, db))
	mux.Handle("GET /v1/users_latest", userhandler.ListRecentUsers(logger, db))
	mux.Handle("PUT /v1/users/{id}", userhandler.UpdateUser(logger, db))
	mux.Handle("DELETE /v1/users/{id}", userhandler.DeleteUser(logger, db))

	// Configuration Items
	mux.Handle("GET /v1/cis", cihandler.Get(logger, db))
	mux.Handle("GET /v1/cis_all", cihandler.GetAll(logger, db))
	mux.Handle("POST /v1/cis", cihandler.Post(logger, db))
	mux.Handle("GET /v1/cis/{id}", cihandler.GetByID(logger, db))
	mux.Handle("GET /v1/cis_latest", cihandler.GetLatest(logger, db))
	mux.Handle("PUT /v1/cis/{id}", cihandler.Put(logger, db))
	mux.Handle("DELETE /v1/cis/{id}", cihandler.Delete(logger, db))
}
