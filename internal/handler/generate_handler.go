package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/service"
)

type GenerateHandler struct {
	TemplateService *service.TemplateService
}

func NewGenerateHandler() *GenerateHandler {
	return &GenerateHandler{
		TemplateService: service.NewTemplateService(),
	}
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
	params := new(service.GenerateParams)
	if err := c.Bind(params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	templateID, err := uuid.Parse(c.Param("template_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	params.TemplateID = templateID
	tmplStr, err := h.TemplateService.Generate(*params)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tmplStr)
}

func (h *GenerateHandler) UpdateByID(c echo.Context) error {
	return nil
}

func (h *GenerateHandler) DeleteOne(c echo.Context) error {
	return nil
}
