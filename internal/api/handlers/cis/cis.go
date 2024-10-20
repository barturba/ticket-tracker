// Package cis provides HTTP handlers for managing Configuration Items (CIs).
package cis

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/json"
	"github.com/barturba/ticket-tracker/internal/utils/httperrors"
	"github.com/barturba/ticket-tracker/pkg/validator"
	"github.com/google/uuid"
)

// Get handles GET requests to retrieve a list of CIs based on query parameters and filters.
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
		input.Filters.SortSafelist = []string{"id", "-id",
			"created_at", "-created_at",
			"updated_at", "-updated_at",
			"name", "-name"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		cis, metadata, err := GetFromDB(r, db, input.Query, input.Filters)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/cis")
		json.RespondWithJSON(w, http.StatusOK, data.Envelope{"cis": cis, "metadata": metadata})
	})
}

// GetAll handles GET requests to retrieve all CIs with a large page size limit.
func GetAll(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Query  string
			Limit  int32
			Offset int32
			data.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Query = data.ReadString(qs, "query", "")

		input.Filters.Page = data.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 10_000_000, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id", "-id",
			"created_at", "-created_at",
			"updated_at", "-updated_at",
			"name", "-name"}

		if data.ValidateFiltersGetAll(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		cis, metadata, err := GetFromDB(r, db, input.Query, input.Filters)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/cis")
		json.RespondWithJSON(w, http.StatusOK, data.Envelope{"cis": cis, "metadata": metadata})
	})
}

// GetLatest handles GET requests to retrieve the latest CIs based on pagination and sorting.
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
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		i, err := GetLatestFromDB(r, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/cis_latest")
		json.RespondWithJSON(w, http.StatusOK, i)
	})
}

// GetByID handles GET requests to retrieve a specific CI by its UUID.
func GetByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		i, err := GetByIDFromDB(r, db, id)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", fmt.Sprintf("GET /v1/cis/%s", id))
		json.RespondWithJSON(w, http.StatusOK, i)
	})
}

// Post handles POST requests to create a new CI.
func Post(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name string `json:"Name"`
		}

		err := json.ReadJSON(w, r, &input)
		if err != nil {
			httperrors.BadRequestResponse(w, r, logger, err)
			return
		}

		ci := &data.CI{
			ID:   uuid.New(),
			Name: input.Name,
		}

		v := validator.New()

		if data.ValidateCI(v, ci); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PostToDB(r, db, *ci)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "POST /v1/ci")
		json.RespondWithJSON(w, http.StatusCreated, i)
	})
}

// Put handles PUT requests to update an existing CI by its UUID.
func Put(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		var input struct {
			Name string `json:"name"`
		}

		err = json.ReadJSON(w, r, &input)
		if err != nil {
			httperrors.BadRequestResponse(w, r, logger, err)
			return
		}

		ci := &data.CI{
			ID:        id,
			UpdatedAt: time.Now(),
			Name:      input.Name,
		}

		v := validator.New()

		if data.ValidateCI(v, ci); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PutToDB(r, db, *ci)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "PUT /v1/cis", "id", id)
		json.RespondWithJSON(w, http.StatusOK, i)
	})
}

// Delete handles DELETE requests to remove a CI by its UUID.
func Delete(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		_, err = DeleteFromDB(r, db, id)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "DELETE /v1/cis", "id", id)
		json.RespondWithJSON(w, http.StatusOK, data.Envelope{"message": "ci successfully deleted"})
	})
}
