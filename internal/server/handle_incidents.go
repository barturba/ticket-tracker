package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) handleIncidents(w http.ResponseWriter, r *http.Request, u database.User) {
	type parameters struct {
		ShortDescription    string    `json:"short_description"`
		Description         string    `json:"description"`
		ConfigurationItemID uuid.UUID `json:"configuration_item_id"`
		CompanyID           uuid.UUID `json:"company_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}
	if params.ShortDescription == "" {
		respondWithError(w, http.StatusInternalServerError, "short_description can't be blank")
		return
	}
	if params.CompanyID == uuid.Nil {
		respondWithError(w, http.StatusInternalServerError, "company_id can't be blank")
		return
	}
	if params.ConfigurationItemID == uuid.Nil {
		respondWithError(w, http.StatusInternalServerError, "configuration_item_id can't be blank")
		return
	}
	company, err := cfg.DB.GetCompanyByID(r.Context(), params.CompanyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find company")
		return
	}
	configurationItem, err := cfg.DB.GetConfigurationItemByID(r.Context(), params.ConfigurationItemID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration item")
		return
	}

	incident, err := cfg.DB.CreateIncident(r.Context(), database.CreateIncidentParams{
		ID:                  uuid.New(),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		ShortDescription:    params.ShortDescription,
		Description:         sql.NullString{String: params.Description, Valid: params.Description != ""},
		State:               database.StateEnumNew,
		ConfigurationItemID: configurationItem.ID,
		CompanyID:           company.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create incident")
		return
	}
	respondWithJSON(w, http.StatusOK, models.DatabaseIncidentToIncident(incident))

}

func (cfg *ApiConfig) getIncidents(w http.ResponseWriter, r *http.Request, u database.User) {
	// Get the organization -- should this be a part of the authentication
	// process? I mean we're using the organization quite a bit.
	//

	incidents, err := cfg.DB.GetIncidents(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseIncidentsRowToIncidents(incidents))

}
