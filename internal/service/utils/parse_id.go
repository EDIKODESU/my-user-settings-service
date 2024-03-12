package utils

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func ParseUserIDFromURL(r *http.Request) (int64, error) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
