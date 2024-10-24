package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/google/uuid"
)

// ListUsers retrieves users from the database based on the provided query and filters.
func ListUsers(logger *slog.Logger, db *database.Queries, ctx context.Context, query string, filters models.Filters) ([]models.User, models.Metadata, error) {
	params := database.ListUsersParams{
		Query:    sql.NullString{String: query, Valid: query != ""},
		Limit:    int32(filters.Limit()),
		Offset:   int32(filters.Offset()),
		OrderBy:  filters.SortColumn(),
		OrderDir: filters.SortDirection(),
	}

	rows, err := db.ListUsers(ctx, params)
	if err != nil {
		logger.Error("failed to retrieve users", "error", err)
		return nil, models.Metadata{}, fmt.Errorf("failed to retrieve users: %w", err)
	}

	users, metadata, err := convertUsersAndMetadata(rows, filters)
	if err != nil {
		return nil, models.Metadata{}, err
	}

	return users, metadata, nil
}

// CountUsers retrieves the count of users from the database based on the provided query.
func CountUsers(r *http.Request, logger *slog.Logger, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.CountUsers(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		logger.Error("failed to count users", "error", err)
		return 0, fmt.Errorf("failed to count users: %w", err)
	}
	return count, nil
}

// ListRecentUsers retrieves the latest users from the database.
func ListRecentUsers(r *http.Request, logger *slog.Logger, db *database.Queries, limit, offset int32) ([]models.User, error) {
	params := database.ListRecentUsersParams{
		Limit:  limit,
		Offset: offset,
	}

	rows, err := db.ListRecentUsers(r.Context(), params)
	if err != nil {
		logger.Error("failed to retrieve recent users", "error", err)
		return nil, fmt.Errorf("failed to retrieve recent users: %w", err)
	}

	return convertManyUsers(rows), nil
}

// GetUser retrieves a user from the database based on the provided user ID.
func GetUser(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.User, error) {
	record, err := db.GetUser(r.Context(), id)
	if err != nil {
		logger.Error("failed to retrieve user",
			"error", err,
			"user_id", id,
			"operation", "GetUser",
			"component", "repository")
		return models.User{}, fmt.Errorf("failed to retrieve user: %w", err)
	}

	return convertUser(record), nil
}

// GetUserByToken retrieves a user from the database based on the provided
// session token.
func GetUserByToken(r *http.Request, logger *slog.Logger, db *database.Queries, token string) (models.User, error) {
	record, err := db.GetUserByTkn(r.Context(), token)
	if err != nil {
		logger.Error("failed to retrieve user by token", "error", err, "token", token)
		return models.User{}, fmt.Errorf("failed to retrieve user by token: %w", err)
	}

	return convertUserByTokenRow(record), nil
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
		return models.User{}, fmt.Errorf("failed to create user: %w", err)
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
		return models.User{}, fmt.Errorf("failed to update user: %w", err)
	}

	return convertUser(record), nil
}

// DeleteUser deletes a user from the database based on the provided user ID.
func DeleteUser(r *http.Request, logger *slog.Logger, db *database.Queries, id uuid.UUID) (models.User, error) {
	record, err := db.DeleteUser(r.Context(), id)
	if err != nil {
		logger.Error("failed to delete user", "error", err, "id", id)
		return models.User{}, fmt.Errorf("failed to delete user: %w", err)
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

// convertUserByTokenRow converts a database.UserByTokenRow to a models.User.
func convertUserByTokenRow(dbUser database.GetUserByTknRow) models.User {
	return models.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		FirstName: dbUser.FirstName.String,
		LastName:  dbUser.LastName.String,
		Email:     dbUser.Email,
		Active:    dbUser.Active.Bool,
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

// convertUsersAndMetadata converts a slice of database User records and filters into a slice of models.User and models.Metadata.
func convertUsersAndMetadata(rows []database.ListUsersRow, filters models.Filters) ([]models.User, models.Metadata, error) {
	if len(rows) == 0 {
		return nil, models.Metadata{}, nil
	}

	// Prevent conversion exploits
	totalRecords, err := models.ConvertInt64to32(rows[0].Count)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to convert total records count: %w", err)
	}

	users := convertManyListUsersRowToUsers(rows)
	metadata, err := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	if err != nil {
		return nil, models.Metadata{}, fmt.Errorf("failed to calculate metadata: %w", err)
	}

	return users, metadata, nil
}

// convertListUsersRowToUser converts a database row of type ListUsersRow to a User model.
func convertListUsersRowToUser(row database.ListUsersRow) models.User {
	return models.User{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		FirstName: row.FirstName.String,
		LastName:  row.LastName.String,
		Email:     row.Email,
		Role:      row.Role,
	}
}

// convertManyListUsersRowToUsers converts a database.ListUsersRow to an array of models.User.
func convertManyListUsersRowToUsers(rows []database.ListUsersRow) []models.User {
	users := make([]models.User, len(rows))
	for i, row := range rows {
		users[i] = convertListUsersRowToUser(row)
	}
	return users
}
