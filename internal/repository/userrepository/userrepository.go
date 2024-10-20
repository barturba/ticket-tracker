package userrepository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/google/uuid"
)

// ListUsers retrieves users from the database based on the provided query and filters.
func ListUsers(logger *slog.Logger, db *database.Queries, ctx context.Context, query string, filters models.Filters) ([]models.User, models.Metadata, error) {
	params := database.GetUsersParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}

	rows, err := db.GetUsers(ctx, params)
	if err != nil {
		logger.Error("failed to retrieve users", "error", err)
		return nil, models.Metadata{}, errors.New("failed to retrieve users")
	}

	users, metadata, err := convertRowsAndMetadata(rows, filters)
	if err != nil {
		return nil, models.Metadata{}, err
	}

	return users, metadata, nil
}

// CountUsers retrieves the count of users from the database based on the provided query.
func CountUsers(r *http.Request, logger *slog.Logger, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.GetUsersCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		logger.Error("failed to count users", "error", err)
		return 0, errors.New("failed to count users")
	}
	return count, nil
}

// GetLatestUsers retrieves the latest users from the database.
func GetLatestUsers(r *http.Request, logger *slog.Logger, db *database.Queries, limit, offset int32) ([]models.User, error) {
	params := database.GetUsersLatestParams{
		Limit:  limit,
		Offset: offset,
	}

	rows, err := db.GetUsersLatest(r.Context(), params)
	if err != nil {
		logger.Error("failed to retrieve recent users", "error", err)
		return nil, errors.New("failed to retrieve recent users")
	}

	return convertManyUsers(rows), nil
}

// GetUserByID retrieves a user from the database based on the provided user ID.
func GetUserByID(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.User, error) {
	record, err := db.GetUserByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to retrieve user", "error", err, "user", id)
		return models.User{}, errors.New("failed to retrieve user")
	}

	return convertUser(record), nil
}

// CreateUser creates a new user in the database.
func CreateUser(r *http.Request, logger *slog.Logger, db *database.Queries, user models.User) (models.User, error) {
	params := database.CreateUserParams{
		ID:        user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FirstName: sql.NullString{String: user.FirstName, Valid: user.FirstName != ""},
		LastName:  sql.NullString{String: user.LastName, Valid: user.LastName != ""},
		Email:     user.Email,
	}

	record, err := db.CreateUser(r.Context(), params)
	if err != nil {
		logger.Error("failed to create user", "error", err)
		return models.User{}, errors.New("failed to create user")
	}

	return convertUser(record), nil
}

// UpdateUser updates an existing user in the database.
func UpdateUser(r *http.Request, logger *slog.Logger, db *database.Queries, user models.User) (models.User, error) {
	params := database.UpdateUserParams{
		ID:        user.ID,
		UpdatedAt: time.Now(),
		FirstName: sql.NullString{String: user.FirstName, Valid: user.FirstName != ""},
		LastName:  sql.NullString{String: user.LastName, Valid: user.LastName != ""},
		Email:     user.Email,
	}

	record, err := db.UpdateUser(r.Context(), params)
	if err != nil {
		logger.Error("failed to update user", "error", err, "id", user.ID)
		return models.User{}, errors.New("failed to update user")
	}

	return convertUser(record), nil
}

// DeleteUser deletes a user from the database based on the provided user ID.
func DeleteUser(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.User, error) {
	record, err := db.DeleteUserByID(r.Context(), id)
	if err != nil {
		logger.Error("failed to delete user", "error", err, "id", id)
		return models.User{}, errors.New("failed to delete user")
	}

	return convertUser(record), nil
}

// convertUser converts a database.User to a models.User.
func convertUser(dbUser database.User) models.User {
	return models.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		FirstName: dbUser.FirstName.String,
		LastName:  dbUser.LastName.String,
		Email:     dbUser.Email,
	}
}

// convertManyUsers transforms a slice of database.User to a slice of models.User.
func convertManyUsers(dbUsers []database.User) []models.User {
	users := make([]models.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = convertUser(dbUser)
	}
	return users
}

// convertRowsAndMetadata converts a slice of database.GetUsersRow to a slice of models.User
// and calculates the metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetUsersRow, filters models.Filters) ([]models.User, models.Metadata, error) {
	var output []models.User
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}

	// Prevent conversion exploits
	v32, err := models.ConvertInt64to32(totalRecords)
	if err != nil {
		return nil, models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(v32, filters.Page, filters.PageSize)
	return output, metadata, nil
}

// convertRowAndCount converts a database.GetUsersRow to a models.User and updates the count.
func convertRowAndCount(row database.GetUsersRow, count *int64) models.User {
	outputRow := models.User{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		FirstName: row.FirstName.String,
		LastName:  row.LastName.String,
	}
	*count = row.Count

	return outputRow
}
