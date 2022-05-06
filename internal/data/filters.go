package data

import (
	"strings"

	"greenlight.arammikayelyan.dev/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafelist []string
}

// sortColumn checks the client-provided Sort field matches one of the entries
// in our safelist and if it does, extract the column name from the Sort field
// by stripping the leading hyphen character (if one exists).
func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort paramter: " + f.Sort)
}

// sortDirection returns the sort direction ("ASC" or "DESC") depending on the
// prefix character of the Sort field.
func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be a maximum of 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")

	// Chech that the sort paramter matches a value in the safelist.
	v.Check(validator.In(f.Sort, f.SortSafelist...), "sort", "invalid sort value")
}
