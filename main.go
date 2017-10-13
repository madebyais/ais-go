package main

import (
	"github.com/madebyais/ais-go/adapter"
	"github.com/madebyais/ais-go/controller"
	"github.com/madebyais/ais-go/entity"
	"github.com/madebyais/ais-go/infrastructure/database"
	"github.com/madebyais/ais-go/infrastructure/gateway"
	"github.com/madebyais/ais-go/infrastructure/repository"
)

func main() {
	// initialize mongodb
	mongo := new(repository.MongoDb)
	mongo.Register(`localhost:27017`)
	mongo.Set(`fonte`, `accounts`)

	// register mongodb to database interface
	var db database.Interface = mongo

	// initialize entity
	account := new(entity.Account).Register(db)

	// register entity into adapter
	adapterSchema := new(adapter.Schema)
	adapterSchema.Account = account

	// initialize controller and register adapter into controller
	accountCtrl := new(controller.Account).Register(adapterSchema)

	// initialize router
	router := new(gateway.Router).Init(`:9000`)

	// register controller to a router
	router.Register(`GET`, `/account`, accountCtrl.GetAccount)

	// start application
	router.Start()
}
