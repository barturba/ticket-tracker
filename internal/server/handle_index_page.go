package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/views"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *ApiConfig) handleIndexPage(w http.ResponseWriter, r *http.Request) {
	// JWT authorization check
	cookie, err := r.Cookie("jwtCookie")
	if err != nil {
		templ.Handler(views.Index(uuid.UUID{})).ServeHTTP(w, r)
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

	user, err := cfg.DB.GetUserByID(r.Context(), id)
	if err != nil {
		user = database.User{}
	}

	if (user != database.User{}) {
		templ.Handler(views.Index(user.ID)).ServeHTTP(w, r)
		return
	} else {
		templ.Handler(views.Index(uuid.UUID{})).ServeHTTP(w, r)
		return
	}

}
