package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) handleConfigurationItems(w http.ResponseWriter, r *http.Request, u database.User) {
	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error decoding parameters")
		return
	}
	if params.Name == "" {
		respondWithError(w, http.StatusInternalServerError, "missing configuration item name")
		return
	}

	configurationItem, err := cfg.DB.CreateConfigurationItem(r.Context(), database.CreateConfigurationItemParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Name:           params.Name,
		OrganizationID: organization.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating configuration item")
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseConfigurationItemToConfigurationItem(configurationItem))
}

func (cfg *ApiConfig) getConfigurationItems(w http.ResponseWriter, r *http.Request, u database.User) {
	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}

	configurationItems, err := cfg.DB.GetConfigurationItemsByOrganizationID(r.Context(), organization.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration items")
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseConfigurationItemsToConfigurationItems(configurationItems))
}