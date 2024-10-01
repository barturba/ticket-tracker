package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/auth"
	"github.com/barturba/ticket-tracker/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type authedHandler func(http.ResponseWriter, *http.Request, models.User)

func (cfg *ApiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check for JWT authorization and API key authorization
		user := models.User{}

		// JWT authorization check
		cookie, err := r.Cookie("jwtCookie")
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		if errors.Is(err, http.ErrNoCookie) {
			// API key authorization check
			authorization, err := auth.GetAPIKey(r.Header)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

			databaseUser, err := cfg.DB.GetUserByAPIKey(r.Context(), authorization)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "error getting user")
				return
			}
			user = models.DatabaseUserToUser(databaseUser)
		} else {
			tokenString := cookie.Value
			claims := jwt.MapClaims{}
			_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JWTSecret), nil
			})
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "bad token")
				return
			}

			idString := claims["sub"]
			if idString == "" {
				respondWithError(w, http.StatusInternalServerError, "invalid claims")
				return
			}
			idString, ok := idString.(string)
			if !ok {
				respondWithError(w, http.StatusInternalServerError, "invalid claims: sub is not a string")
				return
			}
			id, err := uuid.Parse(idString.(string))
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "invalid claims: can't parse uuid")
				return
			}

			databaseUser, err := cfg.DB.GetUserByID(r.Context(), id)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			user = models.DatabaseUserToUser(databaseUser)
		}

		handler(w, r, user)
	}
}

func (cfg *ApiConfig) middlewareAuthPage(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// JWT authorization check
		cookie, err := r.Cookie("jwtCookie")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		tokenString := cookie.Value
		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "bad token")
			return
		}

		idString := claims["sub"]
		if idString == "" {
			respondWithError(w, http.StatusInternalServerError, "invalid claims")
			return
		}
		idString, ok := idString.(string)
		if !ok {
			respondWithError(w, http.StatusInternalServerError, "invalid claims: sub is not a string")
			return
		}
		id, err := uuid.Parse(idString.(string))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "invalid claims: can't parse uuid")
			return
		}

		databaseUser, err := cfg.DB.GetUserByID(r.Context(), id)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		user := models.DatabaseUserToUser(databaseUser)
		handler(w, r, user)

	}

}

func (cfg *ApiConfig) middlewareAuthPageNoRedirect(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// JWT authorization check
		cookie, err := r.Cookie("jwtCookie")
		if err != nil {
			user := models.User{}
			handler(w, r, user)
			return
		}
		tokenString := cookie.Value
		claims := jwt.MapClaims{}
		_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.JWTSecret), nil
		})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "bad token")
			return
		}

		idString := claims["sub"]
		if idString == "" {
			respondWithError(w, http.StatusInternalServerError, "invalid claims")
			return
		}
		idString, ok := idString.(string)
		if !ok {
			respondWithError(w, http.StatusInternalServerError, "invalid claims: sub is not a string")
			return
		}
		id, err := uuid.Parse(idString.(string))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "invalid claims: can't parse uuid")
			return
		}

		databaseUser, err := cfg.DB.GetUserByID(r.Context(), id)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		user := models.DatabaseUserToUser(databaseUser)

		handler(w, r, user)

	}

}

func (cfg *ApiConfig) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("%v", errors))
}
