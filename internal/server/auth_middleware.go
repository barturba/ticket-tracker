package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/barturba/ticket-tracker/internal/auth"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check for JWT authorization and API key authorization
		user := database.User{}

		// JWT authorization check
		cookie, err := r.Cookie("jwtCookie")
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			log.Println(err)
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

			user, err = cfg.DB.GetUserByAPIKey(r.Context(), authorization)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, "error getting user")
				return
			}

		} else {
			tokenString := cookie.Value
			claims := jwt.MapClaims{}
			_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.JWTSecret), nil
			})
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("bad token: %v", err))
				return
			}

			for key, val := range claims {
				fmt.Printf("Key: %v, value: %v\n", key, val)
			}

			log.Printf("getCookieHandler: sub %v\n", claims["sub"])
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

			user, err = cfg.DB.GetUserByID(r.Context(), id)
			if err != nil {
				respondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

		}

		handler(w, r, user)
	}
}
