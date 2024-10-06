package server

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) GetCompaniesSelection(r *http.Request, dst any) error {
	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		return errors.New("couldn't find companies")
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)
	selectOptionsCompany := models.SelectOptions{}

	for _, company := range companies {
		selectOptionsCompany = append(selectOptionsCompany, models.NewSelectOption(company.Name, company.ID.String()))
	}
	*dst.(*models.SelectOptions) = selectOptionsCompany
	return nil
}

func (cfg *ApiConfig) GetCISelection(r *http.Request, dst any) error {
	databaseCIs, err := cfg.DB.GetConfigurationItems(r.Context())
	if err != nil {
		return errors.New("couldn't find configuration items")
	}
	cis := models.DatabaseConfigurationItemsToConfigurationItems(databaseCIs)
	selectOptionsCI := models.SelectOptions{}

	for _, ci := range cis {
		selectOptionsCI = append(selectOptionsCI, models.NewSelectOption(ci.Name, ci.ID.String()))
	}
	*dst.(*models.SelectOptions) = selectOptionsCI
	return nil
}

func (cfg *ApiConfig) GetUsersSelection(r *http.Request, dst any) error {
	databaseUsers, err := cfg.DB.GetUsers(r.Context())
	if err != nil {
		return errors.New("couldn't find users")
	}
	users := models.DatabaseUsersToUsers(databaseUsers)
	selectOptionsUsers := models.SelectOptions{}

	for _, user := range users {
		selectOptionsUsers = append(selectOptionsUsers, models.NewSelectOption(user.Name, user.ID.String()))
	}
	*dst.(*models.SelectOptions) = selectOptionsUsers
	return nil
}

func (cfg *ApiConfig) GetStatesSelection(r *http.Request, dst any) error {
	selectOptionsStates := models.SelectOptions{}
	for _, so := range models.StateOptionsEnum {
		selectOptionsStates = append(selectOptionsStates, models.NewSelectOption(string(so), string(so)))
	}
	*dst.(*models.SelectOptions) = selectOptionsStates
	return nil
}

func (cfg *ApiConfig) GetStateSelection(dst any) {
	stateOptions := models.SelectOptions{}
	for _, so := range models.StateOptionsEnum {
		stateOptions = append(stateOptions, models.NewSelectOption(string(so), string(so)))
	}
	*dst.(*models.SelectOptions) = stateOptions
}

func (cfg *ApiConfig) GetCIs(r *http.Request) ([]models.ConfigurationItem, error) {
	databaseConfigurationItems, err := cfg.DB.GetConfigurationItems(r.Context())
	if err != nil {
		return nil, err
	}
	cis := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)
	return cis, nil
}

func (cfg *ApiConfig) GetCIsByCompanyID(r *http.Request, id uuid.UUID) ([]models.ConfigurationItem, error) {
	databaseConfigurationItems, err := cfg.DB.GetConfigurationItemsByCompanyID(r.Context(), id)
	if err != nil {
		return nil, err
	}
	cis := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)
	return cis, nil
}

func (cfg *ApiConfig) GetCIByID(r *http.Request, id uuid.UUID) (models.ConfigurationItem, error) {
	databaseConfigurationItem, err := cfg.DB.GetConfigurationItemByID(r.Context(), id)
	if err != nil {
		return models.ConfigurationItem{}, errors.New("can't find configuration item")
	}
	ci := models.DatabaseConfigurationItemToConfigurationItem(databaseConfigurationItem)
	return ci, nil
}

func (cfg *ApiConfig) GetIncidentByID(r *http.Request, id uuid.UUID) (models.Incident, error) {
	databaseIncident, err := cfg.DB.GetIncidentByID(r.Context(), id)
	if err != nil {
		return models.Incident{}, errors.New("can't find incident")
	}
	incident := models.DatabaseIncidentToIncident(databaseIncident)
	return incident, nil
}

//	func (cfg *ApiConfig) GetIncidents(r *http.Request, filters filters.Filter) ([]models.Incident, error) {
//		databaseIncidentsAsc := []database.GetIncidentsAscRow{}
//		databaseIncidentsDesc := []database.GetIncidentsDescRow{}
//		var err error
func (cfg *ApiConfig) GetIncidents(r *http.Request) ([]models.Incident, error) {
	databaseIncidents, err := cfg.DB.GetIncidents(r.Context())
	if err != nil {
		return nil, errors.New("couldn't find incidents")
	}
	incidents := models.DatabaseIncidentsRowToIncidents(databaseIncidents)

	for n, i := range incidents {
		ci, err := cfg.DB.GetConfigurationItemByID(r.Context(), i.ConfigurationItemID)
		if err != nil {
			return nil, errors.New("couldn't find configuration item name")
		}
		incidents[n].ConfigurationItemName = ci.Name

	}
	return incidents, nil
}

func (cfg *ApiConfig) GetIncidentsFiltered(r *http.Request, query string, limit, offset int) ([]models.Incident, error) {
	databaseIncidentsFilteredRow, err := cfg.DB.GetIncidentsFiltered(r.Context(), query)
	if err != nil {
		return nil, errors.New("couldn't find incidents")
	}
	incidents := models.DatabaseIncidentsFilteredRowToIncidents(databaseIncidentsFilteredRow)
	return incidents, nil
}
func (cfg *ApiConfig) UpdateIncident(r *http.Request, i models.Incident) (models.Incident, error) {
	dbIncident, err := cfg.DB.UpdateIncident(r.Context(), database.UpdateIncidentParams{
		ID:                  i.ID,
		UpdatedAt:           time.Now(),
		CompanyID:           i.CompanyID,
		ConfigurationItemID: i.ConfigurationItemID,
		Description:         sql.NullString{String: i.Description, Valid: i.Description != ""},
		ShortDescription:    i.ShortDescription,
		State:               i.State,
		AssignedTo:          uuid.NullUUID{UUID: i.AssignedTo, Valid: true},
	})
	if err != nil {
		return models.Incident{}, errors.New("couldn't update incident")
	}
	incident := models.DatabaseIncidentToIncident(dbIncident)

	return incident, nil
}

func (cfg *ApiConfig) GetCompanies(r *http.Request) ([]models.Company, error) {
	databaseCompanies, err := cfg.DB.GetCompanies(r.Context())
	if err != nil {
		return nil, errors.New("couldn't find companiess")
	}
	companies := models.DatabaseCompaniesToCompanies(databaseCompanies)

	return companies, nil
}

func (cfg *ApiConfig) GetUsers(r *http.Request) ([]models.User, error) {
	databaseUsers, err := cfg.DB.GetUsers(r.Context())
	if err != nil {
		return nil, errors.New("couldn't find users")
	}
	users := models.DatabaseUsersToUsers(databaseUsers)

	return users, nil
}

func NewIncidentEmpty() models.Incident {
	return models.Incident{
		ID:                    uuid.New(),
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
		ShortDescription:      "",
		Description:           "",
		State:                 "",
		AssignedTo:            [16]byte{},
		AssignedToName:        "",
		ConfigurationItemID:   [16]byte{},
		ConfigurationItemName: "",
		CompanyID:             [16]byte{},
	}
}

func NewIncident(id uuid.UUID,
	companyID,
	configurationItemID,
	assignedToID uuid.UUID,
	shortDescription,
	description string,
	state database.StateEnum) models.Incident {
	return models.Incident{
		ID:                    id,
		CreatedAt:             time.Time{},
		UpdatedAt:             time.Time{},
		ShortDescription:      shortDescription,
		Description:           description,
		State:                 state,
		AssignedTo:            assignedToID,
		AssignedToName:        "",
		ConfigurationItemID:   configurationItemID,
		ConfigurationItemName: "",
		CompanyID:             companyID,
	}
}
