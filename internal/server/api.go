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

func (cfg *ApiConfig) handleUsers(w http.ResponseWriter, r *http.Request) {
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

	if params.Name == "" {
		respondWithError(w, http.StatusInternalServerError, "missing name")
		return
	}
	if params.OrganizationName == "" {
		respondWithError(w, http.StatusInternalServerError, "missing organization name")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      params.Name,
		})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create user")
		return
	}

	type response struct {
		User models.User `json:"user"`
	}

	resp := response{
		User: models.DatabaseUserToUser(user),
	}

	respondWithJSON(w, http.StatusCreated, resp)
}

func (cfg *ApiConfig) handleCompanies(w http.ResponseWriter, r *http.Request, u models.User) {
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
		respondWithError(w, http.StatusInternalServerError, "missing company name")
		return
	}

	company, err := cfg.DB.CreateCompany(r.Context(), database.CreateCompanyParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create company")
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseCompanyToCompany(company))

}

func (cfg *ApiConfig) handleConfigurationItems(w http.ResponseWriter, r *http.Request, u models.User) {
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
		respondWithError(w, http.StatusInternalServerError, "missing configuration item name")
		return
	}

	configurationItem, err := cfg.DB.CreateConfigurationItem(r.Context(), database.CreateConfigurationItemParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating configuration item")
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseConfigurationItemToConfigurationItem(configurationItem))
}

func (cfg *ApiConfig) getConfigurationItems(w http.ResponseWriter, r *http.Request, u models.User) {

	configurationItems, err := cfg.DB.GetConfigurationItems(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration items")
		return
	}

	respondWithJSON(w, http.StatusOK, models.DatabaseConfigurationItemsToConfigurationItems(configurationItems))
}

func (cfg *ApiConfig) handleIncidents(w http.ResponseWriter, r *http.Request, u models.User) {
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

func (cfg *ApiConfig) getIncidents(w http.ResponseWriter, r *http.Request, u models.User) {
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
