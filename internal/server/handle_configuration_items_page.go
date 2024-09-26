package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/models"
	"github.com/barturba/ticket-tracker/views"
)

func (cfg *ApiConfig) handleViewConfigurationItems(w http.ResponseWriter, r *http.Request, u database.User) {
	fromProtected := false
	if (u != database.User{}) {
		fromProtected = true
	}

	databaseConfigurationItems, err := cfg.DB.GetConfigurationItems(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't find incidents")
		return
	}
	configurationItems := models.DatabaseConfigurationItemsToConfigurationItems(databaseConfigurationItems)

	cIIndex := views.ConfigurationItemsIndex(configurationItems)
	iList := views.ConfigurationItemsList("Configuration Items List",
		fromProtected,
		false,
		"",
		u.Name,
		u.Email,
		cfg.MenuItems,
		cfg.ProfileItems,
		cIIndex)
	templ.Handler(iList).ServeHTTP(w, r)
}
