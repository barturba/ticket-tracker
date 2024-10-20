package users

import (
	"database/sql"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/google/uuid"
)

// GetFromDB retrieves users from the database based on the provided query and filters.
func GetFromDB(r *http.Request, db *database.Queries, logger *slog.Logger, query string, filters models.Filters) ([]models.User, models.Metadata, error) {
	p := database.GetUsersParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}
	rows, err := db.GetUsers(r.Context(), p)
	if err != nil {
		logger.Error("couldn't find users", "error", err)
		return nil, models.Metadata{}, errors.New("couldn't find users")
	}
	users, metadata, err := convertRowsAndMetadata(rows, filters)
	if err != nil {
		return nil, models.Metadata{}, err
	}
	return users, metadata, nil
}

// GetCountFromDB retrieves the count of users from the database based on the provided query.
func GetCountFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.GetUsersCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		return 0, errors.New("couldn't find users")
	}
	return count, nil
}

// GetLatestFromDB retrieves the latest users from the database based on the provided limit and offset.
func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int32) ([]models.User, error) {
	p := database.GetUsersLatestParams{
		Limit:  limit,
		Offset: offset,
	}
	rows, err := db.GetUsersLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find users")
	}
	users := convertMany(rows)
	return users, nil
}

// GetByIDFromDB retrieves a user from the database based on the provided user ID.
func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (models.User, error) {
	record, err := db.GetUserByID(r.Context(), id)
	if err != nil {
		return models.User{}, errors.New("couldn't find user")
	}
	user := convert(record)
	return user, nil
}

// PostToDB creates a new user in the database based on the provided user models.
func PostToDB(r *http.Request, db *database.Queries, user models.User) (models.User, error) {

	i, err := db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: sql.NullString{String: user.FirstName, Valid: user.FirstName != ""},
		LastName:  sql.NullString{String: user.LastName, Valid: user.LastName != ""},
		Email:     user.Email,
	})
	response := convert(i)
	if err != nil {
		return models.User{}, errors.New("couldn't create user")
	}
	return response, nil
}

// PutToDB updates an existing user in the database based on the provided user models.
func PutToDB(r *http.Request, db *database.Queries, user models.User) (models.User, error) {

	i, err := db.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:        user.ID,
		UpdatedAt: time.Now(),
		FirstName: sql.NullString{String: user.FirstName, Valid: user.FirstName != ""},
		LastName:  sql.NullString{String: user.LastName, Valid: user.LastName != ""},
		Email:     user.Email,
	})
	if err != nil {
		log.Printf("put err %v\n", err)
		return models.User{}, errors.New("couldn't update user")
	}

	response := convert(i)

	return response, nil
}

// DeleteFromDB deletes a user from the database based on the provided user ID.
func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (models.User, error) {
	i, err := db.DeleteUserByID(r.Context(), id)
	if err != nil {
		return models.User{}, errors.New("couldn't delete user")
	}

	response := convert(i)

	return response, nil
}
