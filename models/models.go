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
	ID                  uuid.UUID          `json:"id"`
	CreatedAt           time.Time          `json:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at"`
	ShortDescription    string             `json:"short_description"`
	Description         string             `json:"description"`
	State               database.StateEnum `json:"state"`
	AssignedTo          uuid.UUID          `json:"assigned_to"`
	AssignedToName      string             `json:"assigned_to_name"`
	ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
	OrganizationID      uuid.UUID          `json:"organization_id"`
	CompanyID           uuid.UUID          `json:"company_id"`
}
