package cis

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/google/uuid"
)

// GetFromDB retrieves a list of CIs from the database based on the provided query and filters.
// It returns the list of CIs, metadata, and an error if any.
func GetFromDB(r *http.Request, db *database.Queries, query string, filters models.Filters) ([]models.CI, models.Metadata, error) {
	p := database.GetCIsParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    filters.Limit(),
		Offset:   filters.Offset(),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}
	rows, err := db.GetCIs(r.Context(), p)
	if err != nil {
		return nil, models.Metadata{}, errors.New("couldn't find cis")
	}
	cis, metadata, err := convertRowsAndMetadata(rows, filters)
	if err != nil {
		return nil, models.Metadata{}, err
	}
	return cis, metadata, nil
}

// GetLatestFromDB retrieves the latest CIs from the database based on the provided limit and offset.
// It returns the list of CIs and an error if any.
func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int32) ([]models.CI, error) {
	p := database.GetCIsLatestParams{
		Limit:  limit,
		Offset: offset,
	}
	rows, err := db.GetCIsLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find ci")
	}
	cis := convertMany(rows)
	return cis, nil
}

// GetByIDFromDB retrieves a CI from the database based on the provided ID.
// It returns the CI and an error if any.
func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (models.CI, error) {
	record, err := db.GetCIsByID(r.Context(), id)
	if err != nil {
		return models.CI{}, errors.New("couldn't find ci")
	}
	ci := convert(record)
	return ci, nil
}

// PostToDB creates a new CI in the database based on the provided CI models.
// It returns the created CI and an error if any.
func PostToDB(r *http.Request, db *database.Queries, ci models.CI) (models.CI, error) {
	i, err := db.CreateCIs(r.Context(), database.CreateCIsParams{
		ID:        ci.ID,
		Name:      ci.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	response := convert(i)
	if err != nil {
		return models.CI{}, errors.New("couldn't create ci")
	}
	return response, nil
}

// PutToDB updates an existing CI in the database based on the provided CI models.
// It returns the updated CI and an error if any.
func PutToDB(r *http.Request, db *database.Queries, ci models.CI) (models.CI, error) {
	i, err := db.UpdateCIs(r.Context(), database.UpdateCIsParams{
		ID:        ci.ID,
		UpdatedAt: time.Now(),
		Name:      ci.Name,
	})
	if err != nil {
		return models.CI{}, errors.New("couldn't update ci")
	}

	response := convert(i)

	return response, nil
}

// DeleteFromDB deletes a CI from the database based on the provided ID.
// It returns the deleted CI and an error if any.
func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (models.CI, error) {
	i, err := db.DeleteCIs(r.Context(), id)
	if err != nil {
		return models.CI{}, errors.New("couldn't delete ci")
	}

	response := convert(i)

	return response, nil
}
