package database

// DBInterface is an interface for database drivers
// using CRUSH (CREATE, READ, UPDATE, SOFT DEL, HARD DEL) methodology
type DBInterface interface {
	FindOne(query interface{}, fields ...interface{}) (interface{}, error)
	FindAll(query interface{}, fields ...interface{}) ([]interface{}, error)
	Insert(params interface{}) (interface{}, error)
	Update(query interface{}, params interface{}, multi bool) (interface{}, error)
	Delete(query interface{}, hide bool) error

	// For custom query only
	Find(query interface{}, fields ...interface{}) ([]interface{}, error)
}
