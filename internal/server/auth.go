package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	// type response struct {
	// 	ID    uuid.UUID `json:"id"`
	// 	Email string    `json:"email"`
	// 	Token string    `json:"token"`
	// }

	// respondWithJSON(w, http.StatusOK, response{
	// 	ID:    user.ID,
	// 	Email: user.Email,
	// 	Token: string(jwt),
	// })
}

func (cfg *ApiConfig) getCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cookie from the request using its name (which in our case is
	// "exampleCookie"). If no matching cookie is found, this will return a
	// http.ErrNoCookie error. We check for this, and return a 400 Bad Request
	// response to the client.
	cookie, err := r.Cookie("jwtCookie")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}

	tokenString := cookie.Value
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "bad token")
		return
	}

	// Echo out the cookie value in the response body.
	w.Write([]byte(fmt.Sprintf("%v", token)))
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
