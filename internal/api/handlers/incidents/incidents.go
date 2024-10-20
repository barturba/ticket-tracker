package incidenthandlers

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	incidentrepository "github.com/barturba/ticket-tracker/internal/repository/incidents"
	"github.com/barturba/ticket-tracker/internal/utils/httperrors"
	"github.com/barturba/ticket-tracker/internal/utils/json"
	"github.com/barturba/ticket-tracker/internal/utils/validator"
	"github.com/google/uuid"
)

func Get(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Query  string
			Limit  int
			Offset int
			models.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Query = models.ReadString(qs, "query", "")

		input.Filters.Page = models.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = models.ReadInt(qs, "page_size", 10, v)

		input.Filters.Sort = models.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{
			"id", "-id",
			"created_at", "-created_at",
			"updated_at", "-updated_at",
			"short_description", "-short_description",
			"description", "-description",
			"first_name", "-first_name",
			"last_name", "-last_name",
		}

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		incidents, metadata, err := incidentrepository.GetFromDB(r, db, input.Query, input.Filters)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/incidents")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"incidents": incidents, "metadata": metadata})
	})
}

func GetAll(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Query  string
			Limit  int
			Offset int
			models.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Query = models.ReadString(qs, "query", "")

		input.Filters.Page = models.ReadInt(qs, "page", 1, v)

		// Set the page size to a large value
		input.Filters.PageSize = models.ReadInt(qs, "page_size", 10_000_000, v)

		input.Filters.Sort = models.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{
			"id", "-id",
			"created_at", "-created_at",
			"updated_at", "-updated_at",
			"short_description", "-short_description",
			"description", "-description",
			"first_name", "-first_name",
			"last_name", "-last_name",
		}

		// Ignore the usual page size warnings since we're trying to get all values
		if models.ValidateFiltersGetAll(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		incidents, metadata, err := incidentrepository.GetFromDB(r, db, input.Query, input.Filters)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/incidents-all")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"incidents": incidents, "metadata": metadata})
	})
}

func GetLatest(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Limit  int
			Offset int
			models.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Filters.Page = models.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = models.ReadInt(qs, "page_size", 20, v)

		input.Filters.Sort = models.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id"}

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		i, err := incidentrepository.GetLatestFromDB(r, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/incidents_latest")
		json.RespondWithJSON(w, http.StatusOK, i)
	})
}

func GetByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}
		i, err := incidentrepository.GetByIDFromDB(r, db, id)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", fmt.Sprintf("GET /v1/incidents/%s", id))
		json.RespondWithJSON(w, http.StatusOK, i)
	})
}

func Post(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ShortDescription    string             `json:"short_description"`
			Description         string             `json:"description"`
			ConfigurationItemID uuid.UUID          `json:"configuration_item_id"`
			CompanyID           uuid.UUID          `json:"company_id"`
			AssignedToID        uuid.UUID          `json:"assigned_to_id"`
			State               database.StateEnum `json:"state"`
		}

		err := json.ReadJSON(w, r, &input)
		if err != nil {
			httperrors.BadRequestResponse(w, r, logger, err)
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
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := incidentrepository.PostToDB(r, db, *incident)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "POST /v1/incident")
		json.RespondWithJSON(w, http.StatusCreated, i)
	})
}

func Put(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
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

		err = json.ReadJSON(w, r, &input)
		if err != nil {
			httperrors.BadRequestResponse(w, r, logger, err)
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
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := incidentrepository.PutToDB(r, db, *incident)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "PUT /v1/incident", "id", id)
		json.RespondWithJSON(w, http.StatusOK, i)
	})
}

func Delete(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		_, err = incidentrepository.DeleteFromDB(r, db, id)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "DELETE /v1/incident", "id", id)
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"message": "incident successfully deleted"})
	})
}
