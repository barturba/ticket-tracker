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
		Email:     user.Email,
	}
}

func convertMany(users []database.User) []data.User {
	var items []data.User
	for _, item := range users {
		items = append(items, convert(item))
	}
	return items
}

func convertRowsAndMetadata(rows []database.GetUsersRow, filters data.Filters) ([]data.User, data.Metadata) {
	var output []data.User
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}
	metadata := data.CalculateMetadata(int(totalRecords), filters.Page, filters.PageSize)
	return output, metadata
}

func convertRowAndCount(row database.GetUsersRow, count *int64) data.User {
	outputRow := data.User{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		FirstName: row.FirstName.String,
		LastName:  row.LastName.String,
	}
	*count = row.Count

	return outputRow
}
