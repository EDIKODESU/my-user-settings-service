package pg

import (
	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
	"my-user-settings-service/internal/data"
	"my-user-settings-service/internal/service/requests"
)

func NewUsersQ(db *pgdb.DB) data.UsersQ {
	return &usersQ{
		db: db.Clone(),
	}
}

type usersQ struct {
	db *pgdb.DB
}

func (q *usersQ) New() data.UsersQ {
	return NewUsersQ(q.db)
}

func (q *usersQ) Insert(newUsers []requests.NewUser) error {
	// Логіка для вставки користувача у базу даних
	return nil
}

func (q *usersQ) Select(page, perPage int) ([]data.Users, error) {
	offset := (page - 1) * perPage
	query := squirrel.Select("id", "first_name", "second_name", "mail", "login", "password").
		From("users").
		Offset(uint64(offset)).
		Limit(uint64(perPage))

	var users []data.Users
	err := q.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (q *usersQ) Update(updateUser requests.UpdateUser) error {
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

	_, err := q.db.ExecWithResult(query)
	return err
}

func (q *usersQ) Delete(userID int64) error {
	query := squirrel.Delete("users").Where(squirrel.Eq{"id": userID})

	_, err := q.db.ExecWithResult(query)
	if err != nil {
		return err
	}

	return nil
}

func (q *usersQ) Count() (int, error) {
	query := squirrel.Select("COUNT(*)").
		From("users")

	var count int
	err := q.db.Get(&count, query)
	if err != nil {
		return 0, err
	}

	return count, nil
}
