package utils

import (
	"net/url"
	"strings"
)

type ProductFilter struct {
	Search    string
	SortField string
	SortOrder string
}

func ParseProductFilter(values url.Values) ProductFilter {
	sortValue := strings.TrimSpace(values.Get("sort"))
	field := "created_at"
	order := "DESC"

	switch {
	case strings.HasPrefix(sortValue, "-"):
		field = strings.TrimPrefix(sortValue, "-")
		order = "DESC"
	case sortValue != "":
		field = sortValue
		order = "ASC"
	}

	field = sanitizeSortField(field)

	return ProductFilter{
		Search:    strings.TrimSpace(values.Get("search")),
		SortField: field,
		SortOrder: order,
	}
}

func sanitizeSortField(field string) string {
	switch field {
	case "name", "price", "quantity", "created_at":
		return field
	default:
		return "created_at"
	}
}
