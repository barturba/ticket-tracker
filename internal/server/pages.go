package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/google/uuid"
)

// Incidents

func (cfg *ApiConfig) handleIncidentsGet(w http.ResponseWriter, r *http.Request) {

	i, err := cfg.GetIncidents(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get incidents")
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *ApiConfig) handleIncidentByIdGet(w http.ResponseWriter, r *http.Request) {
	var err error

	myUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(myUrl.RawQuery)

	idString := params.Get("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	i, err := cfg.GetIncidentByID(r, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get incidents")
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *ApiConfig) handleFilteredIncidentsGet(w http.ResponseWriter, r *http.Request) {
	var err error

	limit := 0
	offset := 0

	myUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(myUrl.RawQuery)

	query := params.Get("query")

	limitInput := params.Get("limit")

	offsetInput := params.Get("offset")

	// v := validator.New()
	// v.Check(query != "", "query", "must be provided")
	// v.Check(limitInput != "", "limit", "must be provided")
	// v.Check(offsetInput != "", "offset", "must be provided")
	// if !v.Valid() {
	// 	respondToFailedValidation(w, r, v.Errors)
	// 	return
	// }

	if limit, err = strconv.Atoi(limitInput); err != nil {
		respondWithError(w, http.StatusInternalServerError, "the 'limit' parameter is not a number")
		return
	}
	if offset, err = strconv.Atoi(offsetInput); err != nil {
		respondWithError(w, http.StatusInternalServerError, "the 'offset' parameter is not a number")
		return
	}

	i, err := cfg.GetIncidentsFiltered(r, query, limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get incidents")
		return
	}
	log.Printf("handleFilteredIncidentsGet: returning this data: %v\n", i)
	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *ApiConfig) handleFilteredIncidentsCountGet(w http.ResponseWriter, r *http.Request) {
	var err error

	myUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(myUrl.RawQuery)

	query := params.Get("query")

	// v := validator.New()
	// v.Check(query != "", "query", "must be provided")
	// if !v.Valid() {
	// 	respondToFailedValidation(w, r, v.Errors)
	// 	return
	// }

	i, err := cfg.GetIncidentsFilteredCount(r, query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get incidents")
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *ApiConfig) handleIncidentsLatestGet(w http.ResponseWriter, r *http.Request) {
	var err error

	limit := 0
	offset := 0

	myUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(myUrl.RawQuery)

	limitInput := params.Get("limit")

	offsetInput := params.Get("offset")

	// v := validator.New()
	// v.Check(query != "", "query", "must be provided")
	// v.Check(limitInput != "", "limit", "must be provided")
	// v.Check(offsetInput != "", "offset", "must be provided")
	// if !v.Valid() {
	// 	respondToFailedValidation(w, r, v.Errors)
	// 	return
	// }

	if limit, err = strconv.Atoi(limitInput); err != nil {
		respondWithError(w, http.StatusInternalServerError, "the 'limit' parameter is not a number")
		return
	}
	if offset, err = strconv.Atoi(offsetInput); err != nil {
		respondWithError(w, http.StatusInternalServerError, "the 'offset' parameter is not a number")
		return
	}

	i, err := cfg.GetIncidentsLatest(r, limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get incidents")
		return
	}
	log.Printf("handleLatestIncidentsGet: returning this data: %v\n", i)
	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *ApiConfig) handleIncidentsPost(w http.ResponseWriter, r *http.Request) {

	input := models.NewIncidentInput
	err := cfg.readJSON(w, r, &input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	incident := NewIncident(uuid.New(), input.CompanyID, input.ConfigurationItemID, input.AssignedToID, input.ShortDescription, input.Description, input.State)

	errs := models.CheckIncident(incident)
	if errs != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", errs))
		return
	}

	databaseIncident, err := cfg.DB.CreateIncident(r.Context(), database.CreateIncidentParams{
		ID:                  incident.ID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		ShortDescription:    incident.ShortDescription,
		Description:         sql.NullString{String: incident.Description, Valid: incident.Description != ""},
		State:               incident.State,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create incident")
		return
	}
	incident = models.DatabaseIncidentToIncident(databaseIncident)

	respondWithJSON(w, http.StatusOK, incident)
}

func (cfg *ApiConfig) handleIncidentsPut(w http.ResponseWriter, r *http.Request) {

	input := models.IncidentInput

	err := cfg.readJSON(w, r, &input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	incident := NewIncident(input.ID, input.CompanyID, input.ConfigurationItemID, input.AssignedToID, input.ShortDescription, input.Description, input.State)
	log.Printf("editing incident: %v\n", incident.ID)
	log.Printf("editing incident: %v\n", incident)

	errs := models.CheckIncident(incident)
	if errs != nil {
		log.Printf("handleIncidentsPut: err is %v", err)
		respondWithError(w, http.StatusInternalServerError, "couldn't update incident")
		return
	}

	idString := r.PathValue("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}
	incident.ID = id

	incident, err = cfg.UpdateIncident(r, incident)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("handleIncidentsPut: updated the following incident in the database %v", incident)
	respondWithJSON(w, http.StatusOK, incident)
}

func (cfg *ApiConfig) handleIncidentsDelete(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}

	incident, err := cfg.DeleteIncidentByID(r, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("handleIncidentsDelete: deleted the following incident from the database %v", incident)
	respondWithJSON(w, http.StatusOK, incident)
}

// Companies

func (cfg *ApiConfig) handleCompaniesGet(w http.ResponseWriter, r *http.Request) {

	i, err := cfg.GetCompanies(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get companies")
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *ApiConfig) handleFilteredCompaniesGet(w http.ResponseWriter, r *http.Request) {
	var err error

	limit := 0
	offset := 0

	myUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(myUrl.RawQuery)

	query := params.Get("query")

	limitInput := params.Get("limit")

	offsetInput := params.Get("offset")

	// v := validator.New()
	// v.Check(query != "", "query", "must be provided")
	// v.Check(limitInput != "", "limit", "must be provided")
	// v.Check(offsetInput != "", "offset", "must be provided")
	// if !v.Valid() {
	// 	respondToFailedValidation(w, r, v.Errors)
	// 	return
	// }

	if limit, err = strconv.Atoi(limitInput); err != nil {
		respondWithError(w, http.StatusInternalServerError, "the 'limit' parameter is not a number")
		return
	}
	if offset, err = strconv.Atoi(offsetInput); err != nil {
		respondWithError(w, http.StatusInternalServerError, "the 'offset' parameter is not a number")
		return
	}

	i, err := cfg.GetCompaniesFiltered(r, query, limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get companies")
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *ApiConfig) handleFilteredCompaniesCountGet(w http.ResponseWriter, r *http.Request) {
	var err error

	myUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(myUrl.RawQuery)

	query := params.Get("query")

	// v := validator.New()
	// v.Check(query != "", "query", "must be provided")
	// if !v.Valid() {
	// 	respondToFailedValidation(w, r, v.Errors)
	// 	return
	// }

	i, err := cfg.GetCompaniesFilteredCount(r, query)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get companies")
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

// Configuration Items

func (cfg *ApiConfig) handleConfigurationItemsGet(w http.ResponseWriter, r *http.Request) {

	i, err := cfg.GetConfigurationItems(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get configuration items")
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

// Users

func (cfg *ApiConfig) handleUsersGet(w http.ResponseWriter, r *http.Request) {

	i, err := cfg.GetUsers(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get companies")
		return
	}

	respondWithJSON(w, http.StatusOK, i)
}

func (cfg *ApiConfig) handleUsersByCompanyGet(w http.ResponseWriter, r *http.Request) {
	log.Println("called handleUsersByCompanyGet")
	var err error

	limit := 0
	offset := 0

	myUrl, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(myUrl.RawQuery)

	query := params.Get("query")

	// limitInput := params.Get("limit")

	// offsetInput := params.Get("offset")

	// v := validator.New()
	// v.Check(query != "", "query", "must be provided")
	// v.Check(limitInput != "", "limit", "must be provided")
	// v.Check(offsetInput != "", "offset", "must be provided")
	// if !v.Valid() {
	// 	respondToFailedValidation(w, r, v.Errors)
	// 	return
	// }

	// if limit, err = strconv.Atoi(limitInput); err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "the 'limit' parameter is not a number")
	// 	return
	// }
	// if offset, err = strconv.Atoi(offsetInput); err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "the 'offset' parameter is not a number")
	// 	return
	// }

	i, err := cfg.GetUsersByCompany(r, query, limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get incidents")
		return
	}
	log.Printf("handleLatestIncidentsGet: returning this data: %v\n", i)
	respondWithJSON(w, http.StatusOK, i)
}
