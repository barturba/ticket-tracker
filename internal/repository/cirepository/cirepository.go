package cirepository

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

// ListCIs retrieves cis from the database based on the provided query and filters.
func ListCIs(logger *slog.Logger, db *database.Queries, ctx context.Context, query string, filters models.Filters) ([]models.CI, models.Metadata, error) {
	params := database.GetCIsParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}

	rows, err := db.GetCIs(ctx, params)
	if err != nil {
		logger.Error("failed to retrieve cis", "error", err)
		return nil, models.Metadata{}, errors.New("failed to retrieve cis")
	}

	cis, metadata, err := convertCIsAndMetadata(rows, filters)
	if err != nil {
		return nil, models.Metadata{}, err
	}

	return cis, metadata, nil
}

// CountCIs retrieves the count of cis from the database based on the provided query.
func CountCIs(r *http.Request, logger *slog.Logger, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.GetCIsCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		logger.Error("failed to count cis", "error", err)
		return 0, errors.New("failed to count cis")
	}
	return count, nil
}

// GetLatestCIs retrieves the latest cis from the database.
func GetLatestCIs(r *http.Request, logger *slog.Logger, db *database.Queries, limit, offset int32) ([]models.CI, error) {
	params := database.GetCIsLatestParams{
		Limit:  limit,
		Offset: offset,
	}

	rows, err := db.GetCIsLatest(r.Context(), params)
	if err != nil {
		logger.Error("failed to retrieve recent cis", "error", err)
		return nil, errors.New("failed to retrieve recent cis")
	}

	return convertManyCIs(rows), nil
}

// GetCIByID retrieves a ci from the database based on the provided ci ID.
func GetCIByID(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.CI, error) {
	record, err := db.GetCIsByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to retrieve ci", "error", err, "ci", id)
		return models.CI{}, errors.New("failed to retrieve ci")
	}

	return convertCI(record), nil
}

// CreateCI creates a new ci in the database.
func CreateCI(r *http.Request, logger *slog.Logger, db *database.Queries, ci models.CI) (models.CI, error) {
	params := database.CreateCIsParams{
		ID:        ci.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	record, err := db.CreateCIs(r.Context(), params)
	if err != nil {
		logger.Error("failed to create ci", "error", err)
		return models.CI{}, errors.New("failed to create ci")
	}

	return convertCI(record), nil
}

// UpdateCI updates an existing ci in the database.
func UpdateCI(r *http.Request, logger *slog.Logger, db *database.Queries, ci models.CI) (models.CI, error) {
	params := database.UpdateCIsParams{
		ID:        ci.ID,
		UpdatedAt: time.Now(),
		Name:      ci.Name,
	}

	record, err := db.UpdateCIs(r.Context(), params)
	if err != nil {
		logger.Error("failed to update ci", "error", err, "id", ci.ID)
		return models.CI{}, errors.New("failed to update ci")
	}

	return convertCI(record), nil
}

// DeleteCI deletes a ci from the database based on the provided ci ID.
func DeleteCI(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.CI, error) {
	record, err := db.DeleteCIs(r.Context(), id)
	if err != nil {
		logger.Error("failed to delete ci", "error", err, "id", id)
		return models.CI{}, errors.New("failed to delete ci")
	}

	return convertCI(record), nil
}

// convertCI converts a database.ConfigurationItem to a models.CI.
func convertCI(dbCI database.ConfigurationItem) models.CI {
	return models.CI{
		ID:        dbCI.ID,
		CreatedAt: dbCI.CreatedAt,
		UpdatedAt: dbCI.UpdatedAt,
		Name:      dbCI.Name,
	}
}

// convertManyCIs transforms a slice of database.ConfigurationItem to a slice of models.CI.
func convertManyCIs(dbCIs []database.ConfigurationItem) []models.CI {
	cis := make([]models.CI, len(dbCIs))
	for i, dbCI := range dbCIs {
		cis[i] = convertCI(dbCI)
	}
	return cis
}

// convertCIsAndMetadata converts a slice of database CI records and filters into a slice of models.CI and models.Metadata.
func convertCIsAndMetadata(rows []database.GetCIsRow, filters models.Filters) ([]models.CI, models.Metadata, error) {
	if len(rows) == 0 {
		return nil, models.Metadata{}, nil
	}

	// Prevent conversion exploits
	totalRecords, err := models.ConvertInt64to32(rows[0].Count)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to convert total records count: %w", err)
	}

	cis := convertManyGetCIsRowToCIs(rows)
	metadata, err := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to calculate metadata: %w", err)
	}

	return cis, metadata, nil
}

// convertGetCIsRowToCI converts a database row of type GetCIsRow to a CI model.
func convertGetCIsRowToCI(row database.GetCIsRow) models.CI {
	return models.CI{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Name:      row.Name,
	}
}

// convertManyGetCIsRowToCIs converts a database.GetCIsRow to an array of models.CI.
func convertManyGetCIsRowToCIs(rows []database.GetCIsRow) []models.CI {
	cis := make([]models.CI, len(rows))
	for i, row := range rows {
		cis[i] = convertGetCIsRowToCI(row)
	}
	return cis
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
