package repositories

import (
	"errors"

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
	return &User{
		DB: db,
	}
}

// Create is used to create new user
func (u *User) Create(params interface{}) (interface{}, error) {
	opts := params.(map[string]interface{})

	if _, isExist := opts["fullname"]; !isExist {
		return nil, errors.New(`Fullname is required`)
	}

	data, err := u.DB.Insert(opts)
	return data, err
}

// FindByUUID is used to find user data by providing uuid
func (u *User) FindByUUID(uuid string) (interface{}, error) {
	opts := map[string]interface{}{"uuid": uuid}

	d, e := u.DB.FindOne(opts)
	return d, e
}
