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
				Path: "projects",
				Only: []Request{
					READ_ALL,
					READ_BY_ID,
					CREATE_ONE,
					UPDATE_BY_ID,
					DELETE_BY_ID,
				},
				Handler:     handler.NewV1ProjectHandler(),
				Middlewares: nil,
			},
			{
				Path: "templates",
				Only: []Request{
					READ_ALL,
					READ_BY_ID,
					CREATE_ONE,
					UPDATE_BY_ID,
					DELETE_BY_ID,
				},
				Handler:     handler.NewV1TemplateHandler(),
				Middlewares: nil,
			},
		},
	}
	apiV1Router.BuildRoutes()
	return e
}
