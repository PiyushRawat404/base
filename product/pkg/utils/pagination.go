package utils

import (
	"net/url"
	"strconv"
)

type PaginationParams struct {
	Page   int
	Limit  int
	Offset int
}

func ParsePagination(values url.Values) PaginationParams {
	page := parsePositiveInt(values.Get("page"), 1)
	limit := parsePositiveInt(values.Get("limit"), 10)
	if limit > 100 {
		limit = 100
	}

	return PaginationParams{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}

func CalculateTotalPages(total, limit int) int {
	if total == 0 {
		return 0
	}
	return (total + limit - 1) / limit
}

func parsePositiveInt(value string, fallback int) int {
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}
