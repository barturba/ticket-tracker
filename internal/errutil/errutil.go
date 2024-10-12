package errutil

import (
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/helpers"
)

func ErrorResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, status int, message any) {
	env := data.Envelope{"error": message}

	err := helpers.WriteJSON(w, status, env, nil)
	if err != nil {
		LogError(r, logger, err)
		w.WriteHeader(500)
	}
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, errors map[string]string) {
	ErrorResponse(w, r, logger, http.StatusUnprocessableEntity, errors)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "the requested resource could not be found"
	ErrorResponse(w, r, logger, http.StatusNotFound, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	ErrorResponse(w, r, logger, http.StatusBadRequest, err)
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	LogError(r, logger, err)
	message := "the server encountered a problem and couldn't process your request"
	ErrorResponse(w, r, logger, http.StatusInternalServerError, message)
}

func LogError(r *http.Request, logger *slog.Logger, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	logger.Error(err.Error(), "method", method, "uri", uri)
}
