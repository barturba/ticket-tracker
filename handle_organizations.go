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
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error decoding parameters")
		return
	}
	if params.Name == "" {
		respondWithError(w, http.StatusInternalServerError, "missing organization name")
	}

	organization, err := cfg.DB.UpdateOrganizationByUserID(r.Context(),
		database.UpdateOrganizationByUserIDParams{
			UserID:    u.ID,
			UpdatedAt: time.Now(),
			Name:      params.Name,
		})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't update organization")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseOrganizationToOrganization(organization))
}
