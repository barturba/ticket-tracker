package companies

import (
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

func convert(company database.Company) data.Company {
	return data.Company{
		ID:        company.ID,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
		Name:      company.Name,
	}
}
