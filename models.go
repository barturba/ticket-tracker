package main

import (
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
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
	Email     string    `json:"email"`
	APIkey    string    `json:"api_key"`
	Token     string    `json:"token"`
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

func databaseConfigurationItemToConfigurationItem(configurationItem database.ConfigurationItem) models.ConfigurationItem {
	return models.ConfigurationItem{
		ID:             configurationItem.ID,
		CreatedAt:      configurationItem.CreatedAt,
		UpdatedAt:      configurationItem.UpdatedAt,
		Name:           configurationItem.Name,
		OrganizationID: configurationItem.OrganizationID.String(),
	}
}

func databaseConfigurationItemsToConfigurationItems(configurationItems []database.ConfigurationItem) []models.ConfigurationItem {
	var items []models.ConfigurationItem
	for _, configurationItem := range configurationItems {
		items = append(items, databaseConfigurationItemToConfigurationItem(configurationItem))
	}
	return items
}

func databaseCompanyToCompany(company database.Company) models.Company {
	return models.Company{
		ID:             company.ID,
		CreatedAt:      company.CreatedAt,
		UpdatedAt:      company.UpdatedAt,
		Name:           company.Name,
		OrganizationID: company.ID,
	}
}
func databaseCompaniesToCompanies(companies []database.Company) []models.Company {
	var items []models.Company
	for _, item := range companies {
		items = append(items, databaseCompanyToCompany(item))
	}
	return items
}

func databaseIncidentToIncident(incident database.Incident) models.Incident {
	return models.Incident{
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

func databaseIncidentByOrganizationIDRowToIncident(incident database.GetIncidentsByOrganizationIDRow) models.Incident {
	return models.Incident{
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

func databaseGetIncidentsByOrganizationIDRowToIncidents(incidents []database.GetIncidentsByOrganizationIDRow) []models.Incident {
	var items []models.Incident
	for _, item := range incidents {
		items = append(items, databaseIncidentByOrganizationIDRowToIncident(item))
	}
	return items
}

func databaseIncidentsToIncidents(incidents []database.Incident) []models.Incident {
	var items []models.Incident
	for _, item := range incidents {
		items = append(items, databaseIncidentToIncident(item))
	}
	return items
}

type Page struct {
}
