package controller

import (
	"net/http"

	"github.com/ais-go/adapter"
	"github.com/labstack/echo"
)

type (
	// Account is the controller for account
	Account struct {
		Adapter *adapter.Schema
	}

	// AccountInterface is an interface for account controller
	AccountInterface interface {
		Register(adapter *adapter.Schema) AccountInterface

		GetAccount(ctx echo.Context) error
	}
)

// Register is use to initialize account controller
func (Account) Register(adapter *adapter.Schema) AccountInterface {
	return &Account{
		Adapter: adapter,
	}
}

// GetAccount is a controller for get account data
func (a *Account) GetAccount(ctx echo.Context) error {
	data, err := a.Adapter.Account.GetAllAccount(nil)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, data)
}
