package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/barturba/ticket-tracker/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CheckHashPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func respondToFailedValidation(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	respondWithError(w, http.StatusUnprocessableEntity, fmt.Sprintf("%v", errors))
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (cfg *ApiConfig) createJWT(expiresInSeconds int, userID uuid.UUID) (string, error) {
	var (
		key []byte
		t   *jwt.Token
		s   string
	)

	key = []byte(cfg.JWTSecret)

	expires := time.Now().Add(time.Second * time.Duration(JWT_EXPIRES_IN_SECONDS))

	t = jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "ticket-tracker",
			"sub": userID,
			"iat": jwt.NewNumericDate(time.Now()),
			"nbf": jwt.NewNumericDate(time.Now()),
			"exp": jwt.NewNumericDate(expires),
		})
	s, err := t.SignedString(key)
	return s, err
}

func (cfg *ApiConfig) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return errors.New("Error decoding parameters")
	}
	return nil
}

func NewPage(title string, cfg *ApiConfig, u models.User, alert models.Alert) models.Page {
	return models.Page{
		Title:            title,
		Logo:             cfg.Logo,
		Alert:            alert,
		IsLoggedIn:       true,
		IsError:          false,
		Msg:              "",
		User:             u.Name,
		Email:            u.Email,
		ProfilePicture:   cfg.ProfilePicPlaceholder,
		MenuItems:        cfg.MenuItems,
		ProfileMenuItems: cfg.ProfileItems,
	}
}
