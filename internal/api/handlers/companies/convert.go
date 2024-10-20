package companies

import (
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

// convert converts a single database.Company to data.Company.
func convert(company database.Company) data.Company {
	return data.Company{
		ID:        company.ID,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
		Name:      company.Name,
	}
}

// convertMany converts a slice of database.Company to a slice of data.Company.
func convertMany(companies []database.Company) []data.Company {
	var items []data.Company
	for _, item := range companies {
		items = append(items, convert(item))
	}
	return items
}

// convertRowsAndMetadata converts a slice of database.GetCompaniesRow to a slice of data.Company
// and calculates metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetCompaniesRow, filters data.Filters) ([]data.Company, data.Metadata, error) {
	var output []data.Company
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}

	// Prevent conversion exploits
	v32, err := data.ConvertInt64to32(totalRecords)
	if err != nil {
		return nil, data.Metadata{}, err
	}

	metadata := data.CalculateMetadata(v32, filters.Page, filters.PageSize)
	return output, metadata, nil
}

// convertRowAndCount converts a single database.GetCompaniesRow to data.Company and updates the count.
func convertRowAndCount(row database.GetCompaniesRow, count *int64) data.Company {
	outputRow := data.Company{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Name:      row.Name,
	}
	*count = row.Count

	return outputRow
}
