package server

import "net/http"

func (cfg *ApiConfig) Routes() *http.ServeMux {

	mux := http.NewServeMux()

	// Page Endpoints

	// - Incidents

	mux.HandleFunc("GET /incidents", cfg.middlewareAuthPage(cfg.handleViewIncidents))
	mux.HandleFunc("GET /search-incidents", cfg.middlewareAuthPage(cfg.handleSearchIncidents))
	mux.HandleFunc("GET /incidents/add", cfg.middlewareAuthPage(cfg.handleIncidentsAddPage))
	// mux.HandleFunc("POST /incidents", cfg.middlewareAuthPage(cfg.handleIncidentsPostPage))
	mux.HandleFunc("GET /incidents/{id}/edit", cfg.middlewareAuthPage(cfg.handleIncidentsEditPage))
	// mux.HandleFunc("PUT /incidents/{id}", cfg.middlewareAuthPage(cfg.handleIncidentsPutPage))

	mux.HandleFunc("GET /v1/incidents", cfg.handleIncidentsGet)
	mux.HandleFunc("GET /v1/incident_by_id", cfg.handleIncidentByIdGet)
	mux.HandleFunc("GET /v1/filtered_incidents", cfg.handleFilteredIncidentsGet)
	mux.HandleFunc("GET /v1/filtered_incidents_count", cfg.handleFilteredIncidentsCountGet)
	mux.HandleFunc("GET /v1/incidents_latest", cfg.handleIncidentsLatestGet)
	mux.HandleFunc("GET /v1/users_by_company", cfg.handleUsersByCompanyGet)
	mux.HandleFunc("GET /v1/companies", cfg.handleCompaniesGet)
	mux.HandleFunc("GET /v1/users", cfg.handleUsersGet)
	mux.HandleFunc("GET /v1/configuration_items", cfg.handleConfigurationItemsGet)
	mux.HandleFunc("POST /v1/incidents", cfg.handleIncidentsPost)
	mux.HandleFunc("PUT /v1/incidents/{id}", cfg.handleIncidentsPut)
	mux.HandleFunc("DELETE /v1/incidents/{id}", cfg.handleIncidentsDelete)

	// - Companies
	mux.HandleFunc("GET /v1/filtered_companies", cfg.handleFilteredCompaniesGet)
	mux.HandleFunc("GET /v1/filtered_companies_count", cfg.handleFilteredCompaniesCountGet)

	// - Configuration Items

	mux.HandleFunc("GET /configuration-items", cfg.middlewareAuthPage(cfg.handleViewConfigurationItems))
	mux.HandleFunc("GET /configuration-items-select", cfg.middlewareAuthPage(cfg.handleViewConfigurationItemsSelect))
	mux.HandleFunc("GET /configuration-items/{id}/edit", cfg.middlewareAuthPage(cfg.handleConfigurationItemsEditPage))
	mux.HandleFunc("POST /configuration-items", cfg.middlewareAuthPage(cfg.handleConfigurationItemsPostPage))

	// - Companies

	mux.HandleFunc("GET /companies", cfg.middlewareAuthPage(cfg.handleViewCompanies))

	// - Users

	mux.HandleFunc("GET /users", cfg.middlewareAuthPage(cfg.handleViewUsers))

	// Home

	mux.HandleFunc("GET /", cfg.middlewareAuthPageNoRedirect(cfg.handlePageIndex))

	// Login Endpoints
	mux.HandleFunc("POST /v1/login", cfg.handleLogin)
	mux.HandleFunc("GET /login", cfg.handleLoginPage)
	mux.HandleFunc("GET /logout", cfg.handleLogout)

	return mux
}
