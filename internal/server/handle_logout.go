package server

import (
	"net/http"
	"time"
)

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
