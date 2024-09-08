package main

import (
	"net/http"

	"github.com/barturba/ticket-tracker/internal/database"
)

func (cfg *apiConfig) getOrganizations(w http.ResponseWriter, r *http.Request, u database.User) {
	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseOrganizationToOrganization(organization))
}
