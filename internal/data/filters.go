package data

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/barturba/ticket-tracker/validator"
)

// Filters represents the pagination and sorting parameters.
type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

// Limit returns the number of items per page.
func (f Filters) Limit() int {
	return f.PageSize
}

// Offset returns the offset for the current page.
func (f Filters) Offset() int {
	return (f.Page - 1) * f.PageSize
}

// SortColumn returns the sort field if it is in the safelist, otherwise it panics.
func (f Filters) SortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

// SortDirection returns the sort direction (ASC or DESC) based on the sort field prefix.
func (f Filters) SortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

// ReadInt reads an integer query parameter from the URL values. If the parameter is missing or invalid,
// it returns the default value and adds an error to the validator.
func ReadInt(qs url.Values, key string, defaultValue int, v *validator.Validator) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}

	return i
}

// ReadString reads a string query parameter from the URL values. If the parameter is missing,
// it returns the default value.
func ReadString(qs url.Values, key string, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}

// ValidateFilters validates the pagination and sorting parameters in the Filters struct.
func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be no more than ten million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be no more than 100")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}

// ValidateFiltersGetAll validates the pagination and sorting parameters in the Filters struct
// for retrieving all records in one shot.
func ValidateFiltersGetAll(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be no more than ten million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 10_000_000, "page_size", "must be no more than ten million")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
