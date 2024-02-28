package handlers

import (
	"encoding/json"
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	"my-user-settings-service/internal/config"
	"net/http"
)

type User struct {
	ID         int64
	FirstName  string `db:"first_name"`
	SecondName string `db:"second_name"`
	Login      string `db:"login"`
	Email      string `db:"mail"`
	Pass       string `db:"password"`
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Getting user profile")

	// Підключення до бази даних
	// Створюємо екземпляр конфігурації та отримуємо об'єкт бази даних
	cfg := config.New(kv.MustFromEnv())
	db := cfg.DB()

	// Виконання запиту до бази даних для отримання користувачів
	users, err := getAllUsers(db)

	if err != nil {
		log.Errorf("Failed to fetch profiles from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Відправка відповіді з JSON-представленням профілю
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getAllUsers(db *pgdb.DB) ([]User, error) {
	query := squirrel.Select("id", "first_name", "second_name", "mail", "login", "password").From("users")

	var users []User
	err := db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}
