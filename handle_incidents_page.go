package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/views"
	"github.com/google/uuid"
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

func (cfg *apiConfig) handleIncidentsEditPage(w http.ResponseWriter, r *http.Request, u database.User) {
	idString := r.PathValue("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't parse uuid")
		return
	}

	databaseIncident, err := cfg.DB.GetIncidentByID(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "can't find incident")
		return
	}

	incident := databaseIncidentToIncident(databaseIncident)

	templ.Handler(views.IncidentForm(incident)).ServeHTTP(w, r)
}
