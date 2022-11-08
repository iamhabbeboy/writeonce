package handler

import "github.com/labstack/echo/v4"

type PipelineHandler struct {
}

func NewPipelineHandler() *PipelineHandler {
	return &PipelineHandler{}
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
	return nil
}

func (h *PipelineHandler) UpdateByID(c echo.Context) error {
	return nil
}

func (h *PipelineHandler) DeleteOne(c echo.Context) error {
	return nil
}
