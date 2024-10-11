package users

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/auth"
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/errutil"
	"github.com/barturba/ticket-tracker/internal/helpers"
	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// GET

func Get(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Query  string
			Limit  int
			Offset int
			data.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Query = data.ReadString(qs, "query", "")

		input.Filters.Page = data.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 10, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		i, err := GetFromDB(r, db, input.Query, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/users")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func GetFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) ([]data.User, error) {
	p := database.GetUsersParams{
		Query:  sql.NullString{String: query, Valid: query != ""},
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetUsers(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find users")
	}
	users := convertMany(rows)
	return users, nil
}

func GetCount(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Query  string
			Limit  int
			Offset int
			data.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Query = data.ReadString(qs, "query", "")

		input.Filters.Page = data.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 20, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		count, err := GetCountFromDB(r, db, input.Query, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/users_count")
		helpers.RespondWithJSON(w, http.StatusOK, count)
	})
}

func GetCountFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.GetUsersCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		return 0, errors.New("couldn't find users")
	}
	return count, nil
}

func GetLatest(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Limit  int
			Offset int
			data.Filters
		}

		v := validator.New()

		var qs = r.URL.Query()

		input.Filters.Page = data.ReadInt(qs, "page", 1, v)
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 20, v)

		input.Filters.Sort = data.ReadString(qs, "sort", "id")
		input.Filters.SortSafelist = []string{"id"}

		if data.ValidateFilters(v, input.Filters); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}

		i, err := GetLatestFromDB(r, db, input.Filters.Limit(), input.Filters.Offset())
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/users_latest")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
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

func GetByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		count, err := db.GetUserByID(r.Context(), id)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/user/{id}")
		helpers.RespondWithJSON(w, http.StatusOK, count)
	})
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

func Post(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			FirstName           string
			LastName            string
			ConfigurationItemID uuid.UUID
			ApiKey              string
			Email               string
			Password            string
		}

		err := helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		user := &data.User{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			FirstName: input.FirstName,
			LastName:  input.LastName,
			APIkey:    input.ApiKey,
			Password:  input.Password,
		}

		v := validator.New()

		if data.ValidateUser(v, user); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PostToDB(r, db, *user)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "POST /v1/user")
		helpers.RespondWithJSON(w, http.StatusCreated, i)
	})
}

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
		Apikey:    user.APIkey,
		Email:     user.Email,
		Password:  sql.NullString{String: string(password.Hash), Valid: true},
	})
	response := convert(i)
	if err != nil {
		return data.User{}, errors.New("couldn't find user")
	}
	return response, nil
}

// PUT

func Put(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		var input struct {
			CreatedAt time.Time
			UpdatedAt time.Time
			FirstName string
			LastName  string
			APIKey    string
			Email     string
			Password  string
		}

		err = helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		user := &data.User{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			FirstName: input.FirstName,
			LastName:  input.LastName,
			APIkey:    input.APIKey,
			Email:     input.Email,
			Password:  input.Password,
		}

		v := validator.New()

		if data.ValidateUser(v, user); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PutToDB(r, db, *user)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "PUT /v1/user", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

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
		Apikey:    user.APIkey,
		Email:     user.Email,
		Password:  sql.NullString{String: string(password.Hash), Valid: true},
	})
	if err != nil {
		return data.User{}, errors.New("couldn't update user")
	}

	response := convert(i)

	return response, nil
}

// DELETE

func Delete(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		_, err = DeleteFromDB(r, db, id)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "DELETE /v1/user", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"message": "user successfully deleted"})
	})
}

func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.User, error) {
	i, err := db.DeleteUserByID(r.Context(), id)
	if err != nil {
		return data.User{}, errors.New("couldn't delete user")
	}

	response := convert(i)

	return response, nil
}
