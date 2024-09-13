package main

import (
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/google/uuid"
)

type Organization struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	UserID    string    `json:"user_id"`
}

func databaseOrganizationToOrganization(organization database.Organization) Organization {
	return Organization{
		ID:        organization.ID,
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
		Name:      organization.Name,
		UserID:    organization.UserID.String(),
	}

}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIkey    string    `json:"api_key"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		APIkey:    user.Apikey,
	}
}

type ConfigurationItem struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	OrganizationID string    `json:"organization_id"`
}

func databaseConfigurationItemToConfigurationItem(configurationItem database.ConfigurationItem) ConfigurationItem {
	return ConfigurationItem{
		ID:             configurationItem.ID,
		CreatedAt:      configurationItem.CreatedAt,
		UpdatedAt:      configurationItem.UpdatedAt,
		Name:           configurationItem.Name,
		OrganizationID: configurationItem.OrganizationID.String(),
	}
}

func databaseConfigurationItemsToConfigurationItems(configurationItems []database.ConfigurationItem) []ConfigurationItem {
	var items []ConfigurationItem
	for _, configurationItem := range configurationItems {
		items = append(items, databaseConfigurationItemToConfigurationItem(configurationItem))
	}
	return items
}

type Company struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	OrganizationID uuid.UUID `json:"organization_id"`
}

func databaseCompanyToCompany(company database.Company) Company {
	return Company{
		ID:             company.ID,
		CreatedAt:      company.CreatedAt,
		UpdatedAt:      company.UpdatedAt,
		Name:           company.Name,
		OrganizationID: company.ID,
	}
}

type Incident struct {
	ID                  uuid.UUID          `json:"id"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	ShortDescription    string             `json:"short_description"`
	Description         string             `json:"description"`
	State               database.StateEnum `json:"state"`
	AssignedTo          uuid.UUID          `json:"assigned_to"`
	ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
	OrganizationID      uuid.UUID          `json:"organization_id"`
	CompanyID           uuid.UUID          `json:"company_id"`
}

func databaseIncidentToIncident(incident database.Incident) Incident {
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
