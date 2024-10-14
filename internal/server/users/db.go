package users

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/auth"
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/google/uuid"
)

// GET

func GetFromDB(r *http.Request, db *database.Queries, query string, filters data.Filters) ([]data.User, data.Metadata, error) {
	p := database.GetUsersParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}
	rows, err := db.GetUsers(r.Context(), p)
	if err != nil {
		return nil, data.Metadata{}, errors.New("couldn't find users")
	}
	users, metadata := convertRowsAndMetadata(rows, filters)
	return users, metadata, nil
}

func GetCountFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.GetUsersCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		return 0, errors.New("couldn't find users")
	}
	return count, nil
}

func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int) ([]data.User, error) {
	p := database.GetUsersLatestParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetUsersLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find users")
	}
	users := convertMany(rows)
	return users, nil
}

func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.User, error) {
	record, err := db.GetUserByID(r.Context(), id)
	if err != nil {
		return data.User{}, errors.New("couldn't find user")
	}
	user := convert(record)
	return user, nil
}

// POST

func PostToDB(r *http.Request, db *database.Queries, user data.User) (data.User, error) {

	password := auth.Password{}
	err := auth.Set(&password, user.Password)
	if err != nil {
		return data.User{}, errors.New("couldn't set password")
	}

	i, err := db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: sql.NullString{String: user.FirstName, Valid: user.FirstName != ""},
		LastName:  sql.NullString{String: user.LastName, Valid: user.LastName != ""},
		Email:     user.Email,
		// Password:  sql.NullString{String: string(password.Hash), Valid: true},
	})
	response := convert(i)
	if err != nil {
		return data.User{}, errors.New("couldn't create user")
	}
	return response, nil
}

// PUT

func PutToDB(r *http.Request, db *database.Queries, user data.User) (data.User, error) {

	password := auth.Password{}
	err := auth.Set(&password, user.Password)
	if err != nil {
		return data.User{}, errors.New("couldn't set password")
	}

	i, err := db.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:        user.ID,
		UpdatedAt: time.Now(),
		FirstName: sql.NullString{String: user.FirstName, Valid: user.FirstName != ""},
		LastName:  sql.NullString{String: user.LastName, Valid: user.LastName != ""},
		Email:     user.Email,
		// Password:  sql.NullString{String: string(password.Hash), Valid: true},
	})
	if err != nil {
		log.Printf("put err %v\n", err)
		return data.User{}, errors.New("couldn't update user")
	}

	response := convert(i)

	return response, nil
}

// DELETE

func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.User, error) {
	i, err := db.DeleteUserByID(r.Context(), id)
	if err != nil {
		return data.User{}, errors.New("couldn't delete user")
	}

	response := convert(i)

	return response, nil
}
