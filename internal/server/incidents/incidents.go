package incidents

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/errutil"
	"github.com/barturba/ticket-tracker/internal/helpers"
	"github.com/barturba/ticket-tracker/models"
	"github.com/barturba/ticket-tracker/validator"
	"github.com/google/uuid"
)

// GET

// func Get(logger *slog.Logger, db *database.Queries) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			i, err := GetFromDB(r, db)
// 			if err != nil {
// 				errutil.ServerErrorResponse(w, r, logger, err)
// 			}
// 			logger.Info("msg", "handle", "GET Incidents")
// 			helpers.RespondWithJSON(w, http.StatusOK, i)
// 		},
// 	)
// }

// func GetFromDB(r *http.Request, db *database.Queries) ([]models.Incident, error) {
// 	rows, err := db.GetIncidents(r.Context())
// 	if err != nil {
// 		return nil, errors.New("couldn't find incidents")
// 	}
// 	incidents := models.DatabaseIncidentsRowToIncidents(rows)

// 	for n, i := range incidents {
// 		ci, err := db.GetConfigurationItemByID(r.Context(), i.ConfigurationItemID)
// 		if err != nil {
// 			return nil, errors.New("couldn't find configuration item name")
// 		}
// 		incidents[n].ConfigurationItemName = ci.Name
// 	}
// 	return incidents, nil
// }

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
		logger.Info("msg", "handle", "GET /v1/incidents_latest")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func GetLatestFromDB(r *http.Request, db *database.Queries, limit, offset int) ([]models.Incident, error) {
	p := database.GetIncidentsLatestParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetIncidentsLatest(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find incidents")
	}
	incidents := models.DatabaseIncidentsLatestRowToIncidents(rows)
	return incidents, nil
}

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
		input.Filters.PageSize = data.ReadInt(qs, "page_size", 20, v)

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
		logger.Info("msg", "handle", "GET /v1/incidents")
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func GetFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) ([]models.Incident, error) {
	p := database.GetIncidentsFilteredParams{
		Query:  sql.NullString{String: query, Valid: query != ""},
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	rows, err := db.GetIncidentsFiltered(r.Context(), p)
	if err != nil {
		return nil, errors.New("couldn't find incidents")
	}
	incidents := models.DatabaseIncidentsFilteredRowToIncidents(rows)
	return incidents, nil
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
		logger.Info("msg", "handle", "GET /v1/incidents_count")
		helpers.RespondWithJSON(w, http.StatusOK, count)
	})
}

func GetCountFromDB(r *http.Request, db *database.Queries, query string, limit, offset int) (int64, error) {
	count, err := db.GetIncidentsFilteredCount(r.Context(), sql.NullString{String: query, Valid: query != ""})
	if err != nil {
		return 0, errors.New("couldn't find incidents")
	}
	return count, nil
}

func GetByID(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := helpers.ReadUUIDPath(*r)
		if err != nil {
			errutil.NotFoundResponse(w, r, logger)
			return
		}

		count, err := db.GetIncidentByID(r.Context(), id)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Info("msg", "handle", "GET /v1/incident/{id}")
		helpers.RespondWithJSON(w, http.StatusOK, count)
	})
}

func GetByIDFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (models.Incident, error) {
	record, err := db.GetIncidentByID(r.Context(), id)
	if err != nil {
		return models.Incident{}, errors.New("couldn't find incident")
	}
	incident := models.DatabaseIncidentToIncident(record)
	return incident, nil
}

// POST

func Post(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			CreatedAt           time.Time
			UpdatedAt           time.Time
			ShortDescription    string
			Description         string
			ConfigurationItemID uuid.UUID
			CompanyID           uuid.UUID
			AssignedToID        uuid.UUID
			State               database.StateEnum
		}

		err := helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		incident := &data.Incident{
			ID:                  uuid.New(),
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			ShortDescription:    input.ShortDescription,
			Description:         sql.NullString{String: input.Description, Valid: input.Description != ""},
			ConfigurationItemID: input.ConfigurationItemID,
			CompanyID:           input.CompanyID,
			AssignedToID:        uuid.NullUUID{UUID: input.AssignedToID, Valid: input.AssignedToID != uuid.Nil},
			State:               input.State,
		}

		v := validator.New()

		if data.ValidateIncident(v, incident); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PostToDB(r, db, *incident)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "POST /v1/incident")
		helpers.RespondWithJSON(w, http.StatusCreated, i)
	})
}

func PostToDB(r *http.Request, db *database.Queries, incident data.Incident) (data.Incident, error) {
	i, err := db.CreateIncident(r.Context(), database.CreateIncidentParams{
		ID:                  incident.ID,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description,
		State:               incident.State,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
	})
	response := convert(i)
	if err != nil {
		return data.Incident{}, errors.New("couldn't find incident")
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
			CreatedAt           time.Time
			UpdatedAt           time.Time
			ShortDescription    string
			Description         string
			ConfigurationItemID uuid.UUID
			CompanyID           uuid.UUID
			AssignedToID        uuid.UUID
			State               database.StateEnum
		}

		err = helpers.ReadJSON(w, r, &input)
		if err != nil {
			errutil.BadRequestResponse(w, r, logger, err)
			return
		}

		incident := &data.Incident{
			ID:                  id,
			CreatedAt:           time.Now(),
			UpdatedAt:           time.Now(),
			ShortDescription:    input.ShortDescription,
			Description:         sql.NullString{String: input.Description, Valid: input.Description != ""},
			ConfigurationItemID: input.ConfigurationItemID,
			CompanyID:           input.CompanyID,
			AssignedToID:        uuid.NullUUID{UUID: input.AssignedToID, Valid: input.AssignedToID != uuid.Nil},
			State:               input.State,
		}

		v := validator.New()

		if data.ValidateIncident(v, incident); !v.Valid() {
			errutil.FailedValidationResponse(w, r, logger, v.Errors)
			return
		}
		i, err := PutToDB(r, db, *incident)
		if err != nil {
			errutil.ServerErrorResponse(w, r, logger, err)
			return
		}

		logger.Info("msg", "handle", "PUT /v1/incident", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, i)
	})
}

func PutToDB(r *http.Request, db *database.Queries, incident data.Incident) (data.Incident, error) {
	i, err := db.UpdateIncident(r.Context(), database.UpdateIncidentParams{
		ID:                  incident.ID,
		UpdatedAt:           time.Now(),
		ShortDescription:    incident.ShortDescription,
		Description:         incident.Description,
		State:               incident.State,
		ConfigurationItemID: incident.ConfigurationItemID,
		CompanyID:           incident.CompanyID,
		AssignedTo:          incident.AssignedToID,
	})
	if err != nil {
		return data.Incident{}, errors.New("couldn't update incident")
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

		logger.Info("msg", "handle", "DELETE /v1/incident", "id", id)
		helpers.RespondWithJSON(w, http.StatusOK, data.Envelope{"message": "incident successfully deleted"})
	})
}

func DeleteFromDB(r *http.Request, db *database.Queries, id uuid.UUID) (data.Incident, error) {
	i, err := db.DeleteIncidentByID(r.Context(), id)
	if err != nil {
		return data.Incident{}, errors.New("couldn't delete incident")
	}

	response := convert(i)

	return response, nil
}
