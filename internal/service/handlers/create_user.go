package handlers

import (
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	"my-user-settings-service/internal/config"
	"my-user-settings-service/internal/service/requests"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log := Log(r)
	log.Info("Create user profile")

	// Підключення до бази даних
	// Створюємо екземпляр конфігурації та отримуємо об'єкт бази даних
	cfg := config.New(kv.MustFromEnv())
	db := cfg.DB()

	// Отримання даних про користувача з запиту
	newUser, err := requests.NewCreateUserRequest(r)
	if err != nil {
		log.Errorf("Failed to parse create user request: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Вставка нового користувача до бази даних
	err = insertNewUser(db, newUser)
	if err != nil {
		log.Errorf("Failed to insert new user: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Відправка відповіді про успішне створення користувача
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func insertNewUser(db *pgdb.DB, newUser requests.NewUser) error {
	query := squirrel.Insert("users").
		Columns("first_name", "second_name", "mail", "login", "password").
		Values(newUser.FirstName, newUser.SecondName, newUser.Email, newUser.Login, newUser.Password).
		Suffix("RETURNING id")

	_, err := db.ExecWithResult(query)
	return err
}
