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

// GET

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

func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int) ([]data.CI, error) {
	p := database.GetCIsLatestParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetCIsLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find cis")
	}
	cis := convertMany(rows)
	return cis, nil
}

func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.CI, error) {
	record, err := db.GetCIsByID(r.Context(), id)
	if err != nil {
		return data.CI{}, errors.New("couldn't find ci")
	}
	ci := convert(record)
	return ci, nil
}

// POST

func PostToDB(r *http.Request, db *database.Queries, ci data.CI) (data.CI, error) {
	i, err := db.CreateCIs(r.Context(), database.CreateCIsParams{
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

// DELETE

func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.CI, error) {
	i, err := db.DeleteCIs(r.Context(), id)
	if err != nil {
		return data.CI{}, errors.New("couldn't delete ci")
	}

	response := convert(i)

	return response, nil
}
