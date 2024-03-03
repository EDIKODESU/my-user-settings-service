package requests

import (
	"encoding/json"
	"errors"
	"my-user-settings-service/internal/data"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

// NewUser описує структуру нового користувача
//type NewUser struct {
//	FirstName  string `json:"first_name"`
//	SecondName string `json:"second_name"`
//	Login      string `json:"login"`
//	Email      string `json:"mail"`
//	Password   string `json:"password"`
//}

// NewUserRequest описує структуру запиту на створення нового користувача
type NewUserRequest struct {
	Data []struct {
		Attributes struct {
			FirstName  string `json:"first_name"`
			SecondName string `json:"second_name"`
			Login      string `json:"login"`
			Email      string `json:"mail"`
			Password   string `json:"password"`
		} `json:"attributes"`
	} `json:"data"`
}

func NewCreateUserRequest(r *http.Request) ([]data.Users, error) {
	var req NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	var newUsers []data.Users
	for _, userData := range req.Data {
		newUser := data.Users{
			FirstName:  userData.Attributes.FirstName,
			SecondName: userData.Attributes.SecondName,
			Login:      userData.Attributes.Login,
			Email:      userData.Attributes.Email,
			Password:   userData.Attributes.Password,
		}

		err := validateNewUser(newUser)
		if err != nil {
			return nil, err
		}

		hashedPassword, err := utils.HashPassword(newUser.Password)
		if err != nil {
			return nil, err
		}
		newUser.Password = hashedPassword

		newUsers = append(newUsers, newUser)
	}

	return newUsers, nil
}

func validateNewUser(user data.Users) error {
	if user.FirstName == "" {
		return errors.New("first_name field cannot be empty")
	}
	if user.SecondName == "" {
		return errors.New("second_name field cannot be empty")
	}
	if user.Login == "" {
		return errors.New("login field cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email field cannot be empty")
	}
	if user.Password == "" {
		return errors.New("password field cannot be empty")
	}
	return nil
}
