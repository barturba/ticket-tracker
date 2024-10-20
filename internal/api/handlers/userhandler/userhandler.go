// Package userhandler provides functions for managing user resources.
package userhandler

import (
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/repository/userrepository"
	"github.com/barturba/ticket-tracker/internal/utils/httperrors"
	"github.com/barturba/ticket-tracker/internal/utils/json"
	"github.com/barturba/ticket-tracker/internal/utils/validator"
	"github.com/google/uuid"
)

// ListUsers retrieves a list of users with optional filtering, sorting, and pagination.
func ListUsers(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()
		input := parseFilters(r, v)

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		users, metadata, err := userrepository.ListUsers(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/users")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"users": users, "metadata": metadata})
	})
}

// ListAllUsers retrieves all users with optional filtering, sorting, and pagination.
func ListAllUsers(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()

		input := parseFilters(r, v)
		input.Filters.PageSize = 10_000_000

		if models.ValidateFiltersGetAll(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		users, metadata, err := userrepository.ListUsers(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/users_all")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"users": users, "metadata": metadata})
	})
}

// ListRecentUsers retrieves the latest users.
func ListRecentUsers(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()

		input := parseFilters(r, v)
		input.Filters.PageSize = 20

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		latestUsers, err := userrepository.GetLatestUsers(r, logger, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/users_latest")
		json.RespondWithJSON(w, http.StatusOK, latestUsers)
	})
}

// GetUserByID retrieves a single user by their unique identifier.
func GetUserByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		user, err := userrepository.GetUserByID(r, logger, db, id)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/users/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, user)
	})
}

// CreateUser creates a new user.
func CreateUser(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			httperrors.BadRequestResponse(w, r, logger, err)
			return
		}

		user := &models.User{
			ID:        uuid.New(),
			UpdatedAt: time.Now(),
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
		}

		v := validator.New()
		if models.ValidateUser(v, user); !v.Valid() {
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		createdUser, err := userrepository.CreateUser(r, logger, db, *user)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled POST /v1/users")
		json.RespondWithJSON(w, http.StatusCreated, createdUser)
	})
}

// UpdateUser updates an existing user by their unique identifier.
func UpdateUser(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		var input struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			httperrors.BadRequestResponse(w, r, logger, err)
			return
		}

		user := &models.User{
			ID:        id,
			UpdatedAt: time.Now(),
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
		}

		v := validator.New()
		if models.ValidateUser(v, user); !v.Valid() {
			log.Printf("Put %v\n", v.Errors)
			httperrors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		updatedUser, err := userrepository.UpdateUser(r, logger, db, *user)
		if err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled PUT /v1/user/{id}")
		json.RespondWithJSON(w, http.StatusOK, updatedUser)
	})
}

// DeleteUser deletes an existing user by their unique identifier.
func DeleteUser(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			httperrors.NotFoundResponse(w, r, logger)
			return
		}

		if _, err = userrepository.DeleteUser(r, logger, db, id); err != nil {
			httperrors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled DELETE /v1/users/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"message": "user successfully deleted"})
	})
}

func parseFilters(r *http.Request, v *validator.Validator) struct {
	Query   string
	Filters models.Filters
} {
	qs := r.URL.Query()
	return struct {
		Query   string
		Filters models.Filters
	}{
		Query: models.ReadString(qs, "query", ""),
		Filters: models.Filters{
			Page:     models.ReadInt(qs, "page", 1, v),
			PageSize: models.ReadInt(qs, "page_size", 10, v),
			Sort:     models.ReadString(qs, "sort", "id"),
			SortSafelist: []string{
				"id", "-id",
				"created_at", "-created_at",
				"updated_at", "-updated_at",
				"-first_name", "first_name",
				"-last_name", "last_name",
			},
		},
	}
}
