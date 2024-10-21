package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/repository"
	"github.com/barturba/ticket-tracker/internal/utils/errors"
	"github.com/barturba/ticket-tracker/internal/utils/json"
	"github.com/barturba/ticket-tracker/internal/utils/validator"
	"github.com/google/uuid"
)

// ListCIs retrieves a list of cis with optional filtering, sorting, and pagination.
func ListCIs(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()
		input := parseCIFilters(r, v)

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		cis, metadata, err := repository.ListCIs(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/cis")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"cis": cis, "metadata": metadata})
	})
}

// ListAllCIs retrieves all cis with optional filtering, sorting, and pagination.
func ListAllCIs(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()

		input := parseCIFilters(r, v)
		input.Filters.PageSize = 10_000_000

		if models.ValidateFiltersGetAll(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		cis, metadata, err := repository.ListCIs(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/cis_all")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"cis": cis, "metadata": metadata})
	})
}

// ListRecentCIs retrieves the latest cis.
func ListRecentCIs(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()

		input := parseCIFilters(r, v)
		input.Filters.PageSize = 20

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		latestCIs, err := repository.ListRecentCI(r, logger, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/cis_latest")
		json.RespondWithJSON(w, http.StatusOK, latestCIs)
	})
}

// GetCI retrieves a single ci by their unique identifier.
func GetCI(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			errors.NotFoundResponse(w, r, logger)
			return
		}

		ci, err := repository.GetCI(r, logger, db, id)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/cis/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, ci)
	})
}

// CreateCI creates a new ci.
func CreateCI(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name string `json:"name"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			errors.BadRequestResponse(w, r, logger, err)
			return
		}

		ci := &models.CI{
			ID:        uuid.New(),
			UpdatedAt: time.Now(),
			Name:      input.Name,
		}

		v := validator.New()
		if models.ValidateCI(v, ci); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		createdCI, err := repository.CreateCI(r, logger, db, *ci)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled POST /v1/cis")
		json.RespondWithJSON(w, http.StatusCreated, createdCI)
	})
}

// UpdateCI updates an existing ci by their unique identifier.
func UpdateCI(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			errors.NotFoundResponse(w, r, logger)
			return
		}

		var input struct {
			Name string `json:"name"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			errors.BadRequestResponse(w, r, logger, err)
			return
		}

		ci := &models.CI{
			ID:        id,
			UpdatedAt: time.Now(),
			Name:      input.Name,
		}

		v := validator.New()
		if models.ValidateCI(v, ci); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		updatedCI, err := repository.UpdateCI(r, logger, db, *ci)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled PUT /v1/ci/{id}")
		json.RespondWithJSON(w, http.StatusOK, updatedCI)
	})
}

// DeleteCI deletes an existing ci by their unique identifier.
func DeleteCI(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			errors.NotFoundResponse(w, r, logger)
			return
		}

		if _, err = repository.DeleteCI(r, logger, db, id); err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled DELETE /v1/cis/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"message": "ci successfully deleted"})
	})
}

func parseCIFilters(r *http.Request, v *validator.Validator) struct {
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
				"-name", "name",
			},
		},
	}
}
