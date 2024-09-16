package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *apiConfig) handleIncidentsPage(w http.ResponseWriter, r *http.Request, u database.User) {
	organization, err := cfg.DB.GetOrganizationByUserID(r.Context(), u.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find organization")
		return
	}

	databaseIncidents, err := cfg.DB.GetIncidentsByOrganizationID(r.Context(), organization.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}

	// respondWithJSON(w, http.StatusOK, )
	incidents := databaseGetIncidentsByOrganizationIDRowToIncidents(databaseIncidents)
	page := views.NewPage()

	templ.Handler(views.Incidents(page, incidents)).ServeHTTP(w, r)
}
