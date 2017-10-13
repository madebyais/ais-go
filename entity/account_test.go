package entity

import "testing"

type mockDb struct{}

func (m mockDb) FindOne(query interface{}) (interface{}, error) {
	return map[string]interface{}{"username": "ais"}, nil
}

func (m mockDb) FindAll(query interface{}) ([]interface{}, error) {
	var result []interface{}
	result[0] = map[string]interface{}{"username": "test"}
	return result, nil
}

// Insert is use insert one or more document to a collection
func (m mockDb) Insert(docs interface{}) error {
	return nil
}

func TestGetAccountByUsername(t *testing.T) {
	m := new(mockDb)
	account := new(Account).Register(m)
	data, _ := account.GetAccountByUsername(`ais`)

	if data.(map[string]interface{})["username"] != `ais` {
		t.Fatalf(`Expecting 'ais' got '%s'`, data.(map[string]interface{})["username"])
	}
}
