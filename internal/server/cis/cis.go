package cis

import (
	"database/sql"
	"errors"
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
		input.Filters.SortSafelist = []string{"id"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		i, err := GetFromDB(r, db, input.Query, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/cis")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func GetFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) ([]data.CI, error) {
	p := database.GetConfigurationItemsParams{
		Query:  sql.NullString{String: query, Valid: query != ""},
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetConfigurationItems(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find cis")
	}
	cis := convertMany(rows)
	return cis, nil
}

func GetCount(logger *slog.Logger, db *database.Queries) http.Handler {
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
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 20, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		count, err := GetCountFromDB(r, db, input.Query, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/cis_count")
		helpers.RespondWithJSON(w, http.StatusOK, count)
	})
}

func GetCountFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.GetConfigurationItemsCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		return 0, errors.New("couldn't find cis")
	}
	return count, nil
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
		logger.Info("msg", "handle", "GET /v1/cis_latest")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int) ([]data.CI, error) {
	p := database.GetConfigurationItemsLatestParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetConfigurationItemsLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find cis")
	}
	cis := convertMany(rows)
	return cis, nil
}

func GetByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		count, err := db.GetConfigurationItemsByID(r.Context(), id)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/cis/{id}")
		helpers.RespondWithJSON(w, http.StatusOK, count)
	})
}

func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.CI, error) {
	record, err := db.GetConfigurationItemsByID(r.Context(), id)
	if err != nil {
		return data.CI{}, errors.New("couldn't find ci")
	}
	ci := convert(record)
	return ci, nil
}

// POST

func Post(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name string
		}

		err := helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		ci := &data.CI{
			ID:   uuid.New(),
			Name: input.Name,
		}

		v := validator.New()

		if data.ValidateCI(v, ci); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PostToDB(r, db, *ci)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "POST /v1/ci")
		helpers.RespondWithJSON(w, http.StatusCreated, i)
	})
}

func PostToDB(r *http.Request, db *database.Queries, ci data.CI) (data.CI, error) {
	i, err := db.CreateConfigurationItems(r.Context(), database.CreateConfigurationItemsParams{
		ID:        ci.ID,
		Name:      ci.Name,
		UpdatedAt: time.Now(),
	})
	response := convert(i)
	if err != nil {
		return data.CI{}, errors.New("couldn't find ci")
	}
	return response, nil
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
			UpdatedAt time.Time
			Name      string
		}

		err = helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		ci := &data.CI{
			ID:        id,
			UpdatedAt: time.Now(),
			Name:      input.Name,
		}

		v := validator.New()

		if data.ValidateCI(v, ci); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PutToDB(r, db, *ci)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "PUT /v1/cis", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func PutToDB(r *http.Request, db *database.Queries, ci data.CI) (data.CI, error) {
	i, err := db.UpdateConfigurationItems(r.Context(), database.UpdateConfigurationItemsParams{
		ID:        ci.ID,
		UpdatedAt: time.Now(),
		Name:      ci.Name,
	})
	if err != nil {
		return data.CI{}, errors.New("couldn't update ci")
	}

	response := convert(i)

	return response, nil
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

		logger.Info("msg", "handle", "DELETE /v1/cis", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"message": "ci successfully deleted"})
	})
}

func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.CI, error) {
	i, err := db.DeleteConfigurationItems(r.Context(), id)
	if err != nil {
		return data.CI{}, errors.New("couldn't delete ci")
	}

	response := convert(i)

	return response, nil
}
