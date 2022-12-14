package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/entity"
	repo "github.com/theterminalguy/writeonce/internal/repository"
)

type PipeHandler struct {
	PipeRepo *repo.PipeRepository
}

func NewPipeHandler() *PipeHandler {
	return &PipeHandler{
		PipeRepo: repo.NewPipeRepository(),
	}
}

func (h *PipeHandler) Search(c echo.Context) error {
	return nil
}

func (h *PipeHandler) ReadAll(c echo.Context) error {
	return c.JSON(http.StatusOK, h.PipeRepo.GetAll())
}

func (h *PipeHandler) ReadByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	record, err := h.PipeRepo.Get(id)
	if errors.Is(err, entity.ErrNotFound) {
		return c.String(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, record)
}

func (h *PipeHandler) CreateOne(c echo.Context) error {
	params := new(entity.Pipe)
	if err := c.Bind(params); err != nil {
		return err
	}
	record, err := h.PipeRepo.Create(*params)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, record)
}

func (h *PipeHandler) UpdateByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	params := new(entity.Pipe)
	if err := c.Bind(params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	record, err := h.PipeRepo.Update(id, *params)
	if errors.Is(err, entity.ErrNotFound) {
		return c.String(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, record)
}

func (h *PipeHandler) DeleteOne(c echo.Context) error {
	return nil
}
