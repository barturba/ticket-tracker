package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/barturba/ticket-tracker/views"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) handleIncidentsPage(w http.ResponseWriter, r *http.Request, u database.User) {
	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}

	databaseIncidents, err := cfg.DB.GetIncidentsByOrganizationID(r.Context(), organization.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}

	incidents := models.DatabaseGetIncidentsByOrganizationIDRowToIncidents(databaseIncidents)
	incidentsComponent := views.IncidentsList(incidents)
	incidentsButton := views.Button("New", "/incidents/new")

	templ.Handler(views.ContentPage("Incidents", "incidents", incidentsComponent, incidentsButton, true)).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsEditPage(w http.ResponseWriter, r *http.Request, u database.User) {
	idString := r.PathValue("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}

	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}

	databaseIncident, err := cfg.DB.GetIncidentByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't find incident")
		return
	}

	incident := models.DatabaseIncidentToIncident(databaseIncident)
	databaseCompanies, err := cfg.DB.GetCompaniesByOrganizationID(r.Context(), organization.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find companies")
		return
	}

	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)
	incidentsComponent := views.IncidentsEditPage(incident, companies)

	templ.Handler(views.ContentPage("Edit Incident", "edit-incident", incidentsComponent, nil, true)).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsGetPage(w http.ResponseWriter, r *http.Request, u database.User) {
	idString := r.PathValue("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}

	databaseIncident, err := cfg.DB.GetIncidentByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't find incident")
		return
	}

	incident := models.DatabaseIncidentToIncident(databaseIncident)

	templ.Handler(views.IncidentRow(incident)).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsPutPage(w http.ResponseWriter, r *http.Request, u database.User) {

	type parameters struct {
		ShortDescription string `json:"short_description"`
		Description      string `json:"description"`
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

	idString := r.PathValue("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}

	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}

	updatedIncident, err := cfg.DB.UpdateIncident(r.Context(), database.UpdateIncidentParams{
		ID:               id,
		UpdatedAt:        time.Now(),
		Description:      sql.NullString{String: params.Description, Valid: params.Description != ""},
		ShortDescription: params.ShortDescription,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't update incident")
		return
	}

	incident := models.DatabaseIncidentToIncident(updatedIncident)
	databaseCompanies, err := cfg.DB.GetCompaniesByOrganizationID(r.Context(), organization.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find companies")
		return
	}

	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)
	incidentsComponent := views.IncidentsEditPage(incident, companies)

	templ.Handler(views.ContentPage("Edit Incident", "edit-incident", incidentsComponent, nil, true)).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsNewPage(w http.ResponseWriter, r *http.Request, u database.User) {

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
	firstCompany := companies[0]

	configurationItems, err := cfg.DB.GetConfigurationItemsByCompanyID(r.Context(), firstCompany.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration items")
		return
	}

	incidentNewPage := views.IncidentFormNew(companies, models.DatabaseConfigurationItemsToConfigurationItems(configurationItems))

	templ.Handler(views.ContentPage("New Incident", "", incidentNewPage, nil, true)).ServeHTTP(w, r)
}
