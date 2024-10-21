package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/google/uuid"
)

// ListIncidents retrieves incidents from the database based on the provided query and filters.
func ListIncidents(logger *slog.Logger, db *database.Queries, ctx context.Context, query string, filters models.Filters) ([]models.Incident, models.Metadata, error) {
	params := database.ListIncidentsParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}

	rows, err := db.ListIncidents(ctx, params)
	if err != nil {
		logger.Error("failed to retrieve incidents", "error", err)
		return nil, models.Metadata{}, errors.New("failed to retrieve incidents")
	}

	incidents, metadata, err := convertIncidentsAndMetadata(rows, filters)
	if err != nil {
		return nil, models.Metadata{}, err
	}

	return incidents, metadata, nil
}

// CountIncidents retrieves the count of incidents from the database based on the provided query.
func CountIncidents(r *http.Request, logger *slog.Logger, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.CountIncidents(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		logger.Error("failed to count incidents", "error", err)
		return 0, errors.New("failed to count incidents")
	}
	return count, nil
}

// ListRecentIncidents retrieves the latest incidents from the database.
func ListRecentIncidents(r *http.Request, logger *slog.Logger, db *database.Queries, limit, offset int32) ([]models.Incident, error) {
	params := database.ListRecentIncidentsParams{
		Limit:  limit,
		Offset: offset,
	}

	rows, err := db.ListRecentIncidents(r.Context(), params)
	if err != nil {
		logger.Error("failed to retrieve recent incidents", "error", err)
		return nil, errors.New("failed to retrieve recent incidents")
	}

	return convertManyListIncidentsLatestRowToIncidents(rows), nil
}

// GetIncident retrieves a incident from the database based on the provided incident ID.
func GetIncident(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.Incident, error) {
	record, err := db.GetIncident(r.Context(), id)
	if err != nil {
		logger.Error("failed to retrieve incident", "error", err, "incident", id)
		return models.Incident{}, errors.New("failed to retrieve incident")
	}

	return convertIncidentRow(record), nil
}

// CreateIncident creates a new incident in the database.
func CreateIncident(r *http.Request, logger *slog.Logger, db *database.Queries, incident models.Incident) (models.Incident, error) {
	params := database.CreateIncidentParams{
		ID:                  incident.ID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description,
		State:               incident.State,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	}

	record, err := db.CreateIncident(r.Context(), params)
	if err != nil {
		logger.Error("failed to create incident", "error", err)
		return models.Incident{}, errors.New("failed to create incident")
	}

	return convertIncident(record), nil
}

// UpdateIncident updates an existing incident in the database.
func UpdateIncident(r *http.Request, logger *slog.Logger, db *database.Queries, incident models.Incident) (models.Incident, error) {
	params := database.UpdateIncidentParams{
		ID:                  incident.ID,
		UpdatedAt:           time.Now(),
		CompanyID:           incident.CompanyID,
		ConfigurationItemID: incident.ConfigurationItemID,
		Description:         incident.Description,
		ShortDescription:    incident.ShortDescription,
		State:               incident.State,
		AssignedTo:          incident.AssignedToID,
	}

	record, err := db.UpdateIncident(r.Context(), params)
	if err != nil {
		logger.Error("failed to update incident", "error", err, "id", incident.ID)
		return models.Incident{}, errors.New("failed to update incident")
	}

	return convertIncident(record), nil
}

// DeleteIncident deletes a incident from the database based on the provided incident ID.
func DeleteIncident(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.Incident, error) {
	record, err := db.DeleteIncident(r.Context(), id)
	if err != nil {
		logger.Error("failed to delete incident", "error", err, "id", id)
		return models.Incident{}, errors.New("failed to delete incident")
	}

	return convertIncident(record), nil
}

// convertIncident converts a database.Incident to a models.Incident.
func convertIncident(dbIncident database.Incident) models.Incident {
	return models.Incident{
		ID:                  dbIncident.ID,
		CreatedAt:           dbIncident.CreatedAt,
		UpdatedAt:           dbIncident.UpdatedAt,
		ShortDescription:    dbIncident.ShortDescription,
		Description:         dbIncident.Description,
		ConfigurationItemID: dbIncident.ConfigurationItemID,
		CompanyID:           dbIncident.CompanyID,
		AssignedToID:        dbIncident.AssignedTo,
		State:               dbIncident.State,
	}
}

// convertIncidentRow converts a database.GetIncidentRow to a models.Incident.
func convertIncidentRow(dbIncident database.GetIncidentRow) models.Incident {
	return models.Incident{
		ID:                  dbIncident.ID,
		CreatedAt:           dbIncident.CreatedAt,
		UpdatedAt:           dbIncident.UpdatedAt,
		ShortDescription:    dbIncident.ShortDescription,
		Description:         dbIncident.Description,
		ConfigurationItemID: dbIncident.ConfigurationItemID,
		CompanyID:           dbIncident.CompanyID,
		AssignedToID:        dbIncident.AssignedTo,
		AssignedToName:      fmt.Sprintf("%s %s", dbIncident.FirstName.String, dbIncident.LastName.String),
		State:               dbIncident.State,
	}
}

// convertIncidentsAndMetadata converts a slice of database Incident records and filters into a slice of models.Incident and models.Metadata.
func convertIncidentsAndMetadata(rows []database.ListIncidentsRow, filters models.Filters) ([]models.Incident, models.Metadata, error) {
	if len(rows) == 0 {
		return nil, models.Metadata{}, nil
	}

	// Prevent conversion exploits
	totalRecords, err := models.ConvertInt64to32(rows[0].Count)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to convert total records count: %w", err)
	}

	incidents := convertManyListIncidentsRowToIncidents(rows)
	metadata, err := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to calculate metadata: %w", err)
	}

	return incidents, metadata, nil
}

// convertListIncidentsRowToIncident converts a database row of type ListIncidentsRow to a Incident model.
func convertListIncidentsRowToIncident(row database.ListIncidentsRow) models.Incident {
	return models.Incident{
		ID:                  row.ID,
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
		ShortDescription:    row.ShortDescription,
		Description:         row.Description,
		ConfigurationItemID: row.ConfigurationItemID,
		CompanyID:           row.CompanyID,
		AssignedToID:        row.AssignedTo,
		AssignedToName:      row.FirstName.String + " " + row.LastName.String,
		State:               row.State,
	}
}

// convertManyListIncidentsRowToIncidents converts a database.ListIncidentsRow to an array of models.Incident.
func convertManyListIncidentsRowToIncidents(rows []database.ListIncidentsRow) []models.Incident {
	incidents := make([]models.Incident, len(rows))
	for i, row := range rows {
		incidents[i] = convertListIncidentsRowToIncident(row)
	}
	return incidents
}

// convertListIncidentsLatestRowToIncident converts a database row of type ListIncidentsLatestRow to a Incident model.
func convertListIncidentsLatestRowToIncident(row database.ListRecentIncidentsRow) models.Incident {
	return models.Incident{
		ID:                  row.ID,
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
		ShortDescription:    row.ShortDescription,
		Description:         row.Description,
		ConfigurationItemID: row.ConfigurationItemID,
		CompanyID:           row.CompanyID,
		AssignedToID:        row.AssignedTo,
		State:               row.State,
	}
}

// convertManyListIncidentsLatestRowToIncidents converts a database.ListIncidentsLatestRow to an array of models.Incident.
func convertManyListIncidentsLatestRowToIncidents(rows []database.ListRecentIncidentsRow) []models.Incident {
	incidents := make([]models.Incident, len(rows))
	for i, row := range rows {
		incidents[i] = convertListIncidentsLatestRowToIncident(row)
	}
	return incidents
}
