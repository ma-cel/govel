package core

import (
	"github.com/labstack/echo/v4"
	"govel/internal/core/interfaces"
)

type GovelEcho struct {
	*echo.Echo
}

func (g *GovelEcho) GET(path string, controller interfaces.Controller, method string, middleware ...echo.MiddlewareFunc) *echo.Route {
	handler := Handle(controller, method)
	return g.Echo.GET(path, handler, middleware...)
}

func (g *GovelEcho) PUT(path string, controller interfaces.Controller, method string, middleware ...echo.MiddlewareFunc) *echo.Route {
	handler := Handle(controller, method)
	return g.Echo.PUT(path, handler, middleware...)
}

func (g *GovelEcho) PATCH(path string, controller interfaces.Controller, method string, middleware ...echo.MiddlewareFunc) *echo.Route {
	handler := Handle(controller, method)
	return g.Echo.PATCH(path, handler, middleware...)
}

func (g *GovelEcho) POST(path string, controller interfaces.Controller, method string, middleware ...echo.MiddlewareFunc) *echo.Route {
	handler := Handle(controller, method)
	return g.Echo.POST(path, handler, middleware...)
}

func (g *GovelEcho) DELETE(path string, controller interfaces.Controller, method string, middleware ...echo.MiddlewareFunc) *echo.Route {
	handler := Handle(controller, method)
	return g.Echo.DELETE(path, handler, middleware...)
}

func (g *GovelEcho) RESOURCEROUTE(path string, controller interfaces.ResourceController, middleware ...echo.MiddlewareFunc) []*echo.Route {

	getHandler := Handle(controller, "Index")
	postHandler := Handle(controller, "Create")
	deleteHandler := Handle(controller, "Delete")
	patchHandler := Handle(controller, "Update")

	return []*echo.Route{
		g.Echo.GET(path, getHandler, middleware...),
		g.Echo.POST(path, postHandler, middleware...),
		g.Echo.DELETE(path, deleteHandler, middleware...),
		g.Echo.PATCH(path, patchHandler, middleware...),
	}

}
