package server

import (
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *ApiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}
	if params.Email == "" {
		respondWithError(w, http.StatusInternalServerError, "email can't be blank")
		return
	}
	if params.Password == "" {
		respondWithError(w, http.StatusInternalServerError, "password can't be blank")
		return
	}
	if params.ExpiresInSeconds == 0 {
		params.ExpiresInSeconds = JWT_EXPIRES_IN_SECONDS
	}

	user, err := cfg.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't log in")
		return
	}

	err = CheckHashPassword(user.Password.String, params.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid password")
		return
	}

	jwt, err := cfg.createJWT(params.ExpiresInSeconds, user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't create jwt")
		return
	}

	cookie := http.Cookie{
		Name:     "jwtCookie",
		Value:    string(jwt),
		Path:     "/",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("HX-Redirect", "/incidents")
	http.Redirect(w, r, "/incidents", http.StatusFound)
}

func (cfg *ApiConfig) handleLogout(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{
		Name:     "jwtCookie",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("HX-Redirect", "/login")
	http.Redirect(w, r, "/login", http.StatusOK)
}
