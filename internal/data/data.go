package data

// Config represents the application configuration settings.
type Config struct {
	Host     string
	Port     string
	Env      string // Environment (e.g., development, production)
	PageSize int
}

// Envelope is a type alias for a map that can hold any type of value. It is
// used when returning JSON data from an endpoint.
type Envelope map[string]any

// Metadata represents pagination metadata.
type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`  // Current page number
	PageSize     int `json:"page_size,omitempty"`     // Number of items per page
	FirstPage    int `json:"first_page,omitempty"`    // First page number
	LastPage     int `json:"last_page,omitempty"`     // Last page number
	TotalRecords int `json:"total_records,omitempty"` // Total number of records
}

// CalculateMetadata calculates the pagination metadata values given the total number of records,
// current page, and page size values.
func CalculateMetadata(totalRecords, page, pageSize int) Metadata {
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
