package handlers

import (
	"gitlab.com/distributed_lab/ape"
	"my-user-settings-service/internal/service/utils"
	"net/http"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Getting user profile")

	var userID, err = utils.ParseUserIDFromURL(r)
	if err != nil {
		log.Errorf("Failed to parse user ID from URL: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		ape.Render(w, utils.GetBadRequest())
		return
	}

	err = UsersQ(r).Delete(userID)
	if err != nil {
		log.Errorf("Failed to update data of user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		ape.Render(w, utils.GetInternalServerError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
