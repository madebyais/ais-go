package services

import (
	"errors"

	"github.com/madebyais/ais-go/repositories"
)

// User services
type User struct {
	UserRepository repositories.UserInterface
}

// UserInterface is an interface for user service
type UserInterface interface {
	New() UserInterface
	Register(params interface{}) error
}

// New is used to initialize user service
func (u *User) New() UserInterface {
	return u
}

// Register is used to create new user
func (u *User) Register(params interface{}) error {
	opts := params.(map[string]interface{})

	if _, isExist := opts["fullname"]; !isExist {
		return errors.New(`Fullname is required`)
	}

	_, err := u.UserRepository.Create(opts)

	return err
}
