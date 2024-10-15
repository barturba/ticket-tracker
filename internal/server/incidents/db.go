package incidents

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/google/uuid"
)

// GetFromDB retrieves incidents from the database based on the provided query and filters.
func GetFromDB(r *http.Request, db *database.Queries, query string, filters data.Filters) ([]data.Incident, data.Metadata, error) {
	p := database.GetIncidentsParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}
	rows, err := db.GetIncidents(r.Context(), p)
	if err != nil {
		return nil, data.Metadata{}, errors.New("couldn't find incidents")
	}
	incidents, metadata := convertRowsAndMetadata(rows, filters)
	return incidents, metadata, nil
}

// GetAllFromDB retrieves all incidents from the database based on the provided filters.
func GetAllFromDB(r *http.Request, db *database.Queries, filters data.Filters) ([]data.Incident, data.Metadata, error) {
	p := database.GetIncidentsParams{
		Limit:    1_000_000,
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}
	rows, err := db.GetIncidents(r.Context(), p)
	if err != nil {
		return nil, data.Metadata{}, errors.New("couldn't find incidents")
	}
	incidents, metadata := convertRowsAndMetadata(rows, filters)
	return incidents, metadata, nil
}

// GetLatestFromDB retrieves the latest incidents from the database based on the provided limit and offset.
func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int) ([]data.Incident, error) {
	p := database.GetIncidentsLatestParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetIncidentsLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find incidents")
	}
	incidents := convertLatestRowMany(rows)
	return incidents, nil
}

// GetByIDFromDB retrieves an incident from the database based on the provided incident ID.
func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Incident, error) {
	record, err := db.GetIncidentByID(r.Context(), id)
	if err != nil {
		return data.Incident{}, errors.New("couldn't find incident")
	}
	incident := convert(record)
	return incident, nil
}

// PostToDB inserts a new incident into the database.
func PostToDB(r *http.Request, db *database.Queries, incident data.Incident) (data.Incident, error) {
	i, err := db.CreateIncident(r.Context(), database.CreateIncidentParams{
		ID:                  incident.ID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description,
		State:               incident.State,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	})
	response := convert(i)
	if err != nil {
		return data.Incident{}, errors.New("couldn't find incident")
	}
	return response, nil
}

// PutToDB updates an existing incident in the database.
func PutToDB(r *http.Request, db *database.Queries, incident data.Incident) (data.Incident, error) {
	i, err := db.UpdateIncident(r.Context(), database.UpdateIncidentParams{
		ID:                  incident.ID,
		UpdatedAt:           time.Now(),
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description,
		State:               incident.State,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
		AssignedTo:          incident.AssignedToID,
	})
	if err != nil {
		return data.Incident{}, errors.New("couldn't update incident")
	}

	response := convert(i)

	return response, nil
}

// DeleteFromDB deletes an incident from the database based on the provided incident ID.
func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Incident, error) {
	i, err := db.DeleteIncidentByID(r.Context(), id)
	if err != nil {
		return data.Incident{}, errors.New("couldn't delete incident")
	}

	response := convert(i)

	return response, nil
}
