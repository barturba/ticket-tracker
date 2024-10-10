package server

import "net/http"

func (cfg *ApiConfig) Routes() *http.ServeMux {

	mux := http.NewServeMux()

	// Page Endpoints

	// - Incidents

	mux.HandleFunc("GET /v1/incidents", cfg.handleIncidentsGet)
	mux.HandleFunc("GET /v1/filtered_incidents", cfg.handleFilteredIncidentsGet)
	mux.HandleFunc("GET /v1/filtered_incidents_count", cfg.handleFilteredIncidentsCountGet)
	mux.HandleFunc("GET /v1/incident_by_id", cfg.handleIncidentByIdGet)
	mux.HandleFunc("GET /v1/incidents_latest", cfg.handleIncidentsLatestGet)
	mux.HandleFunc("POST /v1/incidents", cfg.handleIncidentsPost)
	mux.HandleFunc("PUT /v1/incidents/{id}", cfg.handleIncidentsPut)
	mux.HandleFunc("DELETE /v1/incidents/{id}", cfg.handleIncidentsDelete)

	// - Companies
	mux.HandleFunc("GET /v1/companies", cfg.handleCompaniesGet)
	mux.HandleFunc("GET /v1/filtered_companies", cfg.handleFilteredCompaniesGet)
	mux.HandleFunc("GET /v1/filtered_companies_count", cfg.handleFilteredCompaniesCountGet)
	mux.HandleFunc("GET /v1/company_by_id", cfg.handleCompanyByIdGet)
	mux.HandleFunc("POST /v1/companies", cfg.handleCompaniesPost)
	mux.HandleFunc("PUT /v1/companies/{id}", cfg.handleCompaniesPut)

	// - Configuration Items
	mux.HandleFunc("GET /v1/configuration_items", cfg.handleConfigurationItemsGet)

	// - Users
	mux.HandleFunc("GET /v1/users", cfg.handleUsersGet)
	mux.HandleFunc("GET /v1/users_by_company", cfg.handleUsersByCompanyGet)

	// Login Endpoints
	mux.HandleFunc("POST /v1/login", cfg.handleLogin)
	mux.HandleFunc("GET /logout", cfg.handleLogout)

	return mux
}
