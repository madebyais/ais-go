package services

import (
	"testing"

	"github.com/madebyais/ais-go/drivers/database"
	"github.com/madebyais/ais-go/repositories"
	"github.com/stretchr/testify/assert"
)

func initUserService() UserInterface {
	dbm := &database.Mongo{
		URL:      `localhost:27017`,
		DBName:   `ais_test_go`,
		CollName: `user`,
	}

	mongo, err := dbm.Dial()
	if err != nil {
		panic(err)
	}

	var dbi database.DBInterface = &mongo
	userRepo := new(repositories.User).New(dbi)

	userService := &User{
		UserRepository: userRepo,
	}

	return userService.New()
}

func TestUserServiceRegisterFailedIfFullnameNotProvided(t *testing.T) {
	userService := initUserService()

	params := map[string]interface{}{"uuid": "uuid"}

	e := userService.Register(params)
	assert.NotNil(t, e)
}

func TestUserServiceRegisterSuccess(t *testing.T) {
	userService := initUserService()

	params := map[string]interface{}{"uuid": "uuid", "fullname": "ais doe"}

	e := userService.Register(params)
	assert.Nil(t, e)
}
