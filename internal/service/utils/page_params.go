package utils

import (
	"net/http"
	"net/url"
	"strconv"
)

type Links struct {
	Self string `json:"self"`
	Next string `json:"next,omitempty"`
	Last string `json:"last,omitempty"`
}

func ParsePaginationParams(r *http.Request) (int, int) {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	// Перетворення рядків параметрів у цілі числа
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 {
		perPage = 10
	}
	return page, perPage
}

func GetNextPage(baseURL string, page, totalPages, perPage int) (string, error) {
	if page < totalPages {
		nextPage := page + 1
		nextLink, err := buildPaginationLink(baseURL, nextPage, perPage)
		return nextLink, err
	}
	return "", nil
}

func GetLastPage(baseURL string, totalPages, perPage int) (string, error) {
	lastLink, err := buildPaginationLink(baseURL, totalPages, perPage)
	return lastLink, err
}

func buildPaginationLink(baseURL string, page, perPage int) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	// Параметри пагінації
	queryParams := u.Query()
	queryParams.Set("page[offset]", strconv.Itoa((page-1)*perPage))
	queryParams.Set("page[limit]", strconv.Itoa(perPage))
	u.RawQuery = queryParams.Encode()

	return u.String(), nil
}
