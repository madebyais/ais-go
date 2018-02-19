package repositories

import (
	"github.com/madebyais/ais-go/drivers/database"
)

// User repository
type User struct {
	DB database.DBInterface
}

// UserInterface is an interface for user repository
type UserInterface interface {
	Create(params interface{}) (interface{}, error)
	FindByUUID(uuid string) (interface{}, error)
}

// New is used to initliaze user repository
func (*User) New(db database.DBInterface) UserInterface {
	db.SetCollection(`user`)

	return &User{
		DB: db,
	}
}

// Create is used to create new user
func (u *User) Create(params interface{}) (interface{}, error) {
	data, err := u.DB.Insert(params)
	return data, err
}

// FindByUUID is used to find user data by providing uuid
func (u *User) FindByUUID(uuid string) (interface{}, error) {
	opts := map[string]interface{}{"uuid": uuid}

	d, e := u.DB.FindOne(opts)
	return d, e
}
