package main

import (
	"github.com/madebyais/ais-go/services"

	"github.com/labstack/echo"
	"github.com/madebyais/ais-go/repositories"
	"github.com/madebyais/ais-go/routers"

	"github.com/madebyais/ais-go/drivers/database"
)

func main() {
	mongo, err := (&database.Mongo{URL: `localhost:27017`, DBName: `ais_go_test`}).Dial()
	if err != nil {
		panic(err)
	}

	var db database.DBInterface = &mongo

	userRepo := new(repositories.User).New(db)

	server := echo.New()
	router := new(routers.Router)
	router.Services.UserService = &services.User{UserRepository: userRepo}
	router.Services.UserService.New()

	router.New(server)
}
