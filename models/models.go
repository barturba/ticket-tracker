package models

import (
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type ConfigurationItem struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type Incident struct {
	ID                    uuid.UUID          `json:"id"`
	CreatedAt             time.Time          `json:"created_at"`
	UpdatedAt             time.Time          `json:"updated_at"`
	ShortDescription      string             `json:"short_description"`
	Description           string             `json:"description"`
	State                 database.StateEnum `json:"state"`
	AssignedTo            uuid.UUID          `json:"assigned_to_id"`
	AssignedToName        string             `json:"assigned_to_name"`
	ConfigurationItemID   uuid.UUID          `json:"configuration_item_id"`
	ConfigurationItemName string             `json:"configuration_item_name"`
	CompanyID             uuid.UUID          `json:"company_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	APIkey    string    `json:"api_key"`
	Token     string    `json:"token"`
}

func DatabaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIkey:    user.Apikey,
	}
}

func DatabaseUsersToUsers(users []database.User) []User {
	var items []User
	for _, item := range users {
		items = append(items, DatabaseUserToUser(item))
	}
	return items
}

func DatabaseConfigurationItemToConfigurationItem(configurationItem database.ConfigurationItem) ConfigurationItem {
	return ConfigurationItem{
		ID:        configurationItem.ID,
		CreatedAt: configurationItem.CreatedAt,
		UpdatedAt: configurationItem.UpdatedAt,
		Name:      configurationItem.Name,
	}
}

func DatabaseConfigurationItemsToConfigurationItems(configurationItems []database.ConfigurationItem) []ConfigurationItem {
	var items []ConfigurationItem
	for _, configurationItem := range configurationItems {
		items = append(items, DatabaseConfigurationItemToConfigurationItem(configurationItem))
	}
	return items
}

func DatabaseCompanyToCompany(company database.Company) Company {
	return Company{
		ID:        company.ID,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
		Name:      company.Name,
	}
}
func DatabaseCompaniesToCompanies(companies []database.Company) []Company {
	var items []Company
	for _, item := range companies {
		items = append(items, DatabaseCompanyToCompany(item))
	}
	return items
}

func DatabaseIncidentToIncident(incident database.Incident) Incident {
	return Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description.String,
		State:               incident.State,
		AssignedTo:          incident.AssignedTo.UUID,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}
}

func DatabaseIncidentRowToIncident(incident database.GetIncidentsRow) Incident {
	return Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description.String,
		State:               incident.State,
		AssignedTo:          incident.AssignedTo.UUID,
		AssignedToName:      incident.Name.String,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}
}

func DatabaseIncidentsRowToIncidents(incidents []database.GetIncidentsRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentRowToIncident(item))
	}
	return items
}

func DatabaseIncidentBySearchTermRowToIncident(incident database.GetIncidentsBySearchTermRow) Incident {
	return Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description.String,
		State:               incident.State,
		AssignedTo:          incident.AssignedTo.UUID,
		AssignedToName:      incident.Name.String,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}
}

func DatabaseIncidentsBySearchTermToIncidents(incidents []database.GetIncidentsBySearchTermRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentBySearchTermRowToIncident(item))
	}
	return items
}

func DatabaseIncidentsToIncidents(incidents []database.Incident) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentToIncident(item))
	}
	return items
}

func DatabaseIncidentBySearchTermLimitOffsetRowToIncident(incident database.GetIncidentsBySearchTermLimitOffsetRow) Incident {
	return Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description.String,
		State:               incident.State,
		AssignedTo:          incident.AssignedTo.UUID,
		AssignedToName:      incident.Name.String,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}
}
func DatabaseIncidentsBySearchTermLimitOffsetRowToIncidents(incidents []database.GetIncidentsBySearchTermLimitOffsetRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentBySearchTermLimitOffsetRowToIncident(item))
	}
	return items
}

type MenuItem struct {
	Name string
	Link string
}

type MenuItems []MenuItem

func NewMenuItem(name, link string) MenuItem {
	return MenuItem{
		Name: name,
		Link: link,
	}
}

type SelectOption struct {
	Name string
	Link string
}

type SelectOptions []SelectOption

func NewSelectOption(name, link string) SelectOption {
	return SelectOption{
		Name: name,
		Link: link,
	}
}

var StateOptionsEnum = []database.StateEnum{
	database.StateEnumNew,
	database.StateEnumAssigned,
	database.StateEnumInProgress,
	database.StateEnumOnHold,
	database.StateEnumResolved,
}

var NewIncidentInput struct {
	ShortDescription    string             `json:"short_description"`
	Description         string             `json:"description"`
	CompanyID           uuid.UUID          `json:"company_id"`
	AssignedToID        uuid.UUID          `json:"assigned_to_id"`
	ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
	State               database.StateEnum `json:"state"`
}

var IncidentInput struct {
	ID                  uuid.UUID          `json:"id"`
	ShortDescription    string             `json:"short_description"`
	Description         string             `json:"description"`
	CompanyID           uuid.UUID          `json:"company_id"`
	AssignedToID        uuid.UUID          `json:"assigned_to_id"`
	ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
	State               database.StateEnum `json:"state"`
}

func NewIncidentEmpty() Incident {
	return Incident{
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

func NewIncident(id uuid.UUID, companyID, configurationItemID, assignedToID uuid.UUID, shortDescription, description string, state database.StateEnum) Incident {
	return Incident{
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

func CheckIncident(i Incident) map[string]string {
	StateEnum := []database.StateEnum{
		database.StateEnumNew,
		database.StateEnumAssigned,
		database.StateEnumInProgress,
		database.StateEnumOnHold,
		database.StateEnumResolved,
	}
	v := validator.New()
	v.Check(i.ID != uuid.UUID{}, "id", "must be provided")
	v.Check(i.CompanyID != uuid.UUID{}, "company_id", "must be provided")
	v.Check(i.ConfigurationItemID != uuid.UUID{}, "configuration_item_id", "must be provided")
	v.Check(validator.PermittedValue(i.State, StateEnum...), "state", "This field must equal New, Assigned, In Progress, On Hold, or Resolved")
	v.Check(i.AssignedTo != uuid.UUID{}, "assigned_to_id", "must be provided")
	v.Check(i.Description != "", "description", "must be provided")
	v.Check(i.ShortDescription != "", "short_description", "must be provided")
	if !v.Valid() {
		return v.Errors
	}
	return nil
}

type AlertEnum string

const (
	AlertEnumError   AlertEnum = "Error"
	AlertEnumWarning AlertEnum = "Warning"
	AlertEnumSuccess AlertEnum = "Success"
	AlertEnumInfo    AlertEnum = "Info"
)

type Alert struct {
	Message   string
	AlertType AlertEnum
	Color     string
}

func NewAlert(message string, alertType AlertEnum, color string) Alert {
	return Alert{
		Message:   message,
		AlertType: alertType,
		Color:     color,
	}
}

func DatabaseIncidentsFilteredRowToIncidents(incidents []database.GetIncidentsRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentFilteredRowToIncident(item))
	}
	return items
}
func DatabaseIncidentFilteredRowToIncident(incident database.GetIncidentsRow) Incident {
	return Incident{
		ID:                    incident.ID,
		CreatedAt:             incident.CreatedAt,
		UpdatedAt:             incident.UpdatedAt,
		ShortDescription:      incident.ShortDescription,
		Description:           incident.Description.String,
		State:                 incident.State,
		AssignedTo:            incident.AssignedTo.UUID,
		AssignedToName:        incident.Name.String,
		ConfigurationItemID:   incident.ConfigurationItemID,
		ConfigurationItemName: "",
		CompanyID:             incident.CompanyID,
	}
}

func DatabaseIncidentsLatestRowToIncidents(incidents []database.GetIncidentsLatestRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentLatestRowToIncident(item))
	}
	return items
}
func DatabaseIncidentLatestRowToIncident(incident database.GetIncidentsLatestRow) Incident {
	return Incident{
		ID:                    incident.ID,
		CreatedAt:             incident.CreatedAt,
		UpdatedAt:             incident.UpdatedAt,
		ShortDescription:      incident.ShortDescription,
		Description:           incident.Description.String,
		State:                 incident.State,
		AssignedTo:            incident.AssignedTo.UUID,
		AssignedToName:        incident.Name.String,
		ConfigurationItemID:   incident.ConfigurationItemID,
		ConfigurationItemName: "",
		CompanyID:             incident.CompanyID,
	}
}

// Companies

var NewCompanyInput struct {
	Name string `json:"name"`
}

var CompanyInput struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewCompanyEmpty() Company {
	return Company{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "",
	}
}

func NewCompany(id uuid.UUID,
	name string,
) Company {
	return Company{
		ID:        id,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      name,
	}
}

func CheckCompany(c Company) map[string]string {
	v := validator.New()
	v.Check(c.ID != uuid.UUID{}, "id", "must be provided")
	v.Check(c.Name != "", "name", "must be provided")
	if !v.Valid() {
		return v.Errors
	}
	return nil
}

func DatabaseCompaniesRowToCompanies(incidents []database.Company) []Company {
	var items []Company
	for _, item := range incidents {
		items = append(items, DatabaseCompanyRowToCompany(item))
	}
	return items
}

func DatabaseCompanyRowToCompany(incident database.Company) Company {
	return Company{
		ID:        incident.ID,
		CreatedAt: incident.CreatedAt,
		UpdatedAt: incident.UpdatedAt,
		Name:      incident.Name,
	}
}

type Count struct {
	Count int
}

func NewUser(id uuid.UUID, name, email, apiKey, token string) User {
	return User{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Email:     email,
		APIkey:    apiKey,
		Token:     token,
	}
}
