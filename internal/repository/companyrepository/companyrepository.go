package companyrepository

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

// ListCompanies retrieves companies from the database based on the provided query and filters.
func ListCompanies(logger *slog.Logger, db *database.Queries, ctx context.Context, query string, filters models.Filters) ([]models.Company, models.Metadata, error) {
	params := database.GetCompaniesParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}

	rows, err := db.GetCompanies(ctx, params)
	if err != nil {
		logger.Error("failed to retrieve companies", "error", err)
		return nil, models.Metadata{}, errors.New("failed to retrieve companies")
	}

	companies, metadata, err := convertCompaniesAndMetadata(rows, filters)
	if err != nil {
		return nil, models.Metadata{}, err
	}

	return companies, metadata, nil
}

// CountCompanies retrieves the count of companies from the database based on the provided query.
func CountCompanies(r *http.Request, logger *slog.Logger, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.GetCompaniesCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		logger.Error("failed to count companies", "error", err)
		return 0, errors.New("failed to count companies")
	}
	return count, nil
}

// GetLatestCompanies retrieves the latest companies from the database.
func GetLatestCompanies(r *http.Request, logger *slog.Logger, db *database.Queries, limit, offset int32) ([]models.Company, error) {
	params := database.GetCompaniesLatestParams{
		Limit:  limit,
		Offset: offset,
	}

	rows, err := db.GetCompaniesLatest(r.Context(), params)
	if err != nil {
		logger.Error("failed to retrieve recent companies", "error", err)
		return nil, errors.New("failed to retrieve recent companies")
	}

	return convertManyCompanies(rows), nil
}

// GetCompanyByID retrieves a company from the database based on the provided company ID.
func GetCompanyByID(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.Company, error) {
	record, err := db.GetCompanyByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to retrieve company", "error", err, "company", id)
		return models.Company{}, errors.New("failed to retrieve company")
	}

	return convertCompany(record), nil
}

// CreateCompany creates a new company in the database.
func CreateCompany(r *http.Request, logger *slog.Logger, db *database.Queries, company models.Company) (models.Company, error) {
	params := database.CreateCompanyParams{
		ID:        company.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      company.Name,
	}

	record, err := db.CreateCompany(r.Context(), params)
	if err != nil {
		logger.Error("failed to create company", "error", err)
		return models.Company{}, errors.New("failed to create company")
	}

	return convertCompany(record), nil
}

// UpdateCompany updates an existing company in the database.
func UpdateCompany(r *http.Request, logger *slog.Logger, db *database.Queries, company models.Company) (models.Company, error) {
	params := database.UpdateCompanyParams{
		ID:        company.ID,
		UpdatedAt: time.Now(),
		Name:      company.Name,
	}

	record, err := db.UpdateCompany(r.Context(), params)
	if err != nil {
		logger.Error("failed to update company", "error", err, "id", company.ID)
		return models.Company{}, errors.New("failed to update company")
	}

	return convertCompany(record), nil
}

// DeleteCompany deletes a company from the database based on the provided company ID.
func DeleteCompany(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.Company, error) {
	record, err := db.DeleteCompanyByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to delete company", "error", err, "id", id)
		return models.Company{}, errors.New("failed to delete company")
	}

	return convertCompany(record), nil
}

// convertCompany converts a database.Company to a models.Company.
func convertCompany(dbCompany database.Company) models.Company {
	return models.Company{
		ID:        dbCompany.ID,
		CreatedAt: dbCompany.CreatedAt,
		UpdatedAt: dbCompany.UpdatedAt,
		Name:      dbCompany.Name,
	}
}

// convertManyCompanies transforms a slice of database.Company to a slice of models.Company.
func convertManyCompanies(dbCompanies []database.Company) []models.Company {
	companies := make([]models.Company, len(dbCompanies))
	for i, dbCompany := range dbCompanies {
		companies[i] = convertCompany(dbCompany)
	}
	return companies
}

// convertCompaniesAndMetadata converts a slice of database Company records and filters into a slice of models.Company and models.Metadata.
func convertCompaniesAndMetadata(rows []database.GetCompaniesRow, filters models.Filters) ([]models.Company, models.Metadata, error) {
	if len(rows) == 0 {
		return nil, models.Metadata{}, nil
	}

	// Prevent conversion exploits
	totalRecords, err := models.ConvertInt64to32(rows[0].Count)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to convert total records count: %w", err)
	}

	companies := convertManyGetCompaniesRowToCompanies(rows)
	metadata, err := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to calculate metadata: %w", err)
	}

	return companies, metadata, nil
}

// convertGetCompaniesRowToCompany converts a database row of type GetCompaniesRow to a Company model.
func convertGetCompaniesRowToCompany(row database.GetCompaniesRow) models.Company {
	return models.Company{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Name:      row.Name,
	}
}

// convertManyGetCompaniesRowToCompanies converts a database.GetCompaniesRow to an array of models.Company.
func convertManyGetCompaniesRowToCompanies(rows []database.GetCompaniesRow) []models.Company {
	companies := make([]models.Company, len(rows))
	for i, row := range rows {
		companies[i] = convertGetCompaniesRowToCompany(row)
	}
	return companies
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
