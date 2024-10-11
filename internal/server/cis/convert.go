package cis

import (
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

func convert(ci database.ConfigurationItem) data.CI {
	return data.CI{
		ID:        ci.ID,
		CreatedAt: ci.CreatedAt,
		UpdatedAt: ci.UpdatedAt,
		Name:      ci.Name,
	}
}

func convertMany(cis []database.ConfigurationItem) []data.CI {
	var items []data.CI
	for _, item := range cis {
		items = append(items, convert(item))
	}
	return items
}
