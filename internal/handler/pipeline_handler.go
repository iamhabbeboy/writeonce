package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/entity"
	"github.com/theterminalguy/writeonce/internal/service"
)

type PipelineHandler struct {
	PipelineService *service.PipelineService
}

func NewPipelineHandler() *PipelineHandler {
	return &PipelineHandler{
		PipelineService: service.NewPipelineService(),
	}
}

func (h *PipelineHandler) Search(c echo.Context) error {
	return nil
}

func (h *PipelineHandler) ReadAll(c echo.Context) error {
	return nil
}

func (h *PipelineHandler) ReadByID(c echo.Context) error {
	return nil
}

func (h *PipelineHandler) CreateOne(c echo.Context) error {
	params := new(entity.Pipeline)
	if err := c.Bind(params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err := h.PipelineService.Run(params)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, params)
}

func (h *PipelineHandler) UpdateByID(c echo.Context) error {
	return nil
}

func (h *PipelineHandler) DeleteOne(c echo.Context) error {
	return nil
}
