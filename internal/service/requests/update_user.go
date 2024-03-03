package requests

import (
	"encoding/json"
	"my-user-settings-service/internal/data"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

func UpdateUserRequest(r *http.Request) (data.Users, error) {
	var req data.Users
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return data.Users{}, err
	}

	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return data.Users{}, err
		}
		req.Password = hashedPassword
	}

	return req, nil
}
