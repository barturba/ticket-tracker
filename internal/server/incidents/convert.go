// Package incidents provides functions to convert database incident records
// into data incident structures used within the application.
package incidents

import (
	"fmt"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

// convert converts a single database.Incident to a data.Incident.
func convert(incident database.Incident) data.Incident {
	return data.Incident{
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

// convertLatestRow converts a single database.GetIncidentsLatestRow to a data.Incident.
func convertLatestRow(incident database.GetIncidentsLatestRow) data.Incident {
	return data.Incident{
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

// convertLatestRowMany converts a slice of database.GetIncidentsLatestRow to a slice of data.Incident.
func convertLatestRowMany(incidents []database.GetIncidentsLatestRow) []data.Incident {
	var items []data.Incident
	for _, item := range incidents {
		items = append(items, convertLatestRow(item))
	}
	return items
}

// convertRowsAndMetadata converts a slice of database.GetIncidentsRow to a slice of data.Incident
// and calculates metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetIncidentsRow, filters data.Filters) ([]data.Incident, data.Metadata) {
	var output []data.Incident
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}
	metadata := data.CalculateMetadata(int(totalRecords), filters.Page, filters.PageSize)
	return output, metadata
}

// convertRowAndCount converts a single database.GetIncidentsRow to a data.Incident
// and updates the total record count.
func convertRowAndCount(row database.GetIncidentsRow, count *int64) data.Incident {
	outputRow := data.Incident{
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
