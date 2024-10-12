package incidents

import (
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/errutil"
	"github.com/barturba/ticket-tracker/internal/helpers"
	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// GET

func Get(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Query  string
			Limit  int
			Offset int
			data.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Query = data.ReadString(qs, "query", "")

		input.Filters.Page = data.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 10, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{
			"id", "-id",
			"created_at", "-created_at",
			"updated_at", "-updated_at",
			"short_description", "-short_description",
			"description", "-description",
			"first_name", "-first_name",
			"last_name", "-last_name",
		}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		incidents, metadata, err := GetFromDB(r, db, input.Query, input.Filters)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/incidents")
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"incidents": incidents, "metadata": metadata})
	})
}

func GetLatest(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Limit  int
			Offset int
			data.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Filters.Page = data.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 20, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		i, err := GetLatestFromDB(r, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/incidents_latest")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func GetByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		count, err := db.GetIncidentByID(r.Context(), id)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/incident/{id}")
		helpers.RespondWithJSON(w, http.StatusOK, count)
	})
}

// POST

func Post(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ShortDescription    string
			Description         string
			ConfigurationItemID uuid.UUID
			CompanyID           uuid.UUID
			AssignedToID        uuid.UUID
			State               database.StateEnum
		}

		err := helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		incident := &data.Incident{
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

		if data.ValidateIncident(v, incident); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PostToDB(r, db, *incident)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "POST /v1/incident")
		helpers.RespondWithJSON(w, http.StatusCreated, i)
	})
}

// PUT

func Put(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		var input struct {
			CreatedAt           time.Time
			UpdatedAt           time.Time
			ShortDescription    string
			Description         string
			ConfigurationItemID uuid.UUID
			CompanyID           uuid.UUID
			AssignedToID        uuid.UUID
			State               database.StateEnum
		}

		err = helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		incident := &data.Incident{
			ID:                  id,
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

		if data.ValidateIncident(v, incident); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PutToDB(r, db, *incident)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "PUT /v1/incident", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

// DELETE

func Delete(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		_, err = DeleteFromDB(r, db, id)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "DELETE /v1/incident", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"message": "incident successfully deleted"})
	})
}
