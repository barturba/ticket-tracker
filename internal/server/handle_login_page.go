package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleLoginPage(w http.ResponseWriter, r *http.Request) {
	fromProtected := false
	lIndexNew := views.LoginIndexNew(fromProtected, cfg.Logo)
	login := views.Login("Login", cfg.Logo, fromProtected, false, "msg", "", cfg.ProfilePicPlaceholder, cfg.MenuItems, cfg.ProfileItems, lIndexNew)
	templ.Handler(login).ServeHTTP(w, r)
}
