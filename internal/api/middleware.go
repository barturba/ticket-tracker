package api

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/repository"
	"github.com/barturba/ticket-tracker/internal/utils/errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func Authenticate(logger *slog.Logger, db *database.Queries, cfg models.Config, next http.Handler) http.Handler {
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

func RequireActiveUser(logger *slog.Logger, db *database.Queries, cfg models.Config, next http.Handler) http.Handler {
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

func WithRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := ContextSetRequestId(r, requestID)
		w.Header().Set("X-Request-ID", requestID)
		next.ServeHTTP(w, ctx)
	})
}
