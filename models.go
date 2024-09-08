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
}

func databaseOrganizationToOrganization(organization database.Organization) Organization {
	return Organization{
		ID:        organization.ID,
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
		Name:      organization.Name,
	}

}

type User struct {
	ID             uuid.UUID `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	APIkey         string    `json:"api_key"`
	OrganizationID string    `json:"organization_id"`
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		Name:           user.Name,
		APIkey:         user.Apikey,
		OrganizationID: user.OrganizationID.String(),
	}
}
