package handlers

import (
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/utils/json"
)

// Declare a handler which writes a plain-text response with information about the
// application status.
func GetHealthcheck(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("msg", "handle", "GET /v1/healthcheck")
		json.RespondWithJSON(w, http.StatusOK, models.Envelope{"status": "available"})
	})
}
