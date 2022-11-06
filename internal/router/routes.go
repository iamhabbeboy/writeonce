package router

import (
	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/handler"
)

func DefineRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", handler.IndexHandler)

	apiV1Router := &Router{
		group:       e.Group("/api"),
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
				Handler:     handler.NewProjectHandler(),
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
				Handler:     handler.NewTemplateHandler(),
				Middlewares: nil,
			},
			{
				Path: "generate/projects/:project_id/templates/:template_id",
				Only: []Request{
					CREATE_ONE,
				},
				Handler: handler.NewGenerateHandler(),
			},
		},
	}
	apiV1Router.BuildRoutes()
	return e
}
