package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	page := views.NewPage()

	templ.Handler(views.Index(page)).ServeHTTP(w, r)
}
