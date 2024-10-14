// Package cis provides functions to interact with the database for Configuration Items (CIs).
// It includes functions to retrieve, create, update, and delete CIs from the database.
package cis

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/google/uuid"
)

// GetFromDB retrieves a list of CIs from the database based on the provided query and filters.
// It returns the list of CIs, metadata, and an error if any.
func GetFromDB(r *http.Request, db *database.Queries, query string, filters data.Filters) ([]data.CI, data.Metadata, error) {
	p := database.GetCIsParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}
	rows, err := db.GetCIs(r.Context(), p)
	if err != nil {
		return nil, data.Metadata{}, errors.New("couldn't find cis")
	}
	cis, metadata := convertRowsAndMetadata(rows, filters)
	return cis, metadata, nil
}

// GetLatestFromDB retrieves the latest CIs from the database based on the provided limit and offset.
// It returns the list of CIs and an error if any.
func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int) ([]data.CI, error) {
	p := database.GetCIsLatestParams{
		Limit:  int32(limit),
		Offset: int32(offset),
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
func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.CI, error) {
	record, err := db.GetCIsByID(r.Context(), id)
	if err != nil {
		return data.CI{}, errors.New("couldn't find ci")
	}
	ci := convert(record)
	return ci, nil
}

// PostToDB creates a new CI in the database based on the provided CI data.
// It returns the created CI and an error if any.
func PostToDB(r *http.Request, db *database.Queries, ci data.CI) (data.CI, error) {
	i, err := db.CreateCIs(r.Context(), database.CreateCIsParams{
		ID:        ci.ID,
		Name:      ci.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	response := convert(i)
	if err != nil {
		return data.CI{}, errors.New("couldn't create ci")
	}
	return response, nil
}

// PutToDB updates an existing CI in the database based on the provided CI data.
// It returns the updated CI and an error if any.
func PutToDB(r *http.Request, db *database.Queries, ci data.CI) (data.CI, error) {
	i, err := db.UpdateCIs(r.Context(), database.UpdateCIsParams{
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

// DeleteFromDB deletes a CI from the database based on the provided ID.
// It returns the deleted CI and an error if any.
func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.CI, error) {
	i, err := db.DeleteCIs(r.Context(), id)
	if err != nil {
		return data.CI{}, errors.New("couldn't delete ci")
	}

	response := convert(i)

	return response, nil
}
