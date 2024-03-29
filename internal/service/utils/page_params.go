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

func GetFirstPage(r *http.Request, perPage int) (string, error) {
	firstLink, err := buildPaginationLink(r.URL.Path, r.Host, 1, perPage)
	return firstLink, err
}

func GetPrevPage(r *http.Request, page, perPage int) (string, error) {
	if page > 1 {
		println(page)
		prevPage := page - 1
		prevLink, err := buildPaginationLink(r.URL.Path, r.Host, prevPage, perPage)
		return prevLink, err
	}
	return "", nil
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
	if totalPages <= 1 {
		return GetSelfPage(r, 1, perPage)
	}
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
