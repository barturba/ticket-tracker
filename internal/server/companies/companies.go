package companies

import (
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
		input.Filters.SortSafelist = []string{"id", "-id",
			"created_at", "-created_at",
			"updated_at", "-updated_at",
			"name", "-name"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		companies, metadata, err := GetFromDB(r, db, input.Query, input.Filters)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/companies")
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"companies": companies, "metadata": metadata})
	})
}

func GetAll(logger *slog.Logger, db *database.Queries) http.Handler {
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
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 10_000_000, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id", "-id",
			"created_at", "-created_at",
			"updated_at", "-updated_at",
			"name", "-name"}

		if data.ValidateFiltersGetAll(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		companies, metadata, err := GetFromDB(r, db, input.Query, input.Filters)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/companies")
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"companies": companies, "metadata": metadata})
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
		logger.Info("msg", "handle", "GET /v1/companies_latest")
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

		i, err := GetByIDFromDB(r, db, id)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/companies/{id}")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
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

		company := &data.Company{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      input.Name,
		}

		v := validator.New()

		if data.ValidateCompany(v, company); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PostToDB(r, db, *company)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "POST /v1/company")
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
			Name string `json:"name"`
		}

		err = helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		company := &data.Company{
			ID:        id,
			UpdatedAt: time.Now(),
			Name:      input.Name,
		}

		v := validator.New()

		if data.ValidateCompany(v, company); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PutToDB(r, db, *company)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "PUT /v1/companies", "id", id)
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

		logger.Info("msg", "handle", "DELETE /v1/companies", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"message": "company successfully deleted"})
	})
}
