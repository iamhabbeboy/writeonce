package handler

import "github.com/labstack/echo/v4"

type TemplateGeneratorHandler struct {
}

func NewTemplateGeneratorHandler() *TemplateGeneratorHandler {
	return &TemplateGeneratorHandler{}
}

func (h *TemplateGeneratorHandler) Create(c echo.Context) error {
	return nil
}
