package main

import (
	"net/http"

	"github.com/barturba/ticket-tracker/internal/auth"
	"github.com/barturba/ticket-tracker/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), authorization)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "error getting user")
			return
		}

		handler(w, r, user)
	}
}
