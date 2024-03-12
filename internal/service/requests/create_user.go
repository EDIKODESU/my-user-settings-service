package requests

import (
	"encoding/json"
	"errors"
	"my-user-settings-service/internal/data"
	"my-user-settings-service/internal/resources"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

func NewCreateUserRequest(r *http.Request) ([]data.Users, error) {
	//var req NewUserRequest
	var req resources.CreateUserListResponse
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
			Email:      userData.Attributes.Mail,
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
