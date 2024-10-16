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

// GetFromDB retrieves a list of companies from the database based on the provided query and filters.
func GetFromDB(r *http.Request, db *database.Queries, query string, filters data.Filters) ([]data.Company, data.Metadata, error) {
	p := database.GetCompaniesParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}
	rows, err := db.GetCompanies(r.Context(), p)
	if err != nil {
		return nil, data.Metadata{}, errors.New("couldn't find companies")
	}

	companies, metadata, err := convertRowsAndMetadata(rows, filters)
	if err != nil {
		return nil, data.Metadata{}, err
	}

	return companies, metadata, nil
}

// GetLatestFromDB retrieves the latest companies from the database based on the provided limit and offset.
func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int32) ([]data.Company, error) {
	p := database.GetCompaniesLatestParams{
		Limit:  limit,
		Offset: offset,
	}
	rows, err := db.GetCompaniesLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find companies")
	}
	companies := convertMany(rows)
	return companies, nil
}

// GetByIDFromDB retrieves a company from the database based on the provided company ID.
func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Company, error) {
	record, err := db.GetCompanyByID(r.Context(), id)
	if err != nil {
		return data.Company{}, errors.New("couldn't find company")
	}
	company := convert(record)
	return company, nil
}

// PostToDB inserts a new company into the database.
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

// PutToDB updates an existing company in the database.
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

// DeleteFromDB deletes a company from the database based on the provided company ID.
func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Company, error) {
	i, err := db.DeleteCompanyByID(r.Context(), id)
	if err != nil {
		return data.Company{}, errors.New("couldn't delete company")
	}

	response := convert(i)

	return response, nil
}
