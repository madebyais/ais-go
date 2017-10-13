package gateway

import "github.com/labstack/echo"

type (
	// Router is the main gateway
	Router struct {
		EchoServer *echo.Echo
		Address    string
	}

	// RouterInterface is an interface for gateway
	RouterInterface interface {
		Init(addr string) RouterInterface
		Start()

		Register(method string, path string, h echo.HandlerFunc)
		RegisterM(method string, path string, h echo.HandlerFunc, m echo.MiddlewareFunc)
	}
)

// Init is use to initialize the gateway
func (Router) Init(addr string) RouterInterface {
	return &Router{
		EchoServer: echo.New(),
		Address:    addr,
	}
}

// Start is use to start the gateway server
func (r *Router) Start() {
	r.EchoServer.Logger.Fatal(r.EchoServer.Start(r.Address))
}

// Register is use to register a route
func (r *Router) Register(method string, path string, h echo.HandlerFunc) {
	switch method {
	case `GET`:
		r.EchoServer.GET(path, h)
	case `POST`:
		r.EchoServer.POST(path, h)
	case `PUT`:
		r.EchoServer.PUT(path, h)
	case `PATCH`:
		r.EchoServer.PATCH(path, h)
	case `DELETE`:
		r.EchoServer.DELETE(path, h)
	}
}

// RegisterM is use to register a route using middlware function
func (r *Router) RegisterM(method string, path string, h echo.HandlerFunc, m echo.MiddlewareFunc) {
	switch method {
	case `GET`:
		r.EchoServer.GET(path, h, m)
	case `POST`:
		r.EchoServer.POST(path, h, m)
	case `PUT`:
		r.EchoServer.PUT(path, h, m)
	case `PATCH`:
		r.EchoServer.PATCH(path, h, m)
	case `DELETE`:
		r.EchoServer.DELETE(path, h, m)
	}
}
