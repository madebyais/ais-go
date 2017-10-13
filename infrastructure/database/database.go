package database

type (
	// Interface for database
	Interface interface {
		FindOne(query interface{}) (interface{}, error)
		FindAll(query interface{}) ([]interface{}, error)

		Insert(doc interface{}) error

		UpdateOne(query interface{}, opts interface{}) error
		UpdateAll(query interface{}, opts interface{}) error

		RemoveOne(query interface{}) error
		RemoveAll(query interface{}) error
	}
)
