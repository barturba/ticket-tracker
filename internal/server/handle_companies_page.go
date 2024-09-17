package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleCompaniesPage(w http.ResponseWriter, r *http.Request, u database.User) {

	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}

	databaseCompanies, err := cfg.DB.GetCompaniesByOrganizationID(r.Context(), organization.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find companies")
		return
	}

	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)

	templ.Handler(views.Companies(companies)).ServeHTTP(w, r)
}
