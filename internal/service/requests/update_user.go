package requests

import (
	"encoding/json"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

type UpdateUser struct {
	ID         int64
	FirstName  string `json:"first_name" db:"first_name"`
	SecondName string `json:"second_name" db:"second_name"`
	Login      string `json:"login" db:"login"`
	Email      string `json:"mail" db:"mail"`
	Password   string `json:"password" db:"password"`
}

func UpdateUserRequest(r *http.Request) (UpdateUser, error) {
	var req UpdateUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return UpdateUser{}, err
	}

	if req.Password != "" {
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			return UpdateUser{}, err
		}
		req.Password = hashedPassword
	}

	return req, nil
}
