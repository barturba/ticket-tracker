package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/repository"
	"github.com/barturba/ticket-tracker/internal/utils/errors"
	"github.com/barturba/ticket-tracker/internal/utils/json"
	"github.com/barturba/ticket-tracker/internal/utils/validator"
	"github.com/google/uuid"
)

// ListUsers retrieves a list of users with optional filtering, sorting, and pagination.
func ListUsers(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := validator.New()
		input := parseUserFilters(r, v)

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		users, metadata, err := repository.ListUsers(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
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

		input := parseUserFilters(r, v)
		input.Filters.PageSize = 10_000_000

		if models.ValidateFiltersGetAll(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		users, metadata, err := repository.ListUsers(logger, db, r.Context(), input.Query, input.Filters)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
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

		input := parseUserFilters(r, v)
		input.Filters.PageSize = 20

		if models.ValidateFilters(v, input.Filters); !v.Valid() {
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		latestUsers, err := repository.GetLatestUsers(r, logger, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled GET /v1/users_latest")
		json.RespondWithJSON(w, http.StatusOK, latestUsers)
	})
}

// GetUser retrieves a single user by their unique identifier.
func GetUser(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := json.ReadUUIDPath(*r)
		if err != nil {
			errors.NotFoundResponse(w, r, logger)
			return
		}

		user, err := repository.GetUser(r, logger, db, id)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
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
			errors.BadRequestResponse(w, r, logger, err)
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
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		createdUser, err := repository.CreateUser(r, logger, db, *user)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
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
			errors.NotFoundResponse(w, r, logger)
			return
		}

		var input struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
		}

		if err := json.ReadJSON(w, r, &input); err != nil {
			errors.BadRequestResponse(w, r, logger, err)
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
			errors.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		updatedUser, err := repository.UpdateUser(r, logger, db, *user)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
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
			errors.NotFoundResponse(w, r, logger)
			return
		}

		if _, err = repository.DeleteUser(r, logger, db, id); err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("handled DELETE /v1/users/{id}", "id", id)
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"message": "user successfully deleted"})
	})
}

func parseUserFilters(r *http.Request, v *validator.Validator) struct {
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
