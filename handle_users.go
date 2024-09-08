package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleUsers(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name             string `json:"name"`
		OrganizationName string `json:"organization_name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	organization, err := cfg.DB.CreateOrganization(r.Context(),
		database.CreateOrganizationParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      params.OrganizationName,
		},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create organization")
	}

	user, err := cfg.DB.CreateUser(r.Context(),
		database.CreateUserParams{
			ID:             uuid.New(),
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			Name:           params.Name,
			OrganizationID: organization.ID,
		})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}
