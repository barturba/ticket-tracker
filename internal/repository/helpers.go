package repository

import (
	"fmt"

	"github.com/barturba/ticket-tracker/internal/models"
)

// calculateMetadata creates a models.Metadata struct based on the total
// records, current page, and page size.
func calculateMetadata(totalRecords, page, pageSize int32) (models.Metadata, error) {
	if totalRecords < 0 || page < 1 || pageSize < 1 {
		return models.Metadata{}, fmt.Errorf("invalid metadata parameters")
	}

	lastPage, err := models.SafeDivide(totalRecords, pageSize)
	if err != nil {
		return models.Metadata{}, fmt.Errorf("failed to calculate the last page: %w", err)
	}

	return models.Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     lastPage,
		TotalRecords: totalRecords,
	}, nil
}
