package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/barturba/ticket-tracker/validator"
	"github.com/barturba/ticket-tracker/views"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	fromProtected := false
	menuItems := models.MenuItems{
		{
			Name: "Login",
			Link: "/login",
		},
	}

	page := NewPage("Login", cfg, models.User{}, models.Alert{})
	page.MenuItems = menuItems

	lIndexNew := views.LoginIndexNew(fromProtected, cfg.Logo)
	login := views.BuildLayout(page, lIndexNew)
	templ.Handler(login).ServeHTTP(w, r)
}

// Configuration Items

func (cfg *ApiConfig) handleViewConfigurationItemsSelect(w http.ResponseWriter, r *http.Request, u models.User) {
	companyID := r.URL.Query().Get("company_id")

	companyUUID, err := uuid.Parse(companyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid 'company_id' parameter")
		return
	}

	v := validator.New()
	v.Check(companyUUID != uuid.Nil, "company_id", "must be provided")

	if !v.Valid() {
		cfg.failedValidationResponse(w, r, v.Errors)
		return
	}

	databaseConfigurationItems, err := cfg.DB.GetConfigurationItemsByCompanyID(r.Context(), companyUUID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}
	configurationItems := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)

	cISelect := views.ConfigurationItems(configurationItems)
	templ.Handler(cISelect).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleViewConfigurationItems(w http.ResponseWriter, r *http.Request, u models.User) {

	cis, err := cfg.GetCIs(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cIIndex := views.ConfigurationItemsIndex(cis)
	page := NewPage("Configuration Items List", cfg, u, models.Alert{})
	cIList := views.BuildLayout(page, cIIndex)

	templ.Handler(cIList).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleConfigurationItemsEditPage(w http.ResponseWriter, r *http.Request, u models.User) {

	idString := r.PathValue("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}

	ci, err := cfg.GetCIByID(r, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var companies models.SelectOptions
	err = cfg.GetCompaniesSelection(r, &companies)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	page := NewPage("Configuration Items - Edit", cfg, u, models.Alert{})

	cEIndexNew := views.ConfigurationItemFormNew(companies, ci)
	cEdit := views.BuildLayout(page, cEIndexNew)
	templ.Handler(cEdit).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleConfigurationItemsPostPage(w http.ResponseWriter, r *http.Request, u models.User) {

	var input struct {
		Name                string    `json:"name"`
		CompanyID           uuid.UUID `json:"company_id"`
		ConfigurationItemID uuid.UUID `json:"configuration_item_id"`
	}

	err := cfg.readJSON(w, r, &input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	v := validator.New()
	v.Check(input.ConfigurationItemID != uuid.UUID{}, "configuration_item_id", "must be provided")
	v.Check(input.CompanyID != uuid.UUID{}, "company_id", "must be provided")

	if !v.Valid() {
		respondToFailedValidation(w, r, v.Errors)
		return
	}

	_, err = cfg.DB.CreateConfigurationItem(r.Context(), database.CreateConfigurationItemParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      input.Name,
		CompanyID: input.CompanyID,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create configuration item")
		return
	}

	w.Header().Set("HX-Location", "/configuration-items")
}

// Incidents

func (cfg *ApiConfig) handleViewIncidents(w http.ResponseWriter, r *http.Request, u models.User) {

	incidents, err := cfg.GetIncidents(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	iIndex := views.IncidentsIndex(incidents)

	page := NewPage("Incidents List", cfg, u, models.Alert{})

	iList := views.BuildLayout(page, iIndex)
	templ.Handler(iList).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleSearchIncidents(w http.ResponseWriter, r *http.Request, u models.User) {

	var err error
	search := r.URL.Query().Get("search")
	limitString := r.URL.Query().Get("limit")
	offsetString := r.URL.Query().Get("offset")

	limit := 0
	offset := 0

	v := validator.New()
	v.Check(limitString != "", "limit_string", "must be provided")
	v.Check(offsetString != "", "offset_string", "must be provided")
	if !v.Valid() {
		respondToFailedValidation(w, r, v.Errors)
		return
	}

	if limit, err = strconv.Atoi(limitString); err != nil {
		respondWithError(w, http.StatusInternalServerError, "the 'limit' parameter is not a number")
		return
	}
	if offset, err = strconv.Atoi(offsetString); err != nil {
		respondWithError(w, http.StatusInternalServerError, "the 'offset' parameter is not a number")
		return
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

func (cfg *ApiConfig) handleIncidentsEditPage(w http.ResponseWriter, r *http.Request, u models.User) {
	idString := r.PathValue("id")

	incidentId, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}

	incident, err := cfg.GetIncidentByID(r, incidentId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	path := fmt.Sprintf("/incidents/%s", incident.ID)
	// alert := models.NewAlert("Incident Updated", models.AlertEnumSuccess, "green")
	layout, err := cfg.BuildIncidentsPage(r, "PUT", "Incidents - Edit", incident, u, path, models.Alert{}, nil)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	templ.Handler(layout).ServeHTTP(w, r)
	return
}

func (cfg *ApiConfig) handleIncidentsPostPage(w http.ResponseWriter, r *http.Request, u models.User) {
	input := models.IncidentInput

	err := cfg.readJSON(w, r, &input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	incident := NewIncident(input.ID, input.CompanyID, input.ConfigurationItemID, input.AssignedToID, input.ShortDescription, input.Description, input.State)

	// here is where you send back a modified version of the form with the
	// invalid fields highlighted

	errs := models.CheckIncident(incident)
	if errs != nil {
		path := fmt.Sprintf("/incidents/add")
		alert := models.NewAlert("Couldn't add incident", models.AlertEnumError, "red")
		layout, err := cfg.BuildIncidentsPage(r, "POST", "Incidents - Add", incident, u, path, alert, errs)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		templ.Handler(layout).ServeHTTP(w, r)
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

	path := fmt.Sprintf("/incidents/%s/edit", incident.ID)
	alert := models.NewAlert("Incident added successfully!", models.AlertEnumSuccess, "green")
	layout, err := cfg.BuildIncidentsPage(r, "PUT", "Incidents - Edit", incident, u, path, alert, errs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("HX-Redirect", path)
	http.Redirect(w, r, path, http.StatusOK)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsAddPage(w http.ResponseWriter, r *http.Request, u models.User) {

	incident := NewIncidentEmpty()

	var companies models.SelectOptions
	err := cfg.GetCompaniesSelection(r, &companies)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var cis models.SelectOptions
	err = cfg.GetCISelection(r, &cis)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var users models.SelectOptions
	err = cfg.GetUsersSelection(r, &users)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var states models.SelectOptions
	err = cfg.GetStatesSelection(r, &states)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	errs := models.CheckIncident(incident)
	fields := MakeIncidentFields(incident, companies, cis, states, users, errs)

	formData := models.NewFormData()
	path := "/incidents"
	cancelPath := "/incidents"
	form := models.NewIncidentForm("POST", path, cancelPath, companies, cis, states, users, incident, formData)

	index := views.NewIncidentForm(form, fields)
	page := NewPage("Incidents - Add", cfg, u, models.Alert{})
	layout := views.BuildLayout(page, index)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsPutPage(w http.ResponseWriter, r *http.Request, u models.User) {
	input := models.IncidentInput

	err := cfg.readJSON(w, r, &input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	incident := NewIncident(input.ID, input.CompanyID, input.ConfigurationItemID, input.AssignedToID, input.ShortDescription, input.Description, input.State)

	errs := models.CheckIncident(incident)
	if errs != nil {
		respondToFailedValidation(w, r, map[string]string{"error": err.Error()})
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

	path := fmt.Sprintf("/incidents/%s/edit", incident.ID)
	alert := models.NewAlert("Incident Updated", models.AlertEnumSuccess, "green")
	layout, err := cfg.BuildIncidentsPage(r, "PUT", "Incidents - Edit", incident, u, path, alert, errs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// w.Header().Set("HX-Redirect", path)
	// http.Redirect(w, r, path, http.StatusOK)
	templ.Handler(layout).ServeHTTP(w, r)
}

// Companies

func (cfg *ApiConfig) handleViewCompanies(w http.ResponseWriter, r *http.Request, u models.User) {

	companies, err := cfg.GetCompanies(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	page := NewPage("Companies List", cfg, u, models.Alert{})

	cIndex := views.CompaniesIndex(companies)
	cList := views.BuildLayout(page, cIndex)
	templ.Handler(cList).ServeHTTP(w, r)
}

// Users
func (cfg *ApiConfig) handleViewUsers(w http.ResponseWriter, r *http.Request, u models.User) {

	users, err := cfg.GetUsers(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	page := NewPage("Users List", cfg, u, models.Alert{})

	cIndex := views.UsersIndex(users)
	cList := views.BuildLayout(page, cIndex)
	templ.Handler(cList).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handlePageIndex(w http.ResponseWriter, r *http.Request, u models.User) {
	fromProtected := false
	if (u != models.User{}) {
		fromProtected = true
	}

	page := NewPage("TicketTracker", cfg, u, models.Alert{})

	hindex := views.HomeIndex(fromProtected)
	home := views.BuildLayout(page, hindex)
	templ.Handler(home).ServeHTTP(w, r)
}
