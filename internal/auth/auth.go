package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(h http.Header) (string, error) {
	authorization := h.Get("Aunthorization")
	if len(authorization) == 0 {
		return "", errors.New("no authorization provided")
	}
	authorization = strings.TrimPrefix(authorization, "ApiKey ")
	return authorization, nil
}
