package handlers

import (
	"database/sql"
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

// ListIncidents retrieves a list of incidents with optional filtering, sorting, and pagination.
func ListIncidents(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()
		input := parseIncidentFilters(r, v)

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		incidents, metadata, err := repository.ListIncidents(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/incidents")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"incidents": incidents, "metadata": metadata})
	})
}

// ListAllIncidents retrieves all incidents with optional filtering, sorting, and pagination.
func ListAllIncidents(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()

		input := parseUserFilters(r, v)
		input.Filters.PageSize = models.MaxPageSize

		if models.ValidateFiltersGetAll(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		incidents, metadata, err := repository.ListIncidents(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/incidents_all")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"incidents": incidents, "metadata": metadata})
	})
}

// ListRecentIncidents retrieves the latest incidents.
func ListRecentIncidents(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()

		input := parseUserFilters(r, v)
		input.Filters.PageSize = models.DefaultPageSize

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		latestIncidents, err := repository.ListRecentIncidents(r, logger, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/incidents_latest")
		json.RespondWithJSON(w, http.StatusOK, latestIncidents)
	})
}

// GetIncidentByID retrieves a single incident by their unique identifier.
func GetIncidentByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			errors.NotFoundResponse(w, r, logger)
			return
		}

		incident, err := repository.GetIncident(r, logger, db, id)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/incidents/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, incident)
	})
}

// CreateIncident creates a new incident.
func CreateIncident(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ShortDescription    string             `json:"short_description"`
			Description         string             `json:"description"`
			ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
			CompanyID           uuid.UUID          `json:"company_id"`
			AssignedToID        uuid.UUID          `json:"assigned_to_id"`
			State               database.StateEnum `json:"state"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			errors.BadRequestResponse(w, r, logger, err)
			return
		}

		incident := &models.Incident{
			ID:                  uuid.New(),
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			ShortDescription:    input.ShortDescription,
			Description:         sql.NullString{String: input.Description, Valid: input.Description != ""},
			ConfigurationItemID: input.ConfigurationItemID,
			CompanyID:           input.CompanyID,
			AssignedToID:        uuid.NullUUID{UUID: input.AssignedToID, Valid: input.AssignedToID != uuid.Nil},
			State:               input.State,
		}

		v := validator.New()
		if models.ValidateIncident(v, incident); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		createdIncident, err := repository.CreateIncident(r, logger, db, *incident)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled POST /v1/incidents")
		json.RespondWithJSON(w, http.StatusCreated, createdIncident)
	})
}

// UpdateIncident updates an existing incident by their unique identifier.
func UpdateIncident(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			errors.NotFoundResponse(w, r, logger)
			return
		}

		var input struct {
			ShortDescription    string             `json:"short_description"`
			Description         string             `json:"description"`
			ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
			CompanyID           uuid.UUID          `json:"company_id"`
			AssignedToID        uuid.UUID          `json:"assigned_to_id"`
			State               database.StateEnum `json:"state"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			errors.BadRequestResponse(w, r, logger, err)
			return
		}

		incident := &models.Incident{
			ID:                  id,
			UpdatedAt:           time.Now(),
			ShortDescription:    input.ShortDescription,
			Description:         sql.NullString{String: input.Description, Valid: input.Description != ""},
			ConfigurationItemID: input.ConfigurationItemID,
			CompanyID:           input.CompanyID,
			AssignedToID:        uuid.NullUUID{UUID: input.AssignedToID, Valid: input.AssignedToID != uuid.Nil},
			State:               input.State,
		}

		v := validator.New()
		if models.ValidateIncident(v, incident); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		updatedIncident, err := repository.UpdateIncident(r, logger, db, *incident)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled PUT /v1/incident/{id}")
		json.RespondWithJSON(w, http.StatusOK, updatedIncident)
	})
}

// DeleteIncident deletes an existing incident by their unique identifier.
func DeleteIncident(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			errors.NotFoundResponse(w, r, logger)
			return
		}

		if _, err = repository.DeleteIncident(r, logger, db, id); err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled DELETE /v1/incidents/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"message": "incident successfully deleted"})
	})
}

func parseIncidentFilters(r *http.Request, v *validator.Validator) struct {
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
				"short_description", "-short_description",
				"description", "-description",
				"first_name", "-first_name",
				"last_name", "-last_name",
			},
		},
	}
}
