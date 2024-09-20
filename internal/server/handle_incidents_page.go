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

func (cfg *ApiConfig) handleViewIncidents(w http.ResponseWriter, r *http.Request, u database.User) {
	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}

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
	iIndex := views.IncidentsIndex(incidents)
	iList := views.IncidentsList2(" | Incidents List",
		fromProtected,
		false,
		"",
		u.Name,
		iIndex)
	templ.Handler(iList).ServeHTTP(w, r)
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

func (cfg *ApiConfig) handleIncidentsPostPage(w http.ResponseWriter, r *http.Request, u database.User) {

	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}

	type parameters struct {
		ShortDescription    string    `json:"short_description"`
		Description         string    `json:"description"`
		CompanyID           uuid.UUID `json:"company_id"`
		ConfigurationItemID uuid.UUID `json:"configuration_item_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}
	if (params.ConfigurationItemID == uuid.UUID{}) {
		respondWithError(w, http.StatusInternalServerError, "configuration_item_id can't be blank")
		return
	}

	if (params.CompanyID == uuid.UUID{}) {
		respondWithError(w, http.StatusInternalServerError, "company_id can't be blank")
		return
	}

	if params.Description == "" {
		respondWithError(w, http.StatusInternalServerError, "description can't be blank")
		return
	}

	if params.ShortDescription == "" {
		respondWithError(w, http.StatusInternalServerError, "short_description can't be blank")
		return
	}
	_, err = cfg.DB.CreateIncident(r.Context(), database.CreateIncidentParams{
		ID:                  uuid.New(),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		ShortDescription:    params.ShortDescription,
		Description:         sql.NullString{String: params.Description, Valid: params.Description != ""},
		State:               "New",
		OrganizationID:      organization.ID,
		ConfigurationItemID: params.ConfigurationItemID,
		CompanyID:           params.CompanyID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create incident")
		return
	}

	w.Header().Set("HX-Redirect", "/incidents")
	http.Redirect(w, r, "/incidents", http.StatusFound)
}
