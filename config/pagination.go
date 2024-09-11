package config

import (
	"strings"
	libraries "web-service-echo/libraries/db"
)

const (
	PerPage = 2
)

func DefaultPagination(perPage, page int, sort string) *libraries.Pagination {
	sortDefault := "id ASC"
	if false == strings.EqualFold(sort, "") {
		sortDefault = sort
	}

	p := libraries.Pagination{
		Sort:    sortDefault,
		PerPage: perPage,
		Page:    page,
	}

	return &p
}
