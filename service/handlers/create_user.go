package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"my-user-settings-service/internal/service/requests"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Create user")

	newUsers, err := requests.NewCreateUserRequest(r)
	if err != nil {
		log.Errorf("Failed to parse create user request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		ape.Render(w, utils.GetBadRequest())
		return
	}

	err = UsersQ(r).Insert(newUsers)
	if err != nil {
		log.Errorf("Failed to insert new user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ape.Render(w, utils.GetInternalServerError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
