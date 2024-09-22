package models

import (
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/google/uuid"
)

type Company struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	OrganizationID uuid.UUID `json:"organization_id"`
}

type ConfigurationItem struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	OrganizationID string    `json:"organization_id"`
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

	OrganizationID uuid.UUID `json:"organization_id"`
	CompanyID      uuid.UUID `json:"company_id"`
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

type Organization struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	UserID    string    `json:"user_id"`
}

func DatabaseOrganizationToOrganization(organization database.Organization) Organization {
	return Organization{
		ID:        organization.ID,
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
		Name:      organization.Name,
		UserID:    organization.UserID.String(),
	}

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

func DatabaseConfigurationItemToConfigurationItem(configurationItem database.ConfigurationItem) ConfigurationItem {
	return ConfigurationItem{
		ID:             configurationItem.ID,
		CreatedAt:      configurationItem.CreatedAt,
		UpdatedAt:      configurationItem.UpdatedAt,
		Name:           configurationItem.Name,
		OrganizationID: configurationItem.OrganizationID.String(),
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
		ID:             company.ID,
		CreatedAt:      company.CreatedAt,
		UpdatedAt:      company.UpdatedAt,
		Name:           company.Name,
		OrganizationID: company.ID,
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
		OrganizationID:      incident.OrganizationID,
		CompanyID:           incident.CompanyID,
	}
}

func DatabaseIncidentByOrganizationIDRowToIncident(incident database.GetIncidentsByOrganizationIDRow) Incident {
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
		OrganizationID:      incident.OrganizationID,
		CompanyID:           incident.CompanyID,
	}
}

func DatabaseIncidentsByOrganizationIDRowToIncidents(incidents []database.GetIncidentsByOrganizationIDRow) []Incident {
	var items []Incident
	for _, item := range incidents {
		items = append(items, DatabaseIncidentByOrganizationIDRowToIncident(item))
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

type Page struct {
}
