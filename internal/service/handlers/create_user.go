package handlers

import (
	"my-user-settings-service/internal/service/requests"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Create user")

	newUsers, err := requests.NewCreateUserRequest(r)
	if err != nil {
		log.Errorf("Failed to parse create user request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = UsersQ(r).Insert(newUsers)
	if err != nil {
		log.Errorf("Failed to insert new user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
