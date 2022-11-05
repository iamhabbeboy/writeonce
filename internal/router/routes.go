package router

import (
	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/handler"
)

func DefineRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", handler.IndexHandler)

	apiV1Router := &Router{
		group:       e.Group("/api/v1"),
		middlewares: []echo.MiddlewareFunc{},
		handlers: []RouteHandler{
			{
				Path:        "projects",
				Only:        []Request{CREATE_ONE, READ_ALL},
				Handler:     handler.NewV1ProjectHandler(),
				Middlewares: nil,
			},
		},
	}
	apiV1Router.BuildRoutes()
	return e
}
