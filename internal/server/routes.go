package server

import "net/http"

func (cfg *ApiConfig) Routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServer(http.FS(static)))

	// API Endpoints

	mux.HandleFunc("POST /v1/users", cfg.handleUsers)
	mux.HandleFunc("POST /v1/configuration-items", cfg.middlewareAuth(cfg.handleConfigurationItems))
	mux.HandleFunc("GET /v1/configuration-items", cfg.middlewareAuth(cfg.getConfigurationItems))
	mux.HandleFunc("POST /v1/incidents", cfg.middlewareAuth(cfg.handleIncidents))
	mux.HandleFunc("GET /v1/incidents", cfg.middlewareAuth(cfg.getIncidents))

	// Page Endpoints

	// - Incidents

	mux.HandleFunc("GET /incidents", cfg.middlewareAuthPage(cfg.handleViewIncidents))
	mux.HandleFunc("GET /search-incidents", cfg.middlewareAuthPage(cfg.handleSearchIncidents))
	// mux.HandleFunc("GET /incidents/new", cfg.middlewareAuthPage(cfg.handleIncidentsNewPage))
	mux.HandleFunc("POST /incidents", cfg.middlewareAuthPage(cfg.handleIncidentsPostPage))
	mux.HandleFunc("GET /incidents/{id}/edit", cfg.middlewareAuthPage(cfg.handleIncidentsEditPage))
	mux.HandleFunc("PUT /incidents/{id}", cfg.middlewareAuthPage(cfg.handleIncidentsPutPage))

	// - Configuration Items

	mux.HandleFunc("GET /configuration-items", cfg.middlewareAuthPage(cfg.handleViewConfigurationItems))
	mux.HandleFunc("GET /configuration-items-select", cfg.middlewareAuthPage(cfg.handleViewConfigurationItemsSelect))
	mux.HandleFunc("GET /configuration-items/{id}/edit", cfg.middlewareAuthPage(cfg.handleConfigurationItemsEditPage))
	mux.HandleFunc("POST /configuration-items", cfg.middlewareAuthPage(cfg.handleConfigurationItemsPostPage))

	// - Companies

	mux.HandleFunc("POST /v1/companies", cfg.middlewareAuth(cfg.handleCompanies))
	mux.HandleFunc("GET /companies", cfg.middlewareAuthPage(cfg.handleViewCompanies))

	// - Users
	mux.HandleFunc("GET /users", cfg.middlewareAuthPage(cfg.handleViewUsers))

	// Login Endpoints

	mux.HandleFunc("GET /", cfg.middlewareAuthPageNoRedirect(cfg.handlePageIndex))
	mux.HandleFunc("POST /v1/login", cfg.handleLogin)
	mux.HandleFunc("GET /login", cfg.handleLoginPage)
	mux.HandleFunc("GET /logout", cfg.handleLogout)
	mux.HandleFunc("GET /get", cfg.getCookieHandler)

	return mux
}
