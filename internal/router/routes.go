package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/handler"
)

func DefineRoutes(e *echo.Echo) *echo.Echo {
	apiRouter := &Router{
		Group: e.Group("api/"),
		Routes: []*Route{
			{
				Path: "projects",
				Methods: []string{
					http.MethodPost,
				},
				Handler: handler.NewProjectHandler(),
			},
		},
	}
	apiRouter.Register(e)

	return e
}
