package server

import "github.com/barturba/ticket-tracker/models"

var MenuItems = models.MenuItems{
	models.MenuItem{
		Name: "Incidents List",
		Link: "/incidents",
	},
	models.MenuItem{
		Name: "Configuration Items List",
		Link: "/configuration-items",
	},
	models.MenuItem{
		Name: "Companies List",
		Link: "/companies",
	},
	models.MenuItem{
		Name: "Users List",

		Link: "/users",
	}}

var ProfileItems = models.MenuItems{
	models.MenuItem{
		Name: "Settings",
		Link: "/settings",
	},
	models.MenuItem{
		Name: "Log Out",
		Link: "/logout",
	}}
