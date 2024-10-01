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

	databaseConfigurationItems, err := cfg.DB.GetConfigurationItems(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}
	cis := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)

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

	databaseConfigurationItem, err := cfg.DB.GetConfigurationItemByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't find incident")
		return
	}
	ci := models.DatabaseConfigurationItemToConfigurationItem(databaseConfigurationItem)

	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find companies")
		return
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)
	selectOptionsCompany := models.SelectOptions{}

	for _, company := range companies {
		selectOptionsCompany = append(selectOptionsCompany, models.NewSelectOption(company.Name, company.ID.String()))
	}

	page := NewPage("Configuration Items - Edit", cfg, u)

	cEIndexNew := views.ConfigurationItemFormNew(selectOptionsCompany, ci)
	cEdit := views.BuildLayout(page, cEIndexNew)
	templ.Handler(cEdit).ServeHTTP(w, r)
}
func (cfg *ApiConfig) handleConfigurationItemsPostPage(w http.ResponseWriter, r *http.Request, u models.User) {

	type parameters struct {
		Name                string    `json:"name"`
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

	_, err = cfg.DB.CreateConfigurationItem(r.Context(), database.CreateConfigurationItemParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		CompanyID: params.CompanyID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create incident")
		return
	}

	w.Header().Set("HX-Location", "/configuration-items")
}

// Incidents

func (cfg *ApiConfig) handleViewIncidents(w http.ResponseWriter, r *http.Request, u models.User) {

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

	page := NewPage("Incidents List", cfg, u)

	iList := views.BuildLayout(page, iIndex)
	templ.Handler(iList).ServeHTTP(w, r)
}
func (cfg *ApiConfig) handleSearchIncidents(w http.ResponseWriter, r *http.Request, u models.User) {

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
func (cfg *ApiConfig) handleIncidentsEditPage(w http.ResponseWriter, r *http.Request, u models.User) {

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
	selectOptionsCompany := models.SelectOptions{}

	for _, company := range companies {
		selectOptionsCompany = append(selectOptionsCompany, models.NewSelectOption(company.Name, company.ID.String()))
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

	formData := models.NewFormData()
	iEPath := fmt.Sprintf("/incidents/%s", incident.ID)
	iEIndexNew := views.IncidentForm("PUT", iEPath, selectOptionsCompany, selectOptionsCI, stateOptions, incident, formData)

	page := NewPage("Incidents - Edit", cfg, u)

	iEdit := views.BuildLayout(page, iEIndexNew)
	templ.Handler(iEdit).ServeHTTP(w, r)
}
func (cfg *ApiConfig) handleIncidentsPostPage(w http.ResponseWriter, r *http.Request, u models.User) {
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

	databaseIncident, err := cfg.DB.CreateIncident(r.Context(), database.CreateIncidentParams{
		ID:                  uuid.New(),
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		ShortDescription:    input.ShortDescription,
		Description:         sql.NullString{String: input.Description, Valid: input.Description != ""},
		State:               input.State,
		ConfigurationItemID: input.ConfigurationItemID,
		CompanyID:           input.CompanyID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create incident")
		return
	}
	incident := models.DatabaseIncidentToIncident(databaseIncident)

	// Add helpers for these functions:
	var selectOptionsCompany models.SelectOptions
	err = cfg.GetCompaniesSelection(r, &selectOptionsCompany)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find companies")
		return
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)
	// selectOptionsCompany := models.SelectOptions{}

	// for _, company := range companies {
	// 	selectOptionsCompany = append(selectOptionsCompany, models.NewSelectOption(company.Name, company.ID.String()))
	// }

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

	// w.Header().Set("HX-Location", "/incidents")
	formData := models.NewFormData()
	iEPath := fmt.Sprintf("/incidents/%s/edit", incident.ID)
	iEIndexNew := views.IncidentForm("PUT", iEPath, selectOptionsCompany, selectOptionsCI, stateOptions, incident, formData)

	page := NewPage("Incidents - Edit", cfg, u)

	iEdit := views.BuildLayout(page, iEIndexNew)
	templ.Handler(iEdit).ServeHTTP(w, r)
}
func (cfg *ApiConfig) handleIncidentsAddPage(w http.ResponseWriter, r *http.Request, u models.User) {

	incident := NewIncident()

	var selectOptionsCompany models.SelectOptions
	err := cfg.GetCompaniesSelection(r, &selectOptionsCompany)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var selectOptionsState models.SelectOptions
	cfg.GetStateSelection(&selectOptionsState)

	selectOptionsCI := models.SelectOptions{}

	formData := models.NewFormData()
	formData.Errors["company_id"] = "Wrong company ID"
	iIPath := "/incidents"
	iIndexNew := views.IncidentForm("POST", iIPath, selectOptionsCompany, selectOptionsCI, selectOptionsState, incident, formData)

	page := NewPage("Configuration Items List", cfg, u)

	iNew := views.BuildLayout(page, iIndexNew)
	templ.Handler(iNew).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handleIncidentsPutPage(w http.ResponseWriter, r *http.Request, u models.User) {

	type parameters struct {
		ShortDescription    string             `json:"short_description"`
		Description         string             `json:"description"`
		CompanyID           uuid.UUID          `json:"company_id"`
		ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
		State               database.StateEnum `json:"state"`
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
		State:            params.State,
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

	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find comanies")
		return
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)

	page := NewPage("Companies List", cfg, u)

	cIndex := views.CompaniesIndex(companies)
	cList := views.BuildLayout(page, cIndex)
	templ.Handler(cList).ServeHTTP(w, r)
}

// Users
func (cfg *ApiConfig) handleViewUsers(w http.ResponseWriter, r *http.Request, u models.User) {

	databaseUsers, err := cfg.DB.GetUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find comanies")
		return
	}
	users := models.DatabaseUsersToUsers(databaseUsers)

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
