package command

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Page         int `json:"page"`
	Limit        int `json:"limit"`
	Offset       int
	SortBy       string `json:"sort_by"`
	Order        string `json:"order"`
	Keyword      string `json:"keyword"`
	IsPagination bool   `json:"is_pagination"`
}

func GetPagination(r *http.Request) Pagination {

	var pagination Pagination
	pagination.Page = 1
	pagination.Limit = 10
	pagination.SortBy = "created_at"
	pagination.Order = "DESC"
	pagination.IsPagination = true

	query := r.URL.Query()

	if v := query.Get("page"); v != "" {
		pagination.Page, _ = strconv.Atoi(v)
	}

	if v := query.Get("limit"); v != "" {
		pagination.Limit, _ = strconv.Atoi(v)
	}

	if v := query.Get("sort_by"); v != "" {
		pagination.SortBy = v
	}

	if v := query.Get("order"); v != "" {
		pagination.Order = v
	}

	if v := query.Get("keyword"); v != "" {
		pagination.Keyword = v
	}

	pagination.Offset = (pagination.Page * pagination.Limit) - pagination.Limit

	return pagination
}
