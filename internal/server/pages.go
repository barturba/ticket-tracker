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

	page := NewPage("Login", cfg, models.User{})
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
	page := NewPage("Configuration Items List", cfg, u)
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

	var selectOptionsCompany models.SelectOptions
	err = cfg.GetCompaniesSelection(r, &selectOptionsCompany)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	page := NewPage("Configuration Items - Edit", cfg, u)

	cEIndexNew := views.ConfigurationItemFormNew(selectOptionsCompany, ci)
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

	page := NewPage("Incidents List", cfg, u)

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

	databaseIncident, err := cfg.DB.GetIncidentByID(r.Context(), incidentId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't find incident")
		return
	}
	incident := models.DatabaseIncidentToIncident(databaseIncident)

	companies, err := cfg.GetCompanies(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var selectOptionsCompany models.SelectOptions
	err = cfg.GetCompaniesSelection(r, &selectOptionsCompany)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	cis, err := cfg.GetCIsByCompanyID(r, companies[0].ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	users, err := cfg.GetUsers(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	selectOptionsCI := models.SelectOptions{}
	for _, ci := range cis {
		selectOptionsCI = append(selectOptionsCI, models.NewSelectOption(ci.Name, ci.ID.String()))
	}

	selectOptionsState := models.SelectOptions{}
	for _, so := range models.StateOptionsEnum {
		selectOptionsState = append(selectOptionsState, models.NewSelectOption(string(so), string(so)))
	}

	assignedToOptions := models.SelectOptions{}
	for _, user := range users {
		assignedToOptions = append(assignedToOptions, models.NewSelectOption(user.Name, user.ID.String()))
	}

	fields := MakeIncidentFields(incident, selectOptionsCompany, selectOptionsCI, selectOptionsState, assignedToOptions)

	formData := models.NewFormData()
	path := fmt.Sprintf("/incidents/%s", incident.ID)
	form := models.NewIncidentForm("PUT", path, selectOptionsCompany, selectOptionsCI, selectOptionsState, assignedToOptions, incident, formData)

	index := views.NewIncidentForm(form, fields)
	page := NewPage("Incidents - Edit", cfg, u)
	layout := views.BuildLayout(page, index)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsPostPage(w http.ResponseWriter, r *http.Request, u models.User) {
	var input struct {
		ID                  uuid.UUID          `json:"id"`
		ShortDescription    string             `json:"short_description"`
		Description         string             `json:"description"`
		CompanyID           uuid.UUID          `json:"company_id"`
		ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
		State               database.StateEnum `json:"state"`
	}
	err := cfg.readJSON(w, r, &input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	incident := NewIncident(input.ID, input.CompanyID, input.ConfigurationItemID, input.ShortDescription, input.Description, input.State)

	err = CheckIncident(incident)

	if err != nil {
		respondToFailedValidation(w, r, map[string]string{"error": err.Error()})
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

	var selectOptionsCompany models.SelectOptions
	err = cfg.GetCompaniesSelection(r, &selectOptionsCompany)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	companies, err := cfg.GetCompanies(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	databaseConfigurationItems, err := cfg.DB.GetConfigurationItemsByCompanyID(r.Context(), companies[0].ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration items")
		return
	}
	configurationItems := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)

	users, err := cfg.GetUsers(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	selectOptionsCI := models.SelectOptions{}
	for _, ci := range configurationItems {
		selectOptionsCI = append(selectOptionsCI, models.NewSelectOption(ci.Name, ci.ID.String()))
	}

	selectOptionsState := models.SelectOptions{}
	for _, so := range models.StateOptionsEnum {
		selectOptionsState = append(selectOptionsState, models.NewSelectOption(string(so), string(so)))
	}

	assignedToOptions := models.SelectOptions{}
	for _, user := range users {
		assignedToOptions = append(assignedToOptions, models.NewSelectOption(user.Name, user.ID.String()))
	}

	fields := MakeIncidentFields(incident, selectOptionsCompany, selectOptionsCI, selectOptionsState, assignedToOptions)

	formData := models.NewFormData()
	path := fmt.Sprintf("/incidents/%s", incident.ID)
	form := models.NewIncidentForm("PUT", path, selectOptionsCompany, selectOptionsCI, selectOptionsState, assignedToOptions, incident, formData)

	index := views.NewIncidentForm(form, fields)
	page := NewPage("Incidents - Edit", cfg, u)
	layout := views.BuildLayout(page, index)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsAddPage(w http.ResponseWriter, r *http.Request, u models.User) {

	incident := NewIncidentEmpty()

	var selectOptionsCompany models.SelectOptions
	err := cfg.GetCompaniesSelection(r, &selectOptionsCompany)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var selectOptionsState models.SelectOptions
	cfg.GetStateSelection(&selectOptionsState)

	companies, err := cfg.GetCompanies(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	databaseConfigurationItems, err := cfg.DB.GetConfigurationItemsByCompanyID(r.Context(), companies[0].ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find configuration items")
		return
	}
	configurationItems := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)
	selectOptionsCI := models.SelectOptions{}
	for _, ci := range configurationItems {
		selectOptionsCI = append(selectOptionsCI, models.NewSelectOption(ci.Name, ci.ID.String()))
	}
	stateOptions := models.SelectOptions{}
	for _, so := range models.StateOptionsEnum {
		stateOptions = append(stateOptions, models.NewSelectOption(string(so), string(so)))
	}

	users, err := cfg.GetUsers(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	assignedToOptions := models.SelectOptions{}
	for _, user := range users {
		assignedToOptions = append(assignedToOptions, models.NewSelectOption(user.Name, user.ID.String()))
	}

	fields := MakeIncidentFields(incident, selectOptionsCompany, selectOptionsCI, selectOptionsState, assignedToOptions)

	formData := models.NewFormData()
	path := "/incidents"
	form := models.NewIncidentForm("POST", path, selectOptionsCompany, selectOptionsCI, stateOptions, assignedToOptions, incident, formData)

	index := views.NewIncidentForm(form, fields)
	page := NewPage("Incidents - Add", cfg, u)
	layout := views.BuildLayout(page, index)
	templ.Handler(layout).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsPutPage(w http.ResponseWriter, r *http.Request, u models.User) {
	var input struct {
		ShortDescription    string             `json:"short_description"`
		Description         string             `json:"description"`
		CompanyID           uuid.UUID          `json:"company_id"`
		ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
		State               database.StateEnum `json:"state"`
	}

	err := cfg.readJSON(w, r, &input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	v := validator.New()
	v.Check(input.ConfigurationItemID != uuid.UUID{}, "configuration_item_id", "must be provided")
	v.Check(input.CompanyID != uuid.UUID{}, "company_id", "must be provided")
	v.Check(input.Description != "", "description", "must be provided")
	v.Check(input.ShortDescription != "", "short_description", "must be provided")

	if !v.Valid() {
		respondToFailedValidation(w, r, v.Errors)
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
		Description:      sql.NullString{String: input.Description, Valid: input.Description != ""},
		ShortDescription: input.ShortDescription,
		State:            input.State,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't update incident")
		return
	}

	w.Header().Set("HX-Redirect", "/incidents")
	http.Redirect(w, r, "/incidents", http.StatusOK)
}

// Companies

func (cfg *ApiConfig) handleViewCompanies(w http.ResponseWriter, r *http.Request, u models.User) {

	companies, err := cfg.GetCompanies(r)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	page := NewPage("Companies List", cfg, u)

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

	page := NewPage("Users List", cfg, u)

	cIndex := views.UsersIndex(users)
	cList := views.BuildLayout(page, cIndex)
	templ.Handler(cList).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handlePageIndex(w http.ResponseWriter, r *http.Request, u models.User) {
	fromProtected := false
	if (u != models.User{}) {
		fromProtected = true
	}

	page := NewPage("TicketTracker", cfg, u)

	hindex := views.HomeIndex(fromProtected)
	home := views.BuildLayout(page, hindex)
	templ.Handler(home).ServeHTTP(w, r)
}
