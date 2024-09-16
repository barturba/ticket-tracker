package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("handle Login: createJWT: %v\n", jwt)

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

	type response struct {
		ID    uuid.UUID `json:"id"`
		Email string    `json:"email"`
		Token string    `json:"token"`
	}

	respondWithJSON(w, http.StatusOK, response{
		ID:    user.ID,
		Email: user.Email,
		Token: string(jwt),
	})
}

func (cfg *apiConfig) getCookieHandler(w http.ResponseWriter, r *http.Request) {
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
			log.Println(err)
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
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("bad token: %v", err))
		return
	}

	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	log.Printf("getCookieHandler: sub %v\n", claims["sub"])
	// Echo out the cookie value in the response body.
	w.Write([]byte(fmt.Sprintf("%v", token)))
}
