package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleCompaniesPage(w http.ResponseWriter, r *http.Request, u database.User) {

	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find companies")
		return
	}

	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)

	templ.Handler(views.Companies(companies)).ServeHTTP(w, r)
}
