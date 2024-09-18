package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	loginForm := views.LoginForm()
	templ.Handler(views.ContentPage("Login", "login", loginForm)).ServeHTTP(w, r)
}
