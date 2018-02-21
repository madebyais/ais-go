package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/madebyais/ais-go/services"
)

// User controller
type User struct {
	UserService services.UserInterface
}

// CreateAccount is a controller handler to create new account
func (u *User) CreateAccount(ctx echo.Context) error {
	params := make(map[string]interface{})

	if errJSONBind := ctx.Bind(&params); errJSONBind != nil {
		return errJSONBind
	}

	err := u.UserService.Register(params)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"status": "success"})
}
