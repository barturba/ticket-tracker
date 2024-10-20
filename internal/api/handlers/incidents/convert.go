// Package incidents provides functions to convert database incident records
// into data incident structures used within the application.
package incidents

import (
	"fmt"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
)

// convert converts a single database.Incident to a models.Incident.
func convert(incident database.Incident) models.Incident {
	return models.Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
		AssignedToID:        incident.AssignedTo,
		State:               incident.State,
	}
}

// convertLatestRow converts a single database.GetIncidentsLatestRow to a models.Incident.
func convertLatestRow(incident database.GetIncidentsLatestRow) models.Incident {
	return models.Incident{
		ID:                  incident.ID,
		CreatedAt:           incident.CreatedAt,
		UpdatedAt:           incident.UpdatedAt,
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
		AssignedToID:        incident.AssignedTo,
		State:               incident.State,
	}
}

// convertLatestRowMany converts a slice of database.GetIncidentsLatestRow to a slice of models.Incident.
func convertLatestRowMany(incidents []database.GetIncidentsLatestRow) []models.Incident {
	var items []models.Incident
	for _, item := range incidents {
		items = append(items, convertLatestRow(item))
	}
	return items
}

// convertRowsAndMetadata converts a slice of database.GetIncidentsRow to a slice of models.Incident
// and calculates metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetIncidentsRow, filters models.Filters) ([]models.Incident, models.Metadata, error) {

	var output []models.Incident
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

// convertRowAndCount converts a single database.GetIncidentsRow to a models.Incident
// and updates the total record count.
func convertRowAndCount(row database.GetIncidentsRow, count *int64) models.Incident {
	outputRow := models.Incident{
		ID:                  row.ID,
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
		ShortDescription:    row.ShortDescription,
		Description:         row.Description,
		ConfigurationItemID: row.ConfigurationItemID,
		CompanyID:           row.CompanyID,
		AssignedToID:        row.AssignedTo,
		AssignedToName:      fmt.Sprintf("%s %s", row.FirstName.String, row.LastName.String),
		State:               row.State,
	}
	*count = row.Count

	return outputRow
}
