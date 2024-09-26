package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	databaseIncidents, err := cfg.DB.GetIncidents(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}
	incidents := models.DatabaseIncidentsRowToIncidents(databaseIncidents)

	for n, i := range incidents {
		ci, err := cfg.DB.GetConfigurationItemByID(r.Context(), i.ConfigurationItemID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "couldn't find configuration item name")
			return
		}
		incidents[n].ConfigurationItemName = ci.Name

	}

	iIndex := views.IncidentsIndex(incidents)
	iList := views.IncidentsList("Incidents List",
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.MenuItems,
		cfg.ProfileItems,
		iIndex)
	templ.Handler(iList).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleSearchIncidents(w http.ResponseWriter, r *http.Request, u database.User) {

	var err error
	search := r.URL.Query().Get("search")

	limitString := r.URL.Query().Get("limit")
	log.Println("limitString: ", limitString)
	limit := 0
	if limitString != "" {
		if limit, err = strconv.Atoi(limitString); err != nil {
			respondWithError(w, http.StatusInternalServerError, "the 'limit' parameter is not a number")
			return
		}
	}

	offsetString := r.URL.Query().Get("offset")
	log.Println("offsetString: ", offsetString)
	offset := 0
	if offsetString != "" {
		if offset, err = strconv.Atoi(offsetString); err != nil {
			respondWithError(w, http.StatusInternalServerError, "the 'offset' parameter is not a number")
			return
		}
	}

	databaseIncidents, err := cfg.DB.GetIncidentsBySearchTermLimitOffset(r.Context(), database.GetIncidentsBySearchTermLimitOffsetParams{
		ShortDescription: fmt.Sprintf("%%%s%%", search),
		Limit:            int32(limit),
		Offset:           int32(offset),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}
	fullCount := 0
	if len(databaseIncidents) > 0 {
		fullCount = int(databaseIncidents[0].FullCount)
	}
	incidents := models.DatabaseIncidentsBySearchTermLimitOffsetRowToIncidents(databaseIncidents)

	for n, i := range incidents {
		ci, err := cfg.DB.GetConfigurationItemByID(r.Context(), i.ConfigurationItemID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "couldn't find configuration item name")
			return
		}
		incidents[n].ConfigurationItemName = ci.Name
	}
	type response struct {
		Count   int               `json:"count"`
		Results []models.Incident `json:"results"`
	}

	respondWithJSON(w, http.StatusOK, response{
		Count:   fullCount,
		Results: incidents,
	})
}

func (cfg *ApiConfig) handleIncidentsEditPage(w http.ResponseWriter, r *http.Request, u database.User) {
	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}

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

	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find companies")
		return
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)

	databaseConfigurationItems, err := cfg.DB.GetConfigurationItemsByCompanyID(r.Context(), companies[0].ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration items")
		return
	}
	configurationItems := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)

	iEIndex := views.IncidentsEditIndex(incident, companies, configurationItems)
	iEdit := views.IncidentsEdit("Incidents - Edit",
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.MenuItems,
		cfg.ProfileItems,
		iEIndex)
	templ.Handler(iEdit).ServeHTTP(w, r)
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

	_, err = cfg.DB.UpdateIncident(r.Context(), database.UpdateIncidentParams{
		ID:               id,
		UpdatedAt:        time.Now(),
		Description:      sql.NullString{String: params.Description, Valid: params.Description != ""},
		ShortDescription: params.ShortDescription,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't update incident")
		return
	}

	w.Header().Set("HX-Redirect", "/incidents")
	http.Redirect(w, r, "/incidents", http.StatusOK)
}

func (cfg *ApiConfig) handleIncidentsNewPage(w http.ResponseWriter, r *http.Request, u database.User) {

	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}

	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find companies")
		return
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)

	databaseConfigurationItems, err := cfg.DB.GetConfigurationItemsByCompanyID(r.Context(), companies[0].ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration items")
		return
	}
	configurationItems := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)

	iNIndex := views.IncidentsNewIndex(companies, configurationItems)
	iNew := views.IncidentsEdit("Incidents - New",
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.MenuItems,
		cfg.ProfileItems,
		iNIndex)
	templ.Handler(iNew).ServeHTTP(w, r)

}

func (cfg *ApiConfig) handleIncidentsPostPage(w http.ResponseWriter, r *http.Request, u database.User) {

	type parameters struct {
		ShortDescription    string    `json:"short_description"`
		Description         string    `json:"description"`
		CompanyID           uuid.UUID `json:"company_id"`
		ConfigurationItemID uuid.UUID `json:"configuration_item_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
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
