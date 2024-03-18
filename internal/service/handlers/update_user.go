package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"my-user-settings-service/internal/service/requests"
	"net/http"

	"my-user-settings-service/internal/service/utils"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Updating user")

	updateUser, err := requests.UpdateUserRequest(r)
	if err != nil {
		log.Errorf("Failed to parse update user request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		ape.Render(w, utils.GetBadRequest())
		return
	}

	updateUser.Id, err = utils.ParseUserIDFromURL(r)
	if err != nil {
		log.Errorf("Failed to parse user ID from URL: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		ape.Render(w, utils.GetBadRequest())
		return
	}

	err = UsersQ(r).Update(updateUser)
	if err != nil {
		log.Errorf("Failed to update data of user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ape.Render(w, utils.GetInternalServerError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
