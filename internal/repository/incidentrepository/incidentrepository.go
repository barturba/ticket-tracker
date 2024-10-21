package incidentrepository

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
	params := database.GetIncidentsParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}

	rows, err := db.GetIncidents(ctx, params)
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
	count, err := db.GetIncidentsCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		logger.Error("failed to count incidents", "error", err)
		return 0, errors.New("failed to count incidents")
	}
	return count, nil
}

// GetLatestIncidents retrieves the latest incidents from the database.
func GetLatestIncidents(r *http.Request, logger *slog.Logger, db *database.Queries, limit, offset int32) ([]models.Incident, error) {
	params := database.GetIncidentsLatestParams{
		Limit:  limit,
		Offset: offset,
	}

	rows, err := db.GetIncidentsLatest(r.Context(), params)
	if err != nil {
		logger.Error("failed to retrieve recent incidents", "error", err)
		return nil, errors.New("failed to retrieve recent incidents")
	}

	return convertManyGetIncidentsLatestRowToIncidents(rows), nil
}

// GetIncidentByID retrieves a incident from the database based on the provided incident ID.
func GetIncidentByID(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.Incident, error) {
	record, err := db.GetIncidentByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to retrieve incident", "error", err, "incident", id)
		return models.Incident{}, errors.New("failed to retrieve incident")
	}

	return convertIncident(record), nil
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
	record, err := db.DeleteIncidentByID(r.Context(), id)
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

// convertIncidentsAndMetadata converts a slice of database Incident records and filters into a slice of models.Incident and models.Metadata.
func convertIncidentsAndMetadata(rows []database.GetIncidentsRow, filters models.Filters) ([]models.Incident, models.Metadata, error) {
	if len(rows) == 0 {
		return nil, models.Metadata{}, nil
	}

	// Prevent conversion exploits
	totalRecords, err := models.ConvertInt64to32(rows[0].Count)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to convert total records count: %w", err)
	}

	incidents := convertManyGetIncidentsRowToIncidents(rows)
	metadata, err := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to calculate metadata: %w", err)
	}

	return incidents, metadata, nil
}

// convertGetIncidentsRowToIncident converts a database row of type GetIncidentsRow to a Incident model.
func convertGetIncidentsRowToIncident(row database.GetIncidentsRow) models.Incident {
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

// convertManyGetIncidentsRowToIncidents converts a database.GetIncidentsRow to an array of models.Incident.
func convertManyGetIncidentsRowToIncidents(rows []database.GetIncidentsRow) []models.Incident {
	incidents := make([]models.Incident, len(rows))
	for i, row := range rows {
		incidents[i] = convertGetIncidentsRowToIncident(row)
	}
	return incidents
}

// convertGetIncidentsLatestRowToIncident converts a database row of type GetIncidentsLatestRow to a Incident model.
func convertGetIncidentsLatestRowToIncident(row database.GetIncidentsLatestRow) models.Incident {
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

// convertManyGetIncidentsLatestRowToIncidents converts a database.GetIncidentsLatestRow to an array of models.Incident.
func convertManyGetIncidentsLatestRowToIncidents(rows []database.GetIncidentsLatestRow) []models.Incident {
	incidents := make([]models.Incident, len(rows))
	for i, row := range rows {
		incidents[i] = convertGetIncidentsLatestRowToIncident(row)
	}
	return incidents
}

// calculateMetadata creates a models.Metadata struct based on the total
// records, current page, and page size.
func calculateMetadata(totalRecords, page, pageSize int32) (models.Metadata, error) {
	if totalRecords < 0 || page < 1 || pageSize < 1 {
		return models.Metadata{}, fmt.Errorf("invalid metadata parameters")
	}

	lastPage, err := models.SafeDivide(totalRecords, pageSize)
	if err != nil {
		return models.Metadata{}, fmt.Errorf("failed to calculate the last page: %w", err)
	}

	return models.Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     lastPage,
		TotalRecords: totalRecords,
	}, nil
}
