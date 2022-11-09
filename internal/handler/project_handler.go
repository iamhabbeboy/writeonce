package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/entity"
	repo "github.com/theterminalguy/writeonce/internal/repository"
)

type ProjectHandler struct {
	ProjectRepo *repo.ProjectRepository
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{
		ProjectRepo: repo.NewProjectRepository(),
	}
}

func (h *ProjectHandler) Search(c echo.Context) error {
	return nil
}

func (h *ProjectHandler) ReadAll(c echo.Context) error {
	return c.JSON(http.StatusOK, h.ProjectRepo.GetAll())
}

func (h *ProjectHandler) ReadByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	record, err := h.ProjectRepo.Get(id)
	if errors.Is(err, entity.ErrNotFound) {
		return c.String(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, record)
}

func (h *ProjectHandler) CreateOne(c echo.Context) error {
	params := new(repo.ProjectParams)
	if err := c.Bind(params); err != nil {
		return err
	}
	record, err := h.ProjectRepo.Create(*params)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, record)
}

func (h *ProjectHandler) UpdateByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	params := new(repo.ProjectParams)
	if err := c.Bind(params); err != nil {
		return err
	}
	record, err := h.ProjectRepo.Update(id, *params)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, record)
}

func (h *ProjectHandler) DeleteOne(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.ProjectRepo.Delete(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}
