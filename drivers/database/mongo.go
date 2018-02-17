package database

import (
	"time"

	"gopkg.in/mgo.v2"
)

// Mongo is a driver database for MongoDB
type Mongo struct {
	URL         string
	DBName      string
	CollName    string
	Session     *mgo.Session
	tempSession *mgo.Session
}

// Dial is used to create a connection to database
func (m *Mongo) Dial() (Mongo, error) {
	session, err := mgo.Dial(m.URL)
	if err != nil {
		return Mongo{}, err
	}
	defer session.Close()

	m.Session = session.Clone()
	return *m, nil
}

func (m *Mongo) getCollection() *mgo.Collection {
	m.tempSession = m.Session.Clone()
	return m.tempSession.DB(m.DBName).C(m.CollName)
}

// FindOne is used to returns a record
func (m *Mongo) FindOne(query interface{}, fields ...interface{}) (interface{}, error) {
	var result map[string]interface{}

	exec := m.getCollection().Find(query)
	if len(fields) > 0 {
		exec = exec.Select(fields[0])
	}

	err := exec.One(&result)
	defer m.tempSession.Close()

	return result, err
}

// FindAll is used to returns more than one record
func (m *Mongo) FindAll(query interface{}, fields ...interface{}) ([]interface{}, error) {
	var result []interface{}

	exec := m.getCollection().Find(query)
	if len(fields) > 0 {
		exec = exec.Select(fields[0])
	}

	err := exec.All(&result)
	defer m.tempSession.Close()
	return result, err
}

// Insert is used to insert a record
func (m *Mongo) Insert(params ...interface{}) (interface{}, error) {
	err := m.getCollection().Insert(params...)
	defer m.tempSession.Close()
	return params, err
}

// Update is used to update a/or more than one record
func (m *Mongo) Update(query interface{}, params interface{}, multi bool) (interface{}, error) {
	opts := map[string]interface{}{"$set": params}

	if multi {
		_, err := m.getCollection().UpdateAll(query, opts)
		defer m.tempSession.Close()
		return nil, err
	}

	err := m.getCollection().Update(query, opts)
	defer m.tempSession.Close()
	return nil, err
}

// Delete is used to delete a/or more than one record
// if hide is true, then it will be a SOFT-DELETE
func (m *Mongo) Delete(query interface{}, hide bool) error {
	if hide {
		params := make(map[string]interface{})
		params["deleted_at"] = time.Now()
		_, err := m.getCollection().UpdateAll(query, map[string]interface{}{"$set": params})
		defer m.tempSession.Close()
		return err
	}

	_, err := m.getCollection().RemoveAll(query)
	defer m.tempSession.Close()
	return err
}

// Find is used to to returns more than one record
// using custom query
func (m *Mongo) Find(query interface{}, fields ...interface{}) ([]interface{}, error) {
	return nil, nil
}
