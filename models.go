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
