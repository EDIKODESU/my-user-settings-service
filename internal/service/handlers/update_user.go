package handlers

import (
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"my-user-settings-service/internal/service/requests"
	"net/http"

	"my-user-settings-service/internal/config"
	"my-user-settings-service/internal/service/utils"

	"gitlab.com/distributed_lab/kit/kv"
)

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Getting user profile")

	cfg := config.New(kv.MustFromEnv())
	db := cfg.DB()

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
	err = UpdateDataUser(db, updateUser)
	if err != nil {
		log.Errorf("Failed to update data of user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User update successfully"))
}

func UpdateDataUser(db *pgdb.DB, updateUser requests.UpdateUser) error {
	query := squirrel.Update("users")

	if updateUser.FirstName != "" {
		query = query.Set("first_name", updateUser.FirstName)
	}

	if updateUser.SecondName != "" {
		query = query.Set("second_name", updateUser.SecondName)
	}

	if updateUser.Login != "" {
		query = query.Set("login", updateUser.Login)
	}

	if updateUser.Email != "" {
		query = query.Set("mail", updateUser.Email)
	}

	if updateUser.Password != "" {
		query = query.Set("password", updateUser.Password)
	}

	query = query.Where(squirrel.Eq{"id": updateUser.ID})

	_, err := db.ExecWithResult(query)
	return err
}
