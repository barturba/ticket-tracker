package users

import (
	"github.com/barturba/ticket-tracker/internal/database"
	"github.com/barturba/ticket-tracker/internal/models"
)

// convert converts a database.User to a models.User.
func convert(user database.User) models.User {
	return models.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		FirstName: user.FirstName.String,
		LastName:  user.LastName.String,
		Email:     user.Email,
		Role:      user.Role,
	}
}

// convertMany converts a slice of database.User to a slice of models.User.
func convertMany(users []database.User) []models.User {
	var items []models.User
	for _, item := range users {
		items = append(items, convert(item))
	}
	return items
}

// convertRowsAndMetadata converts a slice of database.GetUsersRow to a slice of models.User
// and calculates the metadata based on the provided filters.
func convertRowsAndMetadata(rows []database.GetUsersRow, filters models.Filters) ([]models.User, models.Metadata, error) {
	var output []models.User
	var totalRecords int64 = 0
	for _, row := range rows {
		outputRow := convertRowAndCount(row, &totalRecords)
		output = append(output, outputRow)
	}

	// Prevent conversion exploits
	v32, err := models.ConvertInt64to32(totalRecords)
	if err != nil {
		return nil, models.Metadata{}, err
	}

	metadata := models.CalculateMetadata(v32, filters.Page, filters.PageSize)
	return output, metadata, nil
}

// convertRowAndCount converts a database.GetUsersRow to a models.User and updates the count.
func convertRowAndCount(row database.GetUsersRow, count *int64) models.User {
	outputRow := models.User{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		FirstName: row.FirstName.String,
		LastName:  row.LastName.String,
	}
	*count = row.Count

	return outputRow
}
