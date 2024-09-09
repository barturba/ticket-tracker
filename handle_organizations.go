package main

import (
	"encoding/json"
	"net/http"
	"time"

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

func (cfg *apiConfig) updateOrganizations(w http.ResponseWriter, r *http.Request, u database.User) {
	// organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
	// 	return
	// }
	type parameters struct {
		OrganizationName string `json:"organization_name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error decoding parameters")
		return
	}
	if params.OrganizationName == "" {
		respondWithError(w, http.StatusInternalServerError, "missing organization name")
	}

	organization, err := cfg.DB.UpdateOrganizationByUserID(r.Context(),
		database.UpdateOrganizationByUserIDParams{
			UserID:    u.ID,
			UpdatedAt: time.Now(),
			Name:      params.OrganizationName,
		})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't update organization")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseOrganizationToOrganization(organization))
}
