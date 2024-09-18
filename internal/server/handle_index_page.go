package server

import (
	"net/http"

	"github.com/barturba/ticket-tracker/internal/database"
)

func (cfg *ApiConfig) handleIndexPage(w http.ResponseWriter, r *http.Request, u database.User) {
	w.Header().Set("HX-Redirect", "/incidents")
	http.Redirect(w, r, "/incidents", http.StatusFound)
}
