package companies

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/google/uuid"
)

// GET

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

func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Company, error) {
	record, err := db.GetCompanyByID(r.Context(), id)
	if err != nil {
		return data.Company{}, errors.New("couldn't find company")
	}
	company := convert(record)
	return company, nil
}

// POST

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

// DB

func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Company, error) {
	i, err := db.DeleteCompanyByID(r.Context(), id)
	if err != nil {
		return data.Company{}, errors.New("couldn't delete company")
	}

	response := convert(i)

	return response, nil
}
