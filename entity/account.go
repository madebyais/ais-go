package entity

import "github.com/madebyais/ais-go/infrastructure/database"

type (
	// Account entity
	Account struct {
		Db database.Interface
	}

	// AccountInterface is an interface for Account entity
	AccountInterface interface {
		Register(db database.Interface) AccountInterface

		GetAllAccount(query interface{}) ([]interface{}, error)
		GetAccountByUsername(username string) (interface{}, error)
	}

	// AccountSchema for Account entity
	AccountSchema struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		AccessToken string `json:"accessToken"`
	}
)

func (Account) Register(db database.Interface) AccountInterface {
	return &Account{
		Db: db,
	}
}

func (a Account) GetAllAccount(query interface{}) ([]interface{}, error) {
	return a.Db.FindAll(query)
}

func (a Account) GetAccountByUsername(username string) (interface{}, error) {
	return a.Db.FindOne(map[string]interface{}{"username": username})
}
