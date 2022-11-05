package router

import (
	"github.com/labstack/echo/v4"
)

type RouteHandler interface {
	Create(echo.Context) error
}

type Route struct {
	Path        string
	Methods     []string
	Handler     RouteHandler
	Middlewares []echo.MiddlewareFunc
}

func (r *Route) Register(e *echo.Group) {
	e.Add(r.Methods[0], r.Path, r.Handler.Create, r.Middlewares...)
}

type Router struct {
	Group  *echo.Group
	Routes []*Route
}

func (r *Router) Register(e *echo.Echo) {
	for _, route := range r.Routes {
		route.Register(r.Group)
	}
}
