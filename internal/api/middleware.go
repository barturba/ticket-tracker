package api

import (
	"context"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/repository"
	"github.com/barturba/ticket-tracker/internal/utils/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Middleware func(http.Handler) http.Handler

type Middleware2 func(http.HandlerFunc) http.HandlerFunc

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func Auth(logger *slog.Logger, db *database.Queries, cfg models.Config) Middleware {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add the "Vary: Authorization" header. This informs caches that the
			// response may vary based on the contents of the Authorization header.
			w.Header().Add("Vary", "Authorization")

			// Retrieve the value of the Authorization header.
			authorizationHeader := r.Header.Get("Authorization")

			// If the Authorization header is blank, give the request context an
			// anonymous user.
			if authorizationHeader == "" {
				r = ContextSetUser(r, models.AnonymousUser)
				next.ServeHTTP(w, r)
				return
			}

			headerParts := strings.Split(authorizationHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				errors.InvalidAuthenticationTokenResponse(w, r, logger)
				return
			}

			// // Get the actual token
			tokenString := headerParts[1]

			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JWTSecret), nil
			})
			if err != nil {
				errors.ServerErrorResponse(w, r, logger, err)
				return
			}

			sessionTokenClaim, ok := claims["sessionToken"].(string)
			if !ok {
				errors.InvalidAuthenticationTokenResponse(w, r, logger)
				return
			}

			// Retrieve the details of the user associated with the authentication
			// token, again calling the InvalidAuthenticationTokenResponse() helper
			// if no record was found.
			user, err := repository.GetUserByToken(logger, db, r.Context(), sessionTokenClaim)
			if err != nil {
				errors.ServerErrorResponse(w, r, logger, err)
				return
			}

			// Call the ContextSetUser() helper to add the user to the context.
			r = ContextSetUser(r, &user)

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}
func RequireActiveUserMiddleware(logger *slog.Logger, db *database.Queries, cfg models.Config) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Use the ContextGetUser() helper to retrieve the user information.
			user := ContextGetUser(r)

			if models.IsAnonymous(user) {
				errors.AuthenticationRequiredResponse(w, r, logger)
				return
			}

			if !user.Active {
				errors.InactiveAccountResponse(w, r, logger)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func RequireActiveUser(logger *slog.Logger, db *database.Queries, cfg models.Config) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Use the ContextGetUser() helper to retrieve the user information.
			user := ContextGetUser(r)

			if models.IsAnonymous(user) {
				errors.AuthenticationRequiredResponse(w, r, logger)
				return
			}

			if !user.Active {
				errors.InactiveAccountResponse(w, r, logger)
				return
			}

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}

// Log each request to an id
func WithRequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if there's already an id
			requestID := r.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.New().String()
			}

			// Set the request id header
			w.Header().Set("X-Request-ID", requestID)

			// Set the request id in the context
			ctx := ContextSetRequestId(r, requestID)
			next.ServeHTTP(w, ctx)
		})
	}
}

// Custom ResponseWriter to capture status code
type ResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w}
}

func (rw *ResponseWriter) Status() int {
	return rw.status
}

func (rw *ResponseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

// Should use middleware for consistent logging
func Logger(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			requestID := GetRequestID(r.Context())

			// Wrap ResponseWriter to capture status code
			ww := NewResponseWriter(w)

			// Add request ID to all log entries
			loggerWithRequestID := logger.With(
				"request_id", requestID,
				"method", r.Method,
				"path", r.URL.Path,
			)

			// Add logger to context
			ctx := context.WithValue(r.Context(), LoggerContextKey, loggerWithRequestID)
			r = r.WithContext(ctx)

			// Process request
			next.ServeHTTP(ww, r)

			// Log completion
			loggerWithRequestID.Info("request completed",
				"status", ww.Status(),
				"duration", time.Since(start),
			)

		})
	}
}
