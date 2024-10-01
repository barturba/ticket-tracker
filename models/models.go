package models

import (
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
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
	AssignedTo            uuid.UUID          `json:"assigned_to"`
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

type FormData struct {
	Errors map[string]string
	Values map[string]string
}

func NewFormData() FormData {
	return FormData{
		Errors: map[string]string{},
		Values: map[string]string{},
	}
}

type Page struct {
	Title            string
	Logo             string
	FlashMessage     string
	IsLoggedIn       bool
	IsError          bool
	Msg              string
	User             string
	Email            string
	ProfilePicture   string
	MenuItems        MenuItems
	ProfileMenuItems MenuItems
}
