package handlers

import (
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
		return
	}

	err = UsersQ(r).Delete(userID)
	if err != nil {
		log.Errorf("Failed to update data of user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User delete successfully"))
}
