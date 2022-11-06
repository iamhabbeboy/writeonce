package handler

import (
	"github.com/labstack/echo/v4"
)

type GenerateHandler struct {
}

func NewGenerateHandler() *GenerateHandler {
	return &GenerateHandler{}
}

func (h *GenerateHandler) Create(c echo.Context) error {
	return nil
}

func (h *GenerateHandler) Search(c echo.Context) error {
	return nil
}

func (h *GenerateHandler) ReadAll(c echo.Context) error {
	return nil
}

func (h *GenerateHandler) ReadByID(c echo.Context) error {
	return nil
}

func (h *GenerateHandler) CreateOne(c echo.Context) error {
	return c.JSON(200, "Hello World")
}

func (h *GenerateHandler) UpdateByID(c echo.Context) error {
	return nil
}

func (h *GenerateHandler) DeleteOne(c echo.Context) error {
	return nil
}
