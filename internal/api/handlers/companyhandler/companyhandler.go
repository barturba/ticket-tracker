// Package companyhandler provides functions for managing company resources.
package companyhandler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/repository/companyrepository"
	"github.com/barturba/ticket-tracker/internal/utils/httperrors"
	"github.com/barturba/ticket-tracker/internal/utils/json"
	"github.com/barturba/ticket-tracker/internal/utils/validator"
	"github.com/google/uuid"
)

// ListCompanies retrieves a list of companies with optional filtering, sorting, and pagination.
func ListCompanies(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()
		input := parseFilters(r, v)

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		companies, metadata, err := companyrepository.ListCompanies(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/companies")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"companies": companies, "metadata": metadata})
	})
}

// ListAllCompanies retrieves all companies with optional filtering, sorting, and pagination.
func ListAllCompanies(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()

		input := parseFilters(r, v)
		input.Filters.PageSize = 10_000_000

		if models.ValidateFiltersGetAll(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		companies, metadata, err := companyrepository.ListCompanies(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/companies_all")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"companies": companies, "metadata": metadata})
	})
}

// ListRecentCompanies retrieves the latest companies.
func ListRecentCompanies(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()

		input := parseFilters(r, v)
		input.Filters.PageSize = 20

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		latestCompanies, err := companyrepository.GetLatestCompanies(r, logger, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/companies_latest")
		json.RespondWithJSON(w, http.StatusOK, latestCompanies)
	})
}

// GetCompanyByID retrieves a single company by their unique identifier.
func GetCompanyByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		company, err := companyrepository.GetCompanyByID(r, logger, db, id)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/companies/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, company)
	})
}

// CreateCompany creates a new company.
func CreateCompany(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name string `json:"name"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			httperrors.BadRequestResponse(w, r, logger, err)
			return
		}

		company := &models.Company{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      input.Name,
		}

		v := validator.New()
		if models.ValidateCompany(v, company); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		createdCompany, err := companyrepository.CreateCompany(r, logger, db, *company)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled POST /v1/companies")
		json.RespondWithJSON(w, http.StatusCreated, createdCompany)
	})
}

// UpdateCompany updates an existing company by their unique identifier.
func UpdateCompany(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		var input struct {
			Name string `json:"name"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			httperrors.BadRequestResponse(w, r, logger, err)
			return
		}

		company := &models.Company{
			ID:        id,
			UpdatedAt: time.Now(),
			Name:      input.Name,
		}

		v := validator.New()
		if models.ValidateCompany(v, company); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		updatedCompany, err := companyrepository.UpdateCompany(r, logger, db, *company)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled PUT /v1/company/{id}")
		json.RespondWithJSON(w, http.StatusOK, updatedCompany)
	})
}

// DeleteCompany deletes an existing company by their unique identifier.
func DeleteCompany(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		if _, err = companyrepository.DeleteCompany(r, logger, db, id); err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled DELETE /v1/companies/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"message": "company successfully deleted"})
	})
}

func parseFilters(r *http.Request, v *validator.Validator) struct {
	Query   string
	Filters models.Filters
} {
	qs := r.URL.Query()
	return struct {
		Query   string
		Filters models.Filters
	}{
		Query: models.ReadString(qs, "query", ""),
		Filters: models.Filters{
			Page:     models.ReadInt(qs, "page", 1, v),
			PageSize: models.ReadInt(qs, "page_size", 10, v),
			Sort:     models.ReadString(qs, "sort", "id"),
			SortSafelist: []string{
				"id", "-id",
				"created_at", "-created_at",
				"updated_at", "-updated_at",
				"name", "-name",
			},
		},
	}
}
