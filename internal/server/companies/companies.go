package companies

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/errutil"
	"github.com/barturba/ticket-tracker/internal/helpers"
	"github.com/barturba/ticket-tracker/models"
	"github.com/barturba/ticket-tracker/validator"
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
		logger.Info("msg", "handle", "GET /v1/companies")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func GetFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) ([]models.Company, error) {
	p := database.GetCompaniesParams{
		Query:  sql.NullString{String: query, Valid: query != ""},
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetCompanies(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find companies")
	}
	companies := models.DatabaseCompaniesRowToCompanies(rows)
	return companies, nil
}
