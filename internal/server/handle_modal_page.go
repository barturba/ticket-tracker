package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleModalPage(w http.ResponseWriter, r *http.Request, u database.User) {
	modal := views.Modal()
	templ.Handler(modal).ServeHTTP(w, r)
}
