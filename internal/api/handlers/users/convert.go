package users

import (
	"github.com/barturba/ticket-tracker/internal/data"
	"github.com/barturba/ticket-tracker/internal/database"
)

// convert converts a database.User to a data.User.
func convert(user database.User) data.User {
	return data.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Email:     user.Email,
		Role:      user.Role,
	}
}

// convertMany converts a slice of database.User to a slice of data.User.
func convertMany(users []database.User) []data.User {
	var items []data.User
	for _, item := range users {
		items = append(items, convert(item))
	}
	return items
}

// convertRowsAndMetadata converts a slice of database.GetUsersRow to a slice of data.User
// and calculates the metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetUsersRow, filters data.Filters) ([]data.User, data.Metadata, error) {
	var output []data.User
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}

	// Prevent conversion exploits
	v32, err := data.ConvertInt64to32(totalRecords)
	if err != nil {
		return nil, data.Metadata{}, err
	}

	metadata := data.CalculateMetadata(v32, filters.Page, filters.PageSize)
	return output, metadata, nil
}

// convertRowAndCount converts a database.GetUsersRow to a data.User and updates the count.
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
