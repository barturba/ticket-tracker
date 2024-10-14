package cis

import (
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

// The convert function converts a single database.ConfigurationItem to a data.CI.
func convert(ci database.ConfigurationItem) data.CI {
	return data.CI{
		ID:        ci.ID,
		CreatedAt: ci.CreatedAt,
		UpdatedAt: ci.UpdatedAt,
		Name:      ci.Name,
	}
}

// The convertMany function converts a slice of database.ConfigurationItem to a slice of data.CI.
func convertMany(cis []database.ConfigurationItem) []data.CI {
	var items []data.CI
	for _, item := range cis {
		items = append(items, convert(item))
	}
	return items
}

// The convertRowsAndMetadata function converts a slice of database.GetCIsRow to a slice of data.CI
// and calculates metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetCIsRow, filters data.Filters) ([]data.CI, data.Metadata) {
	var output []data.CI
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}
	metadata := data.CalculateMetadata(int(totalRecords), filters.Page, filters.PageSize)
	return output, metadata
}

// The convertRowAndCount function converts a single database.GetCIsRow to a data.CI and updates the count.
func convertRowAndCount(row database.GetCIsRow, count *int64) data.CI {
	outputRow := data.CI{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Name:      row.Name,
	}
	*count = row.Count

	return outputRow
}
