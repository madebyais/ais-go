package repository

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// MongoDb repository
type MongoDb struct {
	Session *mgo.Session
	Db      string
	Coll    string
}

// Register is use to initialize mongodb connection
func (m *MongoDb) Register(url string) {
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println(err)
	}

	m.Session = session
}

// Set is use to select db and collection
func (m *MongoDb) Set(db string, coll string) {
	m.Db = db
	m.Coll = coll
}

// GetCurrentCollection is use to get current state of selected collection
func (m *MongoDb) GetCurrentCollection() *mgo.Collection {
	return m.Session.DB(m.Db).C(m.Coll)
}

// FindOne is use to find a record based on submitted query
func (m *MongoDb) FindOne(query interface{}) (interface{}, error) {
	var result map[string]interface{}
	err := m.GetCurrentCollection().Find(query).One(&result)
	return result, err
}

// FindAll is use to find more than one record based on submitted query
func (m *MongoDb) FindAll(query interface{}) ([]interface{}, error) {
	var tempResult []map[string]interface{}
	err := m.GetCurrentCollection().Find(query).All(&tempResult)

	var result []interface{}
	for _, item := range tempResult {
		result = append(result, item)
	}

	return result, err
}

// Insert is use insert one or more document to a collection
func (m *MongoDb) Insert(doc interface{}) error {
	return m.GetCurrentCollection().Insert(doc)
}

// UpdateOne is use to update a record that match a query
func (m *MongoDb) UpdateOne(query interface{}, opts interface{}) error {
	return m.GetCurrentCollection().Update(query, opts)
}

// UpdateAll is use to update more than one record that match a query
func (m *MongoDb) UpdateAll(query interface{}, opts interface{}) error {
	_, err := m.GetCurrentCollection().UpdateAll(query, opts)
	return err
}

// RemoveOne is use to remove a record that match a query
func (m *MongoDb) RemoveOne(query interface{}) error {
	return m.GetCurrentCollection().Remove(query)
}

// RemoveAll is use to remove more than one record that match a query
func (m *MongoDb) RemoveAll(query interface{}) error {
	_, err := m.GetCurrentCollection().RemoveAll(query)
	return err
}
