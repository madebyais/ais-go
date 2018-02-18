package database_test

import (
	"testing"

	"github.com/madebyais/ais-go/drivers/database"
	"github.com/stretchr/testify/assert"
)

var UUID1 string = "5953802f-ca4a-4c71-bccd-24c50202212c"
var UUID2 string = "2f387a74-9d1f-4738-9ac2-89c8c3469a93"

func mongoDial() (database.Mongo, error) {
	m := database.Mongo{
		URL:      "localhost:27017",
		DBName:   "ais_go_test",
		CollName: "sample",
	}
	return m.Dial()
}

func TestInitMongo(t *testing.T) {
	m := &database.Mongo{
		URL:      "localhost:27017",
		DBName:   "ais_go_test",
		CollName: "sample",
	}

	assert.Equal(t, "ais_go_test", m.DBName)
}

func TestMongoDial(t *testing.T) {
	mongo, _ := mongoDial()
	assert.Equal(t, nil, mongo.Session.Ping())
}

func TestMongoDialError(t *testing.T) {
	m := database.Mongo{
		URL:      "localhost:27018",
		DBName:   "ais_go_test",
		CollName: "sample",
	}

	_, err := m.Dial()
	assert.NotNil(t, err)
}

func TestMongoInsert(t *testing.T) {
	mongo, _ := mongoDial()
	params := map[string]interface{}{"uuid": UUID1, "name": "ais"}
	params2 := map[string]interface{}{"uuid": UUID2, "name": "john doe"}

	data, err := mongo.Insert(params, params2)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, data)
}

func TestMongoFindOne(t *testing.T) {
	mongo, _ := mongoDial()
	query := map[string]interface{}{"name": "ais"}

	data, err := mongo.FindOne(query)
	assert.Equal(t, nil, err)
	assert.NotEmpty(t, data)
}

func TestMongoFindAll(t *testing.T) {
	mongo, _ := mongoDial()
	query := map[string]interface{}{"name": "ais"}

	data, err := mongo.FindAll(query)
	assert.Equal(t, nil, err)
	assert.NotNil(t, data)
	assert.NotEqual(t, 0, len(data))
}

func TestMongoUpdateOne(t *testing.T) {
	mongo, _ := mongoDial()
	query := map[string]interface{}{"uuid": UUID1}
	params := map[string]interface{}{"uuid": UUID1, "name": "ais go"}
	_, err := mongo.Update(query, params, false)

	assert.Nil(t, err)
	query = map[string]interface{}{"uuid": UUID1}

	data, err := mongo.FindOne(query)
	assert.Nil(t, err)

	mapdata := data.(map[string]interface{})
	assert.Equal(t, "ais go", mapdata["name"])
}

func TestMongoUpdateAll(t *testing.T) {
	mongo, _ := mongoDial()
	query := map[string]interface{}{"uuid": UUID1}
	params := map[string]interface{}{"uuid": UUID1, "name": "ais go"}
	_, err := mongo.Update(query, params, true)

	assert.Nil(t, err)
	query = map[string]interface{}{"uuid": UUID1}

	data, err := mongo.FindOne(query)
	assert.Nil(t, err)

	mapdata := data.(map[string]interface{})
	assert.Equal(t, "ais go", mapdata["name"])
}

func TestMongoDeleteSoft(t *testing.T) {
	mongo, _ := mongoDial()
	query := map[string]interface{}{"uuid": UUID1}

	err := mongo.Delete(query, true)
	assert.Nil(t, err)
}

func TestMongoDeleteHard(t *testing.T) {
	mongo, _ := mongoDial()
	query := map[string]interface{}{"uuid": UUID1}

	err := mongo.Delete(query, false)
	assert.Nil(t, err)
}

func TestMongoFind(t *testing.T) {
	mongo, _ := mongoDial()

	data, err := mongo.Find(nil)
	assert.Nil(t, err)
	assert.Nil(t, data)
}
