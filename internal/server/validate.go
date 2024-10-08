package server

import "github.com/barturba/ticket-tracker/validator"

type Filter struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

func validateFilters(v *validator.Validator, f Filter) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be a maximum of ten million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize < 100, "page_size", "must be less than 100")

	v.Check(validator.PermittedValue(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
