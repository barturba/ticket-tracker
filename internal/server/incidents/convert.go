package incidents

import (
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

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

func convertMany(incidents []database.Incident) []data.Incident {
	var items []data.Incident
	for _, item := range incidents {
		items = append(items, convert(item))
	}
	return items
}

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

func convertLatestRowMany(incidents []database.GetIncidentsLatestRow) []data.Incident {
	var items []data.Incident
	for _, item := range incidents {
		items = append(items, convertLatestRow(item))
	}
	return items
}

func convertRow(incident database.GetIncidentsRow) data.Incident {
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

func convertRowMany(incidents []database.GetIncidentsRow) []data.Incident {
	var items []data.Incident
	for _, item := range incidents {
		items = append(items, convertRow(item))
	}
	return items
}
