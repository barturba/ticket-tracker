package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleIndexPage(w http.ResponseWriter, r *http.Request) {

	templ.Handler(views.IndexComponent()).ServeHTTP(w, r)
}
