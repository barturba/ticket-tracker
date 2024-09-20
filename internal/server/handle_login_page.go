package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	fromProtected := false
	lIndex := views.LoginIndex(fromProtected)
	login := views.Login("", fromProtected, false, "msg", lIndex)
	templ.Handler(login).ServeHTTP(w, r)
}
