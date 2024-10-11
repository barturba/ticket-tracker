package users

import (
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

func convert(user database.User) data.User {
	return data.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		APIkey:    user.Apikey,
	}
}

func convertMany(users []database.User) []data.User {
	var items []data.User
	for _, item := range users {
		items = append(items, convert(item))
	}
	return items
}
