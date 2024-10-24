package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
	"github.com/barturba/ticket-tracker/internal/repository"
	"github.com/barturba/ticket-tracker/internal/utils/errors"
	"github.com/golang-jwt/jwt"
)

func Authenticate(logger *slog.Logger, db *database.Queries, cfg models.Config, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add the "Vary: Authorization" header. This informs caches that the
		// response may vary based on the contents of the Authorization header.
		w.Header().Add("Vary", "Authorization")

		// Retrieve the value of the Authorization header.
		authorizationHeader := r.Header.Get("Authorization")
		logger.Info("authorizationHeader", slog.String("value", authorizationHeader))

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
		// logger.Info(fmt.Sprintf("tokenString: %s", tokenString))

		// tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
		claims := jwt.MapClaims{}
		logger.Info("DEBUG")
		logger.Info(fmt.Sprintf("tokenString: %v", tokenString))
		logger.Info("DEBUG")
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}
		// for key, val := range claims {
		// 	fmt.Printf("Key: %v, value: %v\n", key, val)
		// }
		// fmt.Printf("Key: email, value: %v\n", claims["email"])
		fmt.Printf("Key: sessionToken, value: %v\n", claims["sessionToken"])
		sessionTokenClaim, ok := claims["sessionToken"].(string)
		if !ok {
			logger.Info("!ok", slog.String("sessionTokenClaim", sessionTokenClaim))
			errors.InvalidAuthenticationTokenResponse(w, r, logger)
			return
		}

		// headerParts := strings.Split(authorizationHeader, " ")
		// if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		// 	errors.InvalidAuthenticationTokenResponse(w, r, logger)
		// 	return
		// }

		// // Get the actual token
		// tokenString := headerParts[1]
		// logger.Info(fmt.Sprintf("tokenString: %s", tokenString))

		// // Parse the JWT claims
		// claims := jwt.MapClaims{}
		// _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 	return []byte(cfg.JWTSecret), nil
		// })
		// logger.Info("error", slog.String("error", err.Error()))
		// if err != nil {
		// 	errors.ServerErrorResponse(w, r, logger, err)
		// 	return
		// }
		// logger.Info(fmt.Sprintf("claims: %s", claims))

		// sessionTokenClaim, ok := claims["id"].(string)
		// if !ok {
		// 	logger.Info("!ok", slog.String("sessionTokenClaim", sessionTokenClaim))
		// 	errors.InvalidAuthenticationTokenResponse(w, r, logger)
		// 	return
		// }
		// logger.Warn(fmt.Sprintf("sessionTokenClaim: %s", sessionTokenClaim))

		// if sessionTokenClaim == "" {
		// 	errors.InvalidAuthenticationTokenResponse(w, r, logger)
		// 	return
		// }

		// // Retrieve the details of the user associated with the authentication
		// // token, again calling the InvalidAuthenticationTokenResponse() helper
		// // if no record was found.
		user, err := repository.GetUserByToken(r, logger, db, sessionTokenClaim)
		if err != nil {
			errors.ServerErrorResponse(w, r, logger, err)
			return
		}
		logger.Warn(fmt.Sprintf("user: %v", user))

		// Call the ContextSetUser() helper to add the user to the context.
		r = ContextSetUser(r, &user)

		// r = ContextSetUser(r, models.AnonymousUser)

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
