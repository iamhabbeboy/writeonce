package handler

import (
	"github.com/labstack/echo/v4"
)

type V1ProjectHandler struct {
}

func NewV1ProjectHandler() *V1ProjectHandler {
	return &V1ProjectHandler{}
}

func (h *V1ProjectHandler) Search(c echo.Context) error {
	return nil
}

func (h *V1ProjectHandler) ReadAll(c echo.Context) error {
	return nil
}

func (h *V1ProjectHandler) ReadByID(c echo.Context) error {
	return nil
}

func (h *V1ProjectHandler) CreateOne(c echo.Context) error {
	return c.String(200, "CreateOne")
}

func (h *V1ProjectHandler) UpdateByID(c echo.Context) error {
	return nil
}

func (h *V1ProjectHandler) DeleteOne(c echo.Context) error {
	return nil
}
