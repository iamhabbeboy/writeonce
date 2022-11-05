package handler

import "github.com/labstack/echo/v4"

type PipeHandler struct {
}

func NewPipeHandler() *PipeHandler {
	return &PipeHandler{}
}

func (h *PipeHandler) Create(c echo.Context) error {
	return nil
}
