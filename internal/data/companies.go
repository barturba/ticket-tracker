package data

import (
	"time"

	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// Company represents a company entity with an ID, creation and update timestamps, and a name.
type Company struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

// CompanyRow represents a row in a list of companies, including a count of companies.
type CompanyRow struct {
	Count     int64     `json:"count"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

// ValidateCompany validates the fields of a Company struct using the provided validator.
// It checks that the ID is provided, the name is not empty, and the name does not exceed 500 bytes.
func ValidateCompany(v *validator.Validator, company *Company) {
	v.Check(company.ID != uuid.UUID{}, "id", "must be provided")
	v.Check(company.Name != "", "name", "must be provided")
	v.Check(len(company.Name) <= 500, "name", "must not be more than 500 bytes long")
}
