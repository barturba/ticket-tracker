package healthcheck

import (
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/json"
)

// Declare a handler which writes a plain-text response with information about the
// application status.
func Get(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("msg", "handle", "GET /v1/incidents")
		json.RespondWithJSON(w, http.StatusOK, data.Envelope{"status": "available"})
	})
}
