package cis

import (
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
)

// The convert function converts a single database.ConfigurationItem to a models.CI.
func convert(ci database.ConfigurationItem) models.CI {
	return models.CI{
		ID:        ci.ID,
		CreatedAt: ci.CreatedAt,
		UpdatedAt: ci.UpdatedAt,
		Name:      ci.Name,
	}
}

// The convertMany function converts a slice of database.ConfigurationItem to a slice of models.CI.
func convertMany(cis []database.ConfigurationItem) []models.CI {
	var items []models.CI
	for _, item := range cis {
		items = append(items, convert(item))
	}
	return items
}

// The convertRowsAndMetadata function converts a slice of database.GetCIsRow to a slice of models.CI
// and calculates metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetCIsRow, filters models.Filters) ([]models.CI, models.Metadata, error) {
	var output []models.CI
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

// The convertRowAndCount function converts a single database.GetCIsRow to a models.CI and updates the count.
func convertRowAndCount(row database.GetCIsRow, count *int64) models.CI {
	outputRow := models.CI{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Name:      row.Name,
	}
	*count = row.Count

	return outputRow
}
