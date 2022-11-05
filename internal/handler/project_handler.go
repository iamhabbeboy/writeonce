package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	repo "github.com/theterminalguy/writeonce/internal/repository"
)

type V1ProjectHandler struct {
	ProjectRepo *repo.ProjectRepository
}

func NewV1ProjectHandler() *V1ProjectHandler {
	return &V1ProjectHandler{
		ProjectRepo: repo.NewProjectRepository(),
	}
}

func (h *V1ProjectHandler) Search(c echo.Context) error {
	return nil
}

func (h *V1ProjectHandler) ReadAll(c echo.Context) error {
	return c.JSON(http.StatusOK, h.ProjectRepo.GetAll())
}

func (h *V1ProjectHandler) ReadByID(c echo.Context) error {
	return nil
}

func (h *V1ProjectHandler) CreateOne(c echo.Context) error {
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

func (h *V1ProjectHandler) UpdateByID(c echo.Context) error {
	return nil
}

func (h *V1ProjectHandler) DeleteOne(c echo.Context) error {
	return nil
}
