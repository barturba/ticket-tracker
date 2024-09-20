package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleIndexPage(w http.ResponseWriter, r *http.Request, u database.User) {
	w.Header().Set("HX-Redirect", "/incidents")
	http.Redirect(w, r, "/incidents", http.StatusFound)
}
func (cfg *ApiConfig) handlePageIndex(w http.ResponseWriter, r *http.Request, u database.User) {
	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}
	hindex := views.HomeIndex(fromProtected)
	home := views.Home("", fromProtected, false, "msg", hindex)
	templ.Handler(home).ServeHTTP(w, r)
}
