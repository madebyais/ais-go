package repositories

import (
	"testing"

	"github.com/madebyais/ais-go/drivers/database"
	"github.com/stretchr/testify/assert"
)

func initUserRepo() UserInterface {
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

	return new(User).New(dbi)
}

func TestUserRepoCreateSuccess(t *testing.T) {
	userRepo := initUserRepo()

	params := map[string]interface{}{"uuid": "uuid", "fullname": "ais test"}

	d, e := userRepo.Create(params)
	assert.Nil(t, e)
	assert.NotNil(t, d)
}

func TestUserFindByUUIDSuccess(t *testing.T) {
	userRepo := initUserRepo()

	d, e := userRepo.FindByUUID("uuid")
	assert.Nil(t, e)
	assert.NotNil(t, d)
}
