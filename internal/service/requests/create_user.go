package requests

import (
	"encoding/json"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

type NewUser struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	SecondName string `json:"second_name" db:"second_name"`
	Login      string `json:"login" db:"login"`
	Email      string `json:"mail" db:"mail"`
	Password   string `json:"password" db:"password"`
}

func NewCreateUserRequest(r *http.Request) (NewUser, error) {
	var req NewUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return NewUser{}, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return NewUser{}, err
	}
	req.Password = hashedPassword

	return req, nil
}
