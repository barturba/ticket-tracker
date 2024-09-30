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

func (cfg *ApiConfig) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	fromProtected := false
	lIndexNew := views.LoginIndexNew(fromProtected, cfg.Logo)
	login := views.Login("Login", cfg.Logo, fromProtected, false, "msg", "", cfg.ProfilePicPlaceholder, cfg.MenuItems, cfg.ProfileItems, lIndexNew)
	templ.Handler(login).ServeHTTP(w, r)
}

// Configuration Items

func (cfg *ApiConfig) handleViewConfigurationItemsSelect(w http.ResponseWriter, r *http.Request, u database.User) {
	companyID := r.URL.Query().Get("company_id")
	if companyID == "" {
		respondWithError(w, http.StatusInternalServerError, "the 'company_id' parameter can't be blank")
		return
	}

	companyUUID, err := uuid.Parse(companyID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid 'company_id' parameter")
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
func (cfg *ApiConfig) handleViewConfigurationItems(w http.ResponseWriter, r *http.Request, u database.User) {
	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}

	databaseConfigurationItems, err := cfg.DB.GetConfigurationItems(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}
	cis := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)

	cIIndex := views.ConfigurationItemsIndex(cis)
	cIList := views.IncidentsList("Configuration Items List",
		cfg.Logo,
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.ProfilePicPlaceholder,
		cfg.MenuItems,
		cfg.ProfileItems,
		cIIndex)
	templ.Handler(cIList).ServeHTTP(w, r)
}
func (cfg *ApiConfig) handleConfigurationItemsEditPage(w http.ResponseWriter, r *http.Request, u database.User) {
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

	cEIndexNew := views.ConfigurationItemFormNew(selectOptionsCompany, ci)
	cEdit := views.ConfigurationItemsEdit("Configuration Items - Edit",
		cfg.Logo,
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.ProfilePicPlaceholder,
		cfg.MenuItems,
		cfg.ProfileItems,
		cEIndexNew)
	templ.Handler(cEdit).ServeHTTP(w, r)
}
func (cfg *ApiConfig) handleConfigurationItemsPostPage(w http.ResponseWriter, r *http.Request, u database.User) {

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
		cfg.Logo,
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.ProfilePicPlaceholder,
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

	// iEIndex := views.IncidentsEditIndex(incident, companies, configurationItems)
	iEIndexNew := views.IncidentFormNew(selectOptionsCompany, selectOptionsCI, incident)
	iEdit := views.IncidentsEdit("Incidents - Edit",
		cfg.Logo,
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.ProfilePicPlaceholder,
		cfg.MenuItems,
		cfg.ProfileItems,
		iEIndexNew)
	templ.Handler(iEdit).ServeHTTP(w, r)
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

	w.Header().Set("HX-Location", "/incidents")
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

// Companies
func (cfg *ApiConfig) handleViewCompanies(w http.ResponseWriter, r *http.Request, u database.User) {
	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}

	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find comanies")
		return
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)

	cIndex := views.CompaniesIndex(companies)
	cList := views.CompaniesList("Companies List",
		cfg.Logo,
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.ProfilePicPlaceholder,
		cfg.MenuItems,
		cfg.ProfileItems,
		cIndex)
	templ.Handler(cList).ServeHTTP(w, r)
}

// Users
func (cfg *ApiConfig) handleViewUsers(w http.ResponseWriter, r *http.Request, u database.User) {
	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}

	databaseUsers, err := cfg.DB.GetUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find comanies")
		return
	}
	users := models.DatabaseUsersToUsers(databaseUsers)

	cIndex := views.UsersIndex(users)
	cList := views.UsersList("Users List",
		cfg.Logo,
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.ProfilePicPlaceholder,
		cfg.MenuItems,
		cfg.ProfileItems,
		cIndex)
	templ.Handler(cList).ServeHTTP(w, r)
}

func (cfg *ApiConfig) handlePageIndex(w http.ResponseWriter, r *http.Request, u database.User) {
	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}
	hindex := views.HomeIndex(fromProtected)
	home := views.Home("TicketTracker", cfg.Logo, fromProtected, false, "msg", u.Name, u.Email,
		cfg.ProfilePicPlaceholder,
		cfg.MenuItems, cfg.ProfileItems, hindex)
	templ.Handler(home).ServeHTTP(w, r)
}
