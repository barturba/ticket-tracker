package companies

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
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 20, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id"}

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

func GetFromDB(r *http.Request, db *database.Queries, query string, filters data.Filters) ([]data.Company, data.Metadata, error) {
	p := database.GetCompaniesParams{
		Query:  sql.NullString{String: query, Valid: query != ""},
		Limit:  int32(filters.Limit()),
		Offset: int32(filters.Offset()),
	}
	rows, err := db.GetCompanies(r.Context(), p)
	if err != nil {
		return nil, data.Metadata{}, errors.New("couldn't find companies")
	}

	companies, metadata := convertRowsAndMetadata(rows, filters)

	return companies, metadata, nil
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

func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int) ([]data.Company, error) {
	p := database.GetCompaniesLatestParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetCompaniesLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find companies")
	}
	companies := convertMany(rows)
	return companies, nil
}

func GetByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		count, err := db.GetCompanyByID(r.Context(), id)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/company/{id}")
		helpers.RespondWithJSON(w, http.StatusOK, count)
	})
}

func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Company, error) {
	record, err := db.GetCompanyByID(r.Context(), id)
	if err != nil {
		return data.Company{}, errors.New("couldn't find company")
	}
	company := convert(record)
	return company, nil
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

func PostToDB(r *http.Request, db *database.Queries, company data.Company) (data.Company, error) {
	i, err := db.CreateCompany(r.Context(), database.CreateCompanyParams{
		ID:        company.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      company.Name,
	})
	response := convert(i)
	if err != nil {
		return data.Company{}, errors.New("couldn't find company")
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
			CreatedAt time.Time
			UpdatedAt time.Time
			Name      string
		}

		err = helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		company := &data.Company{
			ID:        id,
			CreatedAt: time.Now(),
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

func PutToDB(r *http.Request, db *database.Queries, company data.Company) (data.Company, error) {
	i, err := db.UpdateCompany(r.Context(), database.UpdateCompanyParams{
		ID:        company.ID,
		UpdatedAt: time.Now(),
		Name:      company.Name,
	})
	if err != nil {
		return data.Company{}, errors.New("couldn't update company")
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

		logger.Info("msg", "handle", "DELETE /v1/companies", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"message": "company successfully deleted"})
	})
}

func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Company, error) {
	i, err := db.DeleteCompanyByID(r.Context(), id)
	if err != nil {
		return data.Company{}, errors.New("couldn't delete company")
	}

	response := convert(i)

	return response, nil
}
