package routers

import (
	"github.com/labstack/echo"
	"github.com/madebyais/ais-go/controllers"
	"github.com/madebyais/ais-go/drivers/database"
	"github.com/madebyais/ais-go/services"
)

// Router is main application router
type Router struct {
	Echo     *echo.Echo
	DB       database.DBInterface
	Services RouterServices
}

// RouterServices structure
// Please update the list of services that you want to use below
type RouterServices struct {
	UserService services.UserInterface
}

// New is used to initlize router
func (r *Router) New(echo *echo.Echo) {
	r.Echo = echo
	r.InitializeControllers(r.Services)
	r.StartServer(`:9000`)
}

// InitializeControllers is used to initialize controller
func (r *Router) InitializeControllers(services RouterServices) {
	userController := &controllers.User{UserService: services.UserService}

	r.Echo.POST(`/user`, userController.CreateAccount)
}

// StartServer is used to start application server
func (r *Router) StartServer(port string) {
	r.Echo.Logger.Fatal(r.Echo.Start(port))
}
