package data

type UsersQ interface {
	New() UsersQ
	Insert(newUsers []Users) error
	Select(page, perPage int) ([]Users, error)
	Update(updateUser Users) error
	Delete(userID int64) error
	Count() (int, error)
}

type Users struct {
	ID         int64  `db:"id" json:"id"`
	FirstName  string `db:"first_name" json:"first_name"`
	SecondName string `db:"second_name" json:"second_name"`
	Login      string `db:"login" json:"login"`
	Email      string `db:"mail" json:"mail"`
	Password   string `db:"password" json:"password"`
}
