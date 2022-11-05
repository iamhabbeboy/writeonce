package handler

import "github.com/labstack/echo/v4"

type TemplateHandler struct {
}

func NewTemplateHandler() *TemplateHandler {
	return &TemplateHandler{}
}

func (h *TemplateHandler) Create(c echo.Context) error {
	return nil
}
