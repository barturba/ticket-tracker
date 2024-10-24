package models

import (
	"fmt"
	"math"
)

// Config represents the application configuration settings.
type Config struct {
	Host      string
	Port      string
	Env       string // Environment (e.g., development, production)
	DBURL     string
	JWTSecret string
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port == "" {
		return fmt.Errorf("port is required")
	}
	if c.Env == "" {
		return fmt.Errorf("env is required")
	}
	if c.JWTSecret == "" {
		return fmt.Errorf("JWT secret is required")
	}
	if c.DBURL == "" {
		return fmt.Errorf("database URL is required")
	}
	return nil
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
func ConvertInt64to32(value int64) (int32, error) {
	if value < math.MinInt32 || value > math.MaxInt32 {
		return 0, fmt.Errorf("value %d is out of int32 range", value)
	}
	return int32(value), nil
}

// SafeDivide performs safe division of two int32 values.
func SafeDivide(dividend, divisor int32) (int32, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return 0, fmt.Errorf("integer overflow")
	}
	return dividend / divisor, nil
}
