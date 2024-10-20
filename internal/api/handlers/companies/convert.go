package companies

import (
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
)

// convert converts a single database.Company to models.Company.
func convert(company database.Company) models.Company {
	return models.Company{
		ID:        company.ID,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
		Name:      company.Name,
	}
}

// convertMany converts a slice of database.Company to a slice of models.Company.
func convertMany(companies []database.Company) []models.Company {
	var items []models.Company
	for _, item := range companies {
		items = append(items, convert(item))
	}
	return items
}

// convertRowsAndMetadata converts a slice of database.GetCompaniesRow to a slice of models.Company
// and calculates metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetCompaniesRow, filters models.Filters) ([]models.Company, models.Metadata, error) {
	var output []models.Company
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}

	// Prevent conversion exploits
	v32, err := models.ConvertInt64to32(totalRecords)
	if err != nil {
		return nil, models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(v32, filters.Page, filters.PageSize)
	return output, metadata, nil
}

// convertRowAndCount converts a single database.GetCompaniesRow to models.Company and updates the count.
func convertRowAndCount(row database.GetCompaniesRow, count *int64) models.Company {
	outputRow := models.Company{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Name:      row.Name,
	}
	*count = row.Count

	return outputRow
}
