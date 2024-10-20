package models

import (
	"errors"
	"math"
)

// Config represents the application configuration settings.
type Config struct {
	Host     string
	Port     string
	Env      string // Environment (e.g., development, production)
	PageSize int
	DBURL    string
}

// Envelope is a type alias for a map that can hold any type of value. It is
// used when returning JSON data from an endpoint.
type Envelope map[string]any

// Metadata represents pagination metamodels.
type Metadata struct {
	CurrentPage  int32 `json:"current_page,omitempty"`  // Current page number
	PageSize     int32 `json:"page_size,omitempty"`     // Number of items per page
	FirstPage    int32 `json:"first_page,omitempty"`    // First page number
	LastPage     int32 `json:"last_page,omitempty"`     // Last page number
	TotalRecords int32 `json:"total_records,omitempty"` // Total number of records
}

// CalculateMetadata calculates the pagination metadata values given the total number of records,
// current page, and page size values.
func CalculateMetadata(totalRecords, page, pageSize int32) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     (totalRecords + pageSize - 1) / pageSize,
		TotalRecords: totalRecords,
	}
}

// ConvertInt64to32 converts an int64 value to an int32 value.
// If the input value is negative, it returns 0.
// If the input value exceeds the maximum value of int32, it returns math.MaxInt32.
// Otherwise, it returns the input value cast to int32.
func ConvertInt64to32(x int64) (int32, error) {
	if x < 0 {
		return 0, errors.New("negative value")
	}
	if x > math.MaxInt32 {
		return math.MaxInt32, errors.New("overload value")
	}
	return int32(x), nil
}
