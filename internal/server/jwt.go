package server

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

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
