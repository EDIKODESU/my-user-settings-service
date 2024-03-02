package data

import "my-user-settings-service/internal/service/requests"

type UsersQ interface {
	New() UsersQ
	Insert(newUsers []requests.NewUser) error
	Select(page, perPage int) ([]Users, error)
	Update(updateUser requests.UpdateUser) error
	Delete(userID int64) error
	Count() (int, error)
}

//	type Users struct {
//		ID         int64  `jsonapi:"primary,users"`
//		FirstName  string `db:"first_name" jsonapi:"attr,first_name"`
//		SecondName string `db:"second_name" jsonapi:"attr,second_name"`
//		Login      string `db:"login" jsonapi:"attr,login"`
//		Email      string `db:"mail" jsonapi:"attr,mail"`
//		Pass       string `db:"password" jsonapi:"attr,password"`
//	}
type Users struct {
	ID         int64  `db:"id" json:"id"`
	FirstName  string `db:"first_name" json:"first_name"`
	SecondName string `db:"second_name" json:"second_name"`
	Login      string `db:"login" json:"login"`
	Email      string `db:"mail" json:"mail"`
	Pass       string `db:"password" json:"password"`
}
