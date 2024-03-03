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
	pageStr := r.URL.Query().Get("page[offset]")
	perPageStr := r.URL.Query().Get("page[limit]")

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

func GetSelfPage(r *http.Request, page, perPage int) (string, error) {
	self, err := buildPaginationLink(r.URL.Path, r.Host, page, perPage)
	return self, err
}

func GetNextPage(r *http.Request, page, totalPages, perPage int) (string, error) {
	if page < totalPages {
		nextPage := page + 1
		nextLink, err := buildPaginationLink(r.URL.Path, r.Host, nextPage, perPage)
		return nextLink, err
	}
	return "", nil
}

func GetLastPage(r *http.Request, totalPages, perPage int) (string, error) {
	lastLink, err := buildPaginationLink(r.URL.Path, r.Host, totalPages, perPage)
	return lastLink, err
}

func buildPaginationLink(baseURL, host string, page, perPage int) (string, error) {
	u, err := url.Parse(baseURL)

	if err != nil {
		return "", err
	}

	queryParams := u.Query()
	queryParams.Set("page[offset]", strconv.Itoa(page))
	queryParams.Set("page[limit]", strconv.Itoa(perPage))
	u.RawQuery = queryParams.Encode()

	link := "http://" + host + u.String()

	return link, nil
}
