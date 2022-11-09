package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/theterminalguy/writeonce/internal/entity"
	repo "github.com/theterminalguy/writeonce/internal/repository"
	"github.com/theterminalguy/writeonce/internal/service"
)

type TemplateHandler struct {
	TemplateRepo    *repo.TemplateRepository
	TemplateService *service.TemplateService
}

func NewTemplateHandler() *TemplateHandler {
	return &TemplateHandler{
		TemplateRepo:    repo.NewTemplateRepository(),
		TemplateService: service.NewTemplateService(),
	}
}

func (h *TemplateHandler) Search(c echo.Context) error {
	return nil
}

func (h *TemplateHandler) ReadAll(c echo.Context) error {
	return c.JSON(http.StatusOK, h.TemplateRepo.GetAll())
}

func (h *TemplateHandler) ReadByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	record, err := h.TemplateRepo.Get(id)
	if err == entity.ErrNotFound {
		return c.String(http.StatusNotFound, fmt.Sprintf("Template with ID %s not found", id))
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, record)
}

func (h *TemplateHandler) CreateOne(c echo.Context) error {
	ct := c.Request().Header.Get(echo.HeaderContentType)

	var params repo.TemplateParams
	if strings.HasPrefix(ct, echo.MIMETextPlain) {
		id, err := uuid.Parse(c.FormValue("project_id"))
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		params.ProjectID = id
		params.Name = c.FormValue("name")
		params.Description = c.FormValue("description")

		body := c.Request().Body
		defer body.Close()
		b, err := io.ReadAll(body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		params.Template = string(b)
	} else if err := c.Bind(&params); err != nil {
		return err
	}
	record, err := h.TemplateService.CreateTemplate(params)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, record)
}

func (h *TemplateHandler) UpdateByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	params := new(repo.TemplateParams)
	if err := c.Bind(params); err != nil {
		return err
	}
	record, err := h.TemplateRepo.Update(id, *params)
	if errors.Is(err, entity.ErrNotFound) {
		return c.String(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, record)
}

func (h *TemplateHandler) DeleteOne(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	err = h.TemplateRepo.Delete(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusNoContent, nil)
}
