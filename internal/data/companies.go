package data

import (
	"time"

	"github.com/barturba/ticket-tracker/pkg/validator"
	"github.com/google/uuid"
)

// Company represents a company entity with relevant details.
type Company struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the company.
	CreatedAt time.Time `json:"created_at"` // Timestamp when the company record was created.
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the company record was last updated.
	Name      string    `json:"name"`       // Name of the company.
}

// CompanyRow represents a row in a list of companies, including a count of companies.
type CompanyRow struct {
	Count     int64     `json:"count"`      // Total number of companies.
	ID        uuid.UUID `json:"id"`         // Unique identifier for the company.
	CreatedAt time.Time `json:"created_at"` // Timestamp when the company record was created.
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the company record was last updated.
	Name      string    `json:"name"`       // Name of the company.
}

// ValidateCompany validates the fields of a Company struct using the provided validator.
// It checks that the ID is provided, the name is not empty, and the name does not exceed 500 bytes.
func ValidateCompany(v *validator.Validator, company *Company) {
	v.Check(company.ID != uuid.UUID{}, "id", "must be provided")
	v.Check(company.Name != "", "name", "must be provided")
	v.Check(len(company.Name) <= 500, "name", "must not be more than 500 bytes long")
}
