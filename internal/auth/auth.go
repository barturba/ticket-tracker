package auth

import (
	"errors"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func GetAPIKey(h http.Header) (string, error) {
	authorization := h.Get("Authorization")
	if len(authorization) == 0 {
		return "", errors.New("no authorization provided")
	}
	authorization = strings.TrimPrefix(authorization, "ApiKey ")
	return authorization, nil
}

type Password struct {
	Plaintext *string
	Hash      []byte
}

func Set(p *Password, plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = hash
	return nil
}
