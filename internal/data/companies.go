package data

import (
	"time"

	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// Companies
type Company struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type CompanyRow struct {
	Count     int64     `json:"count"`
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func ValidateCompany(v *validator.Validator, company *Company) {
	v.Check(company.ID != uuid.UUID{}, "id", "must be provided")

	v.Check(company.Name != "", "name", "must be provided")
	v.Check(len(company.Name) <= 500, "name", "must not be more than 500 bytes long")
}
