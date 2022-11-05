package handler

import "github.com/labstack/echo/v4"

type ProjectHandler struct {
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

func (h *ProjectHandler) Create(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
