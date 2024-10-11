package companies

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/errutil"
	"github.com/barturba/ticket-tracker/internal/helpers"
	"github.com/barturba/ticket-tracker/models"
)

// GET

func Get(logger *slog.Logger, db *database.Queries) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			i, err := GetFromDB(r, db)
			if err != nil {
				errutil.ServerErrorResponse(w, r, logger, err)
			}
			logger.Info("msg", "handle", "GET /v1/companies")
			helpers.RespondWithJSON(w, http.StatusOK, i)
		},
	)
}

func GetFromDB(r *http.Request, db *database.Queries) ([]models.Company, error) {
	rows, err := db.GetCompanies(r.Context())
	if err != nil {
		return nil, errors.New("couldn't find companies")
	}
	companies := models.DatabaseCompaniesToCompanies(rows)

	return companies, nil
}