// Package errutil provides utility functions for handling and responding to errors in HTTP handlers.
package httperrors

import (
	"log/slog"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/json"
)

// ErrorResponse sends a JSON response with the specified status code and error message.
// If writing the JSON response fails, it logs the error and sends a 500 Internal Server Error response.
func ErrorResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, status int, message any) {
	env := data.Envelope{"error": message}

	err := json.WriteJSON(w, status, env, nil)
	if err != nil {
		LogError(r, logger, err)
		w.WriteHeader(500)
	}
}

// FailedValidationResponse sends a 422 Unprocessable Entity response with the provided validation errors.
func FailedValidationResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, errors map[string]string) {
	ErrorResponse(w, r, logger, http.StatusUnprocessableEntity, errors)
}

// NotFoundResponse sends a 404 Not Found response with a standard error message indicating the resource could not be found.
func NotFoundResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger) {
	message := "the requested resource could not be found"
	ErrorResponse(w, r, logger, http.StatusNotFound, message)
}

// BadRequestResponse sends a 400 Bad Request response with the provided error message.
func BadRequestResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	ErrorResponse(w, r, logger, http.StatusBadRequest, err)
}

// ServerErrorResponse logs the provided error and sends a 500 Internal Server Error response with a standard error message.
func ServerErrorResponse(w http.ResponseWriter, r *http.Request, logger *slog.Logger, err error) {
	LogError(r, logger, err)
	message := "the server encountered a problem and couldn't process your request"
	ErrorResponse(w, r, logger, http.StatusInternalServerError, message)
}

// LogError logs the provided error along with the HTTP method and request URI.
func LogError(r *http.Request, logger *slog.Logger, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	logger.Error(err.Error(), "method", method, "uri", uri)
}
