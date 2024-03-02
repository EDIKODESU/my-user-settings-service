package handlers

import (
	"my-user-settings-service/internal/service/requests"
	"net/http"

	"my-user-settings-service/internal/service/utils"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Updating user")

	// Отримання нові дані про користувача з запиту
	updateUser, err := requests.UpdateUserRequest(r)
	if err != nil {
		log.Errorf("Failed to parse update user request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	updateUser.ID, err = utils.ParseUserIDFromURL(r)
	if err != nil {
		log.Errorf("Failed to parse user ID from URL: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Оновлення даних користувача у базі даних
	err = UsersQ(r).Update(updateUser)
	if err != nil {
		log.Errorf("Failed to update data of user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User update successfully"))
}
