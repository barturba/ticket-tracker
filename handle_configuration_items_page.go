package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/views"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleConfigurationItemsPage(w http.ResponseWriter, r *http.Request, u database.User) {

	companyID := r.URL.Query().Get("company_id")
	if companyID == "" {
		respondWithError(w, http.StatusInternalServerError, "company_id can't be blank")
		return
	}

	id, err := uuid.Parse(companyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}

	configurationItems, err := cfg.DB.GetConfigurationItemsByCompanyID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration items")
		return
	}
	templ.Handler(views.ConfigurationItems(databaseConfigurationItemsToConfigurationItems(configurationItems))).ServeHTTP(w, r)
}
