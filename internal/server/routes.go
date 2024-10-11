package server

import "net/http"

func (cfg *ApiConfig) Routes() *http.ServeMux {

	mux := http.NewServeMux()

	// Page Endpoints

	// - Companies

	mux.HandleFunc("GET /v1/companies", cfg.handleCompaniesGet)
	mux.HandleFunc("GET /v1/filtered_companies", cfg.handleFilteredCompaniesGet)
	mux.HandleFunc("GET /v1/filtered_companies_count", cfg.handleFilteredCompaniesCountGet)
	mux.HandleFunc("GET /v1/company_by_id", cfg.handleCompanyByIdGet)
	mux.HandleFunc("POST /v1/companies", cfg.handleCompaniesPost)
	mux.HandleFunc("PUT /v1/companies/{id}", cfg.handleCompaniesPut)
	mux.HandleFunc("DELETE /v1/companies/{id}", cfg.handleCompaniesDelete)

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
